package logic

import (
	"context"
	"go-zero-admin/rpc/sms/smsclient"

	"go-zero-admin/api/internal/svc"
	"go-zero-admin/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type HomeRecommendProductDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHomeRecommendProductDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) HomeRecommendProductDeleteLogic {
	return HomeRecommendProductDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HomeRecommendProductDeleteLogic) HomeRecommendProductDelete(req types.DeleteHomeRecommendProductReq) (*types.DeleteHomeRecommendProductResp, error) {
	_, _ = l.svcCtx.Sms.HomeRecommendProductDelete(l.ctx, &smsclient.HomeRecommendProductDeleteReq{
		Id: req.Id,
	})

	return &types.DeleteHomeRecommendProductResp{
		Code:    "000000",
		Message: "",
	}, nil
}
