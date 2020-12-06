package logic

import (
	"context"

	"go-zero-admin/rpc/sms/internal/svc"
	"go-zero-admin/rpc/sms/sms"

	"github.com/tal-tech/go-zero/core/logx"
)

type HomeRecommendProductUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewHomeRecommendProductUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HomeRecommendProductUpdateLogic {
	return &HomeRecommendProductUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *HomeRecommendProductUpdateLogic) HomeRecommendProductUpdate(in *sms.HomeRecommendProductUpdateReq) (*sms.HomeRecommendProductUpdateResp, error) {
	// todo: add your logic here and delete this line

	return &sms.HomeRecommendProductUpdateResp{}, nil
}
