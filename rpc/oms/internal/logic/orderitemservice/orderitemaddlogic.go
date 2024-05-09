package orderitemservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/oms/gen/model"
	"github.com/feihua/zero-admin/rpc/oms/gen/query"
	"github.com/feihua/zero-admin/rpc/oms/internal/svc"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"

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

func (l *OrderItemAddLogic) OrderItemAdd(in *omsclient.OrderItemAddReq) (*omsclient.OrderItemAddResp, error) {

	err := query.OmsOrderItem.WithContext(l.ctx).Create(&model.OmsOrderItem{
		OrderID:           in.OrderId,
		OrderSn:           in.OrderSn,
		ProductID:         in.ProductId,
		ProductPic:        in.ProductPic,
		ProductName:       in.ProductName,
		ProductBrand:      in.ProductBrand,
		ProductSn:         in.ProductSn,
		ProductPrice:      float64(in.ProductPrice),
		ProductQuantity:   in.ProductQuantity,
		ProductSkuID:      in.ProductSkuId,
		ProductSkuCode:    in.ProductSkuCode,
		ProductCategoryID: in.ProductCategoryId,
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

	return &omsclient.OrderItemAddResp{}, nil
}
