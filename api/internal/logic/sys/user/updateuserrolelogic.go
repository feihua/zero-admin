package logic

import (
	"context"
	"go-zero-admin/rpc/sys/sysclient"

	"go-zero-admin/api/internal/svc"
	"go-zero-admin/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type UpdateUserRoleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateUserRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) UpdateUserRoleLogic {
	return UpdateUserRoleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateUserRoleLogic) UpdateUserRole(req types.UpdateUserRoleReq) (*types.UpdateUserRoleResp, error) {
	_, err := l.svcCtx.Sys.UpdateUserRole(l.ctx, &sysclient.UpdateUserRoleReq{
		Id:     req.Id,
		RoleId: req.RoleId,
	})

	if err != nil {
		return nil, err
	}

	return &types.UpdateUserRoleResp{}, nil
}
