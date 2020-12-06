package logic

import (
	"context"

	"go-zero-admin/rpc/sms/internal/svc"
	"go-zero-admin/rpc/sms/sms"

	"github.com/tal-tech/go-zero/core/logx"
)

type HomeRecommendProductListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewHomeRecommendProductListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HomeRecommendProductListLogic {
	return &HomeRecommendProductListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *HomeRecommendProductListLogic) HomeRecommendProductList(in *sms.HomeRecommendProductListReq) (*sms.HomeRecommendProductListResp, error) {
	// todo: add your logic here and delete this line

	return &sms.HomeRecommendProductListResp{}, nil
}
