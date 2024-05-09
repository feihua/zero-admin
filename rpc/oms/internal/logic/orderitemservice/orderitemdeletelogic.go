package orderitemservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/oms/gen/query"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"

	"github.com/feihua/zero-admin/rpc/oms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

// OrderItemDeleteLogic 订单商品
/*
Author: LiuFeiHua
Date: 2024/5/8 9:31
*/
type OrderItemDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOrderItemDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderItemDeleteLogic {
	return &OrderItemDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// OrderItemDelete 删除订单商品
func (l *OrderItemDeleteLogic) OrderItemDelete(in *omsclient.OrderItemDeleteReq) (*omsclient.OrderItemDeleteResp, error) {
	q := query.OmsOrderItem
	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Delete()

	if err != nil {
		return nil, err
	}

	return &omsclient.OrderItemDeleteResp{}, nil
}
