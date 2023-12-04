package brand

import (
	"context"
	"zero-admin/rpc/pms/pmsclient"
	"zero-admin/rpc/sms/smsclient"

	"zero-admin/front-api/internal/svc"
	"zero-admin/front-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// BrandListLogic 查询推荐品牌列表
/*
Author: LiuFeiHua
Date: 2023/12/4 13:34
*/
type BrandListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewBrandListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BrandListLogic {
	return &BrandListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// BrandList 查询推荐品牌列表
func (l *BrandListLogic) BrandList() (*types.BrandListResp, error) {
	homeBrandList, _ := l.svcCtx.HomeBrandService.HomeBrandList(l.ctx, &smsclient.HomeBrandListReq{
		Current:         1,
		PageSize:        100,
		RecommendStatus: 1, //推荐状态：0->不推荐;1->推荐
	})

	var brandIdLists []int64
	for _, item := range homeBrandList.List {
		brandIdLists = append(brandIdLists, item.BrandId)
	}

	brandListResp, _ := l.svcCtx.BrandService.BrandListByIds(l.ctx, &pmsclient.BrandListByIdsReq{Ids: brandIdLists})
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

	var list []types.BrandListData

	for _, item := range brandListResp.List {
		list = append(list, types.BrandListData{
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
		})
	}

	return &types.BrandListResp{
		Code:    0,
		Message: "操作成功",
		Data:    list,
	}, nil
}
