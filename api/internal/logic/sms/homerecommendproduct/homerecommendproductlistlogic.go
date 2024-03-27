package logic

import (
	"context"
	"encoding/json"
	"github.com/feihua/zero-admin/api/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"

	"github.com/feihua/zero-admin/api/internal/svc"
	"github.com/feihua/zero-admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
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
	resp, err := l.svcCtx.HomeRecommendProductService.HomeRecommendProductList(l.ctx, &smsclient.HomeRecommendProductListReq{
		Current:         req.Current,
		PageSize:        req.PageSize,
		ProductName:     req.ProductName,
		RecommendStatus: req.RecommendStatus,
	})

	if err != nil {
		data, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("参数: %s,查询人气推荐商品列表异常:%s", string(data), err.Error())
		return nil, errorx.NewDefaultError("查询人气推荐商品失败")
	}

	var list []*types.ListtHomeRecommendProductData

	for _, item := range resp.List {
		list = append(list, &types.ListtHomeRecommendProductData{
			Id:              item.Id,
			ProductId:       item.ProductId,
			ProductName:     item.ProductName,
			RecommendStatus: item.RecommendStatus,
			Sort:            item.Sort,
		})
	}

	return &types.ListHomeRecommendProductResp{
		Current:  req.Current,
		Data:     list,
		PageSize: req.PageSize,
		Success:  true,
		Total:    resp.Total,
		Code:     "000000",
		Message:  "添加人气推荐商品成功",
	}, nil
}
