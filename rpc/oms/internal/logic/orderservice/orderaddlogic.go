package orderservicelogic

import (
	"context"
	"time"
	"zero-admin/rpc/model/omsmodel"
	"zero-admin/rpc/oms/internal/svc"
	"zero-admin/rpc/oms/omsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// OrderAddLogic
/*
Author: LiuFeiHua
Date: 2023/12/13 9:29
*/
type OrderAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOrderAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderAddLogic {
	return &OrderAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// OrderAdd 创建订单
//1.插入order表
//2.插入order_item表
//2.删除购物车中的下单商品(删除cart_item表)
func (l *OrderAddLogic) OrderAdd(in *omsclient.OrderAddReq) (*omsclient.OrderAddResp, error) {

	//1.插入order表
	_, err := l.svcCtx.OmsOrderModel.Insert(l.ctx, &omsmodel.OmsOrder{
		MemberId:              in.MemberId,
		CouponId:              in.CouponId,
		OrderSn:               in.OrderSn,
		CreateTime:            time.Now(),
		MemberUsername:        in.MemberUsername,
		TotalAmount:           float64(in.TotalAmount),
		PayAmount:             float64(in.PayAmount),
		FreightAmount:         float64(in.FreightAmount),
		PromotionAmount:       float64(in.PromotionAmount),
		IntegrationAmount:     float64(in.IntegrationAmount),
		CouponAmount:          float64(in.CouponAmount),
		DiscountAmount:        float64(in.DiscountAmount),
		PayType:               in.PayType,
		SourceType:            in.SourceType,
		Status:                in.Status,
		OrderType:             in.OrderType,
		DeliveryCompany:       in.DeliveryCompany,
		DeliverySn:            in.DeliverySn,
		AutoConfirmDay:        in.AutoConfirmDay,
		Integration:           in.Integration,
		Growth:                in.Growth,
		PromotionInfo:         in.PromotionInfo,
		BillType:              in.BillType,
		BillHeader:            in.BillHeader,
		BillContent:           in.BillContent,
		BillReceiverPhone:     in.BillReceiverPhone,
		BillReceiverEmail:     in.BillReceiverEmail,
		ReceiverName:          in.ReceiverName,
		ReceiverPhone:         in.ReceiverPhone,
		ReceiverPostCode:      in.ReceiverPostCode,
		ReceiverProvince:      in.ReceiverProvince,
		ReceiverCity:          in.ReceiverCity,
		ReceiverRegion:        in.ReceiverRegion,
		ReceiverDetailAddress: in.ReceiverDetailAddress,
		Note:                  in.Note,
		ConfirmStatus:         in.ConfirmStatus,
		DeleteStatus:          in.DeleteStatus,
		UseIntegration:        in.UseIntegration,
		PaymentTime:           time.Now(),
		DeliveryTime:          time.Now(),
		ReceiveTime:           time.Now(),
		CommentTime:           time.Now(),
		ModifyTime:            time.Now(),
	})
	if err != nil {
		return nil, err
	}
	//2.插入order_item表
	for _, orderItem := range in.OrderItemList {
		_, _ = l.svcCtx.OmsOrderItemModel.Insert(l.ctx, &omsmodel.OmsOrderItem{
			OrderId:           0,
			OrderSn:           "",
			ProductId:         orderItem.ProductId,
			ProductPic:        orderItem.ProductPic,
			ProductName:       orderItem.ProductName,
			ProductBrand:      orderItem.ProductBrand,
			ProductSn:         orderItem.ProductSn,
			ProductPrice:      float64(orderItem.ProductPrice),
			ProductQuantity:   orderItem.ProductQuantity,
			ProductSkuId:      orderItem.ProductSkuId,
			ProductSkuCode:    orderItem.ProductSkuCode,
			ProductCategoryId: orderItem.ProductCategoryId,
			PromotionName:     orderItem.PromotionName,
			PromotionAmount:   float64(orderItem.PromotionAmount),
			CouponAmount:      float64(orderItem.CouponAmount),
			IntegrationAmount: float64(orderItem.IntegrationAmount),
			RealAmount:        float64(orderItem.RealAmount),
			GiftIntegration:   orderItem.GiftIntegration,
			GiftGrowth:        orderItem.GiftGrowth,
			ProductAttr:       orderItem.ProductAttr,
		})

	}
	//2.删除购物车中的下单商品(删除cart_item表)
	err = l.svcCtx.OmsCartItemModel.DeleteByIds(l.ctx, in.MemberId, in.CartItemIds)

	if err != nil {
		return nil, err
	}
	return &omsclient.OrderAddResp{}, nil
}
