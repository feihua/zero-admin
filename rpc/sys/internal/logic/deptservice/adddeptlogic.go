package deptservicelogic

import (
	"context"
	"errors"
	"fmt"
	"github.com/feihua/zero-admin/rpc/sys/gen/model"
	"github.com/feihua/zero-admin/rpc/sys/gen/query"
	"github.com/feihua/zero-admin/rpc/sys/internal/svc"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"
	"github.com/zeromicro/go-zero/core/logc"
	"strings"

	"github.com/zeromicro/go-zero/core/logx"
)

// AddDeptLogic 添加部门信息
/*
Author: LiuFeiHua
Date: 2023/12/18 16:59
*/
type AddDeptLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddDeptLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddDeptLogic {
	return &AddDeptLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// AddDept 添加部门信息
// 1.根据部门名称查询部门是否已存在
// 2.如果部门已存在,则直接返回
// 3.部门不存在时,则直接添加部门
func (l *AddDeptLogic) AddDept(in *sysclient.AddDeptReq) (*sysclient.AddDeptResp, error) {
	q := query.SysDept.WithContext(l.ctx)

	//1.根据部门名称查询部门是否已存在
	deptName := in.DeptName
	count, err := q.Where(query.SysDept.DeptName.Eq(deptName)).Count()

	if err != nil {
		logc.Errorf(l.ctx, "根据部门名称：%s,查询部门信息失败,异常:%s", deptName, err.Error())
		return nil, errors.New(fmt.Sprintf("查询部门信息失败"))
	}

	//2.如果部门已存在,则直接返回
	if count > 0 {
		logc.Errorf(l.ctx, "部门信息已存在：%+v", in)
		return nil, errors.New(fmt.Sprintf("部门：%s,已存在", deptName))
	}

	//3.部门不存在时,则直接添加部门
	dept := &model.SysDept{
		DeptName:   deptName,
		DeptStatus: in.DeptStatus,
		DeptSort:   in.DeptSort,
		ParentID:   in.ParentId,
		Leader:     in.Leader,
		Phone:      in.Phone,
		Email:      in.Email,
		Remark:     in.Remark,
		IsDeleted:  0,
		ParentIds:  strings.Replace(strings.Trim(fmt.Sprint(in.ParentIds), "[]"), " ", ",", -1),
		CreateBy:   in.CreateBy,
	}

	err = q.Create(dept)

	if err != nil {
		logc.Errorf(l.ctx, "添加部门信息失败,参数:%+v,异常:%s", dept, err.Error())
		return nil, errors.New("添加部门信息失败")
	}

	return &sysclient.AddDeptResp{}, nil
}
