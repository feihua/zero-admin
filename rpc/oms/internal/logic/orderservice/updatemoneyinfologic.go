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

// UpdateMoneyInfoLogic 修改订单费用信息
/*
Author: LiuFeiHua
Date: 2024/5/14 17:34
*/
type UpdateMoneyInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateMoneyInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateMoneyInfoLogic {
	return &UpdateMoneyInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateMoneyInfo 修改订单费用信息
func (l *UpdateMoneyInfoLogic) UpdateMoneyInfo(in *omsclient.UpdateMoneyInfoReq) (*omsclient.UpdateMoneyInfoResp, error) {
	//1.修改订单费用信息
	q := query.OmsOrder
	order, err := q.WithContext(l.ctx).Where(q.ID.Eq(in.OrderId)).First()
	if err != nil {
		return nil, err
	}

	order.FreightAmount = in.FreightAmount
	order.DiscountAmount = in.DiscountAmount

	order.ModifyTime = time.Now()

	_, err = q.WithContext(l.ctx).Updates(order)

	if err != nil {
		return nil, err
	}

	//2.添加操作记录
	var note = "修改订单费用信息"
	err = query.OmsOrderOperateHistory.WithContext(l.ctx).Create(&model.OmsOrderOperateHistory{
		OrderID:     in.OrderId,
		OperateMan:  in.OperateMan,
		CreateTime:  time.Now(),
		OrderStatus: in.Status,
		Note:        &note,
	})
	if err != nil {
		return nil, err
	}

	return &omsclient.UpdateMoneyInfoResp{}, nil
}
