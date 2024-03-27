package orderoperatehistorservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"

	"github.com/feihua/zero-admin/rpc/oms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

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

func (l *OrderOperateHistoryDeleteLogic) OrderOperateHistoryDelete(in *omsclient.OrderOperateHistoryDeleteReq) (*omsclient.OrderOperateHistoryDeleteResp, error) {
	err := l.svcCtx.OmsOrderOperateHistoryModel.DeleteByIds(l.ctx, in.Ids)

	if err != nil {
		return nil, err
	}

	return &omsclient.OrderOperateHistoryDeleteResp{}, nil
}
