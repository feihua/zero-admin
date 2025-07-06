package orderservicelogic

import (
	"context"
	"errors"
	"github.com/bytedance/sonic"
	"github.com/feihua/zero-admin/rpc/oms/gen/query"
	"github.com/zeromicro/go-zero/core/logc"
	"gorm.io/gorm"

	"github.com/feihua/zero-admin/rpc/oms/internal/svc"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type ConfirmOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewConfirmOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ConfirmOrderLogic {
	return &ConfirmOrderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// ConfirmOrder 确认收货(app)
func (l *ConfirmOrderLogic) ConfirmOrder(in *omsclient.ConfirmOrderReq) (*omsclient.ConfirmOrderResp, error) {
	order := query.OmsOrderMain
	item, err := order.WithContext(l.ctx).Where(order.ID.Eq(in.OrderId), order.UserID.Eq(in.MemberId), order.IsDeleted.Eq(0), order.OrderStatus.Eq(3)).First()

	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		logc.Errorf(l.ctx, "订单不存在, 请求参数：%+v, 异常信息: %s", in, err.Error())
		return nil, errors.New("订单不存在")
	case err != nil:
		logc.Errorf(l.ctx, "查询订单异常, 请求参数：%+v, 异常信息: %s", in, err.Error())
		return nil, errors.New("查询订单异常")
	}

	item.OrderStatus = 4
	item.ReceiveStatus = 1

	_, err = order.WithContext(l.ctx).Where(order.ID.Eq(in.OrderId), order.UserID.Eq(in.MemberId)).Updates(item)

	if err != nil {
		logc.Errorf(l.ctx, "更新订单失败,参数:%+v,异常:%s", item, err.Error())
		return nil, errors.New("更新订单失败")
	}

	message := map[string]any{"id": in.OrderId}
	body, _ := sonic.Marshal(message)
	err = l.svcCtx.RabbitMQ.SendMessage("order.confirm.exchange", "order.confirm.queue", "order.confirm.key", body)

	return &omsclient.ConfirmOrderResp{}, nil
}
