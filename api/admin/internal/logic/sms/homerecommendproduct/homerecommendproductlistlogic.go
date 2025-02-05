package homerecommendproduct

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"strings"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// HomeRecommendProductListLogic 人气推荐商品
/*
Author: LiuFeiHua
Date: 2024/5/14 9:40
*/
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

// HomeRecommendProductList 查询人气推荐商品
func (l *HomeRecommendProductListLogic) HomeRecommendProductList(req types.ListHomeRecommendProductReq) (*types.ListHomeRecommendProductResp, error) {
	resp, err := l.svcCtx.HomeRecommendProductService.QueryHomeRecommendProductList(l.ctx, &smsclient.QueryHomeRecommendProductListReq{
		PageNum:         req.Current,
		PageSize:        req.PageSize,
		ProductName:     strings.TrimSpace(req.ProductName),
		RecommendStatus: req.RecommendStatus,
	})

	if err != nil {
		logc.Errorf(l.ctx, "参数: %+v,查询人气推荐商品列表异常:%s", req, err.Error())
		return nil, errorx.NewDefaultError("查询人气推荐商品失败")
	}

	var list []*types.ListHomeRecommendProductData

	for _, item := range resp.List {
		list = append(list, &types.ListHomeRecommendProductData{
			Id:              item.Id,              //
			ProductId:       item.ProductId,       // 商品id
			ProductName:     item.ProductName,     // 商品名称
			RecommendStatus: item.RecommendStatus, // 推荐状态：0->不推荐;1->推荐
			Sort:            item.Sort,            // 排序
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
