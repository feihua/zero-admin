package home

import (
	"context"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"

	"github.com/feihua/zero-admin/api/front/internal/svc"
	"github.com/feihua/zero-admin/api/front/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// RecommendNewProductListLogic 分页获取新品推荐商品
/*
Author: LiuFeiHua
Date: 2024/5/16 15:12
*/
type RecommendNewProductListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRecommendNewProductListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RecommendNewProductListLogic {
	return &RecommendNewProductListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// RecommendNewProductList 分页获取新品推荐商品
func (l *RecommendNewProductListLogic) RecommendNewProductList(req *types.RecommendNewProductListReq) (resp *types.RecommendNewProductListResp, err error) {
	homeNewProductList, _ := l.svcCtx.HomeNewProductService.QueryHomeNewProductList(l.ctx, &smsclient.QueryHomeNewProductListReq{
		PageNum:         req.Current,
		PageSize:        req.PageSize,
		RecommendStatus: 1, //推荐状态：0->不推荐;1->推荐
	})

	var productIds []int64
	for _, item := range homeNewProductList.List {
		productIds = append(productIds, item.ProductId)
	}

	return &types.RecommendNewProductListResp{
		Code:    0,
		Message: "分页获取新品推荐商品成功",
		Data:    queryProductList(l.svcCtx.ProductService, productIds, l.ctx),
	}, nil
}
