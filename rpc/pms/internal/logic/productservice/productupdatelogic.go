package productservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/pms/gen/model"
	"github.com/feihua/zero-admin/rpc/pms/gen/query"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"
	"math/rand"
	"strconv"
	"time"

	"github.com/feihua/zero-admin/rpc/pms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProductUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProductUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductUpdateLogic {
	return &ProductUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// ProductUpdate 更新商品
// 逻辑步骤：
// 1.更新商品基本信息
// 2.会员价格
// 3.阶梯价格
// 4.满减价格
// 5.更新sku库存信息
// 6.更新商品参数,添加自定义商品规格
func (l *ProductUpdateLogic) ProductUpdate(in *pmsclient.ProductUpdateReq) (*pmsclient.ProductUpdateResp, error) {
	PromotionStartTime, _ := time.Parse("2006-01-02 15:04:05", in.PromotionStartTime)
	PromotionEndTime, _ := time.Parse("2006-01-02 15:04:05", in.PromotionEndTime)

	productId := in.Id

	q := query.PmsProduct
	_, err := q.WithContext(l.ctx).Updates(&model.PmsProduct{
		ID:                         productId,
		BrandID:                    in.BrandId,
		ProductCategoryID:          in.ProductCategoryId,
		FeightTemplateID:           in.FeightTemplateId,
		ProductAttributeCategoryID: in.ProductAttributeCategoryId,
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
		DetailHTML:                 in.DetailHtml,
		DetailMobileHTML:           in.DetailMobileHtml,
		PromotionStartTime:         PromotionStartTime,
		PromotionEndTime:           PromotionEndTime,
		PromotionPerLimit:          in.PromotionPerLimit,
		PromotionType:              in.PromotionType,
		BrandName:                  in.BrandName,
		ProductCategoryName:        in.ProductCategoryName,
		ProductCategoryIDArray:     in.ProductCategoryIdArray,
	})

	if err != nil {
		return nil, err
	}

	//根据促销类型设置价格：会员价格、阶梯价格、满减价格
	//2.会员价格
	memberPrice := query.PmsMemberPrice.WithContext(l.ctx)
	_, _ = memberPrice.Where(query.PmsMemberPrice.ProductID.Eq(productId)).Delete()
	for _, list := range in.MemberPriceList {
		_ = memberPrice.Create(&model.PmsMemberPrice{
			ProductID:       productId,
			MemberLevelID:   list.MemberLevelId,
			MemberPrice:     float64(list.MemberPrice),
			MemberLevelName: list.MemberLevelName,
		})
	}

	//3.阶梯价格
	ladder := query.PmsProductLadder.WithContext(l.ctx)
	_, _ = ladder.Where(query.PmsProductLadder.ProductID.Eq(productId)).Delete()
	for _, list := range in.ProductLadderList {
		_ = ladder.Create(&model.PmsProductLadder{
			ProductID: productId,
			Count:     list.Count,
			Discount:  float64(list.Discount),
			Price:     float64(list.Price),
		})
	}
	//4.满减价格
	full := query.PmsProductFullReduction.WithContext(l.ctx)
	_, _ = full.Where(query.PmsProductFullReduction.ProductID.Eq(productId)).Delete()
	for _, list := range in.ProductFullReductionList {
		_ = full.Create(&model.PmsProductFullReduction{
			ProductID:   productId,
			FullPrice:   float64(list.FullPrice),
			ReducePrice: float64(list.ReducePrice),
		})
	}

	//5.更新sku库存信息
	sku := query.PmsSkuStock.WithContext(l.ctx)
	_, _ = sku.Where(query.PmsSkuStock.ProductID.Eq(productId)).Delete()
	for _, list := range in.SkuStockList {
		if list.SkuCode == "" {
			list.SkuCode = time.Now().Format("200601021504") + strconv.Itoa(rand.Intn(10))
		}
		_ = sku.Create(&model.PmsSkuStock{
			ProductID:      productId,
			SkuCode:        list.SkuCode,
			Price:          float64(list.Price),
			Stock:          list.Stock,
			LowStock:       list.LowStock,
			Pic:            list.Pic,
			Sale:           list.Sale,
			PromotionPrice: float64(list.PromotionPrice),
			LockStock:      list.LockStock,
			SpData:         list.SpData,
		})
	}
	//6.更新商品参数,添加自定义商品规格
	attr := query.PmsProductAttributeValue.WithContext(l.ctx)
	_, _ = attr.Where(query.PmsProductAttributeValue.ProductID.Eq(productId)).Delete()
	for _, list := range in.ProductAttributeValueList {
		_ = attr.Create(&model.PmsProductAttributeValue{
			ProductID:          productId,
			ProductAttributeID: list.ProductAttributeId,
			Value:              &list.AttributeValues,
		})
	}

	return &pmsclient.ProductUpdateResp{}, nil
}
