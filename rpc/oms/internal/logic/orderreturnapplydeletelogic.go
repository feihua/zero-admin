package logic

import (
	"context"

	"go-zero-admin/rpc/oms/internal/svc"
	"go-zero-admin/rpc/oms/oms"

	"github.com/tal-tech/go-zero/core/logx"
)

type OrderReturnApplyDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOrderReturnApplyDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderReturnApplyDeleteLogic {
	return &OrderReturnApplyDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *OrderReturnApplyDeleteLogic) OrderReturnApplyDelete(in *oms.OrderReturnApplyDeleteReq) (*oms.OrderReturnApplyDeleteResp, error) {
	err := l.svcCtx.OmsOrderReturnApplyModel.Delete(in.Id)

	if err != nil {
		return nil, err
	}

	return &oms.OrderReturnApplyDeleteResp{}, nil
}
