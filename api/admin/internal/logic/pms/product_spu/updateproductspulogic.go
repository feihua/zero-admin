package product_spu

import (
	"context"

	"github.com/feihua/zero-admin/api/admin/internal/common"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/api/admin/internal/common/res"
	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"
	"github.com/feihua/zero-admin/rpc/cms/cmsclient"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc/status"

	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateProductSpuLogic 更新商品SPU
/*
Author: LiuFeiHua
Date: 2025/06/16 14:37:37
*/
type UpdateProductSpuLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateProductSpuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateProductSpuLogic {
	return &UpdateProductSpuLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// UpdateProductSpu 更新商品SPU
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
func (l *UpdateProductSpuLogic) UpdateProductSpu(req *types.UpdateProductSpuReq) (resp *types.BaseResp, err error) {
	productId := req.ProductData.Id
	err = l.updateProductSpuInfo(req)

	if err != nil {
		logc.Errorf(l.ctx, "更新商品信息失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	// 7.关联专题
	updateSubjectProductRelation(req, l, productId)

	// 8.关联优选
	updatePreferredAreaProductRelation(req, l, productId)

	return res.Success()
}

func (l *UpdateProductSpuLogic) updateProductSpuInfo(req *types.UpdateProductSpuReq) error {
	userId, err := common.GetUserId(l.ctx)
	if err != nil {
		return err
	}
	product := req.ProductData
	_, err = l.svcCtx.ProductSpuService.UpdateProductSpu(l.ctx, &pmsclient.ProductSpuReq{
		Id:                        product.Id,                                // 商品SpuId
		Name:                      product.Name,                              // 商品名称
		ProductSn:                 product.ProductSn,                         // 商品货号
		CategoryId:                product.CategoryId,                        // 商品分类ID
		CategoryIds:               product.CategoryIds,                       // 商品分类ID集合
		CategoryName:              product.CategoryName,                      // 商品分类名称
		BrandId:                   product.BrandId,                           // 品牌ID
		BrandName:                 product.BrandName,                         // 品牌名称
		Unit:                      product.Unit,                              // 单位
		Weight:                    product.Weight,                            // 重量(kg)
		Keywords:                  product.Keywords,                          // 关键词
		AlbumPics:                 product.AlbumPics,                         // 画册图片，最多8张，以逗号分割
		MainPic:                   product.MainPic,                           // 主图
		PublishStatus:             product.PublishStatus,                     // 上架状态：0-下架，1-上架
		NewStatus:                 product.NewStatus,                         // 新品状态:0->不是新品；1->新品
		RecommendStatus:           product.RecommendStatus,                   // 推荐状态；0->不推荐；1->推荐
		VerifyStatus:              product.VerifyStatus,                      // 审核状态：0->未审核；1->审核通过
		PreviewStatus:             product.PreviewStatus,                     // 是否为预告商品：0->不是；1->是
		Sort:                      product.Sort,                              // 排序
		NewStatusSort:             product.NewStatusSort,                     // 新品排序
		RecommendStatusSort:       product.RecommendStatusSort,               // 推荐排序
		Stock:                     product.Stock,                             // 库存
		LowStock:                  product.LowStock,                          // 预警库存
		PromotionType:             product.PromotionType,                     // 促销类型：0->没有促销使用原价;1->使用促销价；2->使用会员价；3->使用阶梯价格；4->使用满减价格；5->秒杀
		SubTitle:                  product.SubTitle,                          // 详情标题
		DetailHtml:                product.DetailHtml,                        // 产品详情网页内容
		DetailMobileHtml:          product.DetailMobileHtml,                  // 移动端网页详情
		CreateBy:                  userId,                                    // 更新人ID
		MemberPriceList:           buildUpdateMemberPriceList(req),           // 会员价格
		ProductAttributeValueList: buildUpdateProductAttributeValueList(req), // 商品参数及自定义规格属性
		ProductFullReductionList:  buildUpdateProductFullReductionList(req),  // 满减价格
		ProductLadderList:         buildUpdateProductLadderList(req),         // 阶梯价格
		SkuStockList:              buildUpdateSkuStockList(req),              // sku库存信息
	})
	return err
}

// 商品会员价格设置
func buildUpdateMemberPriceList(req *types.UpdateProductSpuReq) []*pmsclient.MemberPriceList {
	var memberPriceLists []*pmsclient.MemberPriceList
	for _, item := range req.MemberPriceList {
		memberPriceLists = append(memberPriceLists, &pmsclient.MemberPriceList{
			LevelId:   item.MemberLevelId,   // 会员等级id
			Price:     item.MemberPrice,     // 会员价格
			LevelName: item.MemberLevelName, // 会员等级名称
		})
	}
	return memberPriceLists
}

// 商品参数及自定义规格属性
func buildUpdateProductAttributeValueList(req *types.UpdateProductSpuReq) []*pmsclient.ProductAttributeValueList {
	var attributeValueLists []*pmsclient.ProductAttributeValueList
	for _, item := range req.AttributeValueList {
		attributeValueLists = append(attributeValueLists, &pmsclient.ProductAttributeValueList{
			ProductAttributeId: item.AttributeId, // 商品属性id
			AttributeValues:    item.Value,       // 参数值
		})
	}
	return attributeValueLists
}

// 商品满减价格设置
func buildUpdateProductFullReductionList(req *types.UpdateProductSpuReq) []*pmsclient.ProductFullReductionList {
	var fullReductionLists []*pmsclient.ProductFullReductionList
	for _, item := range req.FullList {
		fullReductionLists = append(fullReductionLists, &pmsclient.ProductFullReductionList{
			FullPrice:   item.FullPrice,   // 商品满多少
			ReducePrice: item.ReducePrice, // 商品减多少
		})
	}
	return fullReductionLists
}

// 商品阶梯价格设置
func buildUpdateProductLadderList(req *types.UpdateProductSpuReq) []*pmsclient.ProductLadderList {
	var ladderLists []*pmsclient.ProductLadderList
	for _, item := range req.LadderList {
		ladderLists = append(ladderLists, &pmsclient.ProductLadderList{
			Count:    item.Count,    // 满足的商品数量
			Discount: item.Discount, // 折扣
			Price:    item.Price,    // 折后价格
		})
	}
	return ladderLists
}

// 商品的sku库存信息
func buildUpdateSkuStockList(req *types.UpdateProductSpuReq) []*pmsclient.SkuStockList {
	var skuStockLists []*pmsclient.SkuStockList
	for _, item := range req.SkuList {
		skuStockLists = append(skuStockLists, &pmsclient.SkuStockList{
			Name:               item.Name,               // SKU名称
			MainPic:            item.MainPic,            // 主图
			AlbumPics:          item.AlbumPics,          // 图片集
			Price:              item.Price,              // 价格
			PromotionPrice:     item.PromotionPrice,     // 单品促销价格
			PromotionStartTime: item.PromotionStartTime, // 促销开始时间
			PromotionEndTime:   item.PromotionEndTime,   // 促销结束时间
			Stock:              item.Stock,              // 库存
			LowStock:           item.LowStock,           // 预警库存
			SpecData:           item.SpecData,           // 规格数据
			Weight:             item.Weight,             // 重量(kg)
			PublishStatus:      item.PublishStatus,      // 上架状态：0-下架，1-上架
			VerifyStatus:       item.VerifyStatus,       // 审核状态：0-未审核，1-审核通过，2-审核不通过
			Sort:               item.Sort,               // 排序
		})
	}
	return skuStockLists
}

// 更新专题关联
func updateSubjectProductRelation(req *types.UpdateProductSpuReq, l *UpdateProductSpuLogic, productId int64) {
	// 1.先删除专题的关联
	// _, _ = l.svcCtx.SubjectProductRelationService.SubjectProductRelationDelete(l.ctx, &cmsclient.SubjectProductRelationDeleteReq{Id: productId})

	// 2.重新添加专题的关联
	_, _ = l.svcCtx.SubjectProductRelationService.AddSubjectProductRelation(l.ctx, &cmsclient.AddSubjectProductRelationReq{
		SubjectId: req.SubjectIds, // 专题ID
		ProductId: productId,      // 商品ID
	})
}

// 更新优选商品关联
func updatePreferredAreaProductRelation(req *types.UpdateProductSpuReq, l *UpdateProductSpuLogic, productId int64) {
	// 1.先删除优选商品的关联(由rpc实现)
	// _, _ = l.svcCtx.PrefrenceAreaProductRelationService.PrefrenceAreaProductRelationDelete(l.ctx, &cmsclient.PrefrenceAreaProductRelationDeleteReq{Id: productId})

	// 2.重新添加优选商品的关联
	_, _ = l.svcCtx.PreferredAreaProductRelationService.AddPreferredAreaProductRelation(l.ctx, &cmsclient.AddPreferredAreaProductRelationReq{
		PreferredAreaId: req.PrefrenceAreaIds, // 优选专区ID
		ProductId:       productId,            // 商品ID
	})

}
