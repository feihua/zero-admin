package ordersettingservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/oms/gen/model"
	"github.com/feihua/zero-admin/rpc/oms/gen/query"

	"github.com/feihua/zero-admin/rpc/oms/internal/svc"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// AddOrderSettingLogic 添加订单设置表
/*
Author: LiuFeiHua
Date: 2024/6/12 9:35
*/
type AddOrderSettingLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddOrderSettingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddOrderSettingLogic {
	return &AddOrderSettingLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// AddOrderSetting 添加订单设置表
func (l *AddOrderSettingLogic) AddOrderSetting(in *omsclient.AddOrderSettingReq) (*omsclient.AddOrderSettingResp, error) {
	err := query.OmsOrderSetting.WithContext(l.ctx).Create(&model.OmsOrderSetting{
		FlashOrderOvertime:  in.FinishOvertime,
		NormalOrderOvertime: in.NormalOrderOvertime,
		ConfirmOvertime:     in.ConfirmOvertime,
		FinishOvertime:      in.FinishOvertime,
		CommentOvertime:     in.CommentOvertime,
	})

	if err != nil {
		return nil, err
	}

	return &omsclient.AddOrderSettingResp{}, nil
}
