package brand

import (
	"context"
	"github.com/feihua/zero-admin/api/front/internal/svc"
	"github.com/feihua/zero-admin/api/front/internal/types"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryBrandProductListLogic 分页获取品牌相关商品
/*
Author: LiuFeiHua
Date: 2024/5/16 14:18
*/
type QueryBrandProductListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryBrandProductListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryBrandProductListLogic {
	return &QueryBrandProductListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// QueryBrandProductList 分页获取品牌相关商品
func (l *QueryBrandProductListLogic) QueryBrandProductList(req *types.QueryBrandProductListReq) (resp *types.QueryBrandProductListResp, err error) {

	productListResp, _ := l.svcCtx.ProductService.ProductList(l.ctx, &pmsclient.ProductListReq{
		Current:       req.Current,
		PageSize:      req.PageSize,
		VerifyStatus:  2,
		PublishStatus: 2,
		DeleteStatus:  2,
		BrandId:       req.BrandId,
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
	return &types.QueryBrandProductListResp{
		Code:    0,
		Message: "分页获取品牌相关商品成功",
		Data:    productLists,
	}, nil
}
