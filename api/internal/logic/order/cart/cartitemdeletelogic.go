package logic

import (
	"context"
	"go-zero-admin/api/internal/svc"
	"go-zero-admin/api/internal/types"
	"go-zero-admin/rpc/oms/omsclient"

	"github.com/tal-tech/go-zero/core/logx"
)

type CartItemDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCartItemDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) CartItemDeleteLogic {
	return CartItemDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CartItemDeleteLogic) CartItemDelete(req types.DeleteCartItemReq) (*types.DeleteCartItemResp, error) {
	_, _ = l.svcCtx.Oms.CartItemDelete(l.ctx, &omsclient.CartItemDeleteReq{
		Id: req.Id,
	})

	return &types.DeleteCartItemResp{}, nil
}
