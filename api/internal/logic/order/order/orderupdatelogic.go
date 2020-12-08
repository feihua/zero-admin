package logic

import (
	"context"
	"go-zero-admin/rpc/oms/omsclient"

	"go-zero-admin/api/internal/svc"
	"go-zero-admin/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type OrderUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOrderUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) OrderUpdateLogic {
	return OrderUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OrderUpdateLogic) OrderUpdate(req types.UpdateOrderReq) (*types.UpdateOrderResp, error) {
	_, err := l.svcCtx.Oms.OrderUpdate(l.ctx, &omsclient.OrderUpdateReq{
		Id:                    req.Id,
		MemberId:              req.MemberId,
		CouponId:              req.CouponId,
		OrderSn:               req.OrderSn,
		CreateTime:            req.CreateTime,
		MemberUsername:        req.MemberUsername,
		TotalAmount:           req.TotalAmount,
		PayAmount:             req.PayAmount,
		FreightAmount:         req.FreightAmount,
		PromotionAmount:       req.PromotionAmount,
		IntegrationAmount:     req.IntegrationAmount,
		CouponAmount:          req.CouponAmount,
		DiscountAmount:        req.DiscountAmount,
		PayType:               req.PayType,
		SourceType:            req.SourceType,
		Status:                req.Status,
		OrderType:             req.OrderType,
		DeliveryCompany:       req.DeliveryCompany,
		DeliverySn:            req.DeliverySn,
		AutoConfirmDay:        req.AutoConfirmDay,
		Integration:           req.Integration,
		Growth:                req.Growth,
		PromotionInfo:         req.PromotionInfo,
		BillType:              req.BillType,
		BillHeader:            req.BillHeader,
		BillContent:           req.BillContent,
		BillReceiverPhone:     req.BillReceiverPhone,
		BillReceiverEmail:     req.BillReceiverEmail,
		ReceiverName:          req.ReceiverName,
		ReceiverPhone:         req.ReceiverPhone,
		ReceiverPostCode:      req.ReceiverPostCode,
		ReceiverProvince:      req.ReceiverProvince,
		ReceiverCity:          req.ReceiverCity,
		ReceiverRegion:        req.ReceiverRegion,
		ReceiverDetailAddress: req.ReceiverDetailAddress,
		Note:                  req.Note,
		ConfirmStatus:         req.ConfirmStatus,
		DeleteStatus:          req.DeleteStatus,
		UseIntegration:        req.UseIntegration,
	})

	if err != nil {
		return nil, err
	}

	return &types.UpdateOrderResp{}, nil
}
