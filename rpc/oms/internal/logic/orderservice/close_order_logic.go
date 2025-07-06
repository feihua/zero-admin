package orderservicelogic

import (
	"context"
	"errors"
	"github.com/bytedance/sonic"
	"github.com/feihua/zero-admin/rpc/oms/gen/model"
	"github.com/feihua/zero-admin/rpc/oms/gen/query"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/rpc/oms/internal/svc"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type CloseOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCloseOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CloseOrderLogic {
	return &CloseOrderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// CloseOrder 关闭订单
func (l *CloseOrderLogic) CloseOrder(in *omsclient.CloseOrderReq) (*omsclient.CloseOrderResp, error) {
	q := query.OmsOrderMain

	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Update(q.OrderStatus, 5)

	if err != nil {
		logc.Errorf(l.ctx, "更新订单状态失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("更新订单状态失败")
	}

	var items []*model.OmsOrderOperationLog
	for _, id := range in.Ids {
		items = append(items, &model.OmsOrderOperationLog{
			OrderID:       id,            // 订单ID
			OrderNo:       in.OrderNo,    // 订单编号
			OperationType: 5,             // 操作类型：1-创建订单，2-支付订单，3-发货，4-确认收货，5-取消订单，6-退款
			OperatorID:    in.OperatorId, // 操作人ID
			OperatorType:  3,             // 操作人类型：1-用户，2-系统，3-管理员
			OperatorNote:  in.Note,       // 操作备注
		})
	}

	err = query.OmsOrderOperationLog.WithContext(l.ctx).CreateInBatches(items, len(items))
	if err != nil {
		logc.Errorf(l.ctx, "添加订单操作记录失败,参数:%+v,异常:%s", items, err.Error())
		return nil, errors.New("添加订单操作记录失败")
	}

	for id := range in.Ids {
		message := map[string]any{"id": id}
		body, _ := sonic.Marshal(message)
		err = l.svcCtx.RabbitMQ.SendMessage("order.event.exchange", "order.close.queue", "order.close.key", body)

	}

	return &omsclient.CloseOrderResp{}, nil
}
