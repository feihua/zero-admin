package logic

import (
	"context"
	"go-zero-admin/rpc/model/omsmodel"
	"time"

	"go-zero-admin/rpc/oms/internal/svc"
	"go-zero-admin/rpc/oms/oms"

	"github.com/tal-tech/go-zero/core/logx"
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

func (l *OrderOperateHistoryUpdateLogic) OrderOperateHistoryUpdate(in *oms.OrderOperateHistoryUpdateReq) (*oms.OrderOperateHistoryUpdateResp, error) {
	CreateTime, _ := time.Parse("2006-01-02 15:04:05", in.CreateTime)
	err := l.svcCtx.OmsOrderOperateHistoryModel.Update(omsmodel.OmsOrderOperateHistory{
		OrderId:     in.OrderId,
		OperateMan:  in.OperateMan,
		CreateTime:  CreateTime,
		OrderStatus: in.OrderStatus,
		Note:        in.Note,
	})
	if err != nil {
		return nil, err
	}

	return &oms.OrderOperateHistoryUpdateResp{}, nil
}
