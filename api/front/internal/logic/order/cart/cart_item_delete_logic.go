package cart

import (
	"context"
	"github.com/feihua/zero-admin/api/front/internal/logic/common"
	"github.com/feihua/zero-admin/pkg/errorx"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc/status"

	"github.com/feihua/zero-admin/api/front/internal/svc"
	"github.com/feihua/zero-admin/api/front/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// CartItemDeleteLogic 删除购物车中的某个商品
/*
Author: LiuFeiHua
Date: 2025/6/20 10:25
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
func (l *CartItemDeleteLogic) CartItemDelete(req *types.CartItemDeleteReq) (resp *types.CartItemResp, err error) {
	memberId, err := common.GetMemberId(l.ctx)
	if err != nil {
		return nil, err
	}
	_, err = l.svcCtx.CartItemService.DeleteCartItem(l.ctx, &omsclient.DeleteCartItemReq{MemberId: memberId, Ids: req.Ids})

	if err != nil {
		logc.Errorf(l.ctx, "删除购物车中的某个商品失败,参数: %+v,异常：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}
	return &types.CartItemResp{
		Code:    0,
		Message: "操作成功",
	}, nil
}
