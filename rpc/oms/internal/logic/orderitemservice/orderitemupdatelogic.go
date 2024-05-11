package orderitemservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/oms/gen/model"
	"github.com/feihua/zero-admin/rpc/oms/gen/query"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"

	"github.com/feihua/zero-admin/rpc/oms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrderItemUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOrderItemUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderItemUpdateLogic {
	return &OrderItemUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *OrderItemUpdateLogic) OrderItemUpdate(in *omsclient.OrderItemUpdateReq) (*omsclient.OrderItemUpdateResp, error) {
	q := query.OmsOrderItem
	_, err := q.WithContext(l.ctx).Updates(&model.OmsOrderItem{
		ID:                in.Id,
		OrderID:           in.OrderId,
		OrderSn:           in.OrderSn,
		ProductID:         in.ProductId,
		ProductPic:        in.ProductPic,
		ProductName:       in.ProductName,
		ProductBrand:      in.ProductBrand,
		ProductSn:         in.ProductSn,
		ProductPrice:      in.ProductPrice,
		ProductQuantity:   in.ProductQuantity,
		ProductSkuID:      in.ProductSkuId,
		ProductSkuCode:    in.ProductSkuCode,
		ProductCategoryID: in.ProductCategoryId,
		PromotionName:     in.PromotionName,
		PromotionAmount:   in.PromotionAmount,
		CouponAmount:      in.CouponAmount,
		IntegrationAmount: in.IntegrationAmount,
		RealAmount:        in.RealAmount,
		GiftIntegration:   in.GiftIntegration,
		GiftGrowth:        in.GiftGrowth,
		ProductAttr:       in.ProductAttr,
	})

	if err != nil {
		return nil, err
	}

	return &omsclient.OrderItemUpdateResp{}, nil
}
