package product

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/cms/cmsclient"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProductUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProductUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) ProductUpdateLogic {
	return ProductUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// ProductUpdate 更新商品
// 逻辑步骤：
// 1.更新商品基本信息
// 2.会员价格
// 3.阶梯价格
// 4.满减价格
// 5.更新sku库存信息
// 6.更新商品参数,添加自定义商品规格
// 7.关联专题
// 8.关联优选
// 注意: 步骤1到6是在商品(rpc)中操作,步骤7到8是在cms模块中操作
func (l *ProductUpdateLogic) ProductUpdate(req types.UpdateProductReq) (*types.UpdateProductResp, error) {
	product := req.ProductData
	productId := product.Id
	_, err := l.svcCtx.ProductService.UpdateProduct(l.ctx, &pmsclient.UpdateProductReq{
		Id:                         productId,
		BrandId:                    product.BrandId,                    // 品牌id
		ProductCategoryId:          product.ProductCategoryId,          // 商品分类id
		FeightTemplateId:           product.FeightTemplateId,           // 商品运费模板id
		ProductAttributeCategoryId: product.ProductAttributeCategoryId, // 商品属性分类id
		Name:                       product.Name,                       // 商品名称
		Pic:                        product.Pic,                        // 商品图片
		ProductSn:                  product.ProductSn,                  // 货号
		DeleteStatus:               product.DeleteStatus,               // 删除状态：0->未删除；1->已删除
		PublishStatus:              product.PublishStatus,              // 上架状态：0->下架；1->上架
		NewStatus:                  product.NewStatus,                  // 新品状态:0->不是新品；1->新品
		RecommandStatus:            product.RecommandStatus,            // 推荐状态；0->不推荐；1->推荐
		VerifyStatus:               product.VerifyStatus,               // 审核状态：0->未审核；1->审核通过
		Sort:                       product.Sort,                       // 排序
		Sale:                       product.Sale,                       // 销量
		Price:                      product.Price,                      // 商品价格
		PromotionPrice:             product.PromotionPrice,             // 促销价格
		GiftGrowth:                 product.GiftGrowth,                 // 赠送的成长值
		GiftPoint:                  product.GiftPoint,                  // 赠送的积分
		UsePointLimit:              product.UsePointLimit,              // 限制使用的积分数
		SubTitle:                   product.SubTitle,                   // 副标题
		Description:                product.Description,                // 商品描述
		OriginalPrice:              product.OriginalPrice,              // 市场价
		Stock:                      product.Stock,                      // 库存
		LowStock:                   product.LowStock,                   // 库存预警值
		Unit:                       product.Unit,                       // 单位
		Weight:                     product.Weight,                     // 商品重量，默认为克
		PreviewStatus:              product.PreviewStatus,              // 是否为预告商品：0->不是；1->是
		ServiceIds:                 product.ServiceIds,                 // 以逗号分割的产品服务：1->无忧退货；2->快速退款；3->免费包邮
		Keywords:                   product.Keywords,                   // 搜索关键字
		Note:                       product.Note,                       // 备注
		AlbumPics:                  product.AlbumPics,                  // 画册图片，连产品图片限制为5张，以逗号分割
		DetailTitle:                product.DetailTitle,                // 详情标题
		DetailDesc:                 product.DetailDesc,                 // 详情描述
		DetailHtml:                 product.DetailHtml,                 // 产品详情网页内容
		DetailMobileHtml:           product.DetailMobileHtml,           // 移动端网页详情
		PromotionStartTime:         product.PromotionStartTime,         // 促销开始时间
		PromotionEndTime:           product.PromotionEndTime,           // 促销结束时间
		PromotionPerLimit:          product.PromotionPerLimit,          // 活动限购数量
		PromotionType:              product.PromotionType,              // 促销类型：0->没有促销使用原价;1->使用促销价；2->使用会员价；3->使用阶梯价格；4->使用满减价格；5->限时购
		BrandName:                  product.BrandName,                  // 品牌名称
		ProductCategoryName:        product.ProductCategoryName,        // 商品分类名称
		ProductCategoryIdArray:     product.ProductCategoryIdArray,     // 商品分类id字符串
		MemberPriceList:            buildUpdateMemberPriceList(req),
		ProductAttributeValueList:  buildUpdateProductAttributeValueList(req),
		ProductFullReductionList:   buildUpdateProductFullReductionList(req),
		ProductLadderList:          buildUpdateProductLadderList(req),
		SkuStockList:               buildUpdateSkuStockList(req),
	})

	if err != nil {
		logc.Errorf(l.ctx, "更新商品信息失败,参数：%+v,响应：%s", req, err.Error())
		return nil, errorx.NewDefaultError("更新商品信息失败")
	}

	// 7.关联专题
	updateSubjectProductRelation(req, l, productId)

	// 8.关联优选
	updatePreferredAreaProductRelation(req, l, productId)

	return &types.UpdateProductResp{
		Code:    "000000",
		Message: "更新商品信息成功",
	}, nil
}

// 商品会员价格设置
func buildUpdateMemberPriceList(req types.UpdateProductReq) []*pmsclient.MemberPriceList {
	var memberPriceLists []*pmsclient.MemberPriceList
	for _, item := range req.MemberPriceList {
		memberPriceLists = append(memberPriceLists, &pmsclient.MemberPriceList{
			MemberLevelId:   item.MemberLevelId,
			MemberPrice:     item.MemberPrice,
			MemberLevelName: item.MemberLevelName,
		})
	}
	return memberPriceLists
}

// 商品参数及自定义规格属性
func buildUpdateProductAttributeValueList(req types.UpdateProductReq) []*pmsclient.ProductAttributeValueList {
	var attributeValueLists []*pmsclient.ProductAttributeValueList
	for _, item := range req.ProductAttributeValueList {
		attributeValueLists = append(attributeValueLists, &pmsclient.ProductAttributeValueList{
			ProductAttributeId: item.ProductAttributeId,
			AttributeValues:    item.Value,
		})
	}
	return attributeValueLists
}

// 商品满减价格设置
func buildUpdateProductFullReductionList(req types.UpdateProductReq) []*pmsclient.ProductFullReductionList {
	var fullReductionLists []*pmsclient.ProductFullReductionList
	for _, item := range req.ProductFullReductionList {
		fullReductionLists = append(fullReductionLists, &pmsclient.ProductFullReductionList{
			FullPrice:   item.FullPrice,
			ReducePrice: item.ReducePrice,
		})
	}
	return fullReductionLists
}

// 商品阶梯价格设置
func buildUpdateProductLadderList(req types.UpdateProductReq) []*pmsclient.ProductLadderList {
	var ladderLists []*pmsclient.ProductLadderList
	for _, item := range req.ProductLadderList {
		ladderLists = append(ladderLists, &pmsclient.ProductLadderList{
			Count:    item.Count,
			Discount: item.Discount,
			Price:    item.Price,
		})
	}
	return ladderLists
}

// 商品的sku库存信息
func buildUpdateSkuStockList(req types.UpdateProductReq) []*pmsclient.SkuStockList {
	var skuStockLists []*pmsclient.SkuStockList
	for _, item := range req.SkuStockList {
		skuStockLists = append(skuStockLists, &pmsclient.SkuStockList{
			SkuCode:        item.SkuCode,
			Price:          item.Price,
			Stock:          item.Stock,
			LowStock:       item.LowStock,
			Pic:            item.Pic,
			Sale:           item.Sale,
			PromotionPrice: item.PromotionPrice,
			LockStock:      item.LockStock,
			SpData:         item.SPData,
		})
	}
	return skuStockLists
}

// 更新专题关联
func updateSubjectProductRelation(req types.UpdateProductReq, l *ProductUpdateLogic, productId int64) {
	// 1.先删除专题的关联
	// _, _ = l.svcCtx.SubjectProductRelationService.SubjectProductRelationDelete(l.ctx, &cmsclient.SubjectProductRelationDeleteReq{Id: productId})

	// 2.重新添加专题的关联
	_, _ = l.svcCtx.SubjectProductRelationService.AddSubjectProductRelation(l.ctx, &cmsclient.AddSubjectProductRelationReq{
		SubjectId: req.SubjectProductRelationList,
		ProductId: productId,
	})
}

// 更新优选商品关联
func updatePreferredAreaProductRelation(req types.UpdateProductReq, l *ProductUpdateLogic, productId int64) {
	// 1.先删除优选商品的关联(由rpc实现)
	// _, _ = l.svcCtx.PrefrenceAreaProductRelationService.PrefrenceAreaProductRelationDelete(l.ctx, &cmsclient.PrefrenceAreaProductRelationDeleteReq{Id: productId})

	// 2.重新添加优选商品的关联
	_, _ = l.svcCtx.PreferredAreaProductRelationService.AddPreferredAreaProductRelation(l.ctx, &cmsclient.AddPreferredAreaProductRelationReq{
		PreferredAreaId: req.PrefrenceAreaProductRelationList,
		ProductId:       int64(productId),
	})

}
