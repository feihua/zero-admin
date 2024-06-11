package product

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"
	"github.com/feihua/zero-admin/rpc/cms/cmsclient"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProductAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProductAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) ProductAddLogic {
	return ProductAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// ProductAdd 添加商品
// 逻辑步骤：
// 1.创建商品基本信息
// 2.会员价格
// 3.阶梯价格
// 4.满减价格
// 5.添加sku库存信息
// 6.添加商品参数,添加自定义商品规格
// 7.关联专题
// 8.关联优选
// 注意: 步骤1到6是在商品(rpc)中操作,步骤7到8是在cms模块中操作
func (l *ProductAddLogic) ProductAdd(req types.AddProductReq) (*types.AddProductResp, error) {
	//添加商品成功后返回商品id,然后关联专题和关联优选(因为这二个表在cms模块中,需要通过rpc去调用)
	product := req.ProductData
	result, err := l.svcCtx.ProductService.ProductAdd(l.ctx, &pmsclient.ProductAddReq{
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
		MemberPriceList:            buildMemberPriceList(req),
		ProductAttributeValueList:  buildProductAttributeValueList(req),
		ProductFullReductionList:   buildProductFullReductionList(req),
		ProductLadderList:          buildProductLadderList(req),
		SkuStockList:               buildSkuStockList(req),
	})

	if err != nil {
		logc.Errorf(l.ctx, "添加商品信息失败,参数：%+v,响应：%s", req, err.Error())
		return nil, errorx.NewDefaultError("添加商品信息失败")
	}

	//商品id
	productId := result.ProductId
	//7.关联专题
	addSubjectProductRelation(req, l, productId)

	//8.关联优选
	addPreferredAreaProductRelation(req, l, productId)

	return &types.AddProductResp{
		Code:    "000000",
		Message: "添加商品信息成功",
	}, nil
}

// 商品会员价格设置
func buildMemberPriceList(req types.AddProductReq) []*pmsclient.MemberPriceList {
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
func buildProductAttributeValueList(req types.AddProductReq) []*pmsclient.ProductAttributeValueList {
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
func buildProductFullReductionList(req types.AddProductReq) []*pmsclient.ProductFullReductionList {
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
func buildProductLadderList(req types.AddProductReq) []*pmsclient.ProductLadderList {
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
func buildSkuStockList(req types.AddProductReq) []*pmsclient.SkuStockList {
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

// 添加专题关联
func addSubjectProductRelation(req types.AddProductReq, l *ProductAddLogic, productId int64) {
	_, _ = l.svcCtx.SubjectProductRelationService.AddSubjectProductRelation(l.ctx, &cmsclient.AddSubjectProductRelationReq{
		SubjectId: req.SubjectProductRelationList,
		ProductId: productId,
	})
}

// 添加优选商品关联
func addPreferredAreaProductRelation(req types.AddProductReq, l *ProductAddLogic, productId int64) {
	_, _ = l.svcCtx.PreferredAreaProductRelationService.AddPreferredAreaProductRelation(l.ctx, &cmsclient.AddPreferredAreaProductRelationReq{
		PreferredAreaId: req.PrefrenceAreaProductRelationList,
		ProductId:       productId,
	})

}
