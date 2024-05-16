package product

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"strconv"
	"strings"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryProductListLogic
/*
Author: LiuFeiHua
Date: 2024/5/15 17:49
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
func (l *QueryProductListLogic) QueryProductList(req *types.QueryProductListReq) (*types.QueryProductListResp, error) {
	var resp, err = l.svcCtx.ProductService.ProductList(l.ctx, &pmsclient.ProductListReq{
		Current:           req.Current,
		PageSize:          req.PageSize,
		Name:              req.Name,
		VerifyStatus:      req.VerifyStatus,
		ProductCategoryId: req.ProductCategoryId,
		PublishStatus:     req.PublishStatus,
		DeleteStatus:      req.DeleteStatus,
		BrandId:           req.BrandId,
		ProductSn:         req.ProductSn,
	})

	if err != nil {
		logc.Errorf(l.ctx, "参数: %+v,查询商品信息列表异常:%s", req, err.Error())
		return nil, errorx.NewDefaultError("查询商品信息失败")
	}

	var list []*types.QueryProductListData

	for _, item := range resp.List {
		var productCategoryIdArray []int64
		for _, s := range strings.Split(item.ProductCategoryIdArray, ",") {
			id, _ := strconv.ParseInt(s, 10, 64)
			productCategoryIdArray = append(productCategoryIdArray, id)
		}

		list = append(list, &types.QueryProductListData{
			Id:                         item.Id,
			BrandId:                    item.BrandId,
			ProductCategoryId:          item.ProductCategoryId,
			ProductCategoryIdArray:     productCategoryIdArray,
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
			Description:                item.Description,
			OriginalPrice:              item.OriginalPrice,
			Stock:                      item.Stock,
			LowStock:                   item.LowStock,
			Unit:                       item.Unit,
			Weight:                     item.Weight,
			PreviewStatus:              item.PreviewStatus,
			ServiceIds:                 item.ServiceIds,
			Keywords:                   item.Keywords,
			Note:                       item.Note,
			AlbumPics:                  item.AlbumPics,
			DetailTitle:                item.DetailTitle,
			DetailDesc:                 item.DetailDesc,
			DetailHtml:                 item.DetailHtml,
			DetailMobileHtml:           item.DetailMobileHtml,
			PromotionStartTime:         item.PromotionStartTime,
			PromotionEndTime:           item.PromotionEndTime,
			PromotionPerLimit:          item.PromotionPerLimit,
			PromotionType:              item.PromotionType,
			BrandName:                  item.BrandName,
			ProductCategoryName:        item.ProductCategoryName,
		})
	}

	return &types.QueryProductListResp{
		Current:  req.Current,
		Data:     list,
		PageSize: req.PageSize,
		Success:  true,
		Total:    resp.Total,
		Code:     "000000",
		Message:  "查询商品信息成功",
	}, nil
}
