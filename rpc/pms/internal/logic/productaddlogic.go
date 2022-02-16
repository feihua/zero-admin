package logic

import (
	"context"
	"time"
	"zero-admin/rpc/model/pmsmodel"

	"zero-admin/rpc/pms/internal/svc"
	"zero-admin/rpc/pms/pms"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProductAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProductAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductAddLogic {
	return &ProductAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ProductAddLogic) ProductAdd(in *pms.ProductAddReq) (*pms.ProductAddResp, error) {
	PromotionStartTime, _ := time.Parse("2006-01-02 15:04:05", in.PromotionStartTime)
	PromotionEndTime, _ := time.Parse("2006-01-02 15:04:05", in.PromotionEndTime)
	_, err := l.svcCtx.PmsProductModel.Insert(pmsmodel.PmsProduct{
		BrandId:                    in.BrandId,
		ProductCategoryId:          in.ProductCategoryId,
		FeightTemplateId:           in.FeightTemplateId,
		ProductAttributeCategoryId: in.ProductAttributeCategoryId,
		Name:                       in.Name,
		Pic:                        in.Pic,
		ProductSn:                  in.ProductSn,
		DeleteStatus:               in.DeleteStatus,
		PublishStatus:              in.PublishStatus,
		NewStatus:                  in.NewStatus,
		RecommandStatus:            in.RecommandStatus,
		VerifyStatus:               in.VerifyStatus,
		Sort:                       in.Sort,
		Sale:                       in.Sale,
		Price:                      in.Price,
		PromotionPrice:             in.PromotionPrice,
		GiftGrowth:                 in.GiftGrowth,
		GiftPoint:                  in.GiftPoint,
		UsePointLimit:              in.UsePointLimit,
		SubTitle:                   in.SubTitle,
		Description:                in.Description,
		OriginalPrice:              in.OriginalPrice,
		Stock:                      in.Stock,
		LowStock:                   in.LowStock,
		Unit:                       in.Unit,
		Weight:                     in.Weight,
		PreviewStatus:              in.PreviewStatus,
		ServiceIds:                 in.ServiceIds,
		Keywords:                   in.Keywords,
		Note:                       in.Note,
		AlbumPics:                  in.AlbumPics,
		DetailTitle:                in.DetailTitle,
		DetailDesc:                 in.DetailDesc,
		DetailHtml:                 in.DetailHtml,
		DetailMobileHtml:           in.DetailMobileHtml,
		PromotionStartTime:         PromotionStartTime,
		PromotionEndTime:           PromotionEndTime,
		PromotionPerLimit:          in.PromotionPerLimit,
		PromotionType:              in.PromotionType,
		BrandName:                  in.BrandName,
		ProductCategoryName:        in.ProductCategoryName,
	})
	if err != nil {
		return nil, err
	}

	return &pms.ProductAddResp{}, nil
}
