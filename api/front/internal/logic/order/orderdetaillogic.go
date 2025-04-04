package order

import (
	"context"
	"github.com/feihua/zero-admin/api/front/internal/logic/common"
	"github.com/feihua/zero-admin/pkg/errorx"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc/status"

	"github.com/feihua/zero-admin/api/front/internal/svc"
	"github.com/feihua/zero-admin/api/front/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// OrderDetailLogic
/*
Author: LiuFeiHua
Date: 2023/12/7 15:22
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

// OrderDetail 订单详情
func (l *OrderDetailLogic) OrderDetail(req *types.OrderDetailReq) (*types.OrderDetailResp, error) {
	memberId, err := common.GetMemberId(l.ctx)
	if err != nil {
		return nil, err
	}
	resp, err := l.svcCtx.OrderService.OrderListByMemberId(l.ctx, &omsclient.OrderListByMemberIdReq{
		Id:       req.OrderId,
		MemberId: memberId,
	})

	if err != nil {
		logc.Errorf(l.ctx, "获取订单详情失败,参数: %+v,异常：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	item := resp.Data

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
		ListOrderItem:         queryOrderItems(item),
	}

	return &types.OrderDetailResp{
		Code:    0,
		Message: "操作成功",
		Data:    data,
	}, nil

}
