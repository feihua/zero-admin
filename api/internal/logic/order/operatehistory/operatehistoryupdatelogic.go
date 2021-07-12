package logic

import (
	"context"
	"encoding/json"
	"go-zero-admin/api/internal/common/errorx"
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
		CreateTime:  req.CreateTime,
		OrderStatus: req.OrderStatus,
		Note:        req.Note,
	})

	if err != nil {
		reqStr, _ := json.Marshal(req)
		logx.Errorf("更新订单操作历史参数:%s,异常:%s", reqStr, err.Error())
		return nil, errorx.NewDefaultError("更新订单操作历史失败")
	}

	return &types.UpdateOperateHistoryResp{
		Code:    "000000",
		Message: "更新订单操作历史成功",
	}, nil
}
