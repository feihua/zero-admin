package logic

import (
	"context"

	"zero-admin/rpc/oms/internal/svc"
	"zero-admin/rpc/oms/omsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type CartItemUpdateQuantityLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCartItemUpdateQuantityLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CartItemUpdateQuantityLogic {
	return &CartItemUpdateQuantityLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CartItemUpdateQuantityLogic) CartItemUpdateQuantity(in *omsclient.CartItemUpdateReq) (*omsclient.CartItemUpdateResp, error) {
	// todo: add your logic here and delete this line

	return &omsclient.CartItemUpdateResp{}, nil
}
