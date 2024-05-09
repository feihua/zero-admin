package logic

import (
	"context"
	"encoding/json"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
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
	_, err := l.svcCtx.OrderOperateHistoryService.OrderOperateHistoryAdd(l.ctx, &omsclient.OrderOperateHistoryAddReq{
		OrderId:     req.OrderId,
		OperateMan:  req.OperateMan,
		OrderStatus: req.OrderStatus,
		Note:        req.Note,
	})

	if err != nil {
		reqStr, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("添加订单操作历史信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, errorx.NewDefaultError("添加订单操作历史失败")
	}

	return &types.AddOperateHistoryResp{
		Code:    "000000",
		Message: "添加订单操作历史成功",
	}, nil
}
