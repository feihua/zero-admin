package logic

import (
	"context"
	"encoding/json"
	"zero-admin/api/internal/common/errorx"
	"zero-admin/rpc/pms/pmsclient"

	"zero-admin/api/internal/svc"
	"zero-admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProductAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProductAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) ProductAddLogic {
	return ProductAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ProductAddLogic) ProductAdd(req types.AddProductReq) (*types.AddProductResp, error) {
	_, err := l.svcCtx.Pms.ProductAdd(l.ctx, &pmsclient.ProductAddReq{
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
		reqStr, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("添加商品信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, errorx.NewDefaultError("添加商品信息失败")
	}

	return &types.AddProductResp{
		Code:    "000000",
		Message: "添加商品信息成功",
	}, nil
}
