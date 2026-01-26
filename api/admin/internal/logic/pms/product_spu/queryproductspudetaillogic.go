package product_spu

import (
	"context"

	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"
	"github.com/feihua/zero-admin/rpc/cms/cmsclient"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
)

// QueryProductSpuDetailLogic 查询商品SPU详情
/*
Author: LiuFeiHua
Date: 2025/06/16 14:37:37
*/
type QueryProductSpuDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryProductSpuDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryProductSpuDetailLogic {
	return &QueryProductSpuDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// QueryProductSpuDetail 查询商品SPU详情
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
func (l *QueryProductSpuDetailLogic) QueryProductSpuDetail(req *types.QueryProductSpuDetailReq) (resp *types.QueryProductSpuDetailResp, err error) {

	detail, err := l.svcCtx.ProductSpuService.QueryProductSpuDetail(l.ctx, &pmsclient.QueryProductSpuDetailReq{
		Id: req.Id,
	})

	if err != nil {
		logc.Errorf(l.ctx, "查询商品SPU详情失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	// 8.查询专题关联
	subjectIds := make([]int64, 0)
	res, _ := l.svcCtx.SubjectProductRelationService.QuerySubjectProductRelationList(l.ctx, &cmsclient.QuerySubjectProductRelationListReq{
		ProductId: req.Id,
	})

	// 9.查询优选关联
	areaRelationIds := make([]int64, 0)
	res1, _ := l.svcCtx.PreferredAreaProductRelationService.QueryPreferredAreaProductRelationList(l.ctx, &cmsclient.QueryPreferredAreaProductRelationListReq{
		ProductId: req.Id,
	})

	// 10.商品可用优惠券(根据商品id和分类id查询)
	couponList, _ := l.svcCtx.CouponService.QueryCouponByScopeId(l.ctx, &smsclient.QueryCouponByScopeIdReq{
		ScopeIds: []int64{req.Id, detail.Data.CategoryId},
	})

	return &types.QueryProductSpuDetailResp{
		Code:    "000000",
		Message: "操作成功",
		Data: types.ProductSpuDetailData{
			ProductData:        buildProductData(detail),
			BrandData:          buildBrandData(detail),
			LadderList:         buildProductLadderListData(detail),
			FullList:           buildProductFullReductionListData(detail),
			MemberPriceList:    buildMemberPriceListData(detail),
			SkuList:            buildSkuStockListData(detail),
			AttributeValueList: buildProductAttributeValueListData(detail),
			SubjectIds:         append(subjectIds, res.SubjectIds...),
			PrefrenceAreaIds:   append(areaRelationIds, res1.PreferredAreaIds...),
			CouponList:         buildCouponListData(couponList.List),
		},
	}, nil
}

// 1.获取商品信息
func buildProductData(resp *pmsclient.QueryProductSpuDetailResp) types.QueryProductSpuListData {
	product := resp.Data

	return types.QueryProductSpuListData{
		Id:                  product.Id,                  // 商品SpuId
		Name:                product.Name,                // 商品名称
		ProductSn:           product.ProductSn,           // 商品货号
		CategoryId:          product.CategoryId,          // 商品分类ID
		CategoryIds:         product.CategoryIds,         // 商品分类ID集合
		CategoryName:        product.CategoryName,        // 商品分类名称
		BrandId:             product.BrandId,             // 品牌ID
		BrandName:           product.BrandName,           // 品牌名称
		Unit:                product.Unit,                // 单位
		Weight:              product.Weight,              // 重量(kg)
		Keywords:            product.Keywords,            // 关键词
		AlbumPics:           product.AlbumPics,           // 画册图片，最多8张，以逗号分割
		MainPic:             product.MainPic,             // 主图
		PriceRange:          product.PriceRange,          // 价格区间
		PublishStatus:       product.PublishStatus,       // 上架状态：0-下架，1-上架
		NewStatus:           product.NewStatus,           // 新品状态:0->不是新品；1->新品
		RecommendStatus:     product.RecommendStatus,     // 推荐状态；0->不推荐；1->推荐
		VerifyStatus:        product.VerifyStatus,        // 审核状态：0->未审核；1->审核通过
		PreviewStatus:       product.PreviewStatus,       // 是否为预告商品：0->不是；1->是
		Sort:                product.Sort,                // 排序
		NewStatusSort:       product.NewStatusSort,       // 新品排序
		RecommendStatusSort: product.RecommendStatusSort, // 推荐排序
		Sales:               product.Sales,               // 销量
		Stock:               product.Stock,               // 库存
		LowStock:            product.LowStock,            // 预警库存
		PromotionType:       product.PromotionType,       // 促销类型：0->没有促销使用原价;1->使用促销价；2->使用会员价；3->使用阶梯价格；4->使用满减价格；5->秒杀
		SubTitle:            product.SubTitle,            // 详情标题
		DetailHtml:          product.DetailHtml,          // 产品详情网页内容
		DetailMobileHtml:    product.DetailMobileHtml,    // 移动端网页详情
		CreateBy:            product.CreateBy,            // 创建人ID
		CreateTime:          product.CreateTime,          // 创建时间
		UpdateBy:            product.UpdateBy,            // 更新人ID
		UpdateTime:          product.UpdateTime,          // 更新时间
	}
}

// 2.获取品牌信息
func buildBrandData(resp *pmsclient.QueryProductSpuDetailResp) types.QueryProductBrandDetailData {
	detail := resp.Brand
	return types.QueryProductBrandDetailData{
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
func buildProductLadderListData(resp *pmsclient.QueryProductSpuDetailResp) []types.ProductLadderResp {
	list := make([]types.ProductLadderResp, 0)
	for _, detail := range resp.ProductLadderList {

		list = append(list, types.ProductLadderResp{
			Id:        detail.Id,        //
			ProductId: detail.ProductId, // 商品id
			Count:     detail.Count,     // 满足的商品数量
			Discount:  detail.Discount,  // 折扣
			Price:     detail.Price,     // 折后价格
		})
	}

	return list

}

// 4.商品满减价格设置
func buildProductFullReductionListData(resp *pmsclient.QueryProductSpuDetailResp) []types.ProductFullReductionResp {
	list := make([]types.ProductFullReductionResp, 0)
	for _, detail := range resp.ProductFullReductionList {

		list = append(list, types.ProductFullReductionResp{
			Id:          detail.Id,          //
			ProductId:   detail.ProductId,   // 商品id
			FullPrice:   detail.FullPrice,   // 商品满多少
			ReducePrice: detail.ReducePrice, // 商品减多少
		})
	}
	return list

}

// 5.会员价格
func buildMemberPriceListData(resp *pmsclient.QueryProductSpuDetailResp) []types.MemberPriceResp {
	list := make([]types.MemberPriceResp, 0)
	for _, detail := range resp.MemberPriceList {

		list = append(list, types.MemberPriceResp{
			Id:              detail.Id,              //
			ProductId:       detail.ProductId,       // 商品id
			MemberLevelId:   detail.MemberLevelId,   // 会员等级id
			MemberPrice:     detail.MemberPrice,     // 会员价格
			MemberLevelName: detail.MemberLevelName, // 会员等级名称
		})
	}
	return list

}

// 6.获取商品SKU库存信息
func buildSkuStockListData(resp *pmsclient.QueryProductSpuDetailResp) []types.QueryProductSkuDetailData {
	list := make([]types.QueryProductSkuDetailData, 0)
	for _, detail := range resp.SkuStockList {

		list = append(list, types.QueryProductSkuDetailData{
			Id:                 detail.Id,                 // 商品SpuId
			SpuId:              detail.SpuId,              // 商品SpuId
			Name:               detail.Name,               // SKU名称
			SkuCode:            detail.SkuCode,            // SKU编码
			MainPic:            detail.MainPic,            // 主图
			AlbumPics:          detail.AlbumPics,          // 图片集
			Price:              detail.Price,              // 价格
			PromotionPrice:     detail.PromotionPrice,     // 单品促销价格
			PromotionStartTime: detail.PromotionStartTime, // 促销开始时间
			PromotionEndTime:   detail.PromotionEndTime,   // 促销结束时间
			Stock:              detail.Stock,              // 库存
			LowStock:           detail.LowStock,           // 预警库存
			SpecData:           detail.SpecData,           // 规格数据
			Weight:             detail.Weight,             // 重量(kg)
			PublishStatus:      detail.PublishStatus,      // 上架状态：0-下架，1-上架
			VerifyStatus:       detail.VerifyStatus,       // 审核状态：0-未审核，1-审核通过，2-审核不通过
			Sort:               detail.Sort,               // 排序
			Sales:              detail.Sales,              // 销量
			CreateBy:           detail.CreateBy,           // 创建人ID
			CreateTime:         detail.CreateTime,         // 创建时间
			UpdateBy:           detail.UpdateBy,           // 更新人ID
			UpdateTime:         detail.UpdateTime,         // 更新时间
		})
	}

	return list
}

// 7.获取商品属性值信息
func buildProductAttributeValueListData(resp *pmsclient.QueryProductSpuDetailResp) []types.QueryProductAttributeValueDetailData {
	list := make([]types.QueryProductAttributeValueDetailData, 0)
	for _, detail := range resp.ProductAttributeValueList {
		list = append(list, types.QueryProductAttributeValueDetailData{
			Id:          detail.Id,          // 主键id
			SpuId:       detail.SpuId,       // 商品SPU ID
			AttributeId: detail.AttributeId, // 属性ID
			Value:       detail.Value,       // 属性值
			CreateBy:    detail.CreateBy,    // 创建人ID
			CreateTime:  detail.CreateTime,  // 创建时间
			UpdateBy:    detail.UpdateBy,    // 更新人ID
			UpdateTime:  detail.UpdateTime,  // 更新时间
		})
	}

	return list
}

// 10.商品优惠券
func buildCouponListData(resp []*smsclient.CouponListData) []types.QueryCouponDetailData {
	list := make([]types.QueryCouponDetailData, 0)
	for _, detail := range resp {

		list = append(list, types.QueryCouponDetailData{
			Id:            detail.Id,            // 优惠券ID
			TypeId:        detail.TypeId,        // 优惠券类型ID
			Name:          detail.Name,          // 优惠券名称
			Code:          detail.Code,          // 优惠券码
			Amount:        detail.Amount,        // 优惠金额/折扣率
			MinAmount:     detail.MinAmount,     // 最低使用金额
			StartTime:     detail.StartTime,     // 生效时间
			EndTime:       detail.EndTime,       // 失效时间
			TotalCount:    detail.TotalCount,    // 发放总量
			ReceivedCount: detail.ReceivedCount, // 已领取数量
			UsedCount:     detail.UsedCount,     // 已使用数量
			PerLimit:      detail.PerLimit,      // 每人限领数量
			Status:        detail.Status,        // 状态：0-未开始，1-进行中，2-已结束，3-已取消
			IsEnabled:     detail.IsEnabled,     // 是否启用
			Description:   detail.Description,   // 使用说明

		})
	}
	return list

}
