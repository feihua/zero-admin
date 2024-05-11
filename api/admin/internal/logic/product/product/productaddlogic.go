package logic

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
	result, err := l.svcCtx.ProductService.ProductAdd(l.ctx, &pmsclient.ProductAddReq{
		BrandId:                    req.BrandId,
		ProductCategoryId:          req.ProductCategoryId,
		ProductCategoryIdArray:     req.ProductCategoryIdArray,
		FeightTemplateId:           req.FeightTemplateId,
		ProductAttributeCategoryId: req.ProductAttributeCategoryId,
		Name:                       req.Name,
		Pic:                        req.Pic,
		ProductSn:                  req.ProductSn,
		DeleteStatus:               req.DeleteStatus,
		PublishStatus:              req.PublishStatus,
		NewStatus:                  req.NewStatus,
		RecommandStatus:            req.RecommandStatus,
		VerifyStatus:               req.VerifyStatus,
		Sort:                       req.Sort,
		Sale:                       req.Sale,
		Price:                      req.Price,
		PromotionPrice:             req.PromotionPrice,
		GiftGrowth:                 req.GiftGrowth,
		GiftPoint:                  req.GiftPoint,
		UsePointLimit:              req.UsePointLimit,
		SubTitle:                   req.SubTitle,
		Description:                req.Description,
		OriginalPrice:              req.OriginalPrice,
		Stock:                      req.Stock,
		LowStock:                   req.LowStock,
		Unit:                       req.Unit,
		Weight:                     req.Weight,
		PreviewStatus:              req.PreviewStatus,
		ServiceIds:                 req.ServiceIds,
		Keywords:                   req.Keywords,
		Note:                       req.Note,
		AlbumPics:                  req.AlbumPics,
		DetailTitle:                req.DetailTitle,
		DetailDesc:                 req.DetailDesc,
		DetailHtml:                 req.DetailHtml,
		DetailMobileHtml:           req.DetailMobileHtml,
		PromotionStartTime:         req.PromotionStartTime,
		PromotionEndTime:           req.PromotionEndTime,
		PromotionPerLimit:          req.PromotionPerLimit,
		PromotionType:              req.PromotionType,
		BrandName:                  req.BrandName,
		ProductCategoryName:        req.ProductCategoryName,
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
	addPrefrenceAreaProductRelation(req, l, productId)

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
			MemberLevelId:   item.MemberLevelID,
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
			ProductAttributeId: item.ProductAttributeID,
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
func addSubjectProductRelation(req types.AddProductReq, l *ProductAddLogic, productId int32) {
	_, _ = l.svcCtx.SubjectProductRelationService.SubjectProductRelationAdd(l.ctx, &cmsclient.SubjectProductRelationAddReq{
		SubjectId: req.SubjectProductRelationList,
		ProductId: int64(productId),
	})
}

// 添加优选商品关联
func addPrefrenceAreaProductRelation(req types.AddProductReq, l *ProductAddLogic, productId int32) {
	_, _ = l.svcCtx.PreferredAreaProductRelationService.PreferredAreaProductRelationAdd(l.ctx, &cmsclient.PreferredAreaProductRelationAddReq{
		PreferredAreaId: req.PrefrenceAreaProductRelationList,
		ProductId:       int64(productId),
	})

}
