package cartitemservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/oms/gen/query"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"

	"github.com/feihua/zero-admin/rpc/oms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

// CartItemDeleteLogic
/*
Author: LiuFeiHua
Date: 2023/12/6 14:45
*/
type CartItemDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCartItemDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CartItemDeleteLogic {
	return &CartItemDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// CartItemDelete 删除购物车中的某个商品
func (l *CartItemDeleteLogic) CartItemDelete(in *omsclient.CartItemDeleteReq) (*omsclient.CartItemDeleteResp, error) {
	q := query.OmsCartItem
	_, err := q.WithContext(l.ctx).Where(q.MemberID.Eq(in.MemberId), q.ID.In(in.Ids...)).Delete()

	if err != nil {
		return nil, err
	}

	return &omsclient.CartItemDeleteResp{}, nil
}
