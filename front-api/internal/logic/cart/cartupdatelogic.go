package cart

import (
	"context"

	"zero-admin/front-api/internal/svc"
	"zero-admin/front-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CartUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCartUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) CartUpdateLogic {
	return CartUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CartUpdateLogic) CartUpdate(req types.CartUpdateReq) (resp *types.CartUpdateResp, err error) {
	// todo: add your logic here and delete this line

	return
}
