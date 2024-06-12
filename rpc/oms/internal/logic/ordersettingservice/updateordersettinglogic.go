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
		ID:                  in.Id,
		FlashOrderOvertime:  in.FinishOvertime,
		NormalOrderOvertime: in.NormalOrderOvertime,
		ConfirmOvertime:     in.ConfirmOvertime,
		FinishOvertime:      in.FinishOvertime,
		CommentOvertime:     in.CommentOvertime,
	})

	if err != nil {
		return nil, err
	}

	return &omsclient.UpdateOrderSettingResp{}, nil
}
