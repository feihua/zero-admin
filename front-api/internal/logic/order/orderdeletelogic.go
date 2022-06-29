package order

import (
	"context"
	"encoding/json"
	"zero-admin/rpc/oms/omsclient"

	"zero-admin/front-api/internal/svc"
	"zero-admin/front-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrderDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOrderDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) OrderDeleteLogic {
	return OrderDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OrderDeleteLogic) OrderDelete(req types.OrderDeleteReq) (resp *types.OrderDeleteResp, err error) {
	_, err = l.svcCtx.Oms.OrderDelete(l.ctx, &omsclient.OrderDeleteReq{
		UserId:  req.UserId,
		OrderId: req.OrderId,
	})

	if err != nil {
		reqStr, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("删除订单失败,参数：%s,响应：%s", reqStr, err.Error())
		return &types.OrderDeleteResp{
			Errno:  1,
			Errmsg: err.Error(),
		}, nil
	}

	return &types.OrderDeleteResp{
		Errno:  0,
		Errmsg: "删除订单成功",
	}, nil
}
