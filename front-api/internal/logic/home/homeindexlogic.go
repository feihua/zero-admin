package home

import (
	"context"
	"encoding/json"
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

	//首页轮播图
	banners := queryBanners(l)

	//首页推荐品牌表
	brands := queryBrands(l)

	//新鲜好物表
	newProducts := queryNewProducts(l)

	//人气推荐商品表
	hotProducts := queryHotProducts(l)

	resp = &types.HomeResp{
		Errno: 0,
		Data: types.HomeData{
			NewGoodsList: newProducts,
			Channel:      brands,
			Banner:       banners,
			HotGoodsList: hotProducts,
		},
		Errmsg: "",
	}

	respStr, _ := json.Marshal(resp)
	logx.WithContext(l.ctx).Errorf("HomeIndex 获取首页数据: %s", respStr)
	return resp, nil
}

//人气推荐商品表
func queryHotProducts(l *HomeIndexLogic) []types.GoodsList {
	homeRecommendProductList, _ := l.svcCtx.Sms.HomeRecommendProductList(l.ctx, &smsclient.HomeRecommendProductListReq{
		Current:  1,
		PageSize: 10,
	})

	var goodsLists []types.GoodsList
	for _, item := range homeRecommendProductList.List {
		goodsLists = append(goodsLists, types.GoodsList{
			ID:           item.Id,
			Name:         "",
			Brief:        "",
			PicURL:       "",
			IsNew:        false,
			IsHot:        false,
			CounterPrice: 0,
			RetailPrice:  0,
		})
	}
	return goodsLists
}

//新鲜好物表
func queryNewProducts(l *HomeIndexLogic) []types.GoodsList {
	homeNewProductList, _ := l.svcCtx.Sms.HomeNewProductList(l.ctx, &smsclient.HomeNewProductListReq{
		Current:  1,
		PageSize: 10,
	})

	var goodsList []types.GoodsList
	for _, item := range homeNewProductList.List {
		goodsList = append(goodsList, types.GoodsList{
			ID:           item.Id,
			Name:         "",
			Brief:        "",
			PicURL:       "",
			IsNew:        false,
			IsHot:        false,
			CounterPrice: 0,
			RetailPrice:  0,
		})
	}
	return goodsList
}

//首页推荐品牌表
func queryBrands(l *HomeIndexLogic) []types.Channel {
	homeBrandList, _ := l.svcCtx.Sms.HomeBrandList(l.ctx, &smsclient.HomeBrandListReq{
		Current:  1,
		PageSize: 10,
	})

	var channels []types.Channel
	for _, item := range homeBrandList.List {
		channels = append(channels, types.Channel{
			ID:      item.BrandId,
			Name:    item.BrandName,
			IconURL: item.BrandName,
		})
	}
	return channels
}

//首页轮播图
func queryBanners(l *HomeIndexLogic) []types.Banner {
	homeAdvertiseList, _ := l.svcCtx.Sms.HomeAdvertiseList(l.ctx, &smsclient.HomeAdvertiseListReq{
		Current:  1,
		PageSize: 10,
	})

	var banners []types.Banner
	for _, item := range homeAdvertiseList.List {
		banners = append(banners, types.Banner{
			ID:       item.Id,
			Name:     item.Name,
			Link:     item.Url,
			URL:      item.Url,
			Position: item.Sort,
			//Content:    item.,
			Enabled: item.Status == 1,
			//AddTime:    item,
			//UpdateTime: item,
			//Deleted:    item,
		})
	}
	return banners
}
