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

type UpdateCartItemQuantityLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateCartItemQuantityLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCartItemQuantityLogic {
	return &UpdateCartItemQuantityLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// UpdateCartItemQuantity 修改购物车中某个商品的数量
func (l *UpdateCartItemQuantityLogic) UpdateCartItemQuantity(req *types.UpdateCartItemQuantityReq) (resp *types.CartItemResp, err error) {
	memberId, err := common.GetMemberId(l.ctx)
	if err != nil {
		return nil, err
	}

	_, err = l.svcCtx.CartItemService.UpdateCartItemQuantity(l.ctx, &omsclient.UpdateCartItemQuantityReq{
		Quantity: req.Quantity,
		Id:       req.Id,
		MemberId: memberId,
	})

	if err != nil {
		logc.Errorf(l.ctx, "修改购物车中某个商品的数量失败,参数memberId: %+v,异常：%s", memberId, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	return &types.CartItemResp{
		Code:    0,
		Message: "修改购物车中某个商品的数量成功",
	}, nil
}
