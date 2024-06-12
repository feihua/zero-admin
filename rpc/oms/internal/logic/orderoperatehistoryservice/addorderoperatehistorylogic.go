package orderoperatehistoryservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/oms/gen/model"
	"github.com/feihua/zero-admin/rpc/oms/gen/query"
	"time"

	"github.com/feihua/zero-admin/rpc/oms/internal/svc"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// AddOrderOperateHistoryLogic 添加订单操作历史记录
/*
Author: LiuFeiHua
Date: 2024/6/12 10:13
*/
type AddOrderOperateHistoryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddOrderOperateHistoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddOrderOperateHistoryLogic {
	return &AddOrderOperateHistoryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// AddOrderOperateHistory 添加订单操作历史记录
func (l *AddOrderOperateHistoryLogic) AddOrderOperateHistory(in *omsclient.AddOrderOperateHistoryReq) (*omsclient.AddOrderOperateHistoryResp, error) {
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

	return &omsclient.AddOrderOperateHistoryResp{}, nil
}
