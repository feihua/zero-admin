package cartitemservicelogic

import (
	"context"

	"zero-admin/rpc/oms/internal/svc"
	"zero-admin/rpc/oms/omsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// CartItemClearLogic
/*
Author: LiuFeiHua
Date: 2023/12/6 14:44
*/
type CartItemClearLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCartItemClearLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CartItemClearLogic {
	return &CartItemClearLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// CartItemClear 清空购物车
func (l *CartItemClearLogic) CartItemClear(in *omsclient.CartItemClearReq) (*omsclient.CartItemClearResp, error) {
	err := l.svcCtx.OmsCartItemModel.ClearCartItem(l.ctx, in.MemberId)

	if err != nil {
		return nil, err
	}

	return &omsclient.CartItemClearResp{}, nil
}
