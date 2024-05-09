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

type OrderOperateHistoryAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOrderOperateHistoryAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderOperateHistoryAddLogic {
	return &OrderOperateHistoryAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *OrderOperateHistoryAddLogic) OrderOperateHistoryAdd(in *omsclient.OrderOperateHistoryAddReq) (*omsclient.OrderOperateHistoryAddResp, error) {
	err := query.OmsOrderOperateHistory.WithContext(l.ctx).Create(&model.OmsOrderOperateHistory{
		OrderID:     in.OrderId,
		OperateMan:  in.OperateMan,
		CreateTime:  time.Now(),
		OrderStatus: in.OrderStatus,
		Note:        &in.Note,
	})

	if err != nil {
		return nil, err
	}

	return &omsclient.OrderOperateHistoryAddResp{}, nil
}
