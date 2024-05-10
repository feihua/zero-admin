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

// DeptAddLogic
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
func (l *DeptAddLogic) DeptAdd(in *sysclient.DeptAddReq) (*sysclient.DeptAddResp, error) {
	err := query.SysDept.WithContext(l.ctx).Create(&model.SysDept{
		DeptName:  in.DeptName,
		ParentID:  in.ParentId,
		OrderNum:  in.OrderNum,
		CreateBy:  in.CreateBy,
		DelFlag:   1, //0：已删除  1：正常'
		ParentIds: strings.Replace(strings.Trim(fmt.Sprint(in.ParentIds), "[]"), " ", ",", -1),
	})
	if err != nil {
		return nil, err
	}

	return &sysclient.DeptAddResp{}, nil
}
