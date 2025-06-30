package order_main

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc/status"

	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateOrderMainLogic 更新订单
/*
Author: LiuFeiHua
Date: 2025/07/01 10:06:43
*/
type UpdateOrderMainLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateOrderMainLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateOrderMainLogic {
	return &UpdateOrderMainLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// UpdateOrderMain 更新订单
func (l *UpdateOrderMainLogic) UpdateOrderMain(req *types.UpdateOrderMainReq) (resp *types.BaseResp, err error) {
	// userId, err := common.GetUserId(l.ctx)
	// if err != nil {
	// 	return nil, err
	// }
	_, err = l.svcCtx.OrderService.UpdateOrder(l.ctx, &omsclient.UpdateOrderReq{
		Id:                 req.Id,                     //
		OrderStatus:        req.OrderStatus,            // 订单状态：1-待支付,2-已支付,3-已发货,4-已完成,5-已取消,6-已退款,7-售后中
		FreightAmount:      float32(req.FreightAmount), // 运费金额
		ExpressOrderNumber: req.ExpressOrderNumber,     // 快递单号
		Remark:             req.Remark,                 // 订单备注
	})

	if err != nil {
		logc.Errorf(l.ctx, "更新订单失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	return &types.BaseResp{
		Code:    "000000",
		Message: "更新订单成功",
	}, nil
}
