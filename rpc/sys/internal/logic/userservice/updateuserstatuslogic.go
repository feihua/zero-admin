package userservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/sys/gen/query"
	"github.com/feihua/zero-admin/rpc/sys/internal/logic/common"
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
func (l *UpdateUserStatusLogic) UpdateUserStatus(in *sysclient.UpdateUserStatusReq) (*sysclient.UpdateUserStatusResp, error) {
	// 1.排除超级管理员
	var userIds []int64
	for _, userId := range in.Ids {
		if common.IsAdmin(l.ctx, userId, l.svcCtx.DB) {
			continue
		}
		userIds = append(userIds, userId)
	}

	q := query.SysUser
	_, err := q.WithContext(l.ctx).Where(q.ID.In(userIds...)).Update(q.UserStatus, in.UserStatus)

	if err != nil {
		logc.Errorf(l.ctx, "更新用户状态异常,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("更新用户状态异常")
	}

	return &sysclient.UpdateUserStatusResp{}, nil
}
