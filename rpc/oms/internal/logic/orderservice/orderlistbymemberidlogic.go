package orderservicelogic

import (
	"context"
	"encoding/json"
	"github.com/feihua/zero-admin/rpc/oms/internal/svc"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrderListByMemberIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOrderListByMemberIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderListByMemberIdLogic {
	return &OrderListByMemberIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *OrderListByMemberIdLogic) OrderListByMemberId(in *omsclient.OrderListByMemberIdReq) (*omsclient.OrderListByMemberIdResp, error) {
	order, err := l.svcCtx.OmsOrderModel.FindOneByMemberIdAndOrderId(l.ctx, in.MemberId, in.Id)

	if err != nil {
		reqStr, _ := json.Marshal(in)
		logx.WithContext(l.ctx).Errorf("查询订单列表信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, err
	}

	data := &omsclient.OrderListData{
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
		ItemListData:          queryOrderItemsById(l, order.Id),
	}

	reqStr, _ := json.Marshal(in)
	listStr, _ := json.Marshal(data)

	logx.WithContext(l.ctx).Infof("查询订单列表信息,参数：%s,响应：%s", reqStr, listStr)
	return &omsclient.OrderListByMemberIdResp{
		Total: 0,
		Data:  data,
	}, nil
}

func queryOrderItemsById(l *OrderListByMemberIdLogic, OrderId int64) []*omsclient.OrderItemListData {
	orderItem, _ := l.svcCtx.OmsOrderItemModel.FindAll(l.ctx, 1, 100, OrderId)
	var itemListData []*omsclient.OrderItemListData
	for _, item := range *orderItem {
		itemListData = append(itemListData, &omsclient.OrderItemListData{
			Id:                item.Id,
			OrderId:           item.OrderId,
			OrderSn:           item.OrderSn,
			ProductId:         item.ProductId,
			ProductPic:        item.ProductPic,
			ProductName:       item.ProductName,
			ProductBrand:      item.ProductBrand,
			ProductSn:         item.ProductSn,
			ProductPrice:      float32(item.ProductPrice),
			ProductQuantity:   item.ProductQuantity,
			ProductSkuId:      item.ProductSkuId,
			ProductSkuCode:    item.ProductSkuCode,
			ProductCategoryId: item.ProductCategoryId,
			PromotionName:     item.PromotionName,
			PromotionAmount:   float32(item.PromotionAmount),
			CouponAmount:      float32(item.CouponAmount),
			IntegrationAmount: float32(item.IntegrationAmount),
			RealAmount:        float32(item.RealAmount),
			GiftIntegration:   item.GiftIntegration,
			GiftGrowth:        item.GiftGrowth,
			ProductAttr:       item.ProductAttr,
		})
	}
	return itemListData
}
