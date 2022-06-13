package order

import (
	"context"

	"zero-admin/front-api/internal/svc"
	"zero-admin/front-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrderCancelLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOrderCancelLogic(ctx context.Context, svcCtx *svc.ServiceContext) OrderCancelLogic {
	return OrderCancelLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OrderCancelLogic) OrderCancel(req types.OrderCancelReq) (resp *types.OrderCancelResp, err error) {
	// todo: add your logic here and delete this line

	return
}
