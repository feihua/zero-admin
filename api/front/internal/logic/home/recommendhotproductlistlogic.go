package home

import (
	"context"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"

	"github.com/feihua/zero-admin/api/front/internal/svc"
	"github.com/feihua/zero-admin/api/front/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// RecommendHotProductListLogic 分页获取人气推荐商品
/*
Author: LiuFeiHua
Date: 2024/5/16 15:13
*/
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

// RecommendHotProductList 分页获取人气推荐商品
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
		Message: "分页获取人气推荐商品成功",
		Data:    queryProductList(l.svcCtx.ProductService, productIds, l.ctx),
	}, nil
}
