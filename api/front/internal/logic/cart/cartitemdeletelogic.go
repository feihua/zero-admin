package cart

import (
	"context"
	"encoding/json"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"

	"github.com/feihua/zero-admin/api/front/internal/svc"
	"github.com/feihua/zero-admin/api/front/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// CartItemDeleteLogic
/*
Author: LiuFeiHua
Date: 2023/12/6 14:35
*/
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

// CartItemDelete 删除购物车中的某个商品
func (l *CartItemDeleteLogic) CartItemDelete(req *types.CartItemDeleteReq) (resp *types.CartItemDeleteResp, err error) {
	memberId, _ := l.ctx.Value("memberId").(json.Number).Int64()
	_, _ = l.svcCtx.CartItemService.CartItemDelete(l.ctx, &omsclient.CartItemDeleteReq{MemberId: memberId, Ids: req.Ids})

	return &types.CartItemDeleteResp{
		Code:    0,
		Message: "操作成功",
	}, nil
}
