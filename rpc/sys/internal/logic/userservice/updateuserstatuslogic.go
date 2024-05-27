package userservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/sys/gen/model"
	"github.com/feihua/zero-admin/rpc/sys/gen/query"
	"github.com/feihua/zero-admin/rpc/sys/internal/svc"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateUserStatusLogic 更新用户状态
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
	user := &model.SysUser{
		ID:         in.Id,
		UserStatus: in.UserStatus,
		UpdateBy:   in.UpdateBy,
	}
	_, err := q.WithContext(l.ctx).Updates(user)

	if err != nil {
		logc.Errorf(l.ctx, "更新用户状态失败,参数:%+v,异常:%s", user, err.Error())
		return nil, errors.New("更新用户状态失败")
	}

	return &sysclient.UserStatusResp{}, nil
}
