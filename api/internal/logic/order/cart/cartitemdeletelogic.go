package logic

import (
	"context"
	"zero-admin/api/internal/common/errorx"
	"zero-admin/api/internal/svc"
	"zero-admin/api/internal/types"
	"zero-admin/rpc/oms/omsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type CartItemDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCartItemDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) CartItemDeleteLogic {
	return CartItemDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CartItemDeleteLogic) CartItemDelete(req types.DeleteCartItemReq) (*types.DeleteCartItemResp, error) {
	_, err := l.svcCtx.Oms.CartItemDelete(l.ctx, &omsclient.CartItemDeleteReq{
		Id: req.Id,
	})

	if err != nil {
		logx.WithContext(l.ctx).Errorf("根据Id: %d,删除购物车异常:%s", req.Id, err.Error())
		return nil, errorx.NewDefaultError("删除购物车失败")
	}
	return &types.DeleteCartItemResp{
		Code:    "000000",
		Message: "",
	}, nil
}
