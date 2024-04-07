package logic

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type OperateHistoryDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOperateHistoryDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) OperateHistoryDeleteLogic {
	return OperateHistoryDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OperateHistoryDeleteLogic) OperateHistoryDelete(req types.DeleteOperateHistoryReq) (*types.DeleteOperateHistoryResp, error) {
	_, err := l.svcCtx.OrderOperateHistoryService.OrderOperateHistoryDelete(l.ctx, &omsclient.OrderOperateHistoryDeleteReq{
		Ids: req.Ids,
	})

	if err != nil {
		logx.WithContext(l.ctx).Errorf("根据Id: %d,删除订单操作历史异常:%s", req.Ids, err.Error())
		return nil, errorx.NewDefaultError("删除订单操作历史签失败")
	}

	return &types.DeleteOperateHistoryResp{
		Code:    "000000",
		Message: "",
	}, nil
}
