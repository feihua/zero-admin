package orderoperatehistorservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/oms/gen/query"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"

	"github.com/feihua/zero-admin/rpc/oms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

// OrderOperateHistoryDeleteLogic 订单操作记录
/*
Author: LiuFeiHua
Date: 2024/5/8 9:31
*/
type OrderOperateHistoryDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOrderOperateHistoryDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderOperateHistoryDeleteLogic {
	return &OrderOperateHistoryDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// OrderOperateHistoryDelete 删除订单操作记录
func (l *OrderOperateHistoryDeleteLogic) OrderOperateHistoryDelete(in *omsclient.OrderOperateHistoryDeleteReq) (*omsclient.OrderOperateHistoryDeleteResp, error) {
	q := query.OmsOrderOperateHistory
	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Delete()

	if err != nil {
		return nil, err
	}

	return &omsclient.OrderOperateHistoryDeleteResp{}, nil
}
