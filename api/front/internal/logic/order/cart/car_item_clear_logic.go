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

// CarItemClearLogic 清空购物车
/*
Author: LiuFeiHua
Date: 2025/6/20 10:21
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
	memberId, err := common.GetMemberId(l.ctx)
	if err != nil {
		return nil, err
	}

	_, err = l.svcCtx.CartItemService.DeleteCartItem(l.ctx, &omsclient.DeleteCartItemReq{MemberId: memberId})

	if err != nil {
		logc.Errorf(l.ctx, "清空购物车失败,参数memberId: %+v,异常：%s", memberId, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	return &types.CartItemClearResp{
		Code:    0,
		Message: "清空购物车成功",
	}, nil
}
