package logic

import (
	"context"

	"go-zero-admin/api/internal/svc"
	"go-zero-admin/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type ConfigAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewConfigAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) ConfigAddLogic {
	return ConfigAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ConfigAddLogic) ConfigAdd(req types.AddConfigReq) (*types.AddConfigResp, error) {
	// todo: add your logic here and delete this line

	return &types.AddConfigResp{}, nil
}