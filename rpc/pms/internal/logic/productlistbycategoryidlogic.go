package logic

import (
	"context"
	"encoding/json"
	"zero-admin/rpc/pms/pms"

	"zero-admin/rpc/pms/internal/svc"
	"zero-admin/rpc/pms/pmsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProductListByCategoryIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProductListByCategoryIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductListByCategoryIdLogic {
	return &ProductListByCategoryIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ProductListByCategoryIdLogic) ProductListByCategoryId(in *pmsclient.ProductListByCategoryIdReq) (*pmsclient.ProductListByCategoryIdResp, error) {
	all, err := l.svcCtx.PmsProductModel.FindAll(1, 1000)
	count, _ := l.svcCtx.PmsProductModel.Count()

	if err != nil {
		reqStr, _ := json.Marshal(in)
		logx.WithContext(l.ctx).Errorf("查询商品列表信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, err
	}

	var list []*pms.ProductListByCategoryIdData
	for _, product := range *all {

		list = append(list, &pms.ProductListByCategoryIdData{
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

	reqStr, _ := json.Marshal(in)
	listStr, _ := json.Marshal(list)
	logx.WithContext(l.ctx).Infof("查询商品列表信息,参数：%s,响应：%s", reqStr, listStr)
	return &pms.ProductListByCategoryIdResp{
		Total: count,
		List:  list,
	}, nil

}
