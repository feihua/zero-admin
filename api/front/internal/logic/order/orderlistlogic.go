package order

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"

	"github.com/feihua/zero-admin/api/front/internal/svc"
	"github.com/feihua/zero-admin/api/front/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// OrderListLogic
/*
Author: LiuFeiHua
Date: 2023/12/7 15:22
*/
type OrderListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOrderListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderListLogic {
	return &OrderListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// OrderList 订单查询
func (l *OrderListLogic) OrderList(req *types.OrderListReq) (*types.OrderListResp, error) {
	memberId, _ := l.ctx.Value("memberId").(json.Number).Int64()
	resp, err := l.svcCtx.OrderService.QueryOrderList(l.ctx, &omsclient.QueryOrderListReq{
		Current:  req.Current,
		PageSize: req.PageSize,
		Status:   req.Status,
		MemberId: memberId,
	})

	if err != nil {
		data, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("参数: %s,查询订单信息列表异常:%s", string(data), err.Error())
		return nil, errors.New("查询订单信息失败")
	}

	list := make([]*types.ListOrderData, 0)

	for _, item := range resp.List {

		list = append(list, &types.ListOrderData{
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
		})
	}

	return &types.OrderListResp{
		Code:    0,
		Message: "操作成功",
		Data:    list,
	}, nil
}

// 获取商品项
func queryOrderItems(item1 *omsclient.OrderListData) []types.ListOrderItemData {
	itemListData := make([]types.ListOrderItemData, 0)
	for _, item := range item1.ItemListData {
		itemListData = append(itemListData, types.ListOrderItemData{
			Id:                item.Id,
			OrderId:           item.OrderId,
			OrderSn:           item.OrderSn,
			ProductId:         item.ProductId,
			ProductPic:        item.ProductPic,
			ProductName:       item.ProductName,
			ProductBrand:      item.ProductBrand,
			ProductSn:         item.ProductSn,
			ProductPrice:      item.ProductPrice,
			ProductQuantity:   item.ProductQuantity,
			ProductSkuId:      item.ProductSkuId,
			ProductSkuCode:    item.ProductSkuCode,
			ProductCategoryId: item.ProductCategoryId,
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
