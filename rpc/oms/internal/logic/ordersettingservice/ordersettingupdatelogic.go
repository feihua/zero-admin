package ordersettingservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/model/omsmodel"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"

	"github.com/feihua/zero-admin/rpc/oms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrderSettingUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOrderSettingUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderSettingUpdateLogic {
	return &OrderSettingUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *OrderSettingUpdateLogic) OrderSettingUpdate(in *omsclient.OrderSettingUpdateReq) (*omsclient.OrderSettingUpdateResp, error) {
	err := l.svcCtx.OmsOrderSettingModel.Update(l.ctx, &omsmodel.OmsOrderSetting{
		Id:                  in.Id,
		FlashOrderOvertime:  in.FinishOvertime,
		NormalOrderOvertime: in.NormalOrderOvertime,
		ConfirmOvertime:     in.ConfirmOvertime,
		FinishOvertime:      in.FinishOvertime,
		CommentOvertime:     in.CommentOvertime,
	})
	if err != nil {
		return nil, err
	}

	return &omsclient.OrderSettingUpdateResp{}, nil
}
