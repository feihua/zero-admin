package orderservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/oms/gen/model"
	"github.com/feihua/zero-admin/rpc/oms/gen/query"
	"time"

	"github.com/feihua/zero-admin/rpc/oms/internal/svc"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// CloseOrderLogic 批量关闭订单
/*
Author: LiuFeiHua
Date: 2024/5/14 17:32
*/
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

// CloseOrder 批量关闭订单
func (l *CloseOrderLogic) CloseOrder(in *omsclient.CloseOrderReq) (*omsclient.CloseOrderResp, error) {
	//1.批量关闭订单
	q := query.OmsOrder
	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Update(q.Status, 4)

	if err != nil {
		return nil, err
	}

	//2.添加操作记录
	var note = "订单关闭" + in.Note
	var histories []*model.OmsOrderOperateHistory
	for _, id := range in.Ids {
		histories = append(histories, &model.OmsOrderOperateHistory{
			OrderID:     id,
			OperateMan:  in.OperateMan,
			CreateTime:  time.Now(),
			OrderStatus: 4,
			Note:        &note,
		})
	}

	err = query.OmsOrderOperateHistory.WithContext(l.ctx).CreateInBatches(histories, len(histories))
	if err != nil {
		return nil, err
	}

	return &omsclient.CloseOrderResp{}, nil
}
