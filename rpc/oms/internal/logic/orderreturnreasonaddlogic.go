package logic

import (
	"context"
	"time"
	"zero-admin/rpc/model/omsmodel"
	"zero-admin/rpc/oms/internal/svc"
	"zero-admin/rpc/oms/oms"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrderReturnReasonAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOrderReturnReasonAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderReturnReasonAddLogic {
	return &OrderReturnReasonAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *OrderReturnReasonAddLogic) OrderReturnReasonAdd(in *oms.OrderReturnReasonAddReq) (*oms.OrderReturnReasonAddResp, error) {
	CreateTime, _ := time.Parse("2006-01-02 15:04:05", in.CreateTime)
	_, err := l.svcCtx.OmsOrderReturnReasonModel.Insert(omsmodel.OmsOrderReturnReason{
		Name:       in.Name,
		Sort:       in.Sort,
		Status:     in.Status,
		CreateTime: CreateTime,
	})
	if err != nil {
		return nil, err
	}

	return &oms.OrderReturnReasonAddResp{}, nil
}
