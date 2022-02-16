package logic

import (
	"context"
	"zero-admin/rpc/model/omsmodel"
	"zero-admin/rpc/oms/internal/svc"
	"zero-admin/rpc/oms/oms"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrderItemAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOrderItemAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderItemAddLogic {
	return &OrderItemAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *OrderItemAddLogic) OrderItemAdd(in *oms.OrderItemAddReq) (*oms.OrderItemAddResp, error) {

	_, err := l.svcCtx.OmsOrderItemModel.Insert(omsmodel.OmsOrderItem{
		OrderId:           in.OrderId,
		OrderSn:           in.OrderSn,
		ProductId:         in.ProductId,
		ProductPic:        in.ProductPic,
		ProductName:       in.ProductName,
		ProductBrand:      in.ProductBrand,
		ProductSn:         in.ProductSn,
		ProductPrice:      float64(in.ProductPrice),
		ProductQuantity:   in.ProductQuantity,
		ProductSkuId:      in.ProductSkuId,
		ProductSkuCode:    in.ProductSkuCode,
		ProductCategoryId: in.ProductCategoryId,
		PromotionName:     in.PromotionName,
		PromotionAmount:   float64(in.PromotionAmount),
		CouponAmount:      float64(in.CouponAmount),
		IntegrationAmount: float64(in.IntegrationAmount),
		RealAmount:        float64(in.RealAmount),
		GiftIntegration:   in.GiftIntegration,
		GiftGrowth:        in.GiftGrowth,
		ProductAttr:       in.ProductAttr,
	})
	if err != nil {
		return nil, err
	}

	return &oms.OrderItemAddResp{}, nil
}
