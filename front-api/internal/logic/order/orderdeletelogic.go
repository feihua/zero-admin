package order

import (
	"context"
	"zero-admin/rpc/oms/omsclient"

	"zero-admin/front-api/internal/svc"
	"zero-admin/front-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrderDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOrderDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) OrderDeleteLogic {
	return OrderDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OrderDeleteLogic) OrderDelete(req types.OrderDeleteReq) (resp *types.OrderDeleteResp, err error) {
	l.svcCtx.Oms.OrderDelete(l.ctx, &omsclient.OrderDeleteReq{
		Id: req.OrderId,
	})

	return &types.OrderDeleteResp{
		Errno:  0,
		Errmsg: "",
	}, nil
}
