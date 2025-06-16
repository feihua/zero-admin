package home

import (
	"context"
	"github.com/feihua/zero-admin/api/front/internal/types"
	"github.com/feihua/zero-admin/rpc/pms/client/productspuservice"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"
)

// 根据商品Ids查询商品
func queryProductList(pmsClient productspuservice.ProductSpuService, productIdLists []int64, ctx context.Context) []types.ProductList {
	productListResp, _ := pmsClient.QueryProductSpuListByIds(ctx, &pmsclient.QueryProductSpuByIdsReq{Ids: productIdLists})
	productLists := make([]types.ProductList, 0)
	for _, product := range productListResp.List {
		productLists = append(productLists, types.ProductList{
			Id:                  product.Id,                  // 商品SpuId
			Name:                product.Name,                // 商品名称
			CategoryId:          product.CategoryId,          // 商品分类ID
			CategoryIds:         product.CategoryIds,         // 商品分类ID集合
			CategoryName:        product.CategoryName,        // 商品分类名称
			BrandId:             product.BrandId,             // 品牌ID
			BrandName:           product.BrandName,           // 品牌名称
			Unit:                product.Unit,                // 单位
			Weight:              product.Weight,              // 重量(kg)
			Keywords:            product.Keywords,            // 关键词
			Brief:               product.Brief,               // 简介
			Description:         product.Description,         // 详细描述
			AlbumPics:           product.AlbumPics,           // 画册图片，最多8张，以逗号分割
			MainPic:             product.MainPic,             // 主图
			PriceRange:          product.PriceRange,          // 价格区间
			PublishStatus:       product.PublishStatus,       // 上架状态：0-下架，1-上架
			NewStatus:           product.NewStatus,           // 新品状态:0->不是新品；1->新品
			RecommendStatus:     product.RecommendStatus,     // 推荐状态；0->不推荐；1->推荐
			VerifyStatus:        product.VerifyStatus,        // 审核状态：0->未审核；1->审核通过
			PreviewStatus:       product.PreviewStatus,       // 是否为预告商品：0->不是；1->是
			Sort:                product.Sort,                // 排序
			NewStatusSort:       product.NewStatusSort,       // 新品排序
			RecommendStatusSort: product.RecommendStatusSort, // 推荐排序
			Sales:               product.Sales,               // 销量
			Stock:               product.Stock,               // 库存
			LowStock:            product.LowStock,            // 预警库存
			PromotionType:       product.PromotionType,       // 促销类型：0->没有促销使用原价;1->使用促销价；2->使用会员价；3->使用阶梯价格；4->使用满减价格；5->秒杀
			DetailTitle:         product.DetailTitle,         // 详情标题
			DetailDesc:          product.DetailDesc,          // 详情描述
			DetailHtml:          product.DetailHtml,          // 产品详情网页内容
			DetailMobileHtml:    product.DetailMobileHtml,    // 移动端网页详情
			CreateBy:            product.CreateBy,            // 创建人ID
			CreateTime:          product.CreateTime,          // 创建时间
			UpdateBy:            product.UpdateBy,            // 更新人ID
			UpdateTime:          product.UpdateTime,          // 更新时间
		})
	}
	return productLists
}
