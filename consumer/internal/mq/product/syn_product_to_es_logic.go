package product

import (
	"context"

	"github.com/bytedance/sonic"
	"github.com/feihua/zero-admin/rpc/pms/client/productspuservice"
	"github.com/feihua/zero-admin/rpc/search/search_client"
	"github.com/zeromicro/go-zero/core/logc"
)

// SynProductToEs 同步商品到es
func SynProductToEs(ctx context.Context, body []byte, Search search_client.Search, productSpuService productspuservice.ProductSpuService) {
	logc.Infof(ctx, "需要同步商品的id: %s", body)
	var orderInfo map[string]int64
	err := sonic.Unmarshal(body, &orderInfo)
	if err != nil {
		logc.Errorf(ctx, "序列化 JSON 失败: %v", err)
		return
	}
	id := orderInfo["id"]

	res, err := productSpuService.QueryProductSpuDetail(ctx, &productspuservice.QueryProductSpuDetailReq{
		Id: id,
	})
	if err != nil {
		logc.Errorf(ctx, "查询商品异常,请求参数: %s, 异常信息: %+v", body, err)
		return
	}
	logc.Infof(ctx, "需要同步商品的信息: %v", res)

	product := res.Data
	product1 := &search_client.ProductData{
		Id:           product.Id,           // 商品SpuId
		Name:         product.Name,         // 商品名称
		CategoryId:   product.CategoryId,   // 商品分类ID
		CategoryIds:  product.CategoryIds,  // 商品分类ID集合
		CategoryName: product.CategoryName, // 商品分类名称
		BrandId:      product.BrandId,      // 品牌ID
		BrandName:    product.BrandName,    // 品牌名称
		Unit:         product.Unit,         // 单位
		Weight:       product.Weight,       // 重量(kg)
		Keywords:     product.Keywords,     // 关键词
		// Brief:               product.Brief,               // 简介
		// Description:         product.Description,         // 详细描述
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
		// DetailTitle:         product.DetailTitle,         // 详情标题
		// DetailDesc:          product.DetailDesc,          // 详情描述
		DetailHtml:       product.DetailHtml,       // 产品详情网页内容
		DetailMobileHtml: product.DetailMobileHtml, // 移动端网页详情
		CreateBy:         product.CreateBy,         // 创建人ID
		CreateTime:       product.CreateTime,       // 创建时间
		UpdateBy:         product.UpdateBy,         // 更新人ID
		UpdateTime:       product.UpdateTime,       // 更新时间
	}

	var list []*search_client.ProductData
	list = append(list, product1)
	_, err = Search.Create(ctx, &search_client.CreateReq{
		Data: list,
	})
	if err != nil {
		logc.Errorf(ctx, "同步商品到es失败,请求参数：%+v,错误信息：%+v", body, err)
	}

}
