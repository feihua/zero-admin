package productservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/pms/gen/model"
	"github.com/feihua/zero-admin/rpc/pms/gen/query"
	"math/rand"
	"strconv"
	"time"

	"github.com/feihua/zero-admin/rpc/pms/internal/svc"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// AddProductLogic  添加商品,返回商品id
/*
Author: LiuFeiHua
Date: 2024/6/12 17:14
*/
type AddProductLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddProductLogic {
	return &AddProductLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// AddProduct 添加商品,返回商品id
// 逻辑步骤：
// 1.创建商品基本信息
// 2.会员价格
// 3.阶梯价格
// 4.满减价格
// 5.添加sku库存信息
// 6.添加商品参数,添加自定义商品规格
func (l *AddProductLogic) AddProduct(in *pmsclient.AddProductReq) (*pmsclient.AddProductResp, error) {
	PromotionStartTime := time.Now()
	if len(in.PromotionStartTime) > 0 {
		PromotionStartTime, _ = time.Parse("2006-01-02 15:04:05", in.PromotionStartTime)
	}

	PromotionEndTime := time.Now()
	if len(in.PromotionEndTime) > 0 {
		PromotionEndTime, _ = time.Parse("2006-01-02 15:04:05", in.PromotionEndTime)
	}

	// 1.创建商品
	product := model.PmsProduct{
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
	}
	err := query.PmsProduct.WithContext(l.ctx).Create(&product)

	if err != nil {
		return nil, err
	}

	productId := product.ID
	// 根据促销类型设置价格：会员价格、阶梯价格、满减价格
	// 根据促销类型设置价格：会员价格、阶梯价格、满减价格
	// 2.会员价格
	memberPrice := query.PmsMemberPrice.WithContext(l.ctx)
	for _, list := range in.MemberPriceList {
		_ = memberPrice.Create(&model.PmsMemberPrice{
			ProductID:       productId,
			MemberLevelID:   list.MemberLevelId,
			MemberPrice:     list.MemberPrice,
			MemberLevelName: list.MemberLevelName,
		})
	}

	// 3.阶梯价格
	ladder := query.PmsProductLadder.WithContext(l.ctx)
	for _, list := range in.ProductLadderList {
		_ = ladder.Create(&model.PmsProductLadder{
			ProductID: productId,
			Count:     list.Count,
			Discount:  list.Discount,
			Price:     list.Price,
		})
	}
	// 4.满减价格
	full := query.PmsProductFullReduction.WithContext(l.ctx)
	for _, list := range in.ProductFullReductionList {
		_ = full.Create(&model.PmsProductFullReduction{
			ProductID:   productId,
			FullPrice:   list.FullPrice,
			ReducePrice: list.ReducePrice,
		})
	}

	// 5.更新sku库存信息
	sku := query.PmsSkuStock.WithContext(l.ctx)
	for _, list := range in.SkuStockList {
		if list.SkuCode == "" {
			list.SkuCode = time.Now().Format("200601021504") + strconv.Itoa(rand.Intn(10))
		}
		_ = sku.Create(&model.PmsSkuStock{
			ProductID:      productId,
			SkuCode:        list.SkuCode,
			Price:          list.Price,
			Stock:          list.Stock,
			LowStock:       list.LowStock,
			Pic:            list.Pic,
			Sale:           list.Sale,
			PromotionPrice: list.PromotionPrice,
			LockStock:      list.LockStock,
			SpData:         list.SpData,
		})
	}
	// 6.更新商品参数,添加自定义商品规格
	attr := query.PmsProductAttributeValue.WithContext(l.ctx)
	for _, list := range in.ProductAttributeValueList {
		_ = attr.Create(&model.PmsProductAttributeValue{
			ProductID:          productId,
			ProductAttributeID: list.ProductAttributeId,
			Value:              list.AttributeValues,
		})
	}

	return &pmsclient.AddProductResp{
		ProductId: productId,
	}, nil

}
