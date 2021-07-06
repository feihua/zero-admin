package logic

import (
	"context"
	"go-zero-admin/api/internal/common/errorx"
	"go-zero-admin/rpc/pms/pmsclient"

	"go-zero-admin/api/internal/svc"
	"go-zero-admin/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type ProductUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProductUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) ProductUpdateLogic {
	return ProductUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ProductUpdateLogic) ProductUpdate(req types.UpdateProductReq) (*types.UpdateProductResp, error) {
	_, err := l.svcCtx.Pms.ProductUpdate(l.ctx, &pmsclient.ProductUpdateReq{
		Id:                         req.Id,
		BrandId:                    req.BrandId,
		ProductCategoryId:          req.ProductCategoryId,
		FeightTemplateId:           req.FeightTemplateId,
		ProductAttributeCategoryId: req.ProductAttributeCategoryId,
		Name:                       req.Name,
		Pic:                        req.Pic,
		ProductSn:                  req.ProductSn,
		DeleteStatus:               req.DeleteStatus,
		PublishStatus:              req.PublishStatus,
		NewStatus:                  req.NewStatus,
		RecommandStatus:            req.RecommandStatus,
		VerifyStatus:               req.VerifyStatus,
		Sort:                       req.Sort,
		Sale:                       req.Sale,
		Price:                      req.Price,
		PromotionPrice:             req.PromotionPrice,
		GiftGrowth:                 req.GiftGrowth,
		GiftPoint:                  req.GiftPoint,
		UsePointLimit:              req.UsePointLimit,
		SubTitle:                   req.SubTitle,
		Description:                req.Description,
		OriginalPrice:              req.OriginalPrice,
		Stock:                      req.Stock,
		LowStock:                   req.LowStock,
		Unit:                       req.Unit,
		Weight:                     req.Weight,
		PreviewStatus:              req.PreviewStatus,
		ServiceIds:                 req.ServiceIds,
		Keywords:                   req.Keywords,
		Note:                       req.Note,
		AlbumPics:                  req.AlbumPics,
		DetailTitle:                req.DetailTitle,
		DetailDesc:                 req.DetailDesc,
		DetailHtml:                 req.DetailHtml,
		DetailMobileHtml:           req.DetailMobileHtml,
		PromotionStartTime:         req.PromotionStartTime,
		PromotionEndTime:           req.PromotionEndTime,
		PromotionPerLimit:          req.PromotionPerLimit,
		PromotionType:              req.PromotionType,
		BrandName:                  req.BrandName,
		ProductCategoryName:        req.ProductCategoryName,
	})

	if err != nil {
		return nil, errorx.NewDefaultError("更新商品信息失败")
	}

	return &types.UpdateProductResp{
		Code:    "000000",
		Message: "更新商品信息成功",
	}, nil
}
