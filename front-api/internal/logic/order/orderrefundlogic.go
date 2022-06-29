package order

import (
	"context"
	"encoding/json"
	"zero-admin/rpc/oms/omsclient"

	"zero-admin/front-api/internal/svc"
	"zero-admin/front-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrderRefundLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOrderRefundLogic(ctx context.Context, svcCtx *svc.ServiceContext) OrderRefundLogic {
	return OrderRefundLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OrderRefundLogic) OrderRefund(req types.OrderRefundReq) (resp *types.OrderRefundResp, err error) {
	_, err = l.svcCtx.Oms.OrderRefund(l.ctx, &omsclient.OrderRefundReq{
		UserId:  req.UserId,
		OrderId: req.OrderId,
	})

	if err != nil {
		reqStr, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("订单申请退款失败,参数：%s,响应：%s", reqStr, err.Error())
		return &types.OrderRefundResp{
			Errno:  1,
			Errmsg: err.Error(),
		}, nil
	}

	return &types.OrderRefundResp{
		Errno:  0,
		Errmsg: "订单申请退款成功",
	}, nil
}
