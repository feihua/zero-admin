package cart

import (
	"context"
	"zero-admin/rpc/oms/omsclient"

	"zero-admin/front-api/internal/svc"
	"zero-admin/front-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CartCheckOutLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCartCheckOutLogic(ctx context.Context, svcCtx *svc.ServiceContext) CartCheckOutLogic {
	return CartCheckOutLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CartCheckOutLogic) CartCheckOut(req types.CartCheckOutReq) (resp *types.CartCheckOutResp, err error) {
	l.svcCtx.Oms.CartItemCheckOut(l.ctx, &omsclient.CartItemCheckOutReq{
		Id: 0,
	})

	return &types.CartCheckOutResp{
		Errno:  0,
		Errmsg: "",
	}, nil
}
