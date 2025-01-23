package orderservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/oms/gen/model"
	"github.com/feihua/zero-admin/rpc/oms/gen/query"
	"github.com/feihua/zero-admin/rpc/oms/internal/svc"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
)

// OrderAddLogic
/*
Author: LiuFeiHua
Date: 2023/12/13 9:29
*/
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

// OrderAdd 创建订单
// 1.插入order表
// 2.插入order_item表
// 2.删除cart_item表(删除购物车中的下单商品)
func (l *OrderAddLogic) OrderAdd(in *omsclient.OrderAddReq) (*omsclient.OrderAddResp, error) {

	// 1.插入order表
	orderInfo := buildOrderInfo(in)
	logc.Infof(l.ctx, "插入order表,参数：%+v", orderInfo)
	err := query.OmsOrder.WithContext(l.ctx).Create(orderInfo)
	if err != nil {
		logc.Errorf(l.ctx, "插入order表失败,参数：%+v,异常：%s", orderInfo, err.Error())
		return nil, err
	}

	// 获取订单id
	orderId := orderInfo.ID
	// 2.插入order_item表
	for _, orderItem := range in.OrderItemList {
		buildOrderItem(l, orderInfo, orderItem)
	}

	// 3.删除cart_item表(删除购物车中的下单商品)
	logc.Infof(l.ctx, "删除购物车中的下单商品,参数ids：%+v", in.CartItemIds)
	item := query.OmsCartItem
	_, err = item.WithContext(l.ctx).Where(item.MemberID.Eq(in.MemberId), item.ID.In(in.CartItemIds...)).Delete()

	if err != nil {
		logc.Errorf(l.ctx, "删除购物车中的下单商品失败,参数：%+v,异常：%s", in.CartItemIds, err.Error())
		return nil, err
	}

	return &omsclient.OrderAddResp{
		Id: orderId,
	}, nil
}

// 2.构建下单商品信息
func buildOrderItem(l *OrderAddLogic, orderInfo *model.OmsOrder, orderItem *omsclient.OrderItemData) {
	err := query.OmsOrderItem.WithContext(l.ctx).Create(&model.OmsOrderItem{
		OrderID:           orderInfo.ID,
		OrderSn:           orderInfo.OrderSn,
		ProductID:         orderItem.ProductId,         // 商品id
		ProductPic:        orderItem.ProductPic,        // 商品图片
		ProductName:       orderItem.ProductName,       // 商品名称
		ProductBrand:      orderItem.ProductBrand,      // 商品品牌
		ProductSn:         orderItem.ProductSn,         // 货号
		ProductPrice:      orderItem.ProductPrice,      // 销售价格
		ProductQuantity:   orderItem.ProductQuantity,   // 购买数量
		ProductSkuID:      orderItem.ProductSkuId,      // 商品sku编号
		ProductSkuCode:    orderItem.ProductSkuCode,    // 商品sku条码
		ProductCategoryID: orderItem.ProductCategoryId, // 商品分类id
		PromotionName:     orderItem.PromotionName,     // 商品促销名称
		PromotionAmount:   orderItem.PromotionAmount,   // 商品促销分解金额
		CouponAmount:      orderItem.CouponAmount,      // 优惠券优惠分解金额
		IntegrationAmount: orderItem.IntegrationAmount, // 积分优惠分解金额
		RealAmount:        orderItem.RealAmount,        // 该商品经过优惠后的分解金额
		GiftIntegration:   orderItem.GiftIntegration,   //
		GiftGrowth:        orderItem.GiftGrowth,        //
		ProductAttr:       orderItem.ProductAttr,       // 商品销售属性:[{"key":"颜色","value":"颜色"},{"key":"容量","value":"4G"}]
	})

	if err != nil {
		logc.Errorf(l.ctx, "插入order_item失败,参数：%+v,异常：%s", orderItem, err.Error())
	}
}

// 1.构建订单信息
func buildOrderInfo(in *omsclient.OrderAddReq) *model.OmsOrder {
	return &model.OmsOrder{
		MemberID:              in.MemberId,              // 会员id
		CouponID:              in.CouponId,              // 优惠券id
		OrderSn:               in.OrderSn,               // 订单编号
		MemberUsername:        in.MemberUsername,        // 用户帐号
		TotalAmount:           in.TotalAmount,           // 订单总金额
		PayAmount:             in.PayAmount,             // 应付金额（实际支付金额）
		FreightAmount:         in.FreightAmount,         // 运费金额
		PromotionAmount:       in.PromotionAmount,       // 促销优化金额（促销价、满减、阶梯价）
		IntegrationAmount:     in.IntegrationAmount,     // 积分抵扣金额
		CouponAmount:          in.CouponAmount,          // 优惠券抵扣金额
		DiscountAmount:        in.DiscountAmount,        // 管理员后台调整订单使用的折扣金额
		PayType:               in.PayType,               // 支付方式：0->未支付；1->支付宝；2->微信
		SourceType:            in.SourceType,            // 订单来源：0->PC订单；1->app订单
		Status:                in.Status,                // 订单状态：0->待付款；1->待发货；2->已发货；3->已完成；4->已关闭；5->无效订单
		OrderType:             in.OrderType,             // 订单类型：0->正常订单；1->秒杀订单
		DeliveryCompany:       in.DeliveryCompany,       // 物流公司(配送方式)
		DeliverySn:            in.DeliverySn,            // 物流单号
		AutoConfirmDay:        in.AutoConfirmDay,        // 自动确认时间（天）
		Integration:           in.Integration,           // 可以获得的积分
		Growth:                in.Growth,                // 可以活动的成长值
		PromotionInfo:         in.PromotionInfo,         // 活动信息
		BillType:              in.BillType,              // 发票类型：0->不开发票；1->电子发票；2->纸质发票
		BillHeader:            in.BillHeader,            // 发票抬头
		BillContent:           in.BillContent,           // 发票内容
		BillReceiverPhone:     in.BillReceiverPhone,     // 收票人电话
		BillReceiverEmail:     in.BillReceiverEmail,     // 收票人邮箱
		ReceiverName:          in.ReceiverName,          // 收货人姓名
		ReceiverPhone:         in.ReceiverPhone,         // 收货人电话
		ReceiverPostCode:      in.ReceiverPostCode,      // 收货人邮编
		ReceiverProvince:      in.ReceiverProvince,      // 省份/直辖市
		ReceiverCity:          in.ReceiverCity,          // 城市
		ReceiverRegion:        in.ReceiverRegion,        // 区
		ReceiverDetailAddress: in.ReceiverDetailAddress, // 详细地址
		Note:                  in.Note,                  // 订单备注
		ConfirmStatus:         in.ConfirmStatus,         // 确认收货状态：0->未确认；1->已确认
		DeleteStatus:          in.DeleteStatus,          // 删除状态：0->未删除；1->已删除
		UseIntegration:        in.UseIntegration,        // 下单时使用的积分
	}
}
