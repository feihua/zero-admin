package logic

import (
	"context"
	"time"
	"zero-admin/rpc/model/omsmodel"
	"zero-admin/rpc/oms/internal/svc"
	"zero-admin/rpc/oms/oms"

	"github.com/zeromicro/go-zero/core/logx"
)

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

func (l *OrderAddLogic) OrderAdd(in *oms.OrderAddReq) (*oms.OrderAddResp, error) {

	createTime, _ := time.Parse("2006-01-02 15:04:05", in.CreateTime)
	paymentTime, _ := time.Parse("2006-01-02 15:04:05", in.PaymentTime)
	deliveryTime, _ := time.Parse("2006-01-02 15:04:05", in.DeliveryTime)
	receiveTime, _ := time.Parse("2006-01-02 15:04:05", in.ReceiveTime)
	commentTime, _ := time.Parse("2006-01-02 15:04:05", in.CommentTime)
	modifyTime, _ := time.Parse("2006-01-02 15:04:05", in.ModifyTime)

	_, err := l.svcCtx.OmsOrderModel.Insert(omsmodel.OmsOrder{
		MemberId:              in.MemberId,
		CouponId:              in.CouponId,
		OrderSn:               in.OrderSn,
		CreateTime:            createTime,
		MemberUsername:        in.MemberUsername,
		TotalAmount:           in.TotalAmount,
		PayAmount:             in.PayAmount,
		FreightAmount:         in.FreightAmount,
		PromotionAmount:       in.PromotionAmount,
		IntegrationAmount:     in.IntegrationAmount,
		CouponAmount:          in.CouponAmount,
		DiscountAmount:        in.DiscountAmount,
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
		PaymentTime:           paymentTime,
		DeliveryTime:          deliveryTime,
		ReceiveTime:           receiveTime,
		CommentTime:           commentTime,
		ModifyTime:            modifyTime,
	})
	if err != nil {
		return nil, err
	}

	return &oms.OrderAddResp{}, nil
}
