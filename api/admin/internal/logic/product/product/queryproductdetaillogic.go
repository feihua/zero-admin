package product

import (
	"context"
	"github.com/feihua/zero-admin/rpc/cms/cmsclient"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"
	"strconv"
	"strings"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryProductDetailLogic
/*
Author: LiuFeiHua
Date: 2024/5/15 16:46
*/
type QueryProductDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryProductDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryProductDetailLogic {
	return &QueryProductDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// QueryProductDetail 查询商品详情
// 1.获取商品信息
// 2.获取品牌信息
// 3.商品阶梯价格设置
// 4.商品满减价格设置
// 5.会员价格
// 6.获取商品SKU库存信息
// 7.获取商品属性值信息
// 8.专题和商品关系
// 9.优选专区和商品的关系
// 10.商品优惠券
// 注意: 步骤1到7是在商品模块(rpc),8,9是在内容模块(rpc),10是在营销模块(rpc)
func (l *QueryProductDetailLogic) QueryProductDetail(req *types.QueryProductDetailReq) (resp *types.QueryProductDetailResp, err error) {
	productResp, _ := l.svcCtx.ProductService.ProductDetailById(l.ctx, &pmsclient.ProductDetailByIdReq{
		Id: req.ProductId,
	})

	//8.查询专题关联
	subjectRelationList, _ := l.svcCtx.SubjectProductRelationService.SubjectProductRelationList(l.ctx, &cmsclient.SubjectProductRelationListReq{
		ProductId: req.ProductId,
	})

	//9.查询优选关联
	areaRelationList, _ := l.svcCtx.PreferredAreaProductRelationService.PreferredAreaProductRelationList(l.ctx, &cmsclient.PreferredAreaProductRelationListReq{
		ProductId: req.ProductId,
	})

	//10.商品可用优惠券(根据商品id和分类id查询)
	couponList, _ := l.svcCtx.CouponService.CouponFindByProductIdAndProductCategoryId(l.ctx, &smsclient.CouponFindByProductIdAndProductCategoryIdReq{
		ProductId:         req.ProductId,
		ProductCategoryId: productResp.Product.ProductCategoryId,
	})

	return &types.QueryProductDetailResp{
		Code:    "000000",
		Message: "操作成功",
		Data: types.ProductDetailData{
			ProductData:                      buildProductData(productResp),
			BrandData:                        buildBrandData(productResp),
			ProductLadderList:                buildProductLadderListData(productResp),
			ProductFullReductionList:         buildProductFullReductionListData(productResp),
			MemberPriceList:                  buildMemberPriceListData(productResp),
			SkuStockList:                     buildSkuStockListData(productResp),
			ProductAttributeValueList:        buildProductAttributeValueListData(productResp),
			SubjectProductRelationList:       subjectRelationList.SubjectIds,
			PrefrenceAreaProductRelationList: areaRelationList.PreferredAreaIds,
			CouponList:                       buildCouponListData(couponList),
		},
	}, nil
}

// 1.获取商品信息
func buildProductData(resp *pmsclient.ProductDetailByIdResp) types.ProductData {
	pmsProduct := resp.Product

	var productCategoryIdArray []int64
	for _, s := range strings.Split(pmsProduct.ProductCategoryIdArray, ",") {
		id, _ := strconv.ParseInt(s, 10, 64)
		productCategoryIdArray = append(productCategoryIdArray, id)
	}
	return types.ProductData{
		Id:                         pmsProduct.Id,
		BrandId:                    pmsProduct.BrandId,
		ProductCategoryId:          pmsProduct.ProductCategoryId,
		ProductCategoryIdArray:     productCategoryIdArray,
		FeightTemplateId:           pmsProduct.FeightTemplateId,
		ProductAttributeCategoryId: pmsProduct.ProductAttributeCategoryId,
		Name:                       pmsProduct.Name,
		Pic:                        pmsProduct.Pic,
		ProductSn:                  pmsProduct.ProductSn,
		DeleteStatus:               pmsProduct.DeleteStatus,
		PublishStatus:              pmsProduct.PublishStatus,
		NewStatus:                  pmsProduct.NewStatus,
		RecommandStatus:            pmsProduct.RecommandStatus,
		VerifyStatus:               pmsProduct.VerifyStatus,
		Sort:                       pmsProduct.Sort,
		Sale:                       pmsProduct.Sale,
		Price:                      pmsProduct.Price,
		PromotionPrice:             pmsProduct.PromotionPrice,
		GiftGrowth:                 pmsProduct.GiftGrowth,
		GiftPoint:                  pmsProduct.GiftPoint,
		UsePointLimit:              pmsProduct.UsePointLimit,
		SubTitle:                   pmsProduct.SubTitle,
		Description:                pmsProduct.Description,
		OriginalPrice:              pmsProduct.OriginalPrice,
		Stock:                      pmsProduct.Stock,
		LowStock:                   pmsProduct.LowStock,
		Unit:                       pmsProduct.Unit,
		Weight:                     pmsProduct.Weight,
		PreviewStatus:              pmsProduct.PreviewStatus,
		ServiceIds:                 pmsProduct.ServiceIds,
		Keywords:                   pmsProduct.Keywords,
		Note:                       pmsProduct.Note,
		AlbumPics:                  pmsProduct.AlbumPics,
		DetailTitle:                pmsProduct.DetailTitle,
		DetailDesc:                 pmsProduct.DetailDesc,
		DetailHTML:                 pmsProduct.DetailHtml,
		DetailMobileHTML:           pmsProduct.DetailMobileHtml,
		PromotionStartTime:         pmsProduct.PromotionStartTime,
		PromotionEndTime:           pmsProduct.PromotionEndTime,
		PromotionPerLimit:          pmsProduct.PromotionPerLimit,
		PromotionType:              pmsProduct.PromotionType,
		BrandName:                  pmsProduct.BrandName,
		ProductCategoryName:        pmsProduct.ProductCategoryName,
	}
}

// 2.获取品牌信息
func buildBrandData(resp *pmsclient.ProductDetailByIdResp) types.BrandData {
	item := resp.Brand
	return types.BrandData{
		Id:                  item.Id,
		Name:                item.Name,
		FirstLetter:         item.FirstLetter,
		Sort:                item.Sort,
		FactoryStatus:       item.FactoryStatus,
		ShowStatus:          item.ShowStatus,
		ProductCount:        item.ProductCount,
		ProductCommentCount: item.ProductCommentCount,
		Logo:                item.Logo,
		BigPic:              item.BigPic,
		BrandStory:          item.BrandStory,
	}
}

// 3.商品阶梯价格设置
func buildProductLadderListData(resp *pmsclient.ProductDetailByIdResp) []types.ProductLadderList {
	list := make([]types.ProductLadderList, 0)
	for _, item := range resp.ProductLadderList {

		list = append(list, types.ProductLadderList{
			Id:        item.Id,
			ProductId: item.ProductId,
			Count:     item.Count,
			Discount:  item.Discount,
			Price:     item.Price,
		})
	}

	return list

}

// 4.商品满减价格设置
func buildProductFullReductionListData(resp *pmsclient.ProductDetailByIdResp) []types.ProductFullReductionList {
	list := make([]types.ProductFullReductionList, 0)
	for _, item := range resp.ProductFullReductionList {

		list = append(list, types.ProductFullReductionList{
			Id:          item.Id,
			ProductId:   item.ProductId,
			FullPrice:   item.FullPrice,
			ReducePrice: item.ReducePrice,
		})
	}
	return list

}

// 5.会员价格
func buildMemberPriceListData(resp *pmsclient.ProductDetailByIdResp) []types.MemberPriceList {
	list := make([]types.MemberPriceList, 0)
	for _, item := range resp.MemberPriceList {

		list = append(list, types.MemberPriceList{
			Id:              item.Id,
			ProductId:       item.ProductId,
			MemberLevelId:   item.MemberLevelId,
			MemberPrice:     item.MemberPrice,
			MemberLevelName: item.MemberLevelName,
		})
	}
	return list

}

// 6.获取商品SKU库存信息
func buildSkuStockListData(resp *pmsclient.ProductDetailByIdResp) []types.SkuStockList {
	list := make([]types.SkuStockList, 0)
	for _, item := range resp.SkuStockList {

		list = append(list, types.SkuStockList{
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

// 7.获取商品属性值信息
func buildProductAttributeValueListData(resp *pmsclient.ProductDetailByIdResp) []types.ProductAttributeValueList {
	list := make([]types.ProductAttributeValueList, 0)
	for _, item := range resp.ProductAttributeValueList {
		list = append(list, types.ProductAttributeValueList{
			Id:                 item.Id,
			ProductId:          item.ProductId,
			ProductAttributeId: item.ProductAttributeId,
			Value:              item.Value,
		})
	}

	return list
}

// 10.商品优惠券
func buildCouponListData(resp *smsclient.CouponFindByProductIdAndProductCategoryIdResp) []types.CouponList {
	list := make([]types.CouponList, 0)
	for _, item := range resp.List {

		list = append(list, types.CouponList{
			Id:           item.Id,
			Type:         item.Type,
			Name:         item.Name,
			Platform:     item.Platform,
			Count:        item.Count,
			Amount:       item.Amount,
			PerLimit:     item.PerLimit,
			MinPoint:     item.MinPoint,
			StartTime:    item.StartTime,
			EndTime:      item.EndTime,
			UseType:      item.UseType,
			PublishCount: item.PublishCount,
			UseCount:     item.UseCount,
			ReceiveCount: item.ReceiveCount,
			EnableTime:   item.EnableTime,
		})
	}
	return list

}
