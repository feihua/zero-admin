package product

import (
	"context"
	"zero-admin/rpc/pms/pmsclient"

	"zero-admin/front-api/internal/svc"
	"zero-admin/front-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryProductLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryProductLogic {
	return &QueryProductLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryProductLogic) QueryProduct(req *types.QueryProductReq) (resp *types.QueryProductResp, err error) {
	item, _ := l.svcCtx.Pms.ProductDetailById(l.ctx, &pmsclient.ProductDetailByIdReq{
		Id: req.Id,
	})
	var product types.QueryProductListData = types.QueryProductListData{
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
	}

	return &types.QueryProductResp{
		Code:    0,
		Message: "操作成功",
		Data:    product,
	}, nil
}
