package deptservicelogic

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/feihua/zero-admin/rpc/model/sysmodel"
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
	dept, err := l.svcCtx.DeptModel.FindOne(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}
	err = l.svcCtx.DeptModel.Update(l.ctx, &sysmodel.SysDept{
		Id:         in.Id,
		Name:       in.Name,
		ParentId:   in.ParentId,
		OrderNum:   in.OrderNum,
		CreateBy:   dept.CreateBy,
		CreateTime: dept.CreateTime,
		UpdateBy:   sql.NullString{String: in.LastUpdateBy, Valid: true},
		DelFlag:    in.DelFlag,
		ParentIds:  strings.Replace(strings.Trim(fmt.Sprint(in.ParentIds), "[]"), " ", ",", -1),
	})

	if err != nil {
		return nil, err
	}

	return &sysclient.DeptUpdateResp{}, nil
}
