package cartitemservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/oms/gen/query"

	"github.com/feihua/zero-admin/rpc/oms/internal/svc"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// CartItemUpdateQuantityLogic
/*
Author: LiuFeiHua
Date: 2023/12/6 15:31
*/
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

// CartItemUpdateQuantity 修改购物车中某个商品的数量
func (l *CartItemUpdateQuantityLogic) CartItemUpdateQuantity(in *omsclient.CartItemUpdateReq) (*omsclient.CartItemUpdateResp, error) {
	q := query.OmsCartItem
	_, err := q.WithContext(l.ctx).Where(q.ID.Eq(in.Id), q.MemberID.Eq(in.MemberId)).Update(q.Quantity, in.Quantity)

	if err != nil {
		return nil, err
	}
	return &omsclient.CartItemUpdateResp{}, nil
}
