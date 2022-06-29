package order

import (
	"context"
	"encoding/json"
	"zero-admin/rpc/oms/omsclient"

	"zero-admin/front-api/internal/svc"
	"zero-admin/front-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrderCancelLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOrderCancelLogic(ctx context.Context, svcCtx *svc.ServiceContext) OrderCancelLogic {
	return OrderCancelLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OrderCancelLogic) OrderCancel(req types.OrderCancelReq) (resp *types.OrderCancelResp, err error) {
	_, err = l.svcCtx.Oms.OrderCancel(l.ctx, &omsclient.OrderCancelReq{
		UserId:  req.UserId,
		OrderId: req.OrderId,
	})

	if err != nil {
		reqStr, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("取消订单失败,参数：%s,响应：%s", reqStr, err.Error())
		return &types.OrderCancelResp{
			Errno:  1,
			Errmsg: err.Error(),
		}, nil
	}

	return &types.OrderCancelResp{
		Errno:  0,
		Errmsg: "取消订单成功",
	}, nil
}
