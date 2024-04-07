package home

import (
	"context"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"

	"github.com/feihua/zero-admin/api/front/internal/svc"
	"github.com/feihua/zero-admin/api/front/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

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

func (l *RecommendNewProductListLogic) RecommendNewProductList(req *types.RecommendNewProductListReq) (resp *types.RecommendNewProductListResp, err error) {
	homeNewProductList, _ := l.svcCtx.HomeNewProductService.HomeNewProductList(l.ctx, &smsclient.HomeNewProductListReq{
		Current:         req.Current,
		PageSize:        req.PageSize,
		RecommendStatus: 1, //推荐状态：0->不推荐;1->推荐
	})

	var productIds []int64
	for _, item := range homeNewProductList.List {
		productIds = append(productIds, item.ProductId)
	}

	return &types.RecommendNewProductListResp{
		Code:    0,
		Message: "操作成功",
		Data:    queryProductList(l.svcCtx.ProductService, productIds, l.ctx),
	}, nil
}
