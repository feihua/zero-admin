package cartitemservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/oms/gen/query"
	"github.com/zeromicro/go-zero/core/logc"
	"time"

	"github.com/feihua/zero-admin/rpc/oms/internal/svc"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateCartItemQuantityLogic 修改购物车中某个商品的数量
/*
Author: LiuFeiHua
Date: 2024/6/12 10:09
*/
type UpdateCartItemQuantityLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateCartItemQuantityLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCartItemQuantityLogic {
	return &UpdateCartItemQuantityLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateCartItemQuantity 修改购物车中某个商品的数量
func (l *UpdateCartItemQuantityLogic) UpdateCartItemQuantity(in *omsclient.UpdateCartItemQuantityReq) (*omsclient.CartItemResp, error) {
	q := query.OmsCartItem
	now := time.Now()
	expireTime := now.AddDate(0, 0, l.svcCtx.C.Cart.Timeout)
	_, err := q.WithContext(l.ctx).Where(q.ID.Eq(in.Id), q.MemberID.Eq(in.MemberId)).UpdateSimple(q.Quantity.Value(in.Quantity), q.ExpireTime.Value(expireTime))

	if err != nil {
		logc.Errorf(l.ctx, "修改购物车中某个商品的数量失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("修改购物车中某个商品的数量失败")
	}

	return &omsclient.CartItemResp{}, nil
}
