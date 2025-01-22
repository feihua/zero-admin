package userservicelogic

import (
	"context"
	"errors"
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
// 1.判断是不是超级管理员
// 2.更新用户状态
func (l *UpdateUserStatusLogic) UpdateUserStatus(in *sysclient.UpdateUserStatusReq) (*sysclient.UpdateUserStatusResp, error) {
	ids := in.Ids // 用户id

	for _, roleId := range ids {
		// 1.判断是不是超级管理员
		if roleId == 1 {
			return nil, errors.New("更新用户状态失败,不允许操作超级管理员用户")
		}
	}

	// 2.更新用户状态
	q := query.SysUser
	_, err := q.WithContext(l.ctx).Where(q.ID.In(ids...)).Update(q.UserStatus, in.UserStatus)

	if err != nil {
		logc.Errorf(l.ctx, "更新用户状态异常,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("更新用户状态异常")
	}

	return &sysclient.UpdateUserStatusResp{}, nil
}
