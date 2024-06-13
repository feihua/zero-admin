package productservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/pms/gen/query"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/rpc/pms/internal/svc"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryProductListByIdsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryProductListByIdsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryProductListByIdsLogic {
	return &QueryProductListByIdsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *QueryProductListByIdsLogic) QueryProductListByIds(in *pmsclient.QueryProductByIdsReq) (*pmsclient.QueryProductListResp, error) {
	q := query.PmsProduct
	result, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Find()

	if err != nil {
		logc.Errorf(l.ctx, "查询商品列表信息失败,参数：%+v,异常:%s", in, err.Error())
		return nil, err
	}

	var list []*pmsclient.ProductListData
	for _, product := range result {

		list = append(list, &pmsclient.ProductListData{
			Id:                         product.ID,
			BrandId:                    product.BrandID,
			ProductCategoryId:          product.ProductCategoryID,
			FeightTemplateId:           product.FeightTemplateID,
			ProductAttributeCategoryId: product.ProductAttributeCategoryID,
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
			DetailHtml:                 product.DetailHTML,
			DetailMobileHtml:           product.DetailMobileHTML,
			PromotionStartTime:         product.PromotionStartTime.Format("2006-01-02 15:04:05"),
			PromotionEndTime:           product.PromotionEndTime.Format("2006-01-02 15:04:05"),
			PromotionPerLimit:          product.PromotionPerLimit,
			PromotionType:              product.PromotionType,
			BrandName:                  product.BrandName,
			ProductCategoryName:        product.ProductCategoryName,
		})
	}

	logc.Infof(l.ctx, "查询商品列表信息,参数：%+v,响应：%+v", in, list)
	return &pmsclient.QueryProductListResp{
		Total: 0,
		List:  list,
	}, nil

}
