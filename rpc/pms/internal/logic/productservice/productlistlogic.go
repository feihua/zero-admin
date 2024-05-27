package productservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/pms/gen/query"
	"github.com/feihua/zero-admin/rpc/pms/internal/svc"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/zeromicro/go-zero/core/logx"
)

// ProductListLogic
/*
Author: LiuFeiHua
Date: 2023/11/30 16:44
*/
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

// ProductList 查询商品列表
func (l *ProductListLogic) ProductList(in *pmsclient.ProductListReq) (*pmsclient.ProductListResp, error) {
	q := query.PmsProduct.WithContext(l.ctx)
	if len(in.Name) > 0 {
		q = q.Where(query.PmsProduct.Name.Like("%" + in.Name + "%"))
	}
	if len(in.ProductSn) > 0 {
		q = q.Where(query.PmsProduct.ProductSn.Like("%" + in.ProductSn + "%"))
	}
	if in.VerifyStatus != 2 {
		q = q.Where(query.PmsProduct.VerifyStatus.Eq(in.VerifyStatus))
	}
	if in.ProductCategoryId != 0 {
		q = q.Where(query.PmsProduct.ProductCategoryID.Eq(in.ProductCategoryId))
	}
	if in.BrandId != 0 {
		q = q.Where(query.PmsProduct.BrandID.Eq(in.BrandId))
	}
	if in.PublishStatus != 2 {
		q = q.Where(query.PmsProduct.PublishStatus.Eq(in.PublishStatus))
	}
	if in.DeleteStatus != 2 {
		q = q.Where(query.PmsProduct.DeleteStatus.Eq(in.DeleteStatus))
	}

	result, count, err := q.FindByPage(int((in.Current-1)*in.PageSize), int(in.PageSize))

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
			ProductCategoryIdArray:     product.ProductCategoryIDArray,
		})
	}

	logc.Infof(l.ctx, "查询商品列表信息,参数：%+v,响应：%+v", in, list)
	return &pmsclient.ProductListResp{
		Total: count,
		List:  list,
	}, nil
}
