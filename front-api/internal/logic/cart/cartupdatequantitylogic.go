package cart

import (
	"context"
	"zero-admin/rpc/oms/omsclient"

	"zero-admin/front-api/internal/svc"
	"zero-admin/front-api/internal/types"

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
	_, _ = l.svcCtx.CartItemService.CartItemUpdateQuantity(l.ctx, &omsclient.CartItemUpdateReq{Id: req.Id, Quantity: req.Quantity})

	return &types.CartItemUpdateResp{
		Code:    0,
		Message: "操作成功",
	}, nil
}
