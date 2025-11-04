package orderservicelogic

import (
	"context"
	"errors"

	"github.com/feihua/zero-admin/pkg/time_util"
	"github.com/feihua/zero-admin/rpc/oms/gen/query"
	"github.com/zeromicro/go-zero/core/logc"
	"gorm.io/gorm"

	"github.com/feihua/zero-admin/rpc/oms/internal/svc"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryOrderDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryOrderDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryOrderDetailLogic {
	return &QueryOrderDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryOrderDetail 查询订单详情:订单信息、商品信息、操作记录
func (l *QueryOrderDetailLogic) QueryOrderDetail(in *omsclient.QueryOrderDetailReq) (*omsclient.QueryOrderDetailResp, error) {
	orderMain := query.OmsOrderMain
	q := orderMain.WithContext(l.ctx)
	q = q.Where(orderMain.ID.Eq(in.Id), orderMain.IsDeleted.Eq(0))

	if in.UserId != 0 {
		q = q.Where(orderMain.UserID.Eq(in.UserId))
	}

	item, err := q.First()

	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		logc.Errorf(l.ctx, "订单不存在, 请求参数：%+v, 异常信息: %s", in, err.Error())
		return nil, errors.New("订单不存在")
	case err != nil:
		logc.Errorf(l.ctx, "查询订单异常, 请求参数：%+v, 异常信息: %s", in, err.Error())
		return nil, errors.New("查询订单异常")
	}

	data := &omsclient.OrderListData{
		Id:                 item.ID,                                   //
		OrderNo:            item.OrderNo,                              // 订单编号
		UserId:             item.UserID,                               // 用户ID
		OrderStatus:        item.OrderStatus,                          // 订单状态：1-待支付,2-已支付,3-已发货,4-已完成,5-已取消,6-已退款,7-售后中
		TotalAmount:        float32(item.TotalAmount),                 // 订单总金额
		PromotionAmount:    float32(item.PromotionAmount),             // 促销金额
		CouponAmount:       float32(item.CouponAmount),                // 优惠券金额
		PointsAmount:       float32(item.PointsAmount),                // 积分金额
		DiscountAmount:     float32(item.DiscountAmount),              // 优惠金额
		FreightAmount:      float32(item.FreightAmount),               // 运费金额
		PayAmount:          float32(item.PayAmount),                   // 实付金额
		PayTime:            time_util.TimeToString(item.PayTime),      // 支付时间
		DeliveryTime:       time_util.TimeToString(item.DeliveryTime), // 发货时间
		ReceiveTime:        time_util.TimeToString(item.ReceiveTime),  // 收货时间
		CommentTime:        time_util.TimeToString(item.CommentTime),  // 评价时间
		SourceType:         item.SourceType,                           // 订单来源：1-APP,2-PC,3-小程序
		ExpressOrderNumber: item.ExpressOrderNumber,                   // 快递单号
		UsePoints:          item.UsePoints,                            // 下单时使用的积分
		ReceiveStatus:      item.ReceiveStatus,                        // 是否确认收货：0->否,1->是
		Remark:             item.Remark,                               // 订单备注
		CreateTime:         time_util.TimeToStr(item.CreateTime),      // 提交时间
		UpdateTime:         time_util.TimeToString(item.UpdateTime),   //
	}

	if item.PayType != nil {
		data.PayType = *item.PayType // 支付方式：1-支付宝,2-微信,3-银联
	}

	orderItem := query.OmsOrderItem
	result, err := orderItem.WithContext(l.ctx).Where(orderItem.OrderID.Eq(in.Id)).Find()

	if err != nil {
		logc.Errorf(l.ctx, "查询订单商品列表失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("查询订单商品列表失败")
	}

	var list []*omsclient.OrderItemData

	for _, x := range result {
		list = append(list, &omsclient.OrderItemData{
			Id:              x.ID,                       //
			OrderId:         x.OrderID,                  // 订单ID
			OrderNo:         x.OrderNo,                  // 订单编号
			OrderItemStatus: x.OrderItemStatus,          // 订单商品状态：1-正常,2-退货申请中,3-已退货,4-已拒绝
			SkuId:           x.SkuID,                    // 商品SKU ID
			SkuName:         x.SkuName,                  // 商品名称
			SkuPic:          x.SkuPic,                   // 商品图片
			SkuPrice:        float32(x.SkuPrice),        // 商品单价
			SkuQuantity:     x.SkuQuantity,              // 商品数量
			SpecData:        x.SpecData,                 // 规格数据
			SkuTotalAmount:  float32(x.SkuTotalAmount),  // 商品总金额
			PromotionAmount: float32(x.PromotionAmount), // 促销分摊金额
			CouponAmount:    float32(x.CouponAmount),    // 优惠券分摊金额
			PointsAmount:    float32(x.PointsAmount),    // 积分分摊金额
			DiscountAmount:  float32(x.DiscountAmount),  // 优惠分摊金额
			RealAmount:      float32(x.RealAmount),      // 实付金额

		})
	}

	data.OrderItemData = list

	orderOperationLog := query.OmsOrderOperationLog
	result1, err := orderOperationLog.WithContext(l.ctx).Where(orderOperationLog.OrderID.Eq(in.Id)).Find()

	if err != nil {
		logc.Errorf(l.ctx, "查询订单操作记录列表失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("查询订单操作记录列表失败")
	}

	var list1 []*omsclient.OrderOptLogData

	for _, x := range result1 {
		list1 = append(list1, &omsclient.OrderOptLogData{
			Id:            x.ID,                              // 键ID
			OrderId:       x.OrderID,                         // 订单ID
			OrderNo:       x.OrderNo,                         // 订单编号
			OperationType: x.OperationType,                   // 操作类型：1-创建订单，2-支付订单，3-发货，4-确认收货，5-取消订单，6-退款
			OperatorId:    x.OperatorID,                      // 操作人ID
			OperatorType:  x.OperatorType,                    // 操作人类型：1-用户，2-系统，3-管理员
			OperatorNote:  x.OperatorNote,                    // 操作备注
			CreateTime:    time_util.TimeToStr(x.CreateTime), // 操作时间

		})
	}

	data.OptLogData = list1

	orderPromotion := query.OmsOrderPromotion
	result2, err := orderPromotion.WithContext(l.ctx).Where(orderPromotion.OrderID.Eq(in.Id)).Find()

	if err != nil {
		logc.Errorf(l.ctx, "查询订单优惠信息列表失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("查询订单优惠信息列表失败")
	}

	var list2 []*omsclient.PromotionListData

	for _, x := range result2 {
		y := omsclient.PromotionListData{
			Id:             x.ID,                              // 主键ID
			OrderId:        x.OrderID,                         // 订单ID
			OrderNo:        x.OrderNo,                         // 订单编号
			PromotionType:  x.PromotionType,                   // 优惠类型：1-优惠券，2-积分抵扣，3-会员折扣，4-促销活动
			PromotionName:  x.PromotionName,                   // 优惠名称
			DiscountAmount: float32(x.DiscountAmount),         // 优惠金额
			CreateTime:     time_util.TimeToStr(x.CreateTime), //

		}
		if x.PromotionID != nil {
			y.PromotionId = *x.PromotionID
		}
		list2 = append(list2, &y)
	}

	data.PromotionData = list2

	delivery, err := query.OmsOrderDelivery.WithContext(l.ctx).Where(query.OmsOrderDelivery.OrderID.Eq(item.ID)).First()

	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		logc.Errorf(l.ctx, "订单收货地址不存在, 请求参数：%+v, 异常信息: %s", in, err.Error())
		return nil, errors.New("订单收货地址不存在")
	case err != nil:
		logc.Errorf(l.ctx, "查询订单收货地址异常, 请求参数：%+v, 异常信息: %s", in, err.Error())
		return nil, errors.New("查询订单收货地址异常")
	}

	x := &omsclient.DeliveryData{
		Id:               delivery.ID,                                 //
		OrderId:          delivery.OrderID,                            // 订单ID
		OrderNo:          delivery.OrderNo,                            // 订单编号
		ReceiverName:     delivery.ReceiverName,                       // 收货人姓名
		ReceiverPhone:    delivery.ReceiverPhone,                      // 收货人电话
		ReceiverProvince: delivery.ReceiverProvince,                   // 省份
		ReceiverCity:     delivery.ReceiverCity,                       // 城市
		ReceiverDistrict: delivery.ReceiverDistrict,                   // 区县
		ReceiverAddress:  delivery.ReceiverAddress,                    // 详细地址
		DeliveryCompany:  delivery.DeliveryCompany,                    // 物流公司
		DeliveryNo:       delivery.DeliveryNo,                         // 物流单号
		CreateTime:       time_util.TimeToStr(delivery.CreateTime),    // 创建时间
		UpdateTime:       time_util.TimeToString(delivery.UpdateTime), // 更新时间
	}

	data.DeliveryData = x

	orderPayment := query.OmsOrderPayment
	payment, err := orderPayment.WithContext(l.ctx).Where(orderPayment.OrderID.Eq(in.Id)).Find()

	if err != nil {
		logc.Errorf(l.ctx, "查询订单支付列表失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("查询订单支付列表失败")
	}

	var list3 []*omsclient.PaymentData

	for _, s := range payment {
		list3 = append(list3, &omsclient.PaymentData{
			Id:            s.ID,                                 // 主键ID
			OrderId:       s.OrderID,                            // 订单ID
			OrderNo:       s.OrderNo,                            // 订单编号
			PayType:       s.PayType,                            // 支付方式：1-支付宝，2-微信，3-银联
			TransactionId: s.TransactionID,                      // 支付流水号
			TotalAmount:   float32(s.TotalAmount),               // 订单金额
			PayAmount:     float32(s.PayAmount),                 // 支付金额
			PayStatus:     s.PayStatus,                          // 支付状态：0-待支付，1-支付成功，2-支付失败
			PayTime:       time_util.TimeToString(s.PayTime),    // 支付时间
			CreateTime:    time_util.TimeToStr(s.CreateTime),    // 创建时间
			UpdateTime:    time_util.TimeToString(s.UpdateTime), // 更新时间

		})
	}

	data.PaymentData = list3
	return &omsclient.QueryOrderDetailResp{
		Data: data,
	}, nil
}
