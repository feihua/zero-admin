package deptservicelogic

import (
	"context"
	"errors"
	"fmt"
	"github.com/feihua/zero-admin/rpc/sys/gen/model"
	"github.com/feihua/zero-admin/rpc/sys/gen/query"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"
	"github.com/zeromicro/go-zero/core/logc"
	"strings"

	"github.com/feihua/zero-admin/rpc/sys/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

// DeptAddLogic 添加部门信息
/*
Author: LiuFeiHua
Date: 2023/12/18 16:59
*/
type DeptAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeptAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeptAddLogic {
	return &DeptAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// DeptAdd 添加部门信息
// 1.根据部门名称查询部门是否已存在
// 2.如果部门已存在,则直接返回
// 3.部门不存在时,则直接添加部门
func (l *DeptAddLogic) DeptAdd(in *sysclient.DeptAddReq) (*sysclient.DeptAddResp, error) {
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
		DeptName:  deptName,
		ParentID:  in.ParentId,
		OrderNum:  in.OrderNum,
		CreateBy:  in.CreateBy,
		DelFlag:   1, //0：已删除  1：正常'
		ParentIds: strings.Replace(strings.Trim(fmt.Sprint(in.ParentIds), "[]"), " ", ",", -1),
	}

	err = q.Create(dept)

	if err != nil {
		logc.Errorf(l.ctx, "添加部门信息失败,参数:%+v,异常:%s", dept, err.Error())
		return nil, errors.New("添加部门信息失败")
	}

	return &sysclient.DeptAddResp{}, nil
}
