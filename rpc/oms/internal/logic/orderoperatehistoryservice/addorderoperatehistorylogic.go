package orderoperatehistoryservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/oms/gen/model"
	"github.com/feihua/zero-admin/rpc/oms/gen/query"
	"github.com/feihua/zero-admin/rpc/oms/internal/svc"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// AddOrderOperateHistoryLogic 添加订单操作历史记录
/*
Author: LiuFeiHua
Date: 2024/6/12 10:13
*/
type AddOrderOperateHistoryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddOrderOperateHistoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddOrderOperateHistoryLogic {
	return &AddOrderOperateHistoryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// AddOrderOperateHistory 添加订单操作历史记录
func (l *AddOrderOperateHistoryLogic) AddOrderOperateHistory(in *omsclient.AddOrderOperateHistoryReq) (*omsclient.AddOrderOperateHistoryResp, error) {
	err := query.OmsOrderOperateHistory.WithContext(l.ctx).Create(&model.OmsOrderOperateHistory{
		OrderID:     in.OrderId,     // 订单id
		OperateMan:  in.OperateMan,  // 操作人：用户；系统；后台管理员
		OrderStatus: in.OrderStatus, // 订单状态：0->待付款；1->待发货；2->已发货；3->已完成；4->已关闭；5->无效订单
		Note:        in.Note,        // 备注
	})

	if err != nil {
		return nil, err
	}

	return &omsclient.AddOrderOperateHistoryResp{}, nil
}
