package cart

import (
	"context"
	"encoding/json"
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
	memberId, _ := l.ctx.Value("memberId").(json.Number).Int64()
	_, _ = l.svcCtx.CartItemService.CartItemDelete(l.ctx, &omsclient.CartItemDeleteReq{MemberId: memberId})

	return &types.CartItemClearResp{
		Code:    0,
		Message: "操作成功",
	}, nil
}
