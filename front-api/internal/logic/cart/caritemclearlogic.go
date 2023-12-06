package cart

import (
	"context"
	"encoding/json"
	"zero-admin/rpc/oms/omsclient"

	"zero-admin/front-api/internal/svc"
	"zero-admin/front-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// CarItemClearLogic
/*
Author: LiuFeiHua
Date: 2023/12/6 14:34
*/
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

// CarItemClear 清空购物车
func (l *CarItemClearLogic) CarItemClear() (resp *types.CartItemClearResp, err error) {
	memberId, _ := l.ctx.Value("memberId").(json.Number).Int64()
	_, _ = l.svcCtx.CartItemService.CartItemClear(l.ctx, &omsclient.CartItemClearReq{MemberId: memberId})

	return &types.CartItemClearResp{
		Code:    0,
		Message: "清空购物车成功",
	}, nil
}
