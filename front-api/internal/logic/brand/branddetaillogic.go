package brand

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"

	"github.com/feihua/zero-admin/front-api/internal/svc"
	"github.com/feihua/zero-admin/front-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// BrandDetailLogic
/*
Author: LiuFeiHua
Date: 2023/12/4 13:37
*/
type BrandDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewBrandDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BrandDetailLogic {
	return &BrandDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// BrandDetail 查询品牌详情
func (l *BrandDetailLogic) BrandDetail(req *types.BrandDetailReq) (*types.BrandDetailResp, error) {
	var ids []int64
	ids = append(ids, req.BrandId)
	resp, err := l.svcCtx.BrandService.BrandListByIds(l.ctx, &pmsclient.BrandListByIdsReq{
		Ids: ids,
	})

	if err != nil {
		logx.WithContext(l.ctx).Errorf("查询商品品牌详情异常:%s", err.Error())
		return nil, errors.New("查询商品品牌详情失败")
	}

	item := resp.List[0]
	brandId := item.Id
	productListResp, _ := l.svcCtx.ProductService.ProductList(l.ctx, &pmsclient.ProductListReq{
		Current:           1,
		PageSize:          100,
		Name:              "",
		VerifyStatus:      2,
		ProductCategoryId: 0,
		PublishStatus:     2,
		DeleteStatus:      2,
		BrandId:           brandId,
	})
	productLists := make([]types.BrandProductList, 0)

	for _, item := range productListResp.List {
		productLists = append(productLists, types.BrandProductList{
			Id:                         item.Id,
			BrandId:                    item.BrandId,
			ProductCategoryId:          item.ProductCategoryId,
			FeightTemplateId:           item.FeightTemplateId,
			ProductAttributeCategoryId: item.ProductAttributeCategoryId,
			Name:                       item.Name,
			Pic:                        item.Pic,
			ProductSn:                  item.ProductSn,
			DeleteStatus:               item.DeleteStatus,
			PublishStatus:              item.PublishStatus,
			NewStatus:                  item.NewStatus,
			RecommandStatus:            item.RecommandStatus,
			VerifyStatus:               item.VerifyStatus,
			Sort:                       item.Sort,
			Sale:                       item.Sale,
			Price:                      item.Price,
			PromotionPrice:             item.PromotionPrice,
			GiftGrowth:                 item.GiftGrowth,
			GiftPoint:                  item.GiftPoint,
			UsePointLimit:              item.UsePointLimit,
			SubTitle:                   item.SubTitle,
			OriginalPrice:              item.OriginalPrice,
			Stock:                      item.Stock,
			LowStock:                   item.LowStock,
			Unit:                       item.Unit,
			Weight:                     item.Weight,
			PreviewStatus:              item.PreviewStatus,
			ServiceIDS:                 item.ServiceIds,
			Keywords:                   item.Keywords,
			Note:                       item.Note,
			AlbumPics:                  item.AlbumPics,
			DetailTitle:                item.DetailTitle,
			PromotionStartTime:         item.PromotionStartTime,
			PromotionEndTime:           item.PromotionEndTime,
			PromotionPerLimit:          item.PromotionPerLimit,
			PromotionType:              item.PromotionType,
			BrandName:                  item.BrandName,
			ProductCategoryName:        item.ProductCategoryName,
			Description:                item.Description,
		})
	}
	return &types.BrandDetailResp{
		Code:    0,
		Message: "操作成功",
		Data: types.BrandDetailData{
			Id:                  brandId,
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
			ProductList:         productLists,
		},
	}, nil
}
