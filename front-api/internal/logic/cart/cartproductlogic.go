package cart

import (
	"context"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"

	"github.com/feihua/zero-admin/front-api/internal/svc"
	"github.com/feihua/zero-admin/front-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// CartProductLogic
/*
Author: LiuFeiHua
Date: 2023/12/6 16:07
*/
type CartProductLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCartProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CartProductLogic {
	return &CartProductLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// CartProduct 获取购物车中某个商品的规格,用于重选规格
func (l *CartProductLogic) CartProduct(req *types.CartProductReq) (resp *types.CartProductResp, err error) {
	productResp, _ := l.svcCtx.ProductService.ProductDetailById(l.ctx, &pmsclient.ProductDetailByIdReq{
		Id: req.ProductId,
	})
	return &types.CartProductResp{
		Code:    0,
		Message: "操作成功",
		Data: types.CartProductData{
			ProductAttributeList: buildProductAttributeListData(productResp),
			SkuStockList:         buildSkuStockListData(productResp),
		},
	}, nil
}

// 3.获取商品属性信息
func buildProductAttributeListData(resp *pmsclient.ProductDetailByIdResp) []types.CartItemProductAttributeList {
	var list []types.CartItemProductAttributeList
	for _, item := range resp.ProductAttributeList {
		list = append(list, types.CartItemProductAttributeList{
			Id:                         item.Id,
			ProductAttributeCategoryId: item.ProductAttributeCategoryId,
			Name:                       item.Name,
			SelectType:                 item.SelectType,
			InputType:                  item.InputType,
			InputList:                  item.InputList,
			Sort:                       item.Sort,
			FilterType:                 item.FilterType,
			SearchType:                 item.SearchType,
			RelatedStatus:              item.RelatedStatus,
			HandAddStatus:              item.HandAddStatus,
			Type:                       item.Type,
		})
	}

	return list
}

// 5.获取商品SKU库存信息
func buildSkuStockListData(resp *pmsclient.ProductDetailByIdResp) []types.CartItemSkuStockList {
	var list []types.CartItemSkuStockList
	for _, item := range resp.SkuStockList {

		list = append(list, types.CartItemSkuStockList{
			Id:             item.Id,
			ProductId:      item.ProductId,
			SkuCode:        item.SkuCode,
			Price:          item.Price,
			Stock:          item.Stock,
			LowStock:       item.LowStock,
			Pic:            item.Pic,
			Sale:           item.Sale,
			PromotionPrice: item.PromotionPrice,
			LockStock:      item.LockStock,
			SpData:         item.SpData,
		})
	}

	return list
}
