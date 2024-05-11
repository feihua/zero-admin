package logic

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
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
	_, err := l.svcCtx.OrderOperateHistoryService.OrderOperateHistoryUpdate(l.ctx, &omsclient.OrderOperateHistoryUpdateReq{
		Id:          req.Id,
		OrderId:     req.OrderId,
		OperateMan:  req.OperateMan,
		CreateTime:  req.CreateTime,
		OrderStatus: req.OrderStatus,
		Note:        req.Note,
	})

	if err != nil {
		logc.Errorf(l.ctx, "更新订单操作历史信息失败,参数：%+v,响应：%s", req, err.Error())
		return nil, errorx.NewDefaultError("更新订单操作历史失败")
	}

	return &types.UpdateOperateHistoryResp{
		Code:    "000000",
		Message: "更新订单操作历史成功",
	}, nil
}
