package orderservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/oms/gen/query"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"
	"github.com/zeromicro/go-zero/core/logc"
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
		logc.Errorf(l.ctx, "更新订单失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("更新订单失败")
	}

	order.Status = in.Status
	order.DeliveryCompany = in.DeliveryCompany
	order.DeliverySn = in.DeliverySn
	now := time.Now()
	order.DeliveryTime = &now
	order.Note = in.Note

	_, err = q.WithContext(l.ctx).Updates(order)

	if err != nil {
		logc.Errorf(l.ctx, "更新订单失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("更新订单失败")
	}

	return &omsclient.OrderUpdateResp{}, nil
}
