package cart

import (
	"context"
	"zero-admin/rpc/oms/omsclient"

	"zero-admin/front-api/internal/svc"
	"zero-admin/front-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CartListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCartListLogic(ctx context.Context, svcCtx *svc.ServiceContext) CartListLogic {
	return CartListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CartListLogic) CartList(req types.CartListReq) (resp *types.CartListResp, err error) {
	l.svcCtx.Oms.CartItemList(l.ctx, &omsclient.CartItemListReq{
		Current:  0,
		PageSize: 0,
	})

	return &types.CartListResp{
		Errno:  0,
		Errmsg: "",
	}, nil
}
