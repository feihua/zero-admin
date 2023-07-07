package cart

import (
	"context"
	"zero-admin/rpc/oms/omsclient"

	"zero-admin/front-api/internal/svc"
	"zero-admin/front-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CarItemClearLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCarItemClearLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CarItemClearLogic {
	return &CarItemClearLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CarItemClearLogic) CarItemClear() (resp *types.CartItemClearResp, err error) {
	_, _ = l.svcCtx.Oms.CartItemDelete(l.ctx, &omsclient.CartItemDeleteReq{MemberId: l.ctx.Value("memberId").(int64)})

	return &types.CartItemClearResp{
		Code:    0,
		Message: "操作成功",
	}, nil
}
