package logic

import (
	"context"
	"go-zero-admin/api/internal/common/errorx"
	"go-zero-admin/rpc/oms/omsclient"

	"go-zero-admin/api/internal/svc"
	"go-zero-admin/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type OperateHistoryAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOperateHistoryAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) OperateHistoryAddLogic {
	return OperateHistoryAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OperateHistoryAddLogic) OperateHistoryAdd(req types.AddOperateHistoryReq) (*types.AddOperateHistoryResp, error) {
	_, err := l.svcCtx.Oms.OrderOperateHistoryAdd(l.ctx, &omsclient.OrderOperateHistoryAddReq{
		OrderId:     req.OrderId,
		OperateMan:  req.OperateMan,
		CreateTime:  req.CreateTime,
		OrderStatus: req.OrderStatus,
		Note:        req.Note,
	})

	if err != nil {
		return nil, errorx.NewDefaultError("添加订单操作历史失败")
	}

	return &types.AddOperateHistoryResp{
		Code:    "000000",
		Message: "添加订单操作历史成功",
	}, nil
}
