package home

import (
	"context"
	"github.com/feihua/zero-admin/api/front/internal/svc"
	"github.com/feihua/zero-admin/api/front/internal/types"
	"github.com/feihua/zero-admin/rpc/cms/cmsclient"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// IndexLogic 获取首页数据
/*
Author: LiuFeiHua
Date: 2025/6/19 15:43
*/
type IndexLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewIndexLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IndexLogic {
	return &IndexLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *IndexLogic) Index() (resp *types.HomeResp, err error) {
	return &types.HomeResp{
		Code:    0,
		Message: "操作成功",
		Data: types.Data{
			AdvertiseList:      queryAdvertiseList(l),
			BrandList:          queryBrandList(l),
			HomeFlashPromotion: queryHomeFlashPromotion(l),
			NewProductList:     queryNewProductList(l),
			HotProductList:     queryHotProductList(l),
			SubjectList:        querySubjectList(l), // 推荐专题
		},
	}, nil
}

// 推荐专题
func querySubjectList(l *IndexLogic) []types.SubjectList {
	var list []types.SubjectList
	res, err := l.svcCtx.SubjectService.QuerySubjectList(l.ctx, &cmsclient.QuerySubjectListReq{
		PageNum:         1,
		PageSize:        4,
		Title:           "", // 专题标题
		RecommendStatus: 1,  // 推荐状态：0->不推荐；1->推荐
		ShowStatus:      1,  // 显示状态：0->不显示；1->显示
	})

	// 没有推荐专题的时候返回空数据
	if err != nil {
		return list
	}

	for _, item := range res.List {
		list = append(list, types.SubjectList{
			Id:              item.Id,              // 专题id
			CategoryId:      item.CategoryId,      // 专题分类id
			Title:           item.Title,           // 专题标题
			Pic:             item.Pic,             // 专题主图
			ProductCount:    item.ProductCount,    // 关联产品数量
			RecommendStatus: item.RecommendStatus, // 推荐状态：0->不推荐；1->推荐
			CollectCount:    item.CollectCount,    // 收藏数
			ReadCount:       item.ReadCount,       // 阅读数
			CommentCount:    item.CommentCount,    // 评论数
			AlbumPics:       item.AlbumPics,       // 画册图片用逗号分割
			Description:     item.Description,     // 专题内容
			ShowStatus:      item.ShowStatus,      // 显示状态：0->不显示；1->显示
			Content:         item.Content,         // 专题内容
			ForwardCount:    item.ForwardCount,    // 转发数
			CategoryName:    item.CategoryName,    // 专题分类名称
			Sort:            1,
		})
	}
	return list
}

// 人气推荐
func queryHotProductList(l *IndexLogic) []types.ProductData {
	var resp, _ = l.svcCtx.ProductSpuService.QueryProductSpuList(l.ctx, &pmsclient.QueryProductSpuListReq{
		PageNum:         1,
		PageSize:        4,
		Name:            "",
		VerifyStatus:    2,
		CategoryId:      0,
		PublishStatus:   2,
		BrandId:         0,
		NewStatus:       2,
		RecommendStatus: 1,
	})

	var list []types.ProductData

	for _, detail := range resp.List {

		list = append(list, types.ProductData{
			Id:                  detail.Id,                  // 商品SpuId
			Name:                detail.Name,                // 商品名称
			CategoryId:          detail.CategoryId,          // 商品分类ID
			CategoryIds:         detail.CategoryIds,         // 商品分类ID集合
			CategoryName:        detail.CategoryName,        // 商品分类名称
			BrandId:             detail.BrandId,             // 品牌ID
			BrandName:           detail.BrandName,           // 品牌名称
			Unit:                detail.Unit,                // 单位
			Weight:              detail.Weight,              // 重量(kg)
			Keywords:            detail.Keywords,            // 关键词
			Brief:               detail.Brief,               // 简介
			Description:         detail.Description,         // 详细描述
			AlbumPics:           detail.AlbumPics,           // 画册图片，最多8张，以逗号分割
			MainPic:             detail.MainPic,             // 主图
			PriceRange:          detail.PriceRange,          // 价格区间
			PublishStatus:       detail.PublishStatus,       // 上架状态：0-下架，1-上架
			NewStatus:           detail.NewStatus,           // 新品状态:0->不是新品；1->新品
			RecommendStatus:     detail.RecommendStatus,     // 推荐状态；0->不推荐；1->推荐
			VerifyStatus:        detail.VerifyStatus,        // 审核状态：0->未审核；1->审核通过
			PreviewStatus:       detail.PreviewStatus,       // 是否为预告商品：0->不是；1->是
			Sort:                detail.Sort,                // 排序
			NewStatusSort:       detail.NewStatusSort,       // 新品排序
			RecommendStatusSort: detail.RecommendStatusSort, // 推荐排序
			Sales:               detail.Sales,               // 销量
			Stock:               detail.Stock,               // 库存
			LowStock:            detail.LowStock,            // 预警库存
			PromotionType:       detail.PromotionType,       // 促销类型：0->没有促销使用原价;1->使用促销价；2->使用会员价；3->使用阶梯价格；4->使用满减价格；5->秒杀
			DetailTitle:         detail.DetailTitle,         // 详情标题
			DetailDesc:          detail.DetailDesc,          // 详情描述
			DetailHtml:          detail.DetailHtml,          // 产品详情网页内容
			DetailMobileHtml:    detail.DetailMobileHtml,    // 移动端网页详情
		})
	}
	return list
}

// 新品推荐
func queryNewProductList(l *IndexLogic) []types.ProductData {
	var resp, _ = l.svcCtx.ProductSpuService.QueryProductSpuList(l.ctx, &pmsclient.QueryProductSpuListReq{
		PageNum:         1,
		PageSize:        4,
		Name:            "",
		VerifyStatus:    2,
		CategoryId:      0,
		PublishStatus:   2,
		BrandId:         0,
		NewStatus:       1,
		RecommendStatus: 2,
	})

	var list []types.ProductData

	for _, detail := range resp.List {

		list = append(list, types.ProductData{
			Id:                  detail.Id,                  // 商品SpuId
			Name:                detail.Name,                // 商品名称
			CategoryId:          detail.CategoryId,          // 商品分类ID
			CategoryIds:         detail.CategoryIds,         // 商品分类ID集合
			CategoryName:        detail.CategoryName,        // 商品分类名称
			BrandId:             detail.BrandId,             // 品牌ID
			BrandName:           detail.BrandName,           // 品牌名称
			Unit:                detail.Unit,                // 单位
			Weight:              detail.Weight,              // 重量(kg)
			Keywords:            detail.Keywords,            // 关键词
			Brief:               detail.Brief,               // 简介
			Description:         detail.Description,         // 详细描述
			AlbumPics:           detail.AlbumPics,           // 画册图片，最多8张，以逗号分割
			MainPic:             detail.MainPic,             // 主图
			PriceRange:          detail.PriceRange,          // 价格区间
			PublishStatus:       detail.PublishStatus,       // 上架状态：0-下架，1-上架
			NewStatus:           detail.NewStatus,           // 新品状态:0->不是新品；1->新品
			RecommendStatus:     detail.RecommendStatus,     // 推荐状态；0->不推荐；1->推荐
			VerifyStatus:        detail.VerifyStatus,        // 审核状态：0->未审核；1->审核通过
			PreviewStatus:       detail.PreviewStatus,       // 是否为预告商品：0->不是；1->是
			Sort:                detail.Sort,                // 排序
			NewStatusSort:       detail.NewStatusSort,       // 新品排序
			RecommendStatusSort: detail.RecommendStatusSort, // 推荐排序
			Sales:               detail.Sales,               // 销量
			Stock:               detail.Stock,               // 库存
			LowStock:            detail.LowStock,            // 预警库存
			PromotionType:       detail.PromotionType,       // 促销类型：0->没有促销使用原价;1->使用促销价；2->使用会员价；3->使用阶梯价格；4->使用满减价格；5->秒杀
			DetailTitle:         detail.DetailTitle,         // 详情标题
			DetailDesc:          detail.DetailDesc,          // 详情描述
			DetailHtml:          detail.DetailHtml,          // 产品详情网页内容
			DetailMobileHtml:    detail.DetailMobileHtml,    // 移动端网页详情
		})
	}
	return list
}

// 当前秒杀场次
func queryHomeFlashPromotion(l *IndexLogic) types.HomeFlashPromotion {
	var resp types.HomeFlashPromotion
	// currentDate := time.Now().Format("2006-01-02")
	// flashPromotionList, _ := l.svcCtx.FlashPromotionService.QueryFlashPromotionListByDate(l.ctx, &smsclient.QueryFlashPromotionListByDateReq{
	// 	CurrentDate: currentDate,
	// })
	//
	// // 获取今天是否有活动
	// if len(flashPromotionList.List) > 0 {
	// 	currentTime := time.Now().Format("2006-01-02 15:04:05")
	// 	sessionByTimeResp, _ := l.svcCtx.FlashPromotionSessionService.QueryFlashPromotionSessionListByTime(l.ctx, &smsclient.QueryFlashPromotionSessionListByTimeReq{CurrentTIme: currentTime})
	//
	// 	// 如果今天有活动,则查询今天是否有场次
	// 	sessionListData := sessionByTimeResp.List
	// 	if len(sessionListData) > 0 {
	// 		date := sessionListData[0]
	// 		resp.StartTime = date.StartTime
	// 		resp.EndTime = date.EndTime
	//
	// 		// 查询当前次的下一场时间
	// 		nextSessionByTimeResp, _ := l.svcCtx.FlashPromotionSessionService.QueryFlashPromotionSessionListByTime(l.ctx, &smsclient.QueryFlashPromotionSessionListByTimeReq{CurrentTIme: date.StartTime})
	// 		if len(nextSessionByTimeResp.List) > 0 {
	// 			nextDate := nextSessionByTimeResp.List[0]
	// 			resp.NextStartTime = nextDate.StartTime
	// 			resp.NextEndTime = nextDate.EndTime
	// 		}
	//
	// 		// 查询关联
	// 		_, _ = l.svcCtx.FlashPromotionProductRelationService.QueryFlashPromotionProductRelationList(l.ctx, &smsclient.QueryFlashPromotionProductRelationListReq{
	// 			PageNum:                 1,
	// 			PageSize:                100,
	// 			FlashPromotionId:        flashPromotionList.List[0].Id,
	// 			FlashPromotionSessionId: sessionListData[0].Id,
	// 		})
	//
	// 		// todo 为了测试有数据,这里先注释,用下面模拟的数据
	// 		// todo =====================开始==========================
	// 		// var productIdLists []int64
	// 		// for _, item := range listResp.List {
	// 		//	productIdLists = append(productIdLists, item.ProductId)
	// 		// }
	//
	// 		var productIdLists []int64
	// 		productIdLists = append(productIdLists, 27)
	// 		productIdLists = append(productIdLists, 28)
	// 		productIdLists = append(productIdLists, 29)
	// 		productIdLists = append(productIdLists, 32)
	// 		// todo =====================结束==========================
	// 		// 设置商品
	// 		resp.ProductList = queryProductList(l.svcCtx.ProductService, productIdLists, l.ctx)
	// 	}
	// }
	// var productIdLists []int64
	// productIdLists = append(productIdLists, 27)
	// productIdLists = append(productIdLists, 28)
	// productIdLists = append(productIdLists, 29)
	// productIdLists = append(productIdLists, 32)
	//
	// // 设置商品
	// resp.ProductList = queryProductList(l.svcCtx.ProductService, productIdLists, l.ctx)
	return resp
}

// 推荐品牌
func queryBrandList(l *IndexLogic) []types.BrandData {
	result, _ := l.svcCtx.ProductBrandService.QueryProductBrandList(l.ctx, &pmsclient.QueryProductBrandListReq{
		PageNum:         1,
		PageSize:        6,
		Name:            "", // 品牌名称
		RecommendStatus: 1,  // 推荐状态
		IsEnabled:       1,  // 是否启用
	})

	var list []types.BrandData

	for _, detail := range result.List {
		list = append(list, types.BrandData{
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

		})
	}
	return list
}

// 获取轮播广告
func queryAdvertiseList(l *IndexLogic) []types.AdvertiseList {
	result, _ := l.svcCtx.HomeAdvertiseService.QueryHomeAdvertiseList(l.ctx, &smsclient.QueryHomeAdvertiseListReq{
		PageNum:  1,
		PageSize: 100,
		Type:     1, // 轮播位置：0->PC首页轮播；1->app首页轮播
		Status:   1, // 上下线状态：0->下线；1->上线
	})

	var list []types.AdvertiseList

	for _, detail := range result.List {
		list = append(list, types.AdvertiseList{
			Id:         detail.Id,         // 编号
			Name:       detail.Name,       // 名称
			Type:       detail.Type,       // 轮播位置：0->PC首页轮播；1->app首页轮播
			Pic:        detail.Pic,        // 图片地址
			StartTime:  detail.StartTime,  // 开始时间
			EndTime:    detail.EndTime,    // 结束时间
			Status:     detail.Status,     // 上下线状态：0->下线；1->上线
			ClickCount: detail.ClickCount, // 点击数
			OrderCount: detail.OrderCount, // 下单数
			Url:        detail.Url,        // 链接地址
			Remark:     detail.Remark,     // 备注
			Sort:       detail.Sort,       // 排序

		})
	}

	return list
}
