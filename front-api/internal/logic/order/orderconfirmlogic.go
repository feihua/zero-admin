package order

import (
	"context"
	"zero-admin/rpc/oms/omsclient"

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
	l.svcCtx.Oms.OrderConfirm(l.ctx, &omsclient.OrderConfirmReq{
		Id:       req.OrderId,
		MemberId: req.UserId,
	})

	return &types.OrderConfirmResp{
		Errno:  0,
		Errmsg: "",
	}, nil
}
