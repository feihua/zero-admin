package orderreturnreasonservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/oms/gen/model"
	"github.com/feihua/zero-admin/rpc/oms/gen/query"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"
	"time"

	"github.com/feihua/zero-admin/rpc/oms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

// OrderReturnReasonUpdateLogic 退货原因
/*
Author: LiuFeiHua
Date: 2024/5/14 11:36
*/
type OrderReturnReasonUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOrderReturnReasonUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderReturnReasonUpdateLogic {
	return &OrderReturnReasonUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// OrderReturnReasonUpdate 更新退货原因
func (l *OrderReturnReasonUpdateLogic) OrderReturnReasonUpdate(in *omsclient.OrderReturnReasonUpdateReq) (*omsclient.OrderReturnReasonUpdateResp, error) {
	q := query.OmsOrderReturnReason
	_, err := q.WithContext(l.ctx).Updates(&model.OmsOrderReturnReason{
		ID:         in.Id,
		Name:       in.Name,
		Sort:       in.Sort,
		Status:     in.Status,
		CreateTime: time.Now(),
	})

	if err != nil {
		return nil, err
	}

	return &omsclient.OrderReturnReasonUpdateResp{}, nil
}
