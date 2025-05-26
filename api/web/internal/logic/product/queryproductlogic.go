package product

import (
	"context"
	"github.com/feihua/zero-admin/pkg/errorx"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc/status"

	"github.com/feihua/zero-admin/api/web/internal/svc"
	"github.com/feihua/zero-admin/api/web/internal/types"

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
// 1.获取商品信息
// 2.获取品牌信息
// 3.获取商品属性信息
// 4.获取商品属性值信息
// 5.获取商品SKU库存信息
// 6.商品阶梯价格设置
// 7.商品满减价格设置
// 8.商品可用优惠券
// 注意: 步骤1到7是在商品模块(rpc),8是在营销模块(rpc)
func (l *QueryProductLogic) QueryProduct(req *types.QueryProductReq) (resp *types.QueryProductResp, err error) {
	productResp, err := l.svcCtx.ProductService.QueryProductDetailById(l.ctx, &pmsclient.QueryProductDetailByIdReq{
		Id: req.ProductId,
	})

	if err != nil {
		logc.Errorf(l.ctx, "获取商品详情失败,参数: %+v,异常：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	// 8.商品可用优惠券(根据商品id和分类id查询)
	couponList, err := l.svcCtx.CouponService.QueryCouponFindByProductIdAndProductCategoryId(l.ctx, &smsclient.CouponFindByProductIdAndProductCategoryIdReq{
		ProductId:         req.ProductId,
		ProductCategoryId: productResp.Product.ProductCategoryId,
	})
	if err != nil {
		logc.Errorf(l.ctx, "获取商品可用优惠券失败,参数: %+v,异常：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}
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
			CouponList:                buildCouponListData(couponList),
		},
	}, nil
}

// 1.获取商品信息
func buildProductListData(resp *pmsclient.QueryProductDetailByIdResp) types.Product {
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

// 2.获取品牌信息
func buildBrandListData(resp *pmsclient.QueryProductDetailByIdResp) types.Brand {
	detail := resp.Brand
	return types.Brand{
		Id:                  detail.Id,                  //
		Name:                detail.Name,                // 品牌名称
		Logo:                detail.Logo,                // 品牌logo
		BigPic:              detail.BigPic,              // 专区大图
		Description:         detail.Description,         // 描述
		FirstLetter:         detail.FirstLetter,         // 首字母
		Sort:                detail.Sort,                // 排序
		RecommendStatus:     detail.RecommendStatus,     // 推荐状态
		ProductCount:        detail.ProductCount,        // 产品数量
		ProductCommentCount: detail.ProductCommentCount, // 产品评论数量
		IsEnabled:           detail.IsEnabled,           // 是否启用
		CreateBy:            detail.CreateBy,            // 创建人ID
		CreateTime:          detail.CreateTime,          // 创建时间
		UpdateBy:            detail.UpdateBy,            // 更新人ID
		UpdateTime:          detail.UpdateTime,          // 更新时间
	}
}

// 3.获取商品属性信息
func buildProductAttributeListData(resp *pmsclient.QueryProductDetailByIdResp) []types.ProductAttributeList {
	list := make([]types.ProductAttributeList, 0)
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

// 4.获取商品属性值信息
func buildProductAttributeValueListData(resp *pmsclient.QueryProductDetailByIdResp) []types.ProductAttributeValueList {
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

// 5.获取商品SKU库存信息
func buildSkuStockListData(resp *pmsclient.QueryProductDetailByIdResp) []types.SkuStockList {
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

// 6.商品阶梯价格设置
func buildProductLadderListData(resp *pmsclient.QueryProductDetailByIdResp) []types.ProductLadderList {
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

// 7.商品满减价格设置
func buildProductFullReductionListData(resp *pmsclient.QueryProductDetailByIdResp) []types.ProductFullReductionList {
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

// 8.商品优惠券
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
