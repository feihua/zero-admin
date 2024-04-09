package cart

import (
	"context"

	"github.com/feihua/zero-admin/api/web/internal/svc"
	"github.com/feihua/zero-admin/api/web/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CartProductLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCartProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CartProductLogic {
	return &CartProductLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CartProductLogic) CartProduct(req *types.CartProductReq) (resp *types.CartProductResp, err error) {
	// todo: add your logic here and delete this line

	return
}
