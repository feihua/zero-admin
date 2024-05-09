package orderreturnreasonservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/oms/gen/model"
	"github.com/feihua/zero-admin/rpc/oms/gen/query"
	"github.com/feihua/zero-admin/rpc/oms/internal/svc"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrderReturnReasonAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOrderReturnReasonAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderReturnReasonAddLogic {
	return &OrderReturnReasonAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *OrderReturnReasonAddLogic) OrderReturnReasonAdd(in *omsclient.OrderReturnReasonAddReq) (*omsclient.OrderReturnReasonAddResp, error) {
	err := query.OmsOrderReturnReason.WithContext(l.ctx).Create(&model.OmsOrderReturnReason{
		Name:       in.Name,
		Sort:       in.Sort,
		Status:     in.Status,
		CreateTime: time.Now(),
	})

	if err != nil {
		return nil, err
	}

	return &omsclient.OrderReturnReasonAddResp{}, nil
}
