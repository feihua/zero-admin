package logic

import (
	"context"
	"zero-admin/rpc/model/omsmodel"
	"zero-admin/rpc/oms/internal/svc"
	"zero-admin/rpc/oms/oms"

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

func (l *OrderSettingAddLogic) OrderSettingAdd(in *oms.OrderSettingAddReq) (*oms.OrderSettingAddResp, error) {
	_, err := l.svcCtx.OmsOrderSettingModel.Insert(omsmodel.OmsOrderSetting{
		FlashOrderOvertime:  in.FinishOvertime,
		NormalOrderOvertime: in.NormalOrderOvertime,
		ConfirmOvertime:     in.ConfirmOvertime,
		FinishOvertime:      in.FinishOvertime,
		CommentOvertime:     in.CommentOvertime,
	})
	if err != nil {
		return nil, err
	}

	return &oms.OrderSettingAddResp{}, nil
}
