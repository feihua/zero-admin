package logic

import (
	"context"

	"zero-admin/rpc/oms/internal/svc"
	"zero-admin/rpc/oms/omsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type CartItemCheckOutLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCartItemCheckOutLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CartItemCheckOutLogic {
	return &CartItemCheckOutLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CartItemCheckOutLogic) CartItemCheckOut(in *omsclient.CartItemCheckOutReq) (*omsclient.CartItemCheckOutResp, error) {
	// todo: add your logic here and delete this line

	return &omsclient.CartItemCheckOutResp{}, nil
}
