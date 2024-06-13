package brand

import (
	"context"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"

	"github.com/feihua/zero-admin/api/front/internal/svc"
	"github.com/feihua/zero-admin/api/front/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryBrandListLogic 分页获取推荐品牌
/*
Author: LiuFeiHua
Date: 2024/5/16 14:15
*/
type QueryBrandListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryBrandListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryBrandListLogic {
	return &QueryBrandListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// QueryBrandList 分页获取推荐品牌
func (l *QueryBrandListLogic) QueryBrandList(req *types.BrandListReq) (resp *types.BrandListResp, err error) {
	homeBrandList, _ := l.svcCtx.HomeBrandService.QueryHomeBrandList(l.ctx, &smsclient.QueryHomeBrandListReq{
		PageNum:         req.Current,
		PageSize:        req.PageSize,
		RecommendStatus: 1, //推荐状态：0->不推荐;1->推荐
	})

	var brandIdLists []int64
	for _, item := range homeBrandList.List {
		brandIdLists = append(brandIdLists, item.BrandId)
	}

	brandListResp, _ := l.svcCtx.BrandService.QueryBrandListByIds(l.ctx, &pmsclient.QueryBrandListByIdsReq{Ids: brandIdLists})
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
		Message: "分页获取推荐品牌成功",
		Data:    list,
	}, nil
}
