package home

import (
	"context"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"

	"github.com/feihua/zero-admin/api/front/internal/svc"
	"github.com/feihua/zero-admin/api/front/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RecommendBrandListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRecommendBrandListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RecommendBrandListLogic {
	return &RecommendBrandListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RecommendBrandListLogic) RecommendBrandList(req *types.RecommendBrandListReq) (resp *types.RecommendBrandListResp, err error) {
	homeBrandList, _ := l.svcCtx.HomeBrandService.HomeBrandList(l.ctx, &smsclient.HomeBrandListReq{
		Current:         req.Current,
		PageSize:        req.PageSize,
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

	return &types.RecommendBrandListResp{
		Code:    0,
		Message: "操作成功",
		Data:    brandLists,
	}, nil
}
