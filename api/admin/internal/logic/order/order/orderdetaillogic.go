package order

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// OrderDetailLogic 订单信息
/*
Author: LiuFeiHua
Date: 2024/5/14 17:05
*/
type OrderDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOrderDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderDetailLogic {
	return &OrderDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// OrderDetail 获取订单详情：订单信息、商品信息、操作记录
func (l *OrderDetailLogic) OrderDetail(req *types.OrderDetailReq) (resp *types.OrderDetailResp, err error) {
	result, err := l.svcCtx.OrderService.OrderDetail(l.ctx, &omsclient.OrderDetailReq{
		OrderId: req.Id,
	})

	if err != nil {
		logc.Errorf(l.ctx, "参数: %+v,获取订单详情异常:%s", req, err.Error())
		return nil, errorx.NewDefaultError("获取订单详情失败")
	}

	item := result.Data
	data := types.ListOrderData{
		Id:                    item.Id,
		MemberId:              item.MemberId,
		CouponId:              item.CouponId,
		OrderSn:               item.OrderSn,
		CreateTime:            item.CreateTime,
		MemberUsername:        item.MemberUsername,
		TotalAmount:           item.TotalAmount,
		PayAmount:             item.PayAmount,
		FreightAmount:         item.FreightAmount,
		PromotionAmount:       item.PromotionAmount,
		IntegrationAmount:     item.IntegrationAmount,
		CouponAmount:          item.CouponAmount,
		DiscountAmount:        item.DiscountAmount,
		PayType:               item.PayType,
		SourceType:            item.SourceType,
		Status:                item.Status,
		OrderType:             item.OrderType,
		DeliveryCompany:       item.DeliveryCompany,
		DeliverySn:            item.DeliverySn,
		AutoConfirmDay:        item.AutoConfirmDay,
		Integration:           item.Integration,
		Growth:                item.Growth,
		PromotionInfo:         item.PromotionInfo,
		BillType:              item.BillType,
		BillHeader:            item.BillHeader,
		BillContent:           item.BillContent,
		BillReceiverPhone:     item.BillReceiverPhone,
		BillReceiverEmail:     item.BillReceiverEmail,
		ReceiverName:          item.ReceiverName,
		ReceiverPhone:         item.ReceiverPhone,
		ReceiverPostCode:      item.ReceiverPostCode,
		ReceiverProvince:      item.ReceiverProvince,
		ReceiverCity:          item.ReceiverCity,
		ReceiverRegion:        item.ReceiverRegion,
		ReceiverDetailAddress: item.ReceiverDetailAddress,
		Note:                  item.Note,
		ConfirmStatus:         item.ConfirmStatus,
		DeleteStatus:          item.DeleteStatus,
		UseIntegration:        item.UseIntegration,
		PaymentTime:           item.PaymentTime,
		DeliveryTime:          item.DeliveryTime,
		ReceiveTime:           item.ReceiveTime,
		CommentTime:           item.CommentTime,
		ModifyTime:            item.ModifyTime,
		ListOperateHistory:    queryHistoryRecord(item),
		ListOrderItem:         queryOrderItems(item),
	}

	return &types.OrderDetailResp{
		Data:    data,
		Code:    "000000",
		Message: "获取订单详情成功",
	}, nil
}
