package logic

import (
	"context"
	"time"
	"zero-admin/rpc/model/omsmodel"

	"zero-admin/rpc/oms/internal/svc"
	"zero-admin/rpc/oms/oms"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrderReturnReasonUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOrderReturnReasonUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderReturnReasonUpdateLogic {
	return &OrderReturnReasonUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *OrderReturnReasonUpdateLogic) OrderReturnReasonUpdate(in *oms.OrderReturnReasonUpdateReq) (*oms.OrderReturnReasonUpdateResp, error) {
	CreateTime, _ := time.Parse("2006-01-02 15:04:05", in.CreateTime)
	err := l.svcCtx.OmsOrderReturnReasonModel.Update(omsmodel.OmsOrderReturnReason{
		Id:         in.Id,
		Name:       in.Name,
		Sort:       in.Sort,
		Status:     in.Status,
		CreateTime: CreateTime,
	})
	if err != nil {
		return nil, err
	}

	return &oms.OrderReturnReasonUpdateResp{}, nil
}
