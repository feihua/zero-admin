package order

import (
	"context"
	"zero-admin/rpc/oms/omsclient"

	"zero-admin/front-api/internal/svc"
	"zero-admin/front-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrderListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOrderListLogic(ctx context.Context, svcCtx *svc.ServiceContext) OrderListLogic {
	return OrderListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OrderListLogic) OrderList(req types.OrderListReq) (resp *types.OrderListResp, err error) {
	l.svcCtx.Oms.OrderListByMemberId(l.ctx, &omsclient.OrderListByMemberIdReq{
		MemberId: req.UserId,
	})

	return &types.OrderListResp{
		Errno:  0,
		Errmsg: "",
	}, nil
}
