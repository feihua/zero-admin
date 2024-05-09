package orderoperatehistorservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/oms/gen/model"
	"github.com/feihua/zero-admin/rpc/oms/gen/query"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"
	"time"

	"github.com/feihua/zero-admin/rpc/oms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrderOperateHistoryUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOrderOperateHistoryUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderOperateHistoryUpdateLogic {
	return &OrderOperateHistoryUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *OrderOperateHistoryUpdateLogic) OrderOperateHistoryUpdate(in *omsclient.OrderOperateHistoryUpdateReq) (*omsclient.OrderOperateHistoryUpdateResp, error) {
	q := query.OmsOrderOperateHistory
	_, err := q.WithContext(l.ctx).Updates(&model.OmsOrderOperateHistory{
		ID:          in.Id,
		OrderID:     in.OrderId,
		OperateMan:  in.OperateMan,
		CreateTime:  time.Now(),
		OrderStatus: in.OrderStatus,
		Note:        &in.Note,
	})

	if err != nil {
		return nil, err
	}

	return &omsclient.OrderOperateHistoryUpdateResp{}, nil
}
