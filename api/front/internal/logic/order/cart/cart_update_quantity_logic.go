package cart

import (
	"context"

	"github.com/feihua/zero-admin/api/front/internal/svc"
	"github.com/feihua/zero-admin/api/front/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// CartUpdateQuantityLogic
/*
Author: LiuFeiHua
Date: 2025/6/20 10:26
*/
type CartUpdateQuantityLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCartUpdateQuantityLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CartUpdateQuantityLogic {
	return &CartUpdateQuantityLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// CartUpdateQuantity 修改购物车中某个商品的数量
func (l *CartUpdateQuantityLogic) CartUpdateQuantity(req *types.CartItemUpdateQuantityReq) (resp *types.CartItemResp, err error) {
	// todo: add your logic here and delete this line

	return
}
