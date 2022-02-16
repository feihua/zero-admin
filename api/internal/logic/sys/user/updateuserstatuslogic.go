package logic

import (
	"context"
	"zero-admin/rpc/sys/sysclient"

	"zero-admin/api/internal/svc"
	"zero-admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

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

func (l *UpdateUserStatusLogic) UpdateUserStatus(req types.UserStatusReq) (*types.UserStatusResp, error) {
	_, _ = l.svcCtx.Sys.UpdateUserStatus(l.ctx, &sysclient.UserStatusReq{
		Id:           req.Id,
		Status:       req.Status,
		LastUpdateBy: "admin",
	})

	return &types.UserStatusResp{
		Code:    "000000",
		Message: "",
	}, nil
}
