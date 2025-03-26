package product

import (
	"context"
	"github.com/feihua/zero-admin/pkg/errorx"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc/status"

	"github.com/feihua/zero-admin/api/front/internal/svc"
	"github.com/feihua/zero-admin/api/front/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryProductListLogic
/*
Author: LiuFeiHua
Date: 2023/11/30 16:35
*/
type QueryProductListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryProductListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryProductListLogic {
	return &QueryProductListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// QueryProductList 查询商品列表
func (l *QueryProductListLogic) QueryProductList(req *types.QueryProductListReq) (resp *types.QueryProductListResp, err error) {
	productListResp, err := l.svcCtx.ProductService.QueryProductList(l.ctx, &pmsclient.QueryProductListReq{
		Current:           req.Current,
		PageSize:          req.PageSize,
		VerifyStatus:      1, // 审核状态：0->未审核；1->审核通过
		ProductCategoryId: req.ProductCategoryId,
		PublishStatus:     1, // 上架状态：0->下架；1->上架
		DeleteStatus:      0, // 删除状态：0->未删除；1->已删除
		BrandId:           req.BrandId,
		Name:              req.Keyword,
	})

	if err != nil {
		logc.Errorf(l.ctx, "查询商品列表失败,参数: %+v,异常：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	var productLists []types.QueryProductListData
	for _, item := range productListResp.List {
		productLists = append(productLists, types.QueryProductListData{
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

	return &types.QueryProductListResp{
		Code:    0,
		Message: "操作成功",
		Data:    productLists,
	}, nil
}
