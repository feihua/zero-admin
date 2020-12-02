package logic

import (
	"context"

	"go-zero-admin/api/internal/svc"
	"go-zero-admin/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type HomeRecommendProductListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHomeRecommendProductListLogic(ctx context.Context, svcCtx *svc.ServiceContext) HomeRecommendProductListLogic {
	return HomeRecommendProductListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HomeRecommendProductListLogic) HomeRecommendProductList(req types.ListHomeRecommendProductReq) (*types.ListHomeRecommendProductResp, error) {
	// todo: add your logic here and delete this line

	return &types.ListHomeRecommendProductResp{}, nil
}
