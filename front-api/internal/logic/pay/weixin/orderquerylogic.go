package logic

import (
	"context"
	"go-zero-admin/rpc/pay/payclient"

	"go-zero-admin/front-api/internal/svc"
	"go-zero-admin/front-api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type OrderQueryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOrderQueryLogic(ctx context.Context, svcCtx *svc.ServiceContext) OrderQueryLogic {
	return OrderQueryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OrderQueryLogic) OrderQuery(req types.OrderQueryReq) (*types.OrderQueryResp, error) {
	l.svcCtx.Pay.OrderQuery(l.ctx, &payclient.OrderQueryReq{
		BusinessId: req.BusinessId,
	})

	return &types.OrderQueryResp{
		Code:    "",
		Message: "",
	}, nil
}
