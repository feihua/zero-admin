package logic

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrderAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOrderAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) OrderAddLogic {
	return OrderAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OrderAddLogic) OrderAdd(req types.AddOrderReq) (*types.AddOrderResp, error) {
	_, err := l.svcCtx.OrderService.OrderAdd(l.ctx, &omsclient.OrderAddReq{
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
		logc.Errorf(l.ctx, "添加订单信息失败,参数：%+v,响应：%s", req, err.Error())
		return nil, errorx.NewDefaultError("添加订单信息失败")
	}

	return &types.AddOrderResp{
		Code:    "000000",
		Message: "添加订单信息成功",
	}, nil
}
