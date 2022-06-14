package logic

import (
	"context"

	"zero-admin/rpc/pms/internal/svc"
	"zero-admin/rpc/pms/pmsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type HomeIndexLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewHomeIndexLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HomeIndexLogic {
	return &HomeIndexLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *HomeIndexLogic) HomeIndex(in *pmsclient.HomeIndexReq) (*pmsclient.HomeIndexResp, error) {
	// todo: add your logic here and delete this line

	return &pmsclient.HomeIndexResp{}, nil
}
