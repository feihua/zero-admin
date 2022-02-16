package logic

import (
	"context"
	"time"
	"zero-admin/rpc/model/omsmodel"

	"zero-admin/rpc/oms/internal/svc"
	"zero-admin/rpc/oms/oms"

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

func (l *OrderOperateHistoryAddLogic) OrderOperateHistoryAdd(in *oms.OrderOperateHistoryAddReq) (*oms.OrderOperateHistoryAddResp, error) {
	CreateTime, _ := time.Parse("2006-01-02 15:04:05", in.CreateTime)
	_, err := l.svcCtx.OmsOrderOperateHistoryModel.Insert(omsmodel.OmsOrderOperateHistory{
		OrderId:     in.OrderId,
		OperateMan:  in.OperateMan,
		CreateTime:  CreateTime,
		OrderStatus: in.OrderStatus,
		Note:        in.Note,
	})
	if err != nil {
		return nil, err
	}

	return &oms.OrderOperateHistoryAddResp{}, nil
}
