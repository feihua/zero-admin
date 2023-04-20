package order

import (
	"context"
	"encoding/json"
	"zero-admin/common/ctxdata"
	"zero-admin/rpc/oms/omsclient"

	"zero-admin/front-api/internal/svc"
	"zero-admin/front-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrderConfirmLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOrderConfirmLogic(ctx context.Context, svcCtx *svc.ServiceContext) OrderConfirmLogic {
	return OrderConfirmLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OrderConfirmLogic) OrderConfirm(req types.OrderConfirmReq) (resp *types.OrderConfirmResp, err error) {
	memberId := ctxdata.GetUidFromCtx(l.ctx)
	_, err = l.svcCtx.Oms.OrderConfirm(l.ctx, &omsclient.OrderConfirmReq{
		UserId:  memberId,
		OrderId: req.OrderId,
	})

	if err != nil {
		reqStr, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("确认订单失败,参数：%s,响应：%s", reqStr, err.Error())
		return &types.OrderConfirmResp{
			Errno:  1,
			Errmsg: err.Error(),
		}, nil
	}

	return &types.OrderConfirmResp{
		Errno:  0,
		Errmsg: "确认订单成功",
	}, nil
}
