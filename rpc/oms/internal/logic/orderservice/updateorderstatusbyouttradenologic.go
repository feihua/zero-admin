package orderservicelogic

import (
	"context"

	"zero-admin/rpc/oms/internal/svc"
	"zero-admin/rpc/oms/omsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateOrderStatusByOutTradeNoLogic
/*
Author: LiuFeiHua
Date: 2023/12/14 17:10
*/
type UpdateOrderStatusByOutTradeNoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateOrderStatusByOutTradeNoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateOrderStatusByOutTradeNoLogic {
	return &UpdateOrderStatusByOutTradeNoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateOrderStatusByOutTradeNo 更新订单状态
func (l *UpdateOrderStatusByOutTradeNoLogic) UpdateOrderStatusByOutTradeNo(in *omsclient.UpdateOrderStatusByOutTradeNoReq) (*omsclient.UpdateOrderStatusByOutTradeNoResp, error) {
	// todo: add your logic here and delete this line

	return &omsclient.UpdateOrderStatusByOutTradeNoResp{}, nil
}
