package logic

import (
	"context"

	"go-zero-admin/rpc/oms/internal/svc"
	"go-zero-admin/rpc/oms/oms"

	"github.com/tal-tech/go-zero/core/logx"
)

type CartItemUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCartItemUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CartItemUpdateLogic {
	return &CartItemUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CartItemUpdateLogic) CartItemUpdate(in *oms.CartItemUpdateReq) (*oms.CartItemUpdateResp, error) {
	// todo: add your logic here and delete this line

	return &oms.CartItemUpdateResp{}, nil
}
