package home

import (
	"context"
	"github.com/feihua/zero-admin/api/web/internal/types"
	"github.com/feihua/zero-admin/rpc/pms/client/productservice"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"
)

// 根据商品Ids查询商品
func queryProductList(pmsClient productservice.ProductService, productIdLists []int64, ctx context.Context) []types.ProductList {
	productListResp, _ := pmsClient.ProductListByIds(ctx, &pmsclient.ProductByIdsReq{Ids: productIdLists})
	productLists := make([]types.ProductList, 0)
	for _, item := range productListResp.List {
		productLists = append(productLists, types.ProductList{
			Id:                         item.Id,
			BrandId:                    item.BrandId,
			ProductCategoryId:          item.ProductCategoryId,
			FeightTemplateId:           item.FeightTemplateId,
			ProductAttributeCategoryId: item.ProductAttributeCategoryId,
			Name:                       item.Name,
			Pic:                        item.Pic,
			ProductSn:                  item.ProductSn,
			DeleteStatus:               item.DeleteStatus,
			PublishStatus:              item.PublishStatus,
			NewStatus:                  item.NewStatus,
			RecommandStatus:            item.RecommandStatus,
			VerifyStatus:               item.VerifyStatus,
			Sort:                       item.Sort,
			Sale:                       item.Sale,
			Price:                      item.Price,
			PromotionPrice:             item.PromotionPrice,
			GiftGrowth:                 item.GiftGrowth,
			GiftPoint:                  item.GiftPoint,
			UsePointLimit:              item.UsePointLimit,
			SubTitle:                   item.SubTitle,
			OriginalPrice:              item.OriginalPrice,
			Stock:                      item.Stock,
			LowStock:                   item.LowStock,
			Unit:                       item.Unit,
			Weight:                     item.Weight,
			PreviewStatus:              item.PreviewStatus,
			ServiceIDS:                 item.ServiceIds,
			Keywords:                   item.Keywords,
			Note:                       item.Note,
			AlbumPics:                  item.AlbumPics,
			DetailTitle:                item.DetailTitle,
			PromotionStartTime:         item.PromotionStartTime,
			PromotionEndTime:           item.PromotionEndTime,
			PromotionPerLimit:          item.PromotionPerLimit,
			PromotionType:              item.PromotionType,
			BrandName:                  item.BrandName,
			ProductCategoryName:        item.ProductCategoryName,
			Description:                item.Description,
		})
	}
	return productLists
}
