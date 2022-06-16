package logic

import (
	"context"

	"zero-admin/rpc/oms/internal/svc"
	"zero-admin/rpc/oms/omsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type CartItemFastAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCartItemFastAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CartItemFastAddLogic {
	return &CartItemFastAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CartItemFastAddLogic) CartItemFastAdd(in *omsclient.CartItemFastAddReq) (*omsclient.CartItemFastAddResp, error) {
	// todo: add your logic here and delete this line

	return &omsclient.CartItemFastAddResp{}, nil
}
