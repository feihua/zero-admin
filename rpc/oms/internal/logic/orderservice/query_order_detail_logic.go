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
	item, err := query.OmsOrderMain.WithContext(l.ctx).Where(query.OmsOrderMain.ID.Eq(in.Id), query.OmsOrderMain.UserID.Eq(in.UserId), query.OmsOrderMain.IsDeleted.Eq(0)).First()

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
		PayType:            *item.PayType,                             // 支付方式：1-支付宝,2-微信,3-银联
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

	orderItem := query.OmsOrderItem
	q := orderItem.WithContext(l.ctx).Where(orderItem.OrderID.Eq(in.Id))

	result, err := q.Find()

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
	return &omsclient.QueryOrderDetailResp{
		Data: data,
	}, nil
}
