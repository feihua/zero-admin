package orderoperationlogservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/oms/gen/query"
	"github.com/feihua/zero-admin/rpc/oms/internal/svc"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

// QueryOrderOperationLogDetailLogic 查询订单操作记录详情
/*
Author: LiuFeiHua
Date: 2025/06/30 15:49:27
*/
type QueryOrderOperationLogDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryOrderOperationLogDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryOrderOperationLogDetailLogic {
	return &QueryOrderOperationLogDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryOrderOperationLogDetail 查询订单操作记录详情
func (l *QueryOrderOperationLogDetailLogic) QueryOrderOperationLogDetail(in *omsclient.QueryOrderOperationLogDetailReq) (*omsclient.OrderOperationLogData, error) {
	item, err := query.OmsOrderOperationLog.WithContext(l.ctx).Where(query.OmsOrderOperationLog.ID.Eq(in.Id)).First()

	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		logc.Errorf(l.ctx, "订单操作记录不存在, 请求参数：%+v, 异常信息: %s", in, err.Error())
		return nil, errors.New("订单操作记录不存在")
	case err != nil:
		logc.Errorf(l.ctx, "查询订单操作记录异常, 请求参数：%+v, 异常信息: %s", in, err.Error())
		return nil, errors.New("查询订单操作记录异常")
	}

	data := &omsclient.OrderOperationLogData{
		Id:            item.ID,                                       // 主键ID
		OrderId:       item.OrderID,                                  // 订单ID
		OrderNo:       item.OrderNo,                                  // 订单编号
		OperationType: item.OperationType,                            // 操作类型：1-创建订单，2-支付订单，3-发货，4-确认收货，5-取消订单，6-退款
		OperatorId:    item.OperatorID,                               // 操作人ID
		OperatorType:  item.OperatorType,                             // 操作人类型：1-用户，2-系统，3-管理员
		OperatorNote:  item.OperatorNote,                             // 操作备注
		CreateTime:    item.CreateTime.Format("2006-01-02 15:04:05"), // 操作时间
	}

	return data, nil
}
