package logic

import (
	"context"
	"fmt"
	"go-zero-admin/rpc/oms/internal/svc"
	"go-zero-admin/rpc/oms/oms"

	"github.com/tal-tech/go-zero/core/logx"
)

type OrderListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOrderListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderListLogic {
	return &OrderListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *OrderListLogic) OrderList(in *oms.OrderListReq) (*oms.OrderListResp, error) {
	all, _ := l.svcCtx.OmsOrderModel.FindAll(in.Current, in.PageSize)
	count, _ := l.svcCtx.OmsOrderModel.Count()

	var list []*oms.OrderListData
	for _, order := range *all {

		list = append(list, &oms.OrderListData{
			Id:                    order.Id,
			MemberId:              order.MemberId,
			CouponId:              order.CouponId,
			OrderSn:               order.OrderSn,
			CreateTime:            order.CreateTime.Format("2006-01-02 15:04:05"),
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
			Status:                order.Status,
			OrderType:             order.OrderType,
			DeliveryCompany:       order.DeliveryCompany,
			DeliverySn:            order.DeliverySn,
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
			Note:                  order.Note,
			ConfirmStatus:         order.ConfirmStatus,
			DeleteStatus:          order.DeleteStatus,
			UseIntegration:        order.UseIntegration,
			PaymentTime:           order.PaymentTime.Format("2006-01-02 15:04:05"),
			DeliveryTime:          order.DeliveryTime.Format("2006-01-02 15:04:05"),
			ReceiveTime:           order.ReceiveTime.Format("2006-01-02 15:04:05"),
			CommentTime:           order.CommentTime.Format("2006-01-02 15:04:05"),
			ModifyTime:            order.ModifyTime.Format("2006-01-02 15:04:05"),
		})
	}

	fmt.Println(list)
	return &oms.OrderListResp{
		Total: count,
		List:  list,
	}, nil
}
