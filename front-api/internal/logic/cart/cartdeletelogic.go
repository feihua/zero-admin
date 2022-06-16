package cart

import (
	"context"
	"zero-admin/rpc/oms/omsclient"

	"zero-admin/front-api/internal/svc"
	"zero-admin/front-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CartDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCartDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) CartDeleteLogic {
	return CartDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CartDeleteLogic) CartDelete(req types.CartDeleteReq) (resp *types.CartListResp, err error) {
	l.svcCtx.Oms.CartItemDelete(l.ctx, &omsclient.CartItemDeleteReq{
		Id: 0,
	})

	return &types.CartListResp{
		Errno:  0,
		Errmsg: "",
	}, nil
}
