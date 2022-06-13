package order

import (
	"context"

	"zero-admin/front-api/internal/svc"
	"zero-admin/front-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrderConfirmLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOrderConfirmLogic(ctx context.Context, svcCtx *svc.ServiceContext) OrderConfirmLogic {
	return OrderConfirmLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OrderConfirmLogic) OrderConfirm(req types.OrderConfirmReq) (resp *types.OrderConfirmResp, err error) {
	// todo: add your logic here and delete this line

	return
}
