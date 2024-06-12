package orderitemservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/oms/gen/model"
	"github.com/feihua/zero-admin/rpc/oms/gen/query"

	"github.com/feihua/zero-admin/rpc/oms/internal/svc"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// AddOrderItemLogic 添加订单中所包含的商品
/*
Author: LiuFeiHua
Date: 2024/6/12 10:12
*/
type AddOrderItemLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddOrderItemLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddOrderItemLogic {
	return &AddOrderItemLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// AddOrderItem 添加订单中所包含的商品
func (l *AddOrderItemLogic) AddOrderItem(in *omsclient.AddOrderItemReq) (*omsclient.AddOrderItemResp, error) {
	err := query.OmsOrderItem.WithContext(l.ctx).Create(&model.OmsOrderItem{
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

	return &omsclient.AddOrderItemResp{}, nil
}
