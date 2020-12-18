package logic

import (
	"context"

	"go-zero-admin/front-api/internal/svc"
	"go-zero-admin/front-api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type CartItemAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCartItemAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) CartItemAddLogic {
	return CartItemAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CartItemAddLogic) CartItemAdd(req types.AddCartItemReq) (*types.AddCartItemResp, error) {
	// todo: add your logic here and delete this line

	return &types.AddCartItemResp{}, nil
}
