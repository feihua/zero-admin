package logic

import (
	"context"
	"fmt"
	"go-zero-admin/rpc/pms/internal/svc"
	"go-zero-admin/rpc/pms/pms"

	"github.com/tal-tech/go-zero/core/logx"
)

type ProductListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProductListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductListLogic {
	return &ProductListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ProductListLogic) ProductList(in *pms.ProductListReq) (*pms.ProductListResp, error) {
	all, _ := l.svcCtx.PmsProductModel.FindAll(in.Current, in.PageSize)
	count, _ := l.svcCtx.PmsProductModel.Count()

	var list []*pms.ProductListData
	for _, product := range *all {

		list = append(list, &pms.ProductListData{
			Id:                         product.Id,
			BrandId:                    product.BrandId,
			ProductCategoryId:          product.ProductCategoryId,
			FeightTemplateId:           product.FeightTemplateId,
			ProductAttributeCategoryId: product.ProductAttributeCategoryId,
			Name:                       product.Name,
			Pic:                        product.Pic,
			ProductSn:                  product.ProductSn,
			DeleteStatus:               product.DeleteStatus,
			PublishStatus:              product.PublishStatus,
			NewStatus:                  product.NewStatus,
			RecommandStatus:            product.RecommandStatus,
			VerifyStatus:               product.VerifyStatus,
			Sort:                       product.Sort,
			Sale:                       product.Sale,
			Price:                      product.Price,
			PromotionPrice:             product.PromotionPrice,
			GiftGrowth:                 product.GiftGrowth,
			GiftPoint:                  product.GiftPoint,
			UsePointLimit:              product.UsePointLimit,
			SubTitle:                   product.SubTitle,
			Description:                product.Description,
			OriginalPrice:              product.OriginalPrice,
			Stock:                      product.Stock,
			LowStock:                   product.LowStock,
			Unit:                       product.Unit,
			Weight:                     product.Weight,
			PreviewStatus:              product.PreviewStatus,
			ServiceIds:                 product.ServiceIds,
			Keywords:                   product.Keywords,
			Note:                       product.Note,
			AlbumPics:                  product.AlbumPics,
			DetailTitle:                product.DetailTitle,
			DetailDesc:                 product.DetailDesc,
			DetailHtml:                 product.DetailHtml,
			DetailMobileHtml:           product.DetailMobileHtml,
			PromotionStartTime:         product.PromotionStartTime.Format("2006-01-02 15:04:05"),
			PromotionEndTime:           product.PromotionEndTime.Format("2006-01-02 15:04:05"),
			PromotionPerLimit:          product.PromotionPerLimit,
			PromotionType:              product.PromotionType,
			BrandName:                  product.BrandName,
			ProductCategoryName:        product.ProductCategoryName,
		})
	}

	fmt.Println(list)
	return &pms.ProductListResp{
		Total: count,
		List:  list,
	}, nil
}
