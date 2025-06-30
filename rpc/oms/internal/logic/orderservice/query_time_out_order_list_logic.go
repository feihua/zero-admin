package orderservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/pkg/time_util"
	"github.com/feihua/zero-admin/rpc/oms/gen/model"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/rpc/oms/internal/svc"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryTimeOutOrderListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryTimeOutOrderListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryTimeOutOrderListLogic {
	return &QueryTimeOutOrderListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryTimeOutOrderList 查询超时、未支付的订单及订单详情
func (l *QueryTimeOutOrderListLogic) QueryTimeOutOrderList(in *omsclient.QueryTimeOutOrderListReq) (*omsclient.QueryOrderListResp, error) {
	var result []model.OmsOrderMain
	query := `SELECT o.id,
				   o.order_no,
				   o.user_id,
				   o.use_points,
				   ot.id           as ot_id,
				   ot.sku_name     as sku_name,
				   ot.sku_id       as sku_id,
				   ot.sku_quantity as sku_quantity
			FROM oms_order_main o
					 LEFT JOIN oms_order_item ot ON o.id = ot.order_id
			WHERE o.order_status = 0
            AND o.create_time < date_add(NOW(), INTERVAL -? MINUTE)`
	db := l.svcCtx.DB
	err := db.Where(l.ctx).Raw(query, in.Minute).Find(&result).Error

	if err != nil {
		logc.Errorf(l.ctx, " 查询超时、未支付的订单及订单详情失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New(" 查询超时、未支付的订单及订单详情失败")
	}

	var list []*omsclient.OrderListData
	for _, item := range result {

		list = append(list, &omsclient.OrderListData{
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
		})
	}

	return &omsclient.QueryOrderListResp{
		Total: 0,
		List:  list,
	}, nil
}
