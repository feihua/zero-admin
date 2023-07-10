package cartitemservicelogic

import (
	"context"
	"zero-admin/rpc/oms/omsclient"

	"zero-admin/rpc/oms/internal/svc"

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

func (l *CartItemDeleteLogic) CartItemDelete(in *omsclient.CartItemDeleteReq) (*omsclient.CartItemDeleteResp, error) {
	err := l.svcCtx.OmsCartItemModel.DeleteByIds(l.ctx, in.Ids)

	if err != nil {
		return nil, err
	}

	return &omsclient.CartItemDeleteResp{}, nil
}
