package logic

import (
	"context"
	"zero-admin/rpc/model/pmsmodel"
	"zero-admin/rpc/pms/internal/svc"
	"zero-admin/rpc/pms/pms"
	"zero-admin/rpc/pms/pmsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProductDetailByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProductDetailByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductDetailByIdLogic {
	return &ProductDetailByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// ProductDetailById 获取商品详情
//1.获取商品信息
//2.获取品牌信息
//3.获取商品属性信息
//4.获取商品属性值信息
//5.获取商品SKU库存信息
//6.商品阶梯价格设置
//7.商品满减价格设置
func (l *ProductDetailByIdLogic) ProductDetailById(in *pmsclient.ProductDetailByIdReq) (*pmsclient.ProductDetailByIdResp, error) {
	//1.获取商品信息
	productListData, product := buildProductListData(l, in.Id)

	productAttributeListData, attributeIds := buildProductAttributeListData(l, product)
	return &pmsclient.ProductDetailByIdResp{
		Product:                   productListData,
		Brand:                     buildBrandListData(l, product),
		ProductAttributeList:      productAttributeListData,
		ProductAttributeValueList: buildProductAttributeValueListData(l, product, attributeIds),
		SkuStockList:              buildSkuStockListData(l, product),
		ProductLadderList:         buildProductLadderListData(l, product),
		ProductFullReductionList:  buildProductFullReductionListData(l, product),
	}, nil
}

//1.获取商品信息
func buildProductListData(l *ProductDetailByIdLogic, productId int64) (*pmsclient.ProductListData, *pmsmodel.PmsProduct) {
	pmsProduct, _ := l.svcCtx.PmsProductModel.FindOne(l.ctx, productId)

	return &pmsclient.ProductListData{
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
		PromotionStartTime:         pmsProduct.PromotionStartTime.Format("2006-01-02 15:04:05"),
		PromotionEndTime:           pmsProduct.PromotionEndTime.Format("2006-01-02 15:04:05"),
		PromotionPerLimit:          pmsProduct.PromotionPerLimit,
		PromotionType:              pmsProduct.PromotionType,
		BrandName:                  pmsProduct.BrandName,
		ProductCategoryName:        pmsProduct.ProductCategoryName,
		ProductCategoryIdArray:     pmsProduct.ProductCategoryIdArray,
	}, pmsProduct
}

//2.获取品牌信息
func buildBrandListData(l *ProductDetailByIdLogic, pmsProduct *pmsmodel.PmsProduct) *pmsclient.BrandListData {
	item, _ := l.svcCtx.PmsBrandModel.FindOne(l.ctx, pmsProduct.BrandId)
	return &pmsclient.BrandListData{
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
func buildProductAttributeListData(l *ProductDetailByIdLogic, pmsProduct *pmsmodel.PmsProduct) ([]*pmsclient.ProductAttributeListData, []int64) {
	all, _ := l.svcCtx.PmsProductAttributeModel.FindAll(l.ctx, &pms.ProductAttributeListReq{
		Current:                    1,
		PageSize:                   100,
		Name:                       "",
		Type:                       2,
		ProductAttributeCategoryId: pmsProduct.ProductAttributeCategoryId,
	})
	var list []*pms.ProductAttributeListData
	var attributeIds []int64
	for _, item := range *all {
		attributeIds = append(attributeIds, item.Id)
		list = append(list, &pms.ProductAttributeListData{
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

	return list, attributeIds
}

//4.获取商品属性值信息
func buildProductAttributeValueListData(l *ProductDetailByIdLogic, pmsProduct *pmsmodel.PmsProduct, attributeIds []int64) []*pmsclient.ProductAttributeValueListData {
	if len(attributeIds) > 0 {
		all, _ := l.svcCtx.PmsProductAttributeValueModel.FindAll(l.ctx, pmsProduct.Id)
		var list []*pms.ProductAttributeValueListData
		for _, item := range *all {
			list = append(list, &pms.ProductAttributeValueListData{
				Id:                 item.Id,
				ProductId:          item.ProductId,
				ProductAttributeId: item.ProductAttributeId,
				Value:              item.Value.String,
			})
		}

		return list
	}

	return nil
}

//5.获取商品SKU库存信息
func buildSkuStockListData(l *ProductDetailByIdLogic, pmsProduct *pmsmodel.PmsProduct) []*pmsclient.SkuStockListData {
	all, _ := l.svcCtx.PmsSkuStockModel.FindAll(l.ctx, pmsProduct.Id)
	var list []*pmsclient.SkuStockListData
	for _, item := range *all {

		list = append(list, &pmsclient.SkuStockListData{
			Id:             item.Id,
			ProductId:      item.ProductId,
			SkuCode:        item.SkuCode,
			Price:          float32(item.Price),
			Stock:          item.Stock,
			LowStock:       item.LowStock,
			Pic:            item.Pic,
			Sale:           item.Sale,
			PromotionPrice: float32(item.PromotionPrice),
			LockStock:      item.LockStock,
			SpData:         item.SpData,
		})
	}

	return list
}

//6.商品阶梯价格设置
func buildProductLadderListData(l *ProductDetailByIdLogic, pmsProduct *pmsmodel.PmsProduct) []*pmsclient.ProductLadderListData {
	if pmsProduct.PromotionType == 3 {
		all, _ := l.svcCtx.PmsProductLadderModel.FindAll(l.ctx, pmsProduct.Id)
		var list []*pmsclient.ProductLadderListData
		for _, item := range *all {

			list = append(list, &pmsclient.ProductLadderListData{
				Id:        item.Id,
				ProductId: item.ProductId,
				Count:     item.Count,
				Discount:  int64(item.Discount),
				Price:     float32(item.Price),
			})
		}

		return list
	}

	return nil
}

//7.商品满减价格设置
func buildProductFullReductionListData(l *ProductDetailByIdLogic, pmsProduct *pmsmodel.PmsProduct) []*pmsclient.ProductFullReductionListData {
	if pmsProduct.PromotionType == 4 {
		all, _ := l.svcCtx.PmsProductFullReductionModel.FindAll(l.ctx, pmsProduct.Id)

		var list []*pmsclient.ProductFullReductionListData
		for _, item := range *all {

			list = append(list, &pmsclient.ProductFullReductionListData{
				Id:          item.Id,
				ProductId:   item.ProductId,
				FullPrice:   float32(item.FullPrice),
				ReducePrice: float32(item.ReducePrice),
			})
		}
		return list
	}
	return nil

}
