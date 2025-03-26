package orderoperatehistoryservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/oms/gen/query"
	"github.com/zeromicro/go-zero/core/logc"

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
		logc.Errorf(l.ctx, "删除订单操作历史记录失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("删除订单操作历史记录失败")
	}

	return &omsclient.DeleteOrderOperateHistoryResp{}, nil
}
