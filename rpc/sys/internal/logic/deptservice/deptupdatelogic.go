package deptservicelogic

import (
	"context"
	"fmt"
	"github.com/feihua/zero-admin/rpc/sys/gen/model"
	"github.com/feihua/zero-admin/rpc/sys/gen/query"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"
	"strings"

	"github.com/feihua/zero-admin/rpc/sys/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

// DeptUpdateLogic
/*
Author: LiuFeiHua
Date: 2023/12/18 17:01
*/
type DeptUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeptUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeptUpdateLogic {
	return &DeptUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// DeptUpdate 更新部门信息
func (l *DeptUpdateLogic) DeptUpdate(in *sysclient.DeptUpdateReq) (*sysclient.DeptUpdateResp, error) {
	q := query.SysDept
	dept, err := q.WithContext(l.ctx).Where(q.ID.Eq(in.Id)).First()
	if err != nil {
		return nil, err
	}
	_, err = q.WithContext(l.ctx).Updates(&model.SysDept{
		ID:         in.Id,
		Name:       in.Name,
		ParentID:   in.ParentId,
		OrderNum:   in.OrderNum,
		CreateBy:   dept.CreateBy,
		CreateTime: dept.CreateTime,
		UpdateBy:   &in.UpdateBy,
		DelFlag:    in.DelFlag,
		ParentIds:  strings.Replace(strings.Trim(fmt.Sprint(in.ParentIds), "[]"), " ", ",", -1),
	})

	if err != nil {
		return nil, err
	}

	return &sysclient.DeptUpdateResp{}, nil
}
