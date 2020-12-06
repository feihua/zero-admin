package logic

import (
	"context"
	"go-zero-admin/rpc/oms/omsclient"

	"go-zero-admin/api/internal/svc"
	"go-zero-admin/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type OrderAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOrderAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) OrderAddLogic {
	return OrderAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OrderAddLogic) OrderAdd(req types.AddOrderReq) (*types.AddOrderResp, error) {
	_, err := l.svcCtx.Oms.OrderAdd(l.ctx, &omsclient.OrderAddReq{})

	if err != nil {
		return nil, err
	}

	return &types.AddOrderResp{}, nil
}
