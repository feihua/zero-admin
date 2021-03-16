package logic

import (
	"context"
	"go-zero-admin/rpc/oms/omsclient"

	"go-zero-admin/api/internal/svc"
	"go-zero-admin/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type OperateHistoryUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOperateHistoryUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) OperateHistoryUpdateLogic {
	return OperateHistoryUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OperateHistoryUpdateLogic) OperateHistoryUpdate(req types.UpdateOperateHistoryReq) (*types.UpdateOperateHistoryResp, error) {
	_, err := l.svcCtx.Oms.OrderOperateHistoryUpdate(l.ctx, &omsclient.OrderOperateHistoryUpdateReq{
		Id:          req.Id,
		OrderId:     req.OrderId,
		OperateMan:  req.OperateMan,
		CreateTime:  req.CreateTime.Format("2006-01-02 15:04:05"),
		OrderStatus: req.OrderStatus,
		Note:        req.Note,
	})

	if err != nil {
		return nil, err
	}

	return &types.UpdateOperateHistoryResp{
		Code:    "000000",
		Message: "",
	}, nil
}
