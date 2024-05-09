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
