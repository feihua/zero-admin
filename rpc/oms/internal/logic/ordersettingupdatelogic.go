package logic

import (
	"context"

	"go-zero-admin/rpc/oms/internal/svc"
	"go-zero-admin/rpc/oms/oms"

	"github.com/tal-tech/go-zero/core/logx"
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

func (l *OrderSettingUpdateLogic) OrderSettingUpdate(in *oms.OrderSettingUpdateReq) (*oms.OrderSettingUpdateResp, error) {
	// todo: add your logic here and delete this line

	return &oms.OrderSettingUpdateResp{}, nil
}
