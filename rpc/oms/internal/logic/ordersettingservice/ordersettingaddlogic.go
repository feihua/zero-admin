package ordersettingservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/oms/gen/model"
	"github.com/feihua/zero-admin/rpc/oms/gen/query"
	"github.com/feihua/zero-admin/rpc/oms/internal/svc"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrderSettingAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOrderSettingAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderSettingAddLogic {
	return &OrderSettingAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *OrderSettingAddLogic) OrderSettingAdd(in *omsclient.OrderSettingAddReq) (*omsclient.OrderSettingAddResp, error) {
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

	return &omsclient.OrderSettingAddResp{}, nil
}
