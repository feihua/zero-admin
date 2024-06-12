package orderservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/oms/gen/query"
	"github.com/feihua/zero-admin/rpc/oms/internal/svc"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"
	"github.com/zeromicro/go-zero/core/logc"

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
	q := query.OmsOrder
	//1.查询订单是否存在
	order, err := q.WithContext(l.ctx).Where(q.ID.Eq(in.Id), q.MemberID.Eq(in.MemberId)).First()

	if err != nil {
		logc.Errorf(l.ctx, "查询订单列表信息失败,参数：%+v,异常:%s", in, err.Error())
		return nil, err
	}

	data := &omsclient.OrderListData{
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
		ItemListData:          queryOrderItemsById(l, order.ID),
	}

	logc.Infof(l.ctx, "查询订单列表信息,参数：%+v,响应：%+v", in, data)
	return &omsclient.OrderListByMemberIdResp{
		Total: 0,
		Data:  data,
	}, nil
}

func queryOrderItemsById(l *OrderListByMemberIdLogic, OrderId int64) []*omsclient.OrderItemData {
	result, _ := query.OmsOrderItem.WithContext(l.ctx).Where(query.OmsOrderItem.OrderID.Eq(OrderId)).Find()
	var itemListData []*omsclient.OrderItemData
	for _, item := range result {
		itemListData = append(itemListData, &omsclient.OrderItemData{
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
