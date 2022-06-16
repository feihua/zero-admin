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
	//l.svcCtx.Pms.ProductListByCategoryId(l.ctx, &pmsclient.ProductListByCategoryIdReq{
	//	ProductCategoryId: req.OrderId,
	//})

	return &types.OrderGoodsResp{
		Errno:  0,
		Data:   nil,
		Errmsg: "",
	}, nil
}
