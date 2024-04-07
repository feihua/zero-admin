package home

import (
	"context"

	"github.com/feihua/zero-admin/api/web/internal/svc"
	"github.com/feihua/zero-admin/api/web/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type HomeIndexLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHomeIndexLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HomeIndexLogic {
	return &HomeIndexLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HomeIndexLogic) HomeIndex() (resp *types.HomeResp, err error) {
	// todo: add your logic here and delete this line

	return
}
