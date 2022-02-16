package logic

import (
	"context"

	"zero-admin/rpc/oms/internal/svc"
	"zero-admin/rpc/oms/oms"

	"github.com/zeromicro/go-zero/core/logx"
)

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

func (l *CartItemDeleteLogic) CartItemDelete(in *oms.CartItemDeleteReq) (*oms.CartItemDeleteResp, error) {
	err := l.svcCtx.OmsCartItemModel.Delete(in.Id)

	if err != nil {
		return nil, err
	}

	return &oms.CartItemDeleteResp{}, nil
}
