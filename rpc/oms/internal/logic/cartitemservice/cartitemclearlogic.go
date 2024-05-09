package cartitemservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/oms/gen/query"

	"github.com/feihua/zero-admin/rpc/oms/internal/svc"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// CartItemClearLogic 会员购物车
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

// CartItemClear 根据会员id清空购物车
func (l *CartItemClearLogic) CartItemClear(in *omsclient.CartItemClearReq) (*omsclient.CartItemClearResp, error) {
	q := query.OmsCartItem
	_, err := q.WithContext(l.ctx).Where(q.MemberID.Eq(in.MemberId)).Delete()

	if err != nil {
		return nil, err
	}

	return &omsclient.CartItemClearResp{}, nil
}
