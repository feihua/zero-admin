package productservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/pms/gen/model"
	"github.com/feihua/zero-admin/rpc/pms/gen/query"

	"github.com/feihua/zero-admin/rpc/pms/internal/svc"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryProductDetailByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryProductDetailByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryProductDetailByIdLogic {
	return &QueryProductDetailByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryProductDetailById 获取商品详情
// 1.获取商品信息
// 2.获取品牌信息
// 3.获取商品属性信息
// 4.获取商品属性值信息
// 5.获取商品SKU库存信息
// 6.商品阶梯价格设置
// 7.商品满减价格设置
// 8.获取商品的会员价格
func (l *QueryProductDetailByIdLogic) QueryProductDetailById(in *pmsclient.QueryProductDetailByIdReq) (*pmsclient.QueryProductDetailByIdResp, error) {
	//1.获取商品信息
	productListData, product := buildProductListData(l, in.Id)

	productAttributeListData, attributeIds := buildProductAttributeListData(l, product)
	return &pmsclient.QueryProductDetailByIdResp{
		Product:                   productListData,
		Brand:                     buildBrandListData(l, product),
		ProductAttributeList:      productAttributeListData,
		ProductAttributeValueList: buildProductAttributeValueListData(l, product, attributeIds),
		SkuStockList:              buildSkuStockListData(l, product),
		ProductLadderList:         buildProductLadderListData(l, product),
		ProductFullReductionList:  buildProductFullReductionListData(l, product),
		MemberPriceList:           buildProductMemberListData(l, product),
	}, nil
}

// 1.获取商品信息
func buildProductListData(l *QueryProductDetailByIdLogic, productId int64) (*pmsclient.ProductListData, *model.PmsProduct) {
	pmsProduct, _ := query.PmsProduct.WithContext(l.ctx).Where(query.PmsProduct.ID.Eq(productId)).First()

	return &pmsclient.ProductListData{
		Id:                         pmsProduct.ID,
		BrandId:                    pmsProduct.BrandID,
		ProductCategoryId:          pmsProduct.ProductCategoryID,
		FeightTemplateId:           pmsProduct.FeightTemplateID,
		ProductAttributeCategoryId: pmsProduct.ProductAttributeCategoryID,
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
		DetailHtml:                 pmsProduct.DetailHTML,
		DetailMobileHtml:           pmsProduct.DetailMobileHTML,
		PromotionStartTime:         pmsProduct.PromotionStartTime.Format("2006-01-02 15:04:05"),
		PromotionEndTime:           pmsProduct.PromotionEndTime.Format("2006-01-02 15:04:05"),
		PromotionPerLimit:          pmsProduct.PromotionPerLimit,
		PromotionType:              pmsProduct.PromotionType,
		BrandName:                  pmsProduct.BrandName,
		ProductCategoryName:        pmsProduct.ProductCategoryName,
		ProductCategoryIdArray:     pmsProduct.ProductCategoryIDArray,
	}, pmsProduct
}

// 2.获取品牌信息
func buildBrandListData(l *QueryProductDetailByIdLogic, pmsProduct *model.PmsProduct) *pmsclient.BrandData {
	item, _ := query.PmsBrand.WithContext(l.ctx).Where(query.PmsBrand.ID.Eq(pmsProduct.BrandID)).First()
	return &pmsclient.BrandData{
		Id:                  item.ID,
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

// 3.获取商品属性信息
func buildProductAttributeListData(l *QueryProductDetailByIdLogic, pmsProduct *model.PmsProduct) ([]*pmsclient.ProductAttributeDataList, []int64) {

	q := query.PmsProductAttribute
	result, _ := q.WithContext(l.ctx).Where(q.Type.Eq(2), q.ProductAttributeCategoryID.Eq(pmsProduct.ProductAttributeCategoryID)).Find()

	var list []*pmsclient.ProductAttributeDataList
	var attributeIds []int64
	for _, item := range result {
		attributeIds = append(attributeIds, item.ID)
		list = append(list, &pmsclient.ProductAttributeDataList{
			Id:                         item.ID,
			ProductAttributeCategoryId: item.ProductAttributeCategoryID,
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

// 4.获取商品属性值信息
func buildProductAttributeValueListData(l *QueryProductDetailByIdLogic, pmsProduct *model.PmsProduct, attributeIds []int64) []*pmsclient.ProductAttributeValueData {
	if len(attributeIds) > 0 {
		q := query.PmsProductAttributeValue
		result, _ := q.WithContext(l.ctx).Where(q.ProductID.Eq(pmsProduct.ID)).Find()

		var list []*pmsclient.ProductAttributeValueData
		for _, item := range result {
			list = append(list, &pmsclient.ProductAttributeValueData{
				Id:                 item.ID,
				ProductId:          item.ProductID,
				ProductAttributeId: item.ProductAttributeID,
				Value:              *item.Value,
			})
		}

		return list
	}

	return nil
}

// 5.获取商品SKU库存信息
func buildSkuStockListData(l *QueryProductDetailByIdLogic, pmsProduct *model.PmsProduct) []*pmsclient.SkuStockData {
	q := query.PmsSkuStock
	result, _ := q.WithContext(l.ctx).Where(q.ProductID.Eq(pmsProduct.ID)).Find()
	var list []*pmsclient.SkuStockData
	for _, item := range result {

		list = append(list, &pmsclient.SkuStockData{
			Id:             item.ID,
			ProductId:      item.ProductID,
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

// 6.商品阶梯价格设置
func buildProductLadderListData(l *QueryProductDetailByIdLogic, pmsProduct *model.PmsProduct) []*pmsclient.ProductLadderData {
	if pmsProduct.PromotionType == 3 {
		q := query.PmsProductLadder
		result, _ := q.WithContext(l.ctx).Where(q.ProductID.Eq(pmsProduct.ID)).Find()
		var list []*pmsclient.ProductLadderData
		for _, item := range result {

			list = append(list, &pmsclient.ProductLadderData{
				Id:        item.ID,
				ProductId: item.ProductID,
				Count:     item.Count,
				Discount:  item.Discount,
				Price:     item.Price,
			})
		}

		return list
	}

	return nil
}

// 7.商品满减价格设置
func buildProductFullReductionListData(l *QueryProductDetailByIdLogic, pmsProduct *model.PmsProduct) []*pmsclient.ProductFullReductionData {
	if pmsProduct.PromotionType == 4 {
		q := query.PmsProductFullReduction
		result, _ := q.WithContext(l.ctx).Where(q.ProductID.Eq(pmsProduct.ID)).Find()

		var list []*pmsclient.ProductFullReductionData
		for _, item := range result {

			list = append(list, &pmsclient.ProductFullReductionData{
				Id:          item.ID,
				ProductId:   item.ProductID,
				FullPrice:   item.FullPrice,
				ReducePrice: item.ReducePrice,
			})
		}
		return list
	}
	return nil

}

// 8.获取商品的会员价格
func buildProductMemberListData(l *QueryProductDetailByIdLogic, pmsProduct *model.PmsProduct) []*pmsclient.MemberPriceListData {
	if pmsProduct.PromotionType == 2 {
		q := query.PmsMemberPrice
		result, _ := q.WithContext(l.ctx).Where(q.ProductID.Eq(pmsProduct.ID)).Find()

		var list []*pmsclient.MemberPriceListData
		for _, item := range result {

			list = append(list, &pmsclient.MemberPriceListData{
				Id:              item.ID,
				ProductId:       item.ProductID,
				MemberLevelId:   item.MemberLevelID,
				MemberPrice:     item.MemberPrice,
				MemberLevelName: item.MemberLevelName,
			})
		}
		return list
	}
	return nil

}
