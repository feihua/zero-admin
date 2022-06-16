package cart

import (
	"context"
	"zero-admin/rpc/oms/omsclient"

	"zero-admin/front-api/internal/svc"
	"zero-admin/front-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CartCheckedLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCartCheckedLogic(ctx context.Context, svcCtx *svc.ServiceContext) CartCheckedLogic {
	return CartCheckedLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CartCheckedLogic) CartChecked(req types.CartCheckedReq) (resp *types.CartListResp, err error) {
	l.svcCtx.Oms.CartItemChecked(l.ctx, &omsclient.CartItemCheckedReq{
		Id: 0,
	})

	return &types.CartListResp{
		Errno:  0,
		Errmsg: "",
	}, nil
}
