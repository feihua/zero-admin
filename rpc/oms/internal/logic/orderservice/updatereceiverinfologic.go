package orderservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/oms/gen/model"
	"github.com/feihua/zero-admin/rpc/oms/gen/query"
	"github.com/zeromicro/go-zero/core/logc"
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
	// 1.批量发货
	q := query.OmsOrder
	order, err := q.WithContext(l.ctx).Where(q.ID.Eq(in.OrderId)).First()
	if err != nil {
		logc.Errorf(l.ctx, "修改收货人信息失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("修改收货人信息失败")
	}

	order.ReceiverName = in.ReceiverName
	order.ReceiverPhone = in.ReceiverPhone
	order.ReceiverPostCode = in.ReceiverPostCode
	order.ReceiverDetailAddress = in.ReceiverDetailAddress
	order.ReceiverProvince = in.ReceiverProvince
	order.ReceiverCity = in.ReceiverCity
	order.ReceiverRegion = in.ReceiverRegion
	now := time.Now()
	order.ModifyTime = &now

	_, err = q.WithContext(l.ctx).Updates(order)
	if err != nil {
		logc.Errorf(l.ctx, "修改收货人信息失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("修改收货人信息失败")
	}

	// 2.添加操作记录
	var note = "修改收货人信息"
	err = query.OmsOrderOperateHistory.WithContext(l.ctx).Create(&model.OmsOrderOperateHistory{
		OrderID:     in.OrderId,
		OperateMan:  in.OperateMan,
		CreateTime:  now,
		OrderStatus: in.Status,
		Note:        note,
	})
	if err != nil {
		logc.Errorf(l.ctx, "修改收货人信息失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("修改收货人信息失败")
	}

	return &omsclient.UpdateReceiverInfoResp{}, nil
}
