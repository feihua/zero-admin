package order

import (
	"context"
	"encoding/json"

	"github.com/feihua/zero-admin/consumer/internal/svc"
	"github.com/feihua/zero-admin/consumer/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type TestOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTestOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TestOrderLogic {
	return &TestOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// TestOrder 测试订单相关mq
func (l *TestOrderLogic) TestOrder(req *types.TestOrderReq) (resp *types.Response, err error) {

	// 测试取消订单
	message := map[string]any{"id": req.Id}
	body, _ := json.Marshal(message)
	err = l.svcCtx.RabbitMQ.SendMessage("order.cancel.exchange", "order.cancel.by.user.queue", "order.cancel.key", body)

	// 测试关闭订单
	message = map[string]any{"ids": req.Id}
	body, _ = json.Marshal(message)
	err = l.svcCtx.RabbitMQ.SendMessage("order.close.exchange", "order.close.queue", "order.close.key", body)

	// 测试订单发货
	message = map[string]any{"ids": req.Id}
	body, _ = json.Marshal(message)
	err = l.svcCtx.RabbitMQ.SendMessage("order.delivery.exchange", "order.delivery.queue", "order.delivery.key", body)

	// 测试申请退货
	message = map[string]any{"id": req.Id}
	body, _ = json.Marshal(message)
	err = l.svcCtx.RabbitMQ.SendMessage("order.return.exchange", "order.return.queue", "order.return.key", body)

	// 测试用户确认订单
	message = map[string]any{"id": req.Id}
	body, _ = json.Marshal(message)
	err = l.svcCtx.RabbitMQ.SendMessage("order.confirm.exchange", "order.confirm.queue", "order.confirm.key", body)

	return &types.Response{
		Message: "发送测试订单相关mq消息成功",
	}, nil
}
