package orderoperatehistoryservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/oms/gen/query"

	"github.com/feihua/zero-admin/rpc/oms/internal/svc"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// DeleteOrderOperateHistoryLogic 删除订单操作历史记录
/*
Author: LiuFeiHua
Date: 2024/6/12 10:13
*/
type DeleteOrderOperateHistoryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteOrderOperateHistoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteOrderOperateHistoryLogic {
	return &DeleteOrderOperateHistoryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// DeleteOrderOperateHistory 删除订单操作历史记录
func (l *DeleteOrderOperateHistoryLogic) DeleteOrderOperateHistory(in *omsclient.DeleteOrderOperateHistoryReq) (*omsclient.DeleteOrderOperateHistoryResp, error) {
	q := query.OmsOrderOperateHistory
	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Delete()

	if err != nil {
		return nil, err
	}

	return &omsclient.DeleteOrderOperateHistoryResp{}, nil
}
