package logic

import (
	"context"
	"go-zero-admin/rpc/sys/sysclient"

	"go-zero-admin/api/internal/svc"
	"go-zero-admin/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type LoginLogDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) LoginLogDeleteLogic {
	return LoginLogDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogDeleteLogic) LoginLogDelete(req types.DeleteLoginLogReq) (*types.DeleteLoginLogResp, error) {
	_, err := l.svcCtx.Sys.LoginLogDelete(l.ctx, &sysclient.LoginLogDeleteReq{
		Id: req.Id,
	})

	if err != nil {
		return nil, err
	}

	return &types.DeleteLoginLogResp{
		Code:    "000000",
		Message: "",
	}, nil
}
