package orderservicelogic

import (
	"context"
	"fmt"
	"github.com/bytedance/sonic"
	"github.com/feihua/zero-admin/rpc/oms/gen/model"
	"github.com/feihua/zero-admin/rpc/oms/gen/query"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/rpc/oms/internal/svc"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddOrderLogic {
	return &AddOrderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// AddOrder 添加订单(app)
func (l *AddOrderLogic) AddOrder(in *omsclient.AddOrderReq) (*omsclient.AddOrderResp, error) {
	q := query.OmsOrderMain

	item := &model.OmsOrderMain{
		OrderNo:         in.OrderNo,                  // 订单编号
		UserID:          in.UserId,                   // 用户ID
		OrderStatus:     in.OrderStatus,              // 订单状态：1-待支付,2-已支付,3-已发货,4-已完成,5-已取消,6-已退款,7-售后中
		TotalAmount:     float64(in.TotalAmount),     // 订单总金额
		PromotionAmount: float64(in.PromotionAmount), // 促销金额
		CouponAmount:    float64(in.CouponAmount),    // 优惠券金额
		PointsAmount:    float64(in.PointsAmount),    // 积分金额
		DiscountAmount:  float64(in.DiscountAmount),  // 优惠金额
		FreightAmount:   float64(in.FreightAmount),   // 运费金额
		PayAmount:       float64(in.PayAmount),       // 实付金额
		SourceType:      in.SourceType,               // 订单来源：1-APP,2-PC,3-小程序
		UsePoints:       in.UsePoints,                // 下单时使用的积分
		Remark:          in.Remark,                   // 订单备注
	}

	err := q.WithContext(l.ctx).Create(item)
	if err != nil {
		logc.Errorf(l.ctx, "添加订单失败,参数:%+v,异常:%s", item, err.Error())
		return nil, fmt.Errorf("添加订单失败")
	}

	var orderItems []*model.OmsOrderItem
	var ids []int64
	for _, data := range in.OrderItemData {
		item1 := &model.OmsOrderItem{
			OrderID:         item.ID,                       // 订单ID
			OrderNo:         data.OrderNo,                  // 订单编号
			OrderItemStatus: 1,                             // 订单商品状态：1-正常,2-退货申请中,3-已退货,4-已拒绝
			SkuID:           data.SkuId,                    // 商品SKU ID
			SkuName:         data.SkuName,                  // 商品名称
			SkuPic:          data.SkuPic,                   // 商品图片
			SkuPrice:        float64(data.SkuPrice),        // 商品单价
			SkuQuantity:     data.SkuQuantity,              // 商品数量
			SpecData:        data.SpecData,                 // 规格数据
			SkuTotalAmount:  float64(data.SkuTotalAmount),  // 商品总金额
			PromotionAmount: float64(data.PromotionAmount), // 促销分摊金额
			CouponAmount:    float64(data.CouponAmount),    // 优惠券分摊金额
			PointsAmount:    float64(data.PointsAmount),    // 积分分摊金额
			DiscountAmount:  float64(data.DiscountAmount),  // 优惠分摊金额
			RealAmount:      float64(data.RealAmount),      // 实付金额
		}

		orderItems = append(orderItems, item1)
		ids = append(ids, data.SkuId)

	}
	err = query.OmsOrderItem.WithContext(l.ctx).CreateInBatches(orderItems, len(orderItems))
	if err != nil {
		logc.Errorf(l.ctx, "添加订单商品失败,参数:%+v,异常:%s", orderItems, err.Error())
		return nil, fmt.Errorf("添加订单商品失败")
	}

	_, err = query.OmsCartItem.WithContext(l.ctx).Where(q.ID.In(ids...)).Delete()

	if err != nil {
		logc.Errorf(l.ctx, "删除购物车失败,参数:%+v,异常:%s", in, err.Error())
		return nil, fmt.Errorf("删除购物车失败")
	}

	message := map[string]any{"id": item.ID}
	body, _ := sonic.Marshal(message)
	err = l.svcCtx.RabbitMQ.SendMessage("order.event.exchange", "order.create.queue", "order.create.key", body)

	return &omsclient.AddOrderResp{}, nil
}
