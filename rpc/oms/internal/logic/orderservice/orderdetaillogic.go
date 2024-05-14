package orderservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/oms/gen/model"
	"github.com/feihua/zero-admin/rpc/oms/gen/query"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/rpc/oms/internal/svc"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// OrderDetailLogic 获取订单详情：订单信息、商品信息、操作记录
/*
Author: LiuFeiHua
Date: 2024/5/14 17:33
*/
type OrderDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOrderDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderDetailLogic {
	return &OrderDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// OrderDetail 获取订单详情：订单信息、商品信息、操作记录
func (l *OrderDetailLogic) OrderDetail(in *omsclient.OrderDetailReq) (*omsclient.OrderDetailResp, error) {
	order, err := query.OmsOrder.WithContext(l.ctx).Where(query.OmsOrder.ID.Eq(in.OrderId)).First()

	if err != nil {
		logc.Errorf(l.ctx, "获取订单详情失败,参数：%+v,异常:%s", in, err.Error())
		return nil, err
	}

	result, _ := query.OmsOrderItem.WithContext(l.ctx).Where(query.OmsOrderItem.OrderID.Eq(in.OrderId)).Find()

	histories, _ := query.OmsOrderOperateHistory.WithContext(l.ctx).Where(query.OmsOrderOperateHistory.OrderID.Eq(in.OrderId)).Find()

	orderListData := &omsclient.OrderListData{
		Id:                    order.ID,
		MemberId:              order.MemberID,
		CouponId:              order.CouponID,
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
		ItemListData:          buildOrderItemData(result),
		HistoryListData:       buildOperateHistory(histories),
	}

	logc.Infof(l.ctx, "获取订单详情信息,参数：%+v,响应：%+v", in, orderListData)
	return &omsclient.OrderDetailResp{
		Data: orderListData,
	}, nil

}

func buildOperateHistory(histories []*model.OmsOrderOperateHistory) []*omsclient.OrderOperateHistoryListData {
	var historyListData []*omsclient.OrderOperateHistoryListData

	for _, operateHistory := range histories {
		historyListData = append(historyListData, &omsclient.OrderOperateHistoryListData{
			Id:          operateHistory.ID,
			OrderId:     operateHistory.OrderID,
			OperateMan:  operateHistory.OperateMan,
			CreateTime:  operateHistory.CreateTime.Format("2006-01-02 15:04:05"),
			OrderStatus: operateHistory.OrderStatus,
			Note:        *operateHistory.Note,
		})
	}
	return historyListData
}

func buildOrderItemData(result []*model.OmsOrderItem) []*omsclient.OrderItemListData {
	var itemListData []*omsclient.OrderItemListData
	for _, item := range result {
		itemListData = append(itemListData, &omsclient.OrderItemListData{
			Id:                item.ID,
			OrderId:           item.OrderID,
			OrderSn:           item.OrderSn,
			ProductId:         item.ProductID,
			ProductPic:        item.ProductPic,
			ProductName:       item.ProductName,
			ProductBrand:      item.ProductBrand,
			ProductSn:         item.ProductSn,
			ProductPrice:      item.ProductPrice,
			ProductQuantity:   item.ProductQuantity,
			ProductSkuId:      item.ProductSkuID,
			ProductSkuCode:    item.ProductSkuCode,
			ProductCategoryId: item.ProductCategoryID,
			PromotionName:     item.PromotionName,
			PromotionAmount:   item.PromotionAmount,
			CouponAmount:      item.CouponAmount,
			IntegrationAmount: item.IntegrationAmount,
			RealAmount:        item.RealAmount,
			GiftIntegration:   item.GiftIntegration,
			GiftGrowth:        item.GiftGrowth,
			ProductAttr:       item.ProductAttr,
		})
	}
	return itemListData
}
