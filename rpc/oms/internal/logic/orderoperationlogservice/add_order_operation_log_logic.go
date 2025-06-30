package orderoperationlogservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/oms/gen/model"
	"github.com/feihua/zero-admin/rpc/oms/gen/query"
	"github.com/feihua/zero-admin/rpc/oms/internal/svc"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
)

// AddOrderOperationLogLogic 添加订单操作记录
/*
Author: LiuFeiHua
Date: 2025/06/30 15:49:27
*/
type AddOrderOperationLogLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddOrderOperationLogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddOrderOperationLogLogic {
	return &AddOrderOperationLogLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// AddOrderOperationLog 添加订单操作记录
func (l *AddOrderOperationLogLogic) AddOrderOperationLog(in *omsclient.AddOrderOperationLogReq) (*omsclient.AddOrderOperationLogResp, error) {
	q := query.OmsOrderOperationLog

	item := &model.OmsOrderOperationLog{
		OrderID:       in.OrderId,       // 订单ID
		OrderNo:       in.OrderNo,       // 订单编号
		OperationType: in.OperationType, // 操作类型：1-创建订单，2-支付订单，3-发货，4-确认收货，5-取消订单，6-退款
		OperatorID:    in.OperatorId,    // 操作人ID
		OperatorType:  in.OperatorType,  // 操作人类型：1-用户，2-系统，3-管理员
		OperatorNote:  in.OperatorNote,  // 操作备注
	}

	err := q.WithContext(l.ctx).Create(item)
	if err != nil {
		logc.Errorf(l.ctx, "添加订单操作记录失败,参数:%+v,异常:%s", item, err.Error())
		return nil, errors.New("添加订单操作记录失败")
	}

	return &omsclient.AddOrderOperationLogResp{}, nil
}
