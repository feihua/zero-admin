package orderservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/model/omsmodel"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"
	"time"

	"github.com/feihua/zero-admin/rpc/oms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrderUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOrderUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderUpdateLogic {
	return &OrderUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// OrderUpdate 更新订单
func (l *OrderUpdateLogic) OrderUpdate(in *omsclient.OrderUpdateReq) (*omsclient.OrderUpdateResp, error) {
	order, err := l.svcCtx.OmsOrderModel.FindOne(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}
	err = l.svcCtx.OmsOrderModel.Update(l.ctx, &omsmodel.OmsOrder{
		Id:                    in.Id,
		MemberId:              order.MemberId,
		CouponId:              order.CouponId,
		OrderSn:               order.OrderSn,
		CreateTime:            order.CreateTime,
		MemberUsername:        order.MemberUsername,
		TotalAmount:           order.TotalAmount,
		PayAmount:             order.PayAmount,
		FreightAmount:         order.FreightAmount,
		PromotionAmount:       order.PromotionAmount,
		IntegrationAmount:     order.IntegrationAmount,
		CouponAmount:          order.CouponAmount,
		DiscountAmount:        order.DiscountAmount,
		PayType:               order.PayType,
		SourceType:            order.SourceType,
		Status:                in.Status,
		OrderType:             order.OrderType,
		DeliveryCompany:       in.DeliveryCompany,
		DeliverySn:            in.DeliverySn,
		AutoConfirmDay:        order.AutoConfirmDay,
		Integration:           order.Integration,
		Growth:                order.Growth,
		PromotionInfo:         order.PromotionInfo,
		BillType:              order.BillType,
		BillHeader:            order.BillHeader,
		BillContent:           order.BillContent,
		BillReceiverPhone:     order.BillReceiverPhone,
		BillReceiverEmail:     order.BillReceiverEmail,
		ReceiverName:          order.ReceiverName,
		ReceiverPhone:         order.ReceiverPhone,
		ReceiverPostCode:      order.ReceiverPostCode,
		ReceiverProvince:      order.ReceiverProvince,
		ReceiverCity:          order.ReceiverCity,
		ReceiverRegion:        order.ReceiverRegion,
		ReceiverDetailAddress: order.ReceiverDetailAddress,
		Note:                  in.Note,
		ConfirmStatus:         order.ConfirmStatus,
		DeleteStatus:          order.DeleteStatus,
		UseIntegration:        order.UseIntegration,
		PaymentTime:           order.PaymentTime,
		DeliveryTime:          time.Now(),
		ReceiveTime:           order.ReceiveTime,
		CommentTime:           order.CommentTime,
		ModifyTime:            order.ModifyTime,
	})
	if err != nil {
		return nil, err
	}

	return &omsclient.OrderUpdateResp{}, nil
}
