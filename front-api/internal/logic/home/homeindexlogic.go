package home

import (
	"context"
	"time"
	"zero-admin/rpc/cms/cmsclient"
	"zero-admin/rpc/pms/pmsclient"
	"zero-admin/rpc/sms/smsclient"

	"zero-admin/front-api/internal/svc"
	"zero-admin/front-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type HomeIndexLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHomeIndexLogic(ctx context.Context, svcCtx *svc.ServiceContext) HomeIndexLogic {
	return HomeIndexLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// HomeIndex 获取首页数据
func (l *HomeIndexLogic) HomeIndex() (resp *types.HomeResp, err error) {

	return &types.HomeResp{
		Code:    0,
		Message: "操作成功",
		Data: types.Data{
			AdvertiseList:      queryAdvertiseList(l),
			BrandList:          queryBrandList(l),
			HomeFlashPromotion: queryHomeFlashPromotion(l),
			NewProductList:     queryNewProductList(l),
			HotProductList:     queryHotProductList(l),
			SubjectList:        querySubjectList(l),
		},
	}, nil
}

//推荐专题
func querySubjectList(l *HomeIndexLogic) []types.SubjectList {
	homeRecommendSubjectList, _ := l.svcCtx.Sms.HomeRecommendSubjectList(l.ctx, &smsclient.HomeRecommendSubjectListReq{
		Current:         1,
		PageSize:        4,
		RecommendStatus: 1, //推荐状态：0->不推荐;1->推荐
	})

	var homeRecommendSubjectIdLists []int64
	for _, item := range homeRecommendSubjectList.List {
		homeRecommendSubjectIdLists = append(homeRecommendSubjectIdLists, item.SubjectId)
	}

	subjectListResp, _ := l.svcCtx.Cms.SubjectListByIds(l.ctx, &cmsclient.SubjectListByIdsReq{Ids: homeRecommendSubjectIdLists})
	var subjectLists []types.SubjectList
	for _, item := range subjectListResp.List {
		subjectLists = append(subjectLists, types.SubjectList{
			CategoryId:      item.CategoryId,
			Title:           item.Title,
			Pic:             item.Pic,
			ProductCount:    item.ProductCount,
			RecommendStatus: item.RecommendStatus,
			CollectCount:    item.CollectCount,
			ReadCount:       item.ReadCount,
			CommentCount:    item.CommentCount,
			AlbumPics:       item.AlbumPics,
			Description:     item.Description,
			ShowStatus:      item.ShowStatus,
			Content:         item.Content,
			ForwardCount:    item.ForwardCount,
			CategoryName:    item.CategoryName,
		})
	}
	return subjectLists
}

//人气推荐
func queryHotProductList(l *HomeIndexLogic) []types.ProductList {
	homeRecommendProductList, _ := l.svcCtx.Sms.HomeRecommendProductList(l.ctx, &smsclient.HomeRecommendProductListReq{
		Current:         1,
		PageSize:        4,
		RecommendStatus: 1, //推荐状态：0->不推荐;1->推荐
	})

	var productIds []int64
	for _, item := range homeRecommendProductList.List {
		productIds = append(productIds, item.ProductId)
	}

	return queryProductList(l.svcCtx.Pms, productIds, l.ctx)
}

//新品推荐
func queryNewProductList(l *HomeIndexLogic) []types.ProductList {
	homeNewProductList, _ := l.svcCtx.Sms.HomeNewProductList(l.ctx, &smsclient.HomeNewProductListReq{
		Current:         1,
		PageSize:        4,
		RecommendStatus: 1, //推荐状态：0->不推荐;1->推荐
	})

	var productIds []int64
	for _, item := range homeNewProductList.List {
		productIds = append(productIds, item.ProductId)
	}

	return queryProductList(l.svcCtx.Pms, productIds, l.ctx)
}

//当前秒杀场次
func queryHomeFlashPromotion(l *HomeIndexLogic) types.HomeFlashPromotion {
	var resp types.HomeFlashPromotion
	currentDate := time.Now().Format("2006-01-02")
	flashPromotionList, _ := l.svcCtx.Sms.FlashPromotionListByDate(l.ctx, &smsclient.FlashPromotionListByDateReq{
		CurrentDate: currentDate,
	})

	//获取今天是否有活动
	if len(flashPromotionList.List) > 0 {
		currentTime := time.Now().Format("2006-01-02 15:04:05")
		sessionByTimeResp, _ := l.svcCtx.Sms.FlashPromotionSessionByTime(l.ctx, &smsclient.FlashPromotionSessionByTimeReq{CurrentTIme: currentTime})

		//如果今天有活动,则查询今天是否有场次
		sessionListData := sessionByTimeResp.List
		if len(sessionListData) > 0 {
			date := sessionListData[0]
			resp.StartTime = date.StartTime
			resp.EndTime = date.EndTime

			//查询当前次的下一场时间
			nextSessionByTimeResp, _ := l.svcCtx.Sms.FlashPromotionSessionByTime(l.ctx, &smsclient.FlashPromotionSessionByTimeReq{CurrentTIme: date.StartTime})
			if len(nextSessionByTimeResp.List) > 0 {
				nextDate := nextSessionByTimeResp.List[0]
				resp.NextStartTime = nextDate.StartTime
				resp.NextEndTime = nextDate.EndTime
			}

			//查询关联
			listResp, _ := l.svcCtx.Sms.FlashPromotionProductRelationList(l.ctx, &smsclient.FlashPromotionProductRelationListReq{
				Current:                 1,
				PageSize:                100,
				FlashPromotionId:        flashPromotionList.List[0].Id,
				FlashPromotionSessionId: sessionListData[0].Id,
			})

			var productIdLists []int64
			for _, item := range listResp.List {
				productIdLists = append(productIdLists, item.ProductId)
			}

			//设置商品
			resp.ProductList = queryProductList(l.svcCtx.Pms, productIdLists, l.ctx)
		}
	}

	return resp
}

//推荐品牌
func queryBrandList(l *HomeIndexLogic) []types.BrandList {
	homeBrandList, _ := l.svcCtx.Sms.HomeBrandList(l.ctx, &smsclient.HomeBrandListReq{
		Current:         1,
		PageSize:        6,
		RecommendStatus: 1, //推荐状态：0->不推荐;1->推荐
	})

	var brandIdLists []int64
	for _, item := range homeBrandList.List {
		brandIdLists = append(brandIdLists, item.BrandId)
	}

	brandListResp, _ := l.svcCtx.Pms.BrandListByIds(l.ctx, &pmsclient.BrandListByIdsReq{Ids: brandIdLists})
	var brandLists []types.BrandList
	for _, item := range brandListResp.List {

		brandLists = append(brandLists, types.BrandList{
			ID:                  item.Id,
			Name:                item.Name,
			FirstLetter:         item.FirstLetter,
			Sort:                item.Sort,
			FactoryStatus:       item.FactoryStatus,
			ShowStatus:          item.ShowStatus,
			ProductCount:        item.ProductCount,
			ProductCommentCount: item.ProductCommentCount,
			Logo:                item.Logo,
			BigPic:              item.BigPic,
		})
	}
	return brandLists
}

//获取轮播广告
func queryAdvertiseList(l *HomeIndexLogic) []types.AdvertiseList {
	homeAdvertiseList, _ := l.svcCtx.Sms.HomeAdvertiseList(l.ctx, &smsclient.HomeAdvertiseListReq{
		Current:  1,
		PageSize: 100,
		Type:     1, //轮播位置：0->PC首页轮播；1->app首页轮播
		Status:   1, //上下线状态：0->下线；1->上线
	})

	var advertiseLists []types.AdvertiseList
	for _, item := range homeAdvertiseList.List {
		advertiseLists = append(advertiseLists, types.AdvertiseList{
			ID:         item.Id,
			Name:       item.Name,
			Type:       item.Type,
			Pic:        item.Pic,
			StartTime:  item.StartTime,
			EndTime:    item.EndTime,
			Status:     item.Status,
			ClickCount: item.ClickCount,
			OrderCount: item.OrderCount,
			URL:        item.Url,
			Sort:       item.Sort,
		})
	}
	return advertiseLists
}
