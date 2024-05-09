package userservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/sys/gen/model"
	"github.com/feihua/zero-admin/rpc/sys/gen/query"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"
	"time"

	"github.com/feihua/zero-admin/rpc/sys/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateUserStatusLogic
/*
Author: LiuFeiHua
Date: 2023/12/18 14:13
*/
type UpdateUserStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUserStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserStatusLogic {
	return &UpdateUserStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateUserStatus 更新用户状态
func (l *UpdateUserStatusLogic) UpdateUserStatus(in *sysclient.UserStatusReq) (*sysclient.UserStatusResp, error) {
	q := query.SysUser
	now := time.Now()
	_, err := q.WithContext(l.ctx).Updates(&model.SysUser{
		ID:         in.Id,
		Status:     in.Status,
		UpdateBy:   &in.UpdateBy,
		UpdateTime: &now,
	})

	if err != nil {
		return nil, err
	}

	return &sysclient.UserStatusResp{}, nil
}
