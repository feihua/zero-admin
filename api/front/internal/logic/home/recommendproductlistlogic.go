package home

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/api/front/internal/svc"
	"github.com/feihua/zero-admin/api/front/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// RecommendProductListLogic 分页获取推荐商品
/*
Author: LiuFeiHua
Date: 2024/5/16 15:12
*/
type RecommendProductListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRecommendProductListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RecommendProductListLogic {
	return &RecommendProductListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// RecommendProductList 分页获取推荐商品
func (l *RecommendProductListLogic) RecommendProductList(req *types.RecommendProductListReq) (resp *types.RecommendProductListResp, err error) {
	productListResp, err := l.svcCtx.ProductService.ProductList(l.ctx, &pmsclient.ProductListReq{
		Current:           req.Current,
		PageSize:          req.PageSize,
		VerifyStatus:      1, // 审核状态：0->未审核；1->审核通过
		ProductCategoryId: 0,
		PublishStatus:     1, // 上架状态：0->下架；1->上架
		DeleteStatus:      0, // 删除状态：0->未删除；1->已删除
		BrandId:           0,
	})

	if err != nil {
		logc.Errorf(l.ctx, "参数: %s,查询商品信息列表异常:%s", req, err.Error())
		return nil, errors.New("查询商品信息失败")
	}

	var productLists []types.ProductList
	for _, item := range productListResp.List {
		productLists = append(productLists, types.ProductList{
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

	return &types.RecommendProductListResp{
		Code:    0,
		Message: "分页获取推荐商品成功",
		Data:    productLists,
	}, nil
}
