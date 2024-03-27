package orderitemservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/model/omsmodel"
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
	err := l.svcCtx.OmsOrderItemModel.Update(l.ctx, &omsmodel.OmsOrderItem{
		Id:                in.Id,
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

	return &omsclient.OrderItemUpdateResp{}, nil
}
