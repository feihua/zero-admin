package orderservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/oms/gen/query"
	"github.com/feihua/zero-admin/rpc/oms/internal/svc"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteOrderLogic {
	return &DeleteOrderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// DeleteOrder 删除订单
func (l *DeleteOrderLogic) DeleteOrder(in *omsclient.DeleteOrderReq) (*omsclient.DeleteOrderResp, error) {
	q := query.OmsOrderMain
	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...), q.UserID.Eq(in.MemberId), q.OrderStatus.In(4, 5, 6)).Delete()

	if err != nil {
		return nil, errors.New("用户订单不存在,删除用户订单失败")
	}

	return &omsclient.DeleteOrderResp{}, nil
}
