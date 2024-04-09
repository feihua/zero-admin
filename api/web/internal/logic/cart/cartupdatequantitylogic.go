package cart

import (
	"context"

	"github.com/feihua/zero-admin/api/web/internal/svc"
	"github.com/feihua/zero-admin/api/web/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

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

func (l *CartUpdateQuantityLogic) CartUpdateQuantity(req *types.CartItemUpdateQuantityReq) (resp *types.CartItemUpdateResp, err error) {
	// todo: add your logic here and delete this line

	return
}
