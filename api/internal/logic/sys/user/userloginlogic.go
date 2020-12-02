package logic

import (
	"context"
	"go-zero-admin/rpc/sys/sysclient"

	"go-zero-admin/api/internal/svc"
	"go-zero-admin/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type UserLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) UserLoginLogic {
	return UserLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserLoginLogic) UserLogin() (*types.LoginResp, error) {
	resp, err := l.svcCtx.Sys.Login(l.ctx, &sysclient.LoginReq{})

	if err != nil {
		return nil, err
	}
	return &types.LoginResp{
		Status:           resp.Status,
		CurrentAuthority: resp.CurrentAuthority,
	}, nil
}
