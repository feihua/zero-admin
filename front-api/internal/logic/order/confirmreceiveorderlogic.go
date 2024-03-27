package order

import (
	"context"
	"encoding/json"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"

	"github.com/feihua/zero-admin/front-api/internal/svc"
	"github.com/feihua/zero-admin/front-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// ConfirmReceiveOrderLogic
/*
Author: LiuFeiHua
Date: 2023/12/7 9:35
*/
type ConfirmReceiveOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewConfirmReceiveOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ConfirmReceiveOrderLogic {
	return &ConfirmReceiveOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// ConfirmReceiveOrder 用户确认收货
func (l *ConfirmReceiveOrderLogic) ConfirmReceiveOrder(req *types.ConfirmReceiveOrderReq) (resp *types.ConfirmReceiveOrderResp, err error) {
	memberId, _ := l.ctx.Value("memberId").(json.Number).Int64()
	_, err = l.svcCtx.OrderService.OrderConfirm(l.ctx, &omsclient.OrderConfirmReq{
		MemberId: memberId,
		OrderId:  req.OrderId,
	})

	if err != nil {
		return nil, err
	}
	return &types.ConfirmReceiveOrderResp{
		Code:    0,
		Message: "操作成功",
	}, nil
}
