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
	productResp, _ := l.svcCtx.ProductService.QueryProductDetailById(l.ctx, &pmsclient.QueryProductDetailByIdReq{
		Id: req.ProductId,
	})

	// 8.查询专题关联
	subjectRelationList, _ := l.svcCtx.SubjectProductRelationService.QuerySubjectProductRelationList(l.ctx, &cmsclient.QuerySubjectProductRelationListReq{
		ProductId: req.ProductId,
	})

	// 9.查询优选关联
	areaRelationList, _ := l.svcCtx.PreferredAreaProductRelationService.QueryPreferredAreaProductRelationList(l.ctx, &cmsclient.QueryPreferredAreaProductRelationListReq{
		ProductId: req.ProductId,
	})

	// 10.商品可用优惠券(根据商品id和分类id查询)
	couponList, _ := l.svcCtx.CouponService.QueryCouponFindByProductIdAndProductCategoryId(l.ctx, &smsclient.CouponFindByProductIdAndProductCategoryIdReq{
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
func buildProductData(resp *pmsclient.QueryProductDetailByIdResp) types.ProductData {
	product := resp.Product

	var productCategoryIdArray []int64
	for _, s := range strings.Split(product.ProductCategoryIdArray, ",") {
		id, _ := strconv.ParseInt(s, 10, 64)
		productCategoryIdArray = append(productCategoryIdArray, id)
	}
	return types.ProductData{
		Id:                         product.Id,
		BrandId:                    product.BrandId,                    // 品牌id
		ProductCategoryId:          product.ProductCategoryId,          // 商品分类id
		FeightTemplateId:           product.FeightTemplateId,           // 商品运费模板id
		ProductAttributeCategoryId: product.ProductAttributeCategoryId, // 商品属性分类id
		Name:                       product.Name,                       // 商品名称
		Pic:                        product.Pic,                        // 商品图片
		ProductSn:                  product.ProductSn,                  // 货号
		DeleteStatus:               product.DeleteStatus,               // 删除状态：0->未删除；1->已删除
		PublishStatus:              product.PublishStatus,              // 上架状态：0->下架；1->上架
		NewStatus:                  product.NewStatus,                  // 新品状态:0->不是新品；1->新品
		RecommandStatus:            product.RecommandStatus,            // 推荐状态；0->不推荐；1->推荐
		VerifyStatus:               product.VerifyStatus,               // 审核状态：0->未审核；1->审核通过
		Sort:                       product.Sort,                       // 排序
		Sale:                       product.Sale,                       // 销量
		Price:                      product.Price,                      // 商品价格
		PromotionPrice:             product.PromotionPrice,             // 促销价格
		GiftGrowth:                 product.GiftGrowth,                 // 赠送的成长值
		GiftPoint:                  product.GiftPoint,                  // 赠送的积分
		UsePointLimit:              product.UsePointLimit,              // 限制使用的积分数
		SubTitle:                   product.SubTitle,                   // 副标题
		Description:                product.Description,                // 商品描述
		OriginalPrice:              product.OriginalPrice,              // 市场价
		Stock:                      product.Stock,                      // 库存
		LowStock:                   product.LowStock,                   // 库存预警值
		Unit:                       product.Unit,                       // 单位
		Weight:                     product.Weight,                     // 商品重量，默认为克
		PreviewStatus:              product.PreviewStatus,              // 是否为预告商品：0->不是；1->是
		ServiceIds:                 product.ServiceIds,                 // 以逗号分割的产品服务：1->无忧退货；2->快速退款；3->免费包邮
		Keywords:                   product.Keywords,                   // 搜索关键字
		Note:                       product.Note,                       // 备注
		AlbumPics:                  product.AlbumPics,                  // 画册图片，连产品图片限制为5张，以逗号分割
		DetailTitle:                product.DetailTitle,                // 详情标题
		DetailDesc:                 product.DetailDesc,                 // 详情描述
		DetailHtml:                 product.DetailHtml,                 // 产品详情网页内容
		DetailMobileHtml:           product.DetailMobileHtml,           // 移动端网页详情
		PromotionStartTime:         product.PromotionStartTime,         // 促销开始时间
		PromotionEndTime:           product.PromotionEndTime,           // 促销结束时间
		PromotionPerLimit:          product.PromotionPerLimit,          // 活动限购数量
		PromotionType:              product.PromotionType,              // 促销类型：0->没有促销使用原价;1->使用促销价；2->使用会员价；3->使用阶梯价格；4->使用满减价格；5->限时购
		BrandName:                  product.BrandName,                  // 品牌名称
		ProductCategoryName:        product.ProductCategoryName,        // 商品分类名称
		ProductCategoryIdArray:     productCategoryIdArray,             // 商品分类id字符串
	}
}

// 2.获取品牌信息
func buildBrandData(resp *pmsclient.QueryProductDetailByIdResp) types.BrandData {
	detail := resp.Brand
	return types.BrandData{
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

// 3.商品阶梯价格设置
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

// 4.商品满减价格设置
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

// 5.会员价格
func buildMemberPriceListData(resp *pmsclient.QueryProductDetailByIdResp) []types.MemberPriceList {
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

// 7.获取商品属性值信息
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
