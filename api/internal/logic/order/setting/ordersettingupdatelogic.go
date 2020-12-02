package logic

import (
	"context"

	"go-zero-admin/api/internal/svc"
	"go-zero-admin/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type OrderSettingUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOrderSettingUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) OrderSettingUpdateLogic {
	return OrderSettingUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OrderSettingUpdateLogic) OrderSettingUpdate(req types.UpdateOrderSettingReq) (*types.UpdateOrderSettingResp, error) {
	// todo: add your logic here and delete this line

	return &types.UpdateOrderSettingResp{}, nil
}
