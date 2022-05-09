package home

import (
	"context"

	"zero-admin/front-api/internal/svc"
	"zero-admin/front-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type HomeIndexLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHomeIndexLogic(ctx context.Context, svcCtx *svc.ServiceContext) HomeIndexLogic {
	return HomeIndexLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HomeIndexLogic) HomeIndex() (resp *types.ListHomeDisplayResp, err error) {
	// todo: add your logic here and delete this line

	return
}
