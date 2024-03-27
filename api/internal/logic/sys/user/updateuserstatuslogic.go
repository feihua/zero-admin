package user

import (
	"context"
	"github.com/feihua/zero-admin/api/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"

	"github.com/feihua/zero-admin/api/internal/svc"
	"github.com/feihua/zero-admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateUserStatusLogic
/*
Author: LiuFeiHua
Date: 2023/12/18 13:56
*/
type UpdateUserStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateUserStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) UpdateUserStatusLogic {
	return UpdateUserStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// UpdateUserStatus 更新用户状态
func (l *UpdateUserStatusLogic) UpdateUserStatus(req types.UserStatusReq) (*types.UserStatusResp, error) {

	_, err := l.svcCtx.UserService.UpdateUserStatus(l.ctx, &sysclient.UserStatusReq{
		Id:           req.Id,
		Status:       req.Status,
		LastUpdateBy: l.ctx.Value("userName").(string),
	})

	if err != nil {
		return nil, errorx.NewDefaultError("更新用户状态异常")
	}

	return &types.UserStatusResp{
		Code:    "000000",
		Message: "更新用户状态成功",
	}, nil
}
