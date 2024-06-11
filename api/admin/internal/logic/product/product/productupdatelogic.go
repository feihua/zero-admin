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
	_, err := l.svcCtx.ProductService.ProductUpdate(l.ctx, &pmsclient.ProductUpdateReq{
		Id:                         productId,
		BrandId:                    product.BrandId,
		ProductCategoryId:          product.ProductCategoryId,
		ProductCategoryIdArray:     product.ProductCategoryIdArray,
		FeightTemplateId:           product.FeightTemplateId,
		ProductAttributeCategoryId: product.ProductAttributeCategoryId,
		Name:                       product.Name,
		Pic:                        product.Pic,
		ProductSn:                  product.ProductSn,
		DeleteStatus:               product.DeleteStatus,
		PublishStatus:              product.PublishStatus,
		NewStatus:                  product.NewStatus,
		RecommandStatus:            product.RecommandStatus,
		VerifyStatus:               product.VerifyStatus,
		Sort:                       product.Sort,
		Sale:                       product.Sale,
		Price:                      product.Price,
		PromotionPrice:             product.PromotionPrice,
		GiftGrowth:                 product.GiftGrowth,
		GiftPoint:                  product.GiftPoint,
		UsePointLimit:              product.UsePointLimit,
		SubTitle:                   product.SubTitle,
		Description:                product.Description,
		OriginalPrice:              product.OriginalPrice,
		Stock:                      product.Stock,
		LowStock:                   product.LowStock,
		Unit:                       product.Unit,
		Weight:                     product.Weight,
		PreviewStatus:              product.PreviewStatus,
		ServiceIds:                 product.ServiceIds,
		Keywords:                   product.Keywords,
		Note:                       product.Note,
		AlbumPics:                  product.AlbumPics,
		DetailTitle:                product.DetailTitle,
		DetailDesc:                 product.DetailDesc,
		DetailHtml:                 product.DetailHTML,
		DetailMobileHtml:           product.DetailMobileHTML,
		PromotionStartTime:         product.PromotionStartTime,
		PromotionEndTime:           product.PromotionEndTime,
		PromotionPerLimit:          product.PromotionPerLimit,
		PromotionType:              product.PromotionType,
		BrandName:                  product.BrandName,
		ProductCategoryName:        product.ProductCategoryName,
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

	//7.关联专题
	updateSubjectProductRelation(req, l, productId)

	//8.关联优选
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
	//1.先删除专题的关联
	//_, _ = l.svcCtx.SubjectProductRelationService.SubjectProductRelationDelete(l.ctx, &cmsclient.SubjectProductRelationDeleteReq{Id: productId})

	//2.重新添加专题的关联
	_, _ = l.svcCtx.SubjectProductRelationService.AddSubjectProductRelation(l.ctx, &cmsclient.AddSubjectProductRelationReq{
		SubjectId: req.SubjectProductRelationList,
		ProductId: productId,
	})
}

// 更新优选商品关联
func updatePreferredAreaProductRelation(req types.UpdateProductReq, l *ProductUpdateLogic, productId int64) {
	//1.先删除优选商品的关联(由rpc实现)
	//_, _ = l.svcCtx.PrefrenceAreaProductRelationService.PrefrenceAreaProductRelationDelete(l.ctx, &cmsclient.PrefrenceAreaProductRelationDeleteReq{Id: productId})

	//2.重新添加优选商品的关联
	_, _ = l.svcCtx.PreferredAreaProductRelationService.AddPreferredAreaProductRelation(l.ctx, &cmsclient.AddPreferredAreaProductRelationReq{
		PreferredAreaId: req.PrefrenceAreaProductRelationList,
		ProductId:       int64(productId),
	})

}
