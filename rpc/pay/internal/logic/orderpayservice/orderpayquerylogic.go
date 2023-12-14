package orderpayservicelogic

import (
	"context"

	"zero-admin/rpc/pay/internal/svc"
	"zero-admin/rpc/pay/payclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// OrderPayQueryLogic
/*
Author: LiuFeiHua
Date: 2023/12/14 14:13
*/
type OrderPayQueryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOrderPayQueryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderPayQueryLogic {
	return &OrderPayQueryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// OrderPayQuery 订单状态查询
func (l *OrderPayQueryLogic) OrderPayQuery(in *payclient.OrderPayQueryReq) (*payclient.OrderPayQueryResp, error) {
	// todo: add your logic here and delete this line

	return &payclient.OrderPayQueryResp{}, nil
}
