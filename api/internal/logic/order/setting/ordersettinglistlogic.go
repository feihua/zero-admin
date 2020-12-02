package logic

import (
	"context"

	"go-zero-admin/api/internal/svc"
	"go-zero-admin/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type OrderSettingListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOrderSettingListLogic(ctx context.Context, svcCtx *svc.ServiceContext) OrderSettingListLogic {
	return OrderSettingListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OrderSettingListLogic) OrderSettingList(req types.ListOrderSettingReq) (*types.ListOrderSettingResp, error) {
	// todo: add your logic here and delete this line

	return &types.ListOrderSettingResp{}, nil
}
