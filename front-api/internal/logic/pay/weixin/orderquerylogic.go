package logic

import (
	"context"
	"zero-admin/rpc/pay/payclient"

	"zero-admin/front-api/internal/svc"
	"zero-admin/front-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrderQueryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOrderQueryLogic(ctx context.Context, svcCtx *svc.ServiceContext) OrderQueryLogic {
	return OrderQueryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OrderQueryLogic) OrderQuery(req types.OrderQueryReq) (*types.OrderQueryResp, error) {

	queryResp, err := l.svcCtx.Pay.OrderQuery(l.ctx, &payclient.OrderQueryReq{
		BusinessId: req.BusinessId,
		MerId:      req.MerId,
		PayType:    req.PayType,
	})

	if err != nil {
		return nil, err
	}

	return &types.OrderQueryResp{
		Code:       "000000",
		Message:    "查询成功",
		PayMessage: queryResp.PayMessage,
		PayStatus:  queryResp.PayStatus,
	}, nil
}
