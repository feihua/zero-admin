package logic

import (
	"context"
	"go-zero-admin/rpc/oms/omsclient"

	"go-zero-admin/api/internal/svc"
	"go-zero-admin/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type OrderSettingDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOrderSettingDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) OrderSettingDeleteLogic {
	return OrderSettingDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OrderSettingDeleteLogic) OrderSettingDelete(req types.DeleteOrderSettingReq) (*types.DeleteOrderSettingResp, error) {
	_, _ = l.svcCtx.Oms.OrderSettingDelete(l.ctx, &omsclient.OrderSettingDeleteReq{
		Id: req.Id,
	})

	return &types.DeleteOrderSettingResp{
		Code:    "000000",
		Message: "",
	}, nil
}
