package orderdeliveryservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/oms/gen/query"
	"github.com/feihua/zero-admin/rpc/oms/internal/svc"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
)

// DeleteOrderDeliveryLogic 删除订单收货地址
/*
Author: LiuFeiHua
Date: 2025/05/26 15:21:44
*/
type DeleteOrderDeliveryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteOrderDeliveryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteOrderDeliveryLogic {
	return &DeleteOrderDeliveryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// DeleteOrderDelivery 删除订单收货地址
func (l *DeleteOrderDeliveryLogic) DeleteOrderDelivery(in *omsclient.DeleteOrderDeliveryReq) (*omsclient.DeleteOrderDeliveryResp, error) {
	q := query.OmsOrderDelivery

	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Delete()

	if err != nil {
		logc.Errorf(l.ctx, "删除订单收货地址失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("删除订单收货地址失败")
	}

	return &omsclient.DeleteOrderDeliveryResp{}, nil
}
