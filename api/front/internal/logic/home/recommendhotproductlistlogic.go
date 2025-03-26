package home

import (
	"context"
	"github.com/feihua/zero-admin/pkg/errorx"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc/status"

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
	homeRecommendProductList, err := l.svcCtx.HomeRecommendProductService.QueryHomeRecommendProductList(l.ctx, &smsclient.QueryHomeRecommendProductListReq{
		PageNum:         req.Current,
		PageSize:        req.PageSize,
		RecommendStatus: 1, // 推荐状态：0->不推荐;1->推荐
	})
	if err != nil {
		logc.Errorf(l.ctx, "分页获取人气推荐商品失败,参数: %+v,异常：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

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
