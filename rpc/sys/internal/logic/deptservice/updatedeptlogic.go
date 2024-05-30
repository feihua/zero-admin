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

// UpdateDeptLogic 更新部门信息
/*
Author: LiuFeiHua
Date: 2023/12/18 17:01
*/
type UpdateDeptLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateDeptLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateDeptLogic {
	return &UpdateDeptLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateDept 更新部门信息
// 1.根据部门id查询部门是否已存在
// 2.部门存在时,则直接更新部门
func (l *UpdateDeptLogic) UpdateDept(in *sysclient.UpdateDeptReq) (*sysclient.UpdateDeptResp, error) {
	q := query.SysDept.WithContext(l.ctx)

	// 1.根据部门id查询部门是否已存在
	_, err := q.Where(query.SysDept.ID.Eq(in.Id)).First()

	if err != nil {
		logc.Errorf(l.ctx, "根据部门id：%d,查询部门信息失败,异常:%s", in.Id, err.Error())
		return nil, errors.New(fmt.Sprintf("查询部门信息失败"))
	}

	dept := &model.SysDept{
		ID:         in.Id,
		DeptName:   in.DeptName,
		DeptStatus: in.DeptStatus,
		DeptSort:   in.DeptSort,
		ParentID:   in.ParentId,
		Leader:     in.Leader,
		Phone:      in.Phone,
		Email:      in.Email,
		Remark:     in.Remark,
		ParentIds:  strings.Replace(strings.Trim(fmt.Sprint(in.ParentIds), "[]"), " ", ",", -1),
		UpdateBy:   in.UpdateBy,
	}

	// 2.部门存在时,则直接更新部门
	_, err = q.Where(query.SysDept.ID.Eq(in.Id)).Updates(dept)

	if err != nil {
		logc.Errorf(l.ctx, "更新部门信息失败,参数:%+v,异常:%s", dept, err.Error())
		return nil, errors.New(fmt.Sprintf("更新部门信息失败"))
	}

	return &sysclient.UpdateDeptResp{}, nil
}
