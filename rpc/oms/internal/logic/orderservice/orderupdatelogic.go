package orderservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/oms/gen/query"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"
	"time"

	"github.com/feihua/zero-admin/rpc/oms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrderUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOrderUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderUpdateLogic {
	return &OrderUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// OrderUpdate 更新订单
func (l *OrderUpdateLogic) OrderUpdate(in *omsclient.OrderUpdateReq) (*omsclient.OrderUpdateResp, error) {
	q := query.OmsOrder
	order, err := q.WithContext(l.ctx).Where(q.ID.Eq(in.Id)).First()
	if err != nil {
		return nil, err
	}

	order.Status = in.Status
	order.DeliveryCompany = in.DeliveryCompany
	order.DeliverySn = in.DeliverySn
	order.DeliveryTime = time.Now()
	order.Note = in.Note

	_, err = q.WithContext(l.ctx).Updates(order)

	if err != nil {
		return nil, err
	}

	return &omsclient.OrderUpdateResp{}, nil
}
