package product

import (
	"context"
	"zero-admin/rpc/pms/pmsclient"

	"zero-admin/front-api/internal/svc"
	"zero-admin/front-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryProductLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryProductLogic {
	return &QueryProductLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// QueryProduct 获取商品详情
//1.获取商品信息
//2.获取品牌信息
//3.获取商品属性信息
//4.获取商品属性值信息
//5.获取商品SKU库存信息
//6.商品阶梯价格设置
//7.商品满减价格设置
//8.商品可用优惠券
//注意: 步骤1到7是在商品模块(rpc),8是在营销模块(rpc)
func (l *QueryProductLogic) QueryProduct(req *types.QueryProductReq) (resp *types.QueryProductResp, err error) {

	productResp, _ := l.svcCtx.Pms.ProductDetailById(l.ctx, &pmsclient.ProductDetailByIdReq{
		Id: req.Id,
	})

	//8.商品可用优惠券
	//l.svcCtx.Sms.CouponList()
	return &types.QueryProductResp{
		Code:    0,
		Message: "操作成功",
		Data: types.ProductData{
			Product:                   buildProductListData(productResp),
			Brand:                     buildBrandListData(productResp),
			ProductAttributeList:      buildProductAttributeListData(productResp),
			ProductAttributeValueList: buildProductAttributeValueListData(productResp),
			SkuStockList:              buildSkuStockListData(productResp),
			ProductLadderList:         buildProductLadderListData(productResp),
			ProductFullReductionList:  buildProductFullReductionListData(productResp),
			CouponList:                nil,
		},
	}, nil
}

//1.获取商品信息
func buildProductListData(resp *pmsclient.ProductDetailByIdResp) types.Product {
	pmsProduct := resp.Product

	return types.Product{
		Id:                         pmsProduct.Id,
		BrandId:                    pmsProduct.BrandId,
		ProductCategoryId:          pmsProduct.ProductCategoryId,
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
		DetailHtml:                 pmsProduct.DetailHtml,
		DetailMobileHtml:           pmsProduct.DetailMobileHtml,
		PromotionStartTime:         pmsProduct.PromotionStartTime,
		PromotionEndTime:           pmsProduct.PromotionEndTime,
		PromotionPerLimit:          pmsProduct.PromotionPerLimit,
		PromotionType:              pmsProduct.PromotionType,
		BrandName:                  pmsProduct.BrandName,
		ProductCategoryName:        pmsProduct.ProductCategoryName,
	}
}

//2.获取品牌信息
func buildBrandListData(resp *pmsclient.ProductDetailByIdResp) types.Brand {
	item := resp.Brand
	return types.Brand{
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

//3.获取商品属性信息
func buildProductAttributeListData(resp *pmsclient.ProductDetailByIdResp) []types.ProductAttributeList {
	var list []types.ProductAttributeList
	for _, item := range resp.ProductAttributeList {
		list = append(list, types.ProductAttributeList{
			Id:                         item.Id,
			ProductAttributeCategoryId: item.ProductAttributeCategoryId,
			Name:                       item.Name,
			SelectType:                 item.SelectType,
			InputType:                  item.InputType,
			InputList:                  item.InputList,
			Sort:                       item.Sort,
			FilterType:                 item.FilterType,
			SearchType:                 item.SearchType,
			RelatedStatus:              item.RelatedStatus,
			HandAddStatus:              item.HandAddStatus,
			Type:                       item.Type,
		})
	}

	return list
}

//4.获取商品属性值信息
func buildProductAttributeValueListData(resp *pmsclient.ProductDetailByIdResp) []types.ProductAttributeValueList {
	var list []types.ProductAttributeValueList
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

//5.获取商品SKU库存信息
func buildSkuStockListData(resp *pmsclient.ProductDetailByIdResp) []types.SkuStockList {
	var list []types.SkuStockList
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

//6.商品阶梯价格设置
func buildProductLadderListData(resp *pmsclient.ProductDetailByIdResp) []types.ProductLadderList {
	var list []types.ProductLadderList
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

//7.商品满减价格设置
func buildProductFullReductionListData(resp *pmsclient.ProductDetailByIdResp) []types.ProductFullReductionList {

	var list []types.ProductFullReductionList
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
