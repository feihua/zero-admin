package cart

import (
	"context"
	"zero-admin/rpc/oms/omsclient"

	"zero-admin/front-api/internal/svc"
	"zero-admin/front-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CartItemDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCartItemDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CartItemDeleteLogic {
	return &CartItemDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CartItemDeleteLogic) CartItemDelete(req *types.CartItemDeleteReq) (resp *types.CartItemDeleteResp, err error) {
	_, _ = l.svcCtx.CartItemService.CartItemDelete(l.ctx, &omsclient.CartItemDeleteReq{MemberId: l.ctx.Value("memberId").(int64), Ids: req.Ids})

	return &types.CartItemDeleteResp{
		Code:    0,
		Message: "操作成功",
	}, nil
}
