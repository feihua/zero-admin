package ordersettingservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/oms/gen/model"
	"github.com/feihua/zero-admin/rpc/oms/gen/query"

	"github.com/feihua/zero-admin/rpc/oms/internal/svc"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateOrderSettingLogic 更新订单设置表
/*
Author: LiuFeiHua
Date: 2024/6/12 9:40
*/
type UpdateOrderSettingLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateOrderSettingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateOrderSettingLogic {
	return &UpdateOrderSettingLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateOrderSetting 更新订单设置表
func (l *UpdateOrderSettingLogic) UpdateOrderSetting(in *omsclient.UpdateOrderSettingReq) (*omsclient.UpdateOrderSettingResp, error) {
	q := query.OmsOrderSetting
	_, err := q.WithContext(l.ctx).Updates(&model.OmsOrderSetting{
		ID:                  in.Id,                  //
		FlashOrderOvertime:  in.FlashOrderOvertime,  // 秒杀订单超时关闭时间(分)
		NormalOrderOvertime: in.NormalOrderOvertime, // 正常订单超时时间(分)
		ConfirmOvertime:     in.ConfirmOvertime,     // 发货后自动确认收货时间（天）
		FinishOvertime:      in.FinishOvertime,      // 自动完成交易时间，不能申请售后（天）
		Status:              in.Status,              // 状态：0->禁用；1->启用
		IsDefault:           in.IsDefault,           // 是否默认：0->否；1->是
		CommentOvertime:     in.CommentOvertime,     // 订单完成后自动好评时间（天）
	})

	if err != nil {
		return nil, err
	}

	return &omsclient.UpdateOrderSettingResp{}, nil
}
