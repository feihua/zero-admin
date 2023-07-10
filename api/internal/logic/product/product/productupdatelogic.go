package logic

import (
	"context"
	"encoding/json"
	"zero-admin/api/internal/common/errorx"
	"zero-admin/rpc/cms/cmsclient"
	"zero-admin/rpc/pms/pmsclient"

	"zero-admin/api/internal/svc"
	"zero-admin/api/internal/types"

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
//逻辑步骤：
//1.更新商品基本信息
//2.会员价格
//3.阶梯价格
//4.满减价格
//5.更新sku库存信息
//6.更新商品参数,添加自定义商品规格
//7.关联专题
//8.关联优选
//注意: 步骤1到6是在商品(rpc)中操作,步骤7到8是在cms模块中操作
func (l *ProductUpdateLogic) ProductUpdate(req types.UpdateProductReq) (*types.UpdateProductResp, error) {
	productId := req.Id
	_, err := l.svcCtx.ProductService.ProductUpdate(l.ctx, &pmsclient.ProductUpdateReq{
		Id:                         productId,
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
		MemberPriceList:            buildUpdateMemberPriceList(req),
		ProductAttributeValueList:  buildUpdateProductAttributeValueList(req),
		ProductFullReductionList:   buildUpdateProductFullReductionList(req),
		ProductLadderList:          buildUpdateProductLadderList(req),
		SkuStockList:               buildUpdateSkuStockList(req),
	})

	if err != nil {
		reqStr, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("更新商品信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, errorx.NewDefaultError("更新商品信息失败")
	}

	//7.关联专题
	updateSubjectProductRelation(req, l, productId)

	//8.关联优选
	updatePrefrenceAreaProductRelation(req, l, productId)

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
			MemberLevelId:   item.MemberLevelID,
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
			ProductAttributeId: item.ProductAttributeID,
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
	_, _ = l.svcCtx.SubjectProductRelationService.SubjectProductRelationDelete(l.ctx, &cmsclient.SubjectProductRelationDeleteReq{Id: productId})

	//2.重新添加专题的关联
	for _, subjectId := range req.SubjectProductRelationList {
		_, _ = l.svcCtx.SubjectProductRelationService.SubjectProductRelationAdd(l.ctx, &cmsclient.SubjectProductRelationAddReq{
			SubjectId: subjectId,
			ProductId: productId,
		})
	}
}

// 更新优选商品关联
func updatePrefrenceAreaProductRelation(req types.UpdateProductReq, l *ProductUpdateLogic, productId int64) {
	//1.先删除优选商品的关联
	_, _ = l.svcCtx.PrefrenceAreaProductRelationService.PrefrenceAreaProductRelationDelete(l.ctx, &cmsclient.PrefrenceAreaProductRelationDeleteReq{Id: productId})

	//2.重新添加优选商品的关联
	for _, prefrenceAreaId := range req.PrefrenceAreaProductRelationList {
		_, _ = l.svcCtx.PrefrenceAreaProductRelationService.PrefrenceAreaProductRelationAdd(l.ctx, &cmsclient.PrefrenceAreaProductRelationAddReq{
			PrefrenceAreaId: prefrenceAreaId,
			ProductId:       productId,
		})
	}
}
