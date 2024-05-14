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

// DeliveryLogic 批量发货
/*
Author: LiuFeiHua
Date: 2024/5/14 17:33
*/
type DeliveryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeliveryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeliveryLogic {
	return &DeliveryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// Delivery 批量发货
func (l *DeliveryLogic) Delivery(in *omsclient.DeliveryReq) (*omsclient.DeliveryResp, error) {
	//1.批量发货
	q := query.OmsOrder
	order, err := q.WithContext(l.ctx).Where(q.ID.Eq(in.OrderId)).First()
	if err != nil {
		return nil, err
	}

	order.Status = 2
	order.DeliveryCompany = in.DeliveryCompany
	order.DeliverySn = in.DeliverySn
	order.DeliveryTime = time.Now()

	_, err = q.WithContext(l.ctx).Updates(order)

	if err != nil {
		return nil, err
	}

	//2.添加操作记录
	var note = "完成发货"
	err = query.OmsOrderOperateHistory.WithContext(l.ctx).Create(&model.OmsOrderOperateHistory{
		OrderID:     in.OrderId,
		OperateMan:  in.OperateMan,
		CreateTime:  time.Now(),
		OrderStatus: 2,
		Note:        &note,
	})
	if err != nil {
		return nil, err
	}

	return &omsclient.DeliveryResp{}, nil
}
