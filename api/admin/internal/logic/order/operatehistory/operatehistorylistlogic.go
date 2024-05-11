package logic

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/zeromicro/go-zero/core/logx"
)

type OperateHistoryListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOperateHistoryListLogic(ctx context.Context, svcCtx *svc.ServiceContext) OperateHistoryListLogic {
	return OperateHistoryListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OperateHistoryListLogic) OperateHistoryList(req types.ListOperateHistoryReq) (*types.ListOperateHistoryResp, error) {
	resp, err := l.svcCtx.OrderOperateHistoryService.OrderOperateHistoryList(l.ctx, &omsclient.OrderOperateHistoryListReq{
		Current:  req.Current,
		PageSize: req.PageSize,
	})

	if err != nil {
		logc.Errorf(l.ctx, "参数: %+v,查询订单操作历史列表异常:%s", req, err.Error())
		return nil, errorx.NewDefaultError("查询订单操作历史失败")
	}

	var list []*types.ListtOperateHistoryData

	for _, item := range resp.List {
		list = append(list, &types.ListtOperateHistoryData{
			Id:          item.Id,
			OrderId:     item.OrderId,
			OperateMan:  item.OperateMan,
			CreateTime:  item.CreateTime,
			OrderStatus: item.OrderStatus,
			Note:        item.Note,
		})
	}

	return &types.ListOperateHistoryResp{
		Current:  req.Current,
		Data:     list,
		PageSize: req.PageSize,
		Success:  true,
		Total:    resp.Total,
		Code:     "000000",
		Message:  "查询订单操作历史成功",
	}, nil
}
