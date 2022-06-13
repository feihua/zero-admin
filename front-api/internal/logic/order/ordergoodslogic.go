package order

import (
	"context"

	"zero-admin/front-api/internal/svc"
	"zero-admin/front-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrderGoodsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOrderGoodsLogic(ctx context.Context, svcCtx *svc.ServiceContext) OrderGoodsLogic {
	return OrderGoodsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OrderGoodsLogic) OrderGoods(req types.OrderGoodsReq) (resp *types.OrderGoodsResp, err error) {
	// todo: add your logic here and delete this line

	return
}
