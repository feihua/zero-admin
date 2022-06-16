package order

import (
	"context"
	"zero-admin/rpc/oms/omsclient"

	"zero-admin/front-api/internal/svc"
	"zero-admin/front-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrderRefundLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOrderRefundLogic(ctx context.Context, svcCtx *svc.ServiceContext) OrderRefundLogic {
	return OrderRefundLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OrderRefundLogic) OrderRefund(req types.OrderRefundReq) (resp *types.OrderRefundResp, err error) {
	l.svcCtx.Oms.OrderRefund(l.ctx, &omsclient.OrderRefundReq{
		Id:       req.OrderId,
		MemberId: req.UserId,
	})

	return &types.OrderRefundResp{
		Errno:  0,
		Errmsg: "",
	}, nil
}
