package userservicelogic

import (
	"context"
	"database/sql"
	"time"
	"zero-admin/rpc/model/sysmodel"
	"zero-admin/rpc/sys/sysclient"

	"zero-admin/rpc/sys/internal/svc"

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
	_ = l.svcCtx.UserModel.Update(l.ctx, &sysmodel.SysUser{
		Id:         in.Id,
		Status:     in.Status,
		UpdateBy:   sql.NullString{String: in.LastUpdateBy},
		UpdateTime: sql.NullTime{Time: time.Now()},
	})

	return &sysclient.UserStatusResp{}, nil
}
