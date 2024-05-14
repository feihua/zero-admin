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

// UpdateReceiverInfoLogic 修改收货人信息
/*
Author: LiuFeiHua
Date: 2024/5/14 17:34
*/
type UpdateReceiverInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateReceiverInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateReceiverInfoLogic {
	return &UpdateReceiverInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateReceiverInfo 修改收货人信息
func (l *UpdateReceiverInfoLogic) UpdateReceiverInfo(in *omsclient.UpdateReceiverInfoReq) (*omsclient.UpdateReceiverInfoResp, error) {
	//1.批量发货
	q := query.OmsOrder
	order, err := q.WithContext(l.ctx).Where(q.ID.Eq(in.OrderId)).First()
	if err != nil {
		return nil, err
	}

	order.ReceiverName = in.ReceiverName
	order.ReceiverPhone = in.ReceiverPhone
	order.ReceiverPostCode = in.ReceiverPostCode
	order.ReceiverDetailAddress = in.ReceiverDetailAddress
	order.ReceiverProvince = in.ReceiverProvince
	order.ReceiverCity = in.ReceiverCity
	order.ReceiverRegion = in.ReceiverRegion
	order.ModifyTime = time.Now()

	_, err = q.WithContext(l.ctx).Updates(order)

	if err != nil {
		return nil, err
	}

	//2.添加操作记录
	var note = "修改收货人信息"
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

	return &omsclient.UpdateReceiverInfoResp{}, nil
}
