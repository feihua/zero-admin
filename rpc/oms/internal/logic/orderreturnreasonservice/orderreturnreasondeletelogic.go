package orderreturnreasonservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"

	"github.com/feihua/zero-admin/rpc/oms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrderReturnReasonDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOrderReturnReasonDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderReturnReasonDeleteLogic {
	return &OrderReturnReasonDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *OrderReturnReasonDeleteLogic) OrderReturnReasonDelete(in *omsclient.OrderReturnReasonDeleteReq) (*omsclient.OrderReturnReasonDeleteResp, error) {
	err := l.svcCtx.OmsOrderReturnReasonModel.DeleteByIds(l.ctx, in.Ids)

	if err != nil {
		return nil, err
	}

	return &omsclient.OrderReturnReasonDeleteResp{}, nil
}
