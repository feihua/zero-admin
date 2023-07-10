package home

import (
	"context"
	"zero-admin/rpc/sms/smsclient"

	"zero-admin/front-api/internal/svc"
	"zero-admin/front-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RecommendHotProductListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRecommendHotProductListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RecommendHotProductListLogic {
	return &RecommendHotProductListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RecommendHotProductListLogic) RecommendHotProductList(req *types.RecommendHotProductListReq) (resp *types.RecommendHotProductListResp, err error) {
	homeRecommendProductList, _ := l.svcCtx.HomeRecommendProductService.HomeRecommendProductList(l.ctx, &smsclient.HomeRecommendProductListReq{
		Current:         req.Current,
		PageSize:        req.PageSize,
		RecommendStatus: 1, //推荐状态：0->不推荐;1->推荐
	})

	var productIds []int64
	for _, item := range homeRecommendProductList.List {
		productIds = append(productIds, item.ProductId)
	}

	return &types.RecommendHotProductListResp{
		Code:    0,
		Message: "操作成功",
		Data:    queryProductList(l.svcCtx.ProductService, productIds, l.ctx),
	}, nil
}
