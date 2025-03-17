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
		Id:                    item.Id,                    // 订单id
		MemberId:              item.MemberId,              // 会员id
		CouponId:              item.CouponId,              // 优惠券id
		OrderSn:               item.OrderSn,               // 订单编号
		CreateTime:            item.CreateTime,            // 提交时间
		MemberUsername:        item.MemberUsername,        // 用户帐号
		TotalAmount:           item.TotalAmount,           // 订单总金额
		PayAmount:             item.PayAmount,             // 应付金额（实际支付金额）
		FreightAmount:         item.FreightAmount,         // 运费金额
		PromotionAmount:       item.PromotionAmount,       // 促销优化金额（促销价、满减、阶梯价）
		IntegrationAmount:     item.IntegrationAmount,     // 积分抵扣金额
		CouponAmount:          item.CouponAmount,          // 优惠券抵扣金额
		DiscountAmount:        item.DiscountAmount,        // 管理员后台调整订单使用的折扣金额
		PayType:               item.PayType,               // 支付方式：0->未支付；1->支付宝；2->微信
		SourceType:            item.SourceType,            // 订单来源：0->PC订单；1->app订单
		Status:                item.Status,                // 订单状态：0->待付款；1->待发货；2->已发货；3->已完成；4->已关闭；5->无效订单
		OrderType:             item.OrderType,             // 订单类型：0->正常订单；1->秒杀订单
		DeliveryCompany:       item.DeliveryCompany,       // 物流公司(配送方式)
		DeliverySn:            item.DeliverySn,            // 物流单号
		AutoConfirmDay:        item.AutoConfirmDay,        // 自动确认时间（天）
		Integration:           item.Integration,           // 可以获得的积分
		Growth:                item.Growth,                // 可以活动的成长值
		PromotionInfo:         item.PromotionInfo,         // 活动信息
		BillType:              item.BillType,              // 发票类型：0->不开发票；1->电子发票；2->纸质发票
		BillHeader:            item.BillHeader,            // 发票抬头
		BillContent:           item.BillContent,           // 发票内容
		BillReceiverPhone:     item.BillReceiverPhone,     // 收票人电话
		BillReceiverEmail:     item.BillReceiverEmail,     // 收票人邮箱
		ReceiverName:          item.ReceiverName,          // 收货人姓名
		ReceiverPhone:         item.ReceiverPhone,         // 收货人电话
		ReceiverPostCode:      item.ReceiverPostCode,      // 收货人邮编
		ReceiverProvince:      item.ReceiverProvince,      // 省份/直辖市
		ReceiverCity:          item.ReceiverCity,          // 城市
		ReceiverRegion:        item.ReceiverRegion,        // 区
		ReceiverDetailAddress: item.ReceiverDetailAddress, // 详细地址
		Note:                  item.Note,                  // 订单备注
		ConfirmStatus:         item.ConfirmStatus,         // 确认收货状态：0->未确认；1->已确认
		DeleteStatus:          item.DeleteStatus,          // 删除状态：0->未删除；1->已删除
		UseIntegration:        item.UseIntegration,        // 下单时使用的积分
		PaymentTime:           item.PaymentTime,           // 支付时间
		DeliveryTime:          item.DeliveryTime,          // 发货时间
		ReceiveTime:           item.ReceiveTime,           // 确认收货时间
		CommentTime:           item.CommentTime,           // 评价时间
		ModifyTime:            item.ModifyTime,            // 修改时间
		ListOperateHistory:    queryHistoryRecord(item),
		ListOrderItem:         queryOrderItems(item),
	}

	return &types.OrderDetailResp{
		Data:    data,
		Code:    "000000",
		Message: "获取订单详情成功",
	}, nil
}
