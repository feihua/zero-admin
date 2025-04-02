package cart

import (
	"context"
	"github.com/feihua/zero-admin/api/web/internal/logic/common"
	"github.com/feihua/zero-admin/pkg/errorx"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc/status"

	"github.com/feihua/zero-admin/api/web/internal/svc"
	"github.com/feihua/zero-admin/api/web/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// CartUpdateQuantityLogic
/*
Author: LiuFeiHua
Date: 2023/12/6 15:15
*/
type CartUpdateQuantityLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCartUpdateQuantityLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CartUpdateQuantityLogic {
	return &CartUpdateQuantityLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// CartUpdateQuantity 修改购物车中某个商品的数量
func (l *CartUpdateQuantityLogic) CartUpdateQuantity(req *types.CartItemUpdateQuantityReq) (resp *types.CartItemUpdateResp, err error) {
	memberId, err := common.GetMemberId(l.ctx)
	if err != nil {
		return nil, err
	}
	_, err = l.svcCtx.CartItemService.UpdateCartItemQuantity(l.ctx, &omsclient.UpdateCartItemQuantityReq{Id: req.Id, Quantity: req.Quantity, MemberId: memberId})

	if err != nil {
		logc.Errorf(l.ctx, "修改购物车中某个商品的数量失败,参数: %+v,异常：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	return &types.CartItemUpdateResp{
		Code:    0,
		Message: "操作成功",
	}, nil
}
