package order

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/feihua/zero-admin/api/front/internal/logic/common"
	"github.com/feihua/zero-admin/api/front/internal/logic/order/cart"
	"github.com/feihua/zero-admin/pkg/errorx"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc/status"

	"github.com/feihua/zero-admin/api/front/internal/svc"
	"github.com/feihua/zero-admin/api/front/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// GenerateOrderLogic
/*
Author: LiuFeiHua
Date: 2023/12/12 18:04
*/
type GenerateOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGenerateOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GenerateOrderLogic {
	return &GenerateOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// GenerateOrder 根据提交信息生成订单
// 1.获取购物车及优惠信息
// 2.生成下单商品信息
// 3.判断购物车中商品是否都有库存
// 4.判断是否使用了优惠券
// 5.判断是否使用积分
// 6.计算order_item的实付金额
// 7.进行库存锁定
// 8.根据商品合计、运费、活动优惠、优惠券、积分计算应付金额
// 9.转化为订单信息并插入数据库(删除购物车中的下单商品)
// 10.如果使用优惠券,更新优惠券使用状态
// 11.如果使用积分,需要扣除积分
// 12.发送延迟消息取消订单
func (l *GenerateOrderLogic) GenerateOrder(req *types.GenerateOrderReq) (*types.GenerateOrderResp, error) {
	memberId, err := common.GetMemberId(l.ctx)
	if err != nil {
		return nil, err
	}
	memberInfo, _ := l.svcCtx.MemberService.QueryMemberInfoDetail(l.ctx, &umsclient.QueryMemberInfoDetailReq{MemberId: memberId})
	// 1.获取购物车及优惠信息
	cartPromotionItemList, err := cart.QueryCartListPromotion(req.CartIds, l.ctx, l.svcCtx)
	if err != nil {
		return nil, err
	}
	if len(cartPromotionItemList) == 0 {
		return result(1, "购物车还没有商品,请先添加商品到购物车!"), nil
	}
	// 2.生成下单商品信息
	var flag = false // 用于判断库存
	// var totalAmount float32 = 0
	orderItemList := make([]*omsclient.OrderItemData, 0)
	cartItemIds := make([]int64, 0) // 方便提交订单成功,后删除购物车中的商品
	for _, item := range cartPromotionItemList {
		orderItemList = append(orderItemList, &omsclient.OrderItemData{
			// OrderItemStatus: 1, //订单商品状态：1-正常,2-退货申请中,3-已退货,4-已拒绝
			// SkuId: item.ProductSkuId, //商品SKU ID
			// SkuName: item.ProductName, //商品名称
			// SkuPic: item.ProductPic, //商品图片
			// SkuPrice: item.Price, //商品单价
			// SkuQuantity: item.Quantity, //商品数量
			// SkuTotalAmount: item.ReduceAmount, //商品总金额
			// PromotionAmount: item.PromotionAmount, //促销分摊金额
			// CouponAmount: item.CouponAmount, //优惠券分摊金额
			// PointsAmount: item.PointsAmount, //积分分摊金额
			// DiscountAmount: item.DiscountAmount, //优惠分摊金额
			// RealAmount: item.RealAmount, //实付金额
		})

		// totalAmount = totalAmount + item.Price
		if item.RealStock <= 0 || item.RealStock < item.Quantity {
			flag = true
		}

		cartItemIds = append(cartItemIds, item.Id)
	}
	// 3.判断购物车中商品是否都有库存
	if flag {
		return result(1, "库存不足，无法下单"), nil
	}
	// 4.判断是否使用了优惠券
	// 因为nt64传参数默认为0,所以判断为0的时候,表示没有使用优惠券
	if req.CouponId > 0 {
		// var couponHistoryDetailListData *smsclient.QueryCouponRecordDetailResp
		// enableList, _ := coupon.QueryCouponList(l.svcCtx, l.ctx, cartPromotionItemList)
		// for _, item := range enableList {
		// 	if item.CouponListData.Id == req.CouponId {
		// 		couponHistoryDetailListData = item
		// 	}
		// }
		// var couponData = couponHistoryDetailListData.CouponListData
		//
		// if couponData == nil {
		// 	return result(1, "该优惠券不可用"), nil
		// }
		// useType := couponData.UseType
		// if useType == 0 {
		// 	// 全场通用
		// 	// (商品价格/可用商品总价)*优惠券面额
		// 	for _, item := range orderItemList {
		// 		item.CouponAmount = (item.ProductPrice / totalAmount) * couponData.Amount
		// 	}
		// } else if useType == 1 {
		// 	// 指定分类
		// 	categoryIdList := make(map[int64]int64, 0)
		// 	for _, item := range couponHistoryDetailListData.CategoryRelationList {
		// 		categoryIdList[item.ProductCategoryId] = item.ProductCategoryId
		// 	}
		// 	orderItemListByCategoryId := make([]*omsclient.OrderItemData, 0)
		// 	var totalAmountByCategoryId int64
		// 	for _, item := range orderItemList {
		// 		_, ok := categoryIdList[item.ProductCategoryId]
		// 		if ok {
		// 			totalAmountByCategoryId = totalAmountByCategoryId + item.ProductPrice
		// 			orderItemListByCategoryId = append(orderItemListByCategoryId, item)
		// 		}
		// 	}
		// 	for _, item := range orderItemListByCategoryId {
		// 		item.CouponAmount = (item.ProductPrice / totalAmountByCategoryId) * couponData.Amount
		// 	}
		// } else if useType == 2 {
		// 	// 指定商品
		// 	productIdList := make(map[int64]int64)
		// 	for _, item := range couponHistoryDetailListData.ProductRelationList {
		// 		productIdList[item.ProductId] = item.ProductId
		// 	}
		// 	orderItemListByProductId := make([]*omsclient.OrderItemData, 0)
		// 	var totalAmountByProductId int64
		// 	for _, item := range orderItemList {
		// 		_, ok := productIdList[item.ProductCategoryId]
		// 		if ok {
		// 			orderItemListByProductId = append(orderItemListByProductId, item)
		// 		}
		// 	}
		//
		// 	for _, item := range orderItemListByProductId {
		// 		item.CouponAmount = (item.ProductPrice / totalAmountByProductId) * couponData.Amount
		// 	}
		// }

	}
	// 5.判断是否使用积分
	if req.UseIntegration > 0 {
		// 1.先计算出积分的总金额
		// 1.1判断用户是否有这么多积分
		if req.UseIntegration > memberInfo.Points {
			return result(1, "积分大于用户拥有有积分"), nil
		}
		// 1.2根据积分使用规则判断是否可用
		// 是否可与优惠券共用
		consumeSetting, _ := l.svcCtx.MemberConsumeSettingService.QueryMemberConsumeSettingDetail(l.ctx, &umsclient.QueryMemberConsumeSettingDetailReq{Id: 1})
		if req.CouponId > 0 && consumeSetting.CouponStatus == 0 {
			return result(1, "不可与优惠券共用"), nil
		}
		// 1.3是否达到最低使用积分门槛
		if req.UseIntegration < consumeSetting.UseUnit {
			return result(1, "未达到最低使用积分门槛"), nil
		}
		// 1.4是否超过订单抵用最高百分比
		// integrationAmount := int64(req.UseIntegration / consumeSetting.UseUnit)
		// var maxPercent = int64(consumeSetting.MaxPercentPerOrder / 100)
		// if integrationAmount > maxPercent*totalAmount {
		// 	return result(1, "超过订单抵用最高百分比"), nil
		// }

		// 2.可用情况下分摊到可用商品中
		// for _, item := range orderItemList {
		// 	item.PointsAmount = (item.SkuPrice / totalAmount) * integrationAmount
		// }
	}
	// 6.计算order_item的实付金额
	// 原价-促销优惠-优惠券抵扣-积分抵扣
	for _, item := range orderItemList {
		item.RealAmount = item.SkuPrice - item.PromotionAmount - item.CouponAmount - item.PointsAmount
	}
	// 7.进行库存锁定
	var data []*pmsclient.UpdateSkuStockData
	for _, item := range orderItemList {
		data = append(data, &pmsclient.UpdateSkuStockData{
			Id:              item.SkuId,
			ProductQuantity: item.SkuQuantity,
		})
	}
	_, err = l.svcCtx.ProductSkuService.LockSkuStockLock(l.ctx, &pmsclient.UpdateSkuStockReq{
		Data: data,
	})

	if err != nil {
		logc.Errorf(l.ctx, "锁定库存异常,参数: %+v,异常：%s", data, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}
	// 8.根据商品合计、运费、活动优惠、优惠券、积分计算应付金额
	// 计算订单优惠(促销价、满减、阶梯价)
	// var promotionAmount int64
	// // 获取订单促销信息
	// var promotionInfo string
	// // 计算订单优惠券抵扣金额
	// var couponAmount int64
	// // 计算积分抵扣金额
	// var integrationAmount int64
	// // 计算该订单赠送的积分
	// var giftIntegration int32
	// // 计算赠送成长值
	// var giftGrowth int32
	// for _, item := range orderItemList {
	// promotionAmount = promotionAmount + item.PromotionAmount*int64(item.SkuQuantity)
	// promotionInfo = promotionInfo
	// couponAmount = couponAmount + item.CouponAmount*int64(item.SkuQuantity)
	// integrationAmount = integrationAmount + item.IntegrationAmount*int64(item.SkuQuantity)
	// giftIntegration = giftIntegration + item.GiftIntegration*item.SkuQuantity
	// giftGrowth = giftGrowth + item.GiftGrowth*item.SkuQuantity
	// }
	// 计算订单应付金额
	// payAmount := totalAmount - promotionAmount - couponAmount - integrationAmount
	// address, err := l.svcCtx.MemberAddressService.QueryMemberAddressDetail(l.ctx, &umsclient.QueryMemberAddressDetailReq{
	// 	MemberId: memberId,
	// 	Id:       req.MemberReceiveAddressId,
	// })
	// if err != nil {
	// 	return result(1, "查询会员地址异常"), nil
	// }
	// 设置自动收货天数
	// orderSettingList, err := l.svcCtx.OrderSettingService.QueryOrderSettingList(l.ctx, &omsclient.QueryOrderSettingListReq{
	// 	PageNum:  1,
	// 	PageSize: 10,
	// })
	// if err != nil {
	// 	return result(1, "查询订单设置异常"), nil
	// }
	orderInfo := &omsclient.AddOrderReq{
		// UserId: req.UserId, //用户ID
		// OrderStatus: req.OrderStatus, //订单状态：1-待支付,2-已支付,3-已发货,4-已完成,5-已取消,6-已退款,7-售后中
		// TotalAmount: req.TotalAmount, //订单总金额
		// PromotionAmount: req.PromotionAmount, //促销金额
		// CouponAmount: req.CouponAmount, //优惠券金额
		// PointsAmount: req.PointsAmount, //积分金额
		// DiscountAmount: req.DiscountAmount, //优惠金额
		// FreightAmount: req.FreightAmount, //运费金额
		// PayAmount: req.PayAmount, //实付金额
		// PayType: req.PayType, //支付方式：1-支付宝,2-微信,3-银联
		// PayTime: req.PayTime, //支付时间
		// DeliveryTime: req.DeliveryTime, //发货时间
		// ReceiveTime: req.ReceiveTime, //收货时间
		// CommentTime: req.CommentTime, //评价时间
		// SourceType: req.SourceType, //订单来源：1-APP,2-PC,3-小程序
		// ExpressOrderNumber: req.ExpressOrderNumber, //快递单号
		// UsePoints: req.UsePoints, //下单时使用的积分
		// ReceiveStatus: req.ReceiveStatus, //是否确认收货：0->否,1->是
		// Remark: req.Remark, //订单备注
		// OrderItemList:         orderItemList,                            //
	}
	// 9.转化为订单信息并插入数据库(插入order表和order_item表)
	orderAddResp, err := l.svcCtx.OrderService.AddOrder(l.ctx, orderInfo)
	if err != nil {
		return result(1, "提交订单异常"), nil
	}
	// 10.如果使用优惠券,更新优惠券使用状态
	// 10.1修改sms_coupon_history表的use_status字段
	// 10.2记得修改sms_coupon的use_count字段,下单的时候要加1,取消订单的时候,要减1
	// if couponAmount > 0 {
	// 	_, err = l.svcCtx.CouponRecordService.UpdateCouponRecord(l.ctx, &smsclient.UpdateCouponRecordReq{
	// 		MemberId: memberId,
	// 		Status:   1,
	// 		// CouponIds: req.CouponId,
	// 	})
	// 	if err != nil {
	// 		return result(1, "更新优惠券使用状态异常"), nil
	// 	}
	// }
	// 11.如果使用积分,需要扣除积分
	i := memberInfo.Points - req.UseIntegration
	_, err = l.svcCtx.MemberService.UpdateMemberPoints(l.ctx, &umsclient.UpdateMemberPointsReq{MemberId: memberId, Points: i})
	if err != nil {
		return result(1, "扣除积分异常"), nil
	}

	orderId := orderAddResp.Id
	// 12.发送延迟消息取消订单
	err = l.sendMsg(orderId, memberId)
	if err != nil {
		return nil, err
	}
	return &types.GenerateOrderResp{
		Code:    0,
		Message: "下单成功",
		Data: types.GenerateOrderData{
			Id:       orderId,
			MemberId: memberId,
			// MemberUsername:    memberInfo.Nickname,
			// TotalAmount:       orderInfo.TotalAmount,
			// PayAmount:         orderInfo.PayAmount,
			// FreightAmount:     orderInfo.FreightAmount,
			// PromotionAmount:   orderInfo.PromotionAmount,
			// IntegrationAmount: orderInfo.IntegrationAmount,
			// CouponAmount:      orderInfo.CouponAmount,
			// DiscountAmount:    orderInfo.DiscountAmount,
			// PayType:           orderInfo.PayType,
			// SourceType:        orderInfo.SourceType,
			// Status:            orderInfo.Status,
			// OrderType:         orderInfo.OrderType,
			// Integration:       orderInfo.Integration,
			// Growth:            orderInfo.Growth,
			// PromotionInfo:     orderInfo.PromotionInfo,
			// Note:              orderInfo.Note,
			// UseIntegration:    orderInfo.UseIntegration,
		},
	}, nil
}

// 发送延迟消息取消订单
func (l *GenerateOrderLogic) sendMsg(orderId, memberId int64) error {
	delayMinutes := 30 // 延迟时间(分钟)
	message := map[string]any{"orderId": orderId, "memberId": memberId}
	body, err := json.Marshal(message)
	if err != nil {
		logc.Errorf(l.ctx, "序列化 JSON 失败: %v", err)
		return errorx.NewDefaultError("序列化 JSON 错误")
	}
	err = l.svcCtx.RabbitMQ.SendDelayMessage("order.delay.exchange", "order.cancel.queue", "order.cancel", body, delayMinutes)

	if err != nil {
		logc.Errorf(l.ctx, "订单 %d 延时取消,消息发送失败,,异常:%s", orderId, err.Error())
		return errorx.NewDefaultError(fmt.Sprintf("订单 %d 延时取消,消息发送失败,,异常:%s", orderId, err.Error()))
	}

	logx.Errorf("订单 %d 延时取消,消息已发送，将在 %d 分钟后处理", orderId, delayMinutes)
	return nil
}

func result(code int64, message string) *types.GenerateOrderResp {
	return &types.GenerateOrderResp{
		Code:    code,
		Message: message,
	}
}
