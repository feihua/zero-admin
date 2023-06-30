package logic

import (
	"context"
	"database/sql"
	"math/rand"
	"strconv"
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

// ProductAdd 添加商品,返回商品id
//逻辑步骤：
//1.创建商品基本信息
//2.会员价格
//3.阶梯价格
//4.满减价格
//5.添加sku库存信息
//6.添加商品参数,添加自定义商品规格
func (l *ProductAddLogic) ProductAdd(in *pms.ProductAddReq) (*pms.ProductAddResp, error) {
	PromotionStartTime := time.Now()
	if len(in.PromotionStartTime) > 0 {
		PromotionStartTime, _ = time.Parse("2006-01-02 15:04:05", in.PromotionStartTime)
	}

	PromotionEndTime := time.Now()
	if len(in.PromotionEndTime) > 0 {
		PromotionEndTime, _ = time.Parse("2006-01-02 15:04:05", in.PromotionEndTime)
	}

	//1.创建商品
	result, err := l.svcCtx.PmsProductModel.Insert(l.ctx, &pmsmodel.PmsProduct{
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
		ProductCategoryIdArray:     in.ProductCategoryIdArray,
	})
	if err != nil {
		return nil, err
	}

	productId, _ := result.LastInsertId()
	//根据促销类型设置价格：会员价格、阶梯价格、满减价格
	//2.会员价格
	for _, list := range in.MemberPriceList {
		_, _ = l.svcCtx.PmsMemberPriceModel.Insert(l.ctx, &pmsmodel.PmsMemberPrice{
			ProductId:       productId,
			MemberLevelId:   list.MemberLevelId,
			MemberPrice:     float64(list.MemberPrice),
			MemberLevelName: list.MemberLevelName,
		})
	}

	//3.阶梯价格
	for _, list := range in.ProductLadderList {
		_, _ = l.svcCtx.PmsProductLadderModel.Insert(l.ctx, &pmsmodel.PmsProductLadder{
			ProductId: productId,
			Count:     list.Count,
			Discount:  float64(list.Discount),
			Price:     float64(list.Price),
		})
	}
	//4.满减价格
	for _, list := range in.ProductFullReductionList {
		_, _ = l.svcCtx.PmsProductFullReductionModel.Insert(l.ctx, &pmsmodel.PmsProductFullReduction{
			ProductId:   productId,
			FullPrice:   float64(list.FullPrice),
			ReducePrice: float64(list.ReducePrice),
		})
	}
	//5.添加sku库存信息
	for _, list := range in.SkuStockList {
		_, _ = l.svcCtx.PmsSkuStockModel.Insert(l.ctx, &pmsmodel.PmsSkuStock{
			ProductId:      productId,
			SkuCode:        time.Now().Format("200601021504") + strconv.Itoa(rand.Intn(10)),
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
	//6.添加商品参数,添加自定义商品规格
	for _, list := range in.ProductAttributeValueList {
		_, _ = l.svcCtx.PmsProductAttributeValueModel.Insert(l.ctx, &pmsmodel.PmsProductAttributeValue{
			ProductId:          productId,
			ProductAttributeId: list.ProductAttributeId,
			Value:              sql.NullString{String: list.AttributeValues, Valid: true},
		})
	}
	return &pms.ProductAddResp{
		ProductId: int32(productId),
	}, nil
}
