package logic

import (
	"context"
	"go-zero-admin/rpc/sms/smsclient"

	"go-zero-admin/api/internal/svc"
	"go-zero-admin/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type HomeRecommendProductAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHomeRecommendProductAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) HomeRecommendProductAddLogic {
	return HomeRecommendProductAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HomeRecommendProductAddLogic) HomeRecommendProductAdd(req types.AddHomeRecommendProductReq) (*types.AddHomeRecommendProductResp, error) {
	_, err := l.svcCtx.Sms.HomeRecommendProductAdd(l.ctx, &smsclient.HomeRecommendProductAddReq{
		ProductId:       req.ProductId,
		ProductName:     req.ProductName,
		RecommendStatus: req.RecommendStatus,
		Sort:            req.Sort,
	})

	if err != nil {
		return nil, err
	}

	return &types.AddHomeRecommendProductResp{
		Code:    "000000",
		Message: "",
	}, nil
}
