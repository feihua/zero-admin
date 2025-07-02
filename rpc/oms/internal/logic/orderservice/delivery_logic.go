package orderservicelogic

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/feihua/zero-admin/rpc/oms/gen/model"
	"github.com/feihua/zero-admin/rpc/oms/gen/query"
	"github.com/zeromicro/go-zero/core/logc"
	"gorm.io/gorm"

	"github.com/feihua/zero-admin/rpc/oms/internal/svc"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeliveryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeliveryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeliveryLogic {
	return &DeliveryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// Delivery 订单发货
func (l *DeliveryLogic) Delivery(in *omsclient.DeliveryReq) (*omsclient.DeliveryResp, error) {
	q := query.OmsOrderMain.WithContext(l.ctx)

	item, err := q.Where(query.OmsOrderMain.ID.Eq(in.OrderId)).First()

	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		logc.Errorf(l.ctx, "订单不存在, 请求参数：%+v, 异常信息: %s", in, err.Error())
		return nil, errors.New("订单不存在")
	case err != nil:
		logc.Errorf(l.ctx, "查询订单异常, 请求参数：%+v, 异常信息: %s", in, err.Error())
		return nil, errors.New("查询订单异常")
	}

	item.OrderStatus = 3
	item.ExpressOrderNumber = in.DeliverySn
	_, err = q.Where(query.OmsOrderMain.ID.Eq(in.OrderId)).Updates(item)

	if err != nil {
		logc.Errorf(l.ctx, "更新订单失败,参数:%+v,异常:%s", item, err.Error())
		return nil, errors.New("更新订单失败")
	}

	log := model.OmsOrderOperationLog{
		OrderID:       in.OrderId,    // 订单ID
		OrderNo:       in.OrderNo,    // 订单编号
		OperationType: 3,             // 操作类型：1-创建订单，2-支付订单，3-发货，4-确认收货，5-取消订单，6-退款
		OperatorID:    in.OperatorId, // 操作人ID
		OperatorType:  3,             // 操作人类型：1-用户，2-系统，3-管理员
		OperatorNote:  "完成发货",        // 操作备注
	}

	err = query.OmsOrderOperationLog.WithContext(l.ctx).Create(&log)
	if err != nil {
		logc.Errorf(l.ctx, "添加订单操作记录失败,参数:%+v,异常:%s", item, err.Error())
		return nil, errors.New("添加订单操作记录失败")
	}

	message := map[string]any{"ids": in.OrderId}
	body, _ := json.Marshal(message)
	err = l.svcCtx.RabbitMQ.SendMessage("order.event.exchange", "order.delivery.queue", "order.delivery.key", body)

	return &omsclient.DeliveryResp{}, nil
}
