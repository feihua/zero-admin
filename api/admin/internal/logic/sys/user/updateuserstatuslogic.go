package user

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"
	"google.golang.org/grpc/status"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateUserStatusLogic 更新用户状态
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
func (l *UpdateUserStatusLogic) UpdateUserStatus(req *types.UpdateUserStatusReq) (*types.UpdateUserStatusResp, error) {

	_, err := l.svcCtx.UserService.UpdateUserStatus(l.ctx, &sysclient.UpdateUserStatusReq{
		Ids:      req.Ids,    // 用户id
		Status:   req.Status, // 用户状态
		UpdateBy: l.ctx.Value("userName").(string),
	})

	if err != nil {
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	return &types.UpdateUserStatusResp{
		Code:    "000000",
		Message: "更新用户状态成功",
	}, nil
}
