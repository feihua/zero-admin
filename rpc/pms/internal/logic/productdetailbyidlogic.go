package logic

import (
	"context"
	"encoding/json"
	"zero-admin/rpc/pms/internal/svc"
	"zero-admin/rpc/pms/pmsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProductDetailByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProductDetailByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductDetailByIdLogic {
	return &ProductDetailByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ProductDetailByIdLogic) ProductDetailById(in *pmsclient.ProductDetailByIdReq) (*pmsclient.ProductDetailByIdResp, error) {
	product, err := l.svcCtx.PmsProductModel.FindOne(in.Id)

	if err != nil {
		reqStr, _ := json.Marshal(in)
		logx.WithContext(l.ctx).Errorf("查询商品信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, err
	}

	data := &pmsclient.ProductDetailByIdResp{
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
	}

	reqStr, _ := json.Marshal(in)
	listStr, _ := json.Marshal(data)
	logx.WithContext(l.ctx).Infof("查询商品列表信息,参数：%s,响应：%s", reqStr, listStr)

	return data, nil
}
