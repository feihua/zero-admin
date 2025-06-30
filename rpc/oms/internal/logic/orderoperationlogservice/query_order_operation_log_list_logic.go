package orderoperationlogservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/oms/gen/query"
	"github.com/feihua/zero-admin/rpc/oms/internal/svc"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
)

// QueryOrderOperationLogListLogic 查询订单操作记录列表
/*
Author: LiuFeiHua
Date: 2025/06/30 15:49:27
*/
type QueryOrderOperationLogListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryOrderOperationLogListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryOrderOperationLogListLogic {
	return &QueryOrderOperationLogListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryOrderOperationLogList 查询订单操作记录列表
func (l *QueryOrderOperationLogListLogic) QueryOrderOperationLogList(in *omsclient.QueryOrderOperationLogListReq) (*omsclient.QueryOrderOperationLogListResp, error) {
	orderOperationLog := query.OmsOrderOperationLog
	q := orderOperationLog.WithContext(l.ctx)
	if in.OrderId != 2 {
		q = q.Where(orderOperationLog.OrderID.Eq(in.OrderId))
	}
	if len(in.OrderNo) > 0 {
		q = q.Where(orderOperationLog.OrderNo.Like("%" + in.OrderNo + "%"))
	}
	if in.OperationType != 2 {
		q = q.Where(orderOperationLog.OperationType.Eq(in.OperationType))
	}
	if in.OperatorId != 2 {
		q = q.Where(orderOperationLog.OperatorID.Eq(in.OperatorId))
	}
	if in.OperatorType != 2 {
		q = q.Where(orderOperationLog.OperatorType.Eq(in.OperatorType))
	}

	result, count, err := q.FindByPage(int((in.PageNum-1)*in.PageSize), int(in.PageSize))

	if err != nil {
		logc.Errorf(l.ctx, "查询订单操作记录列表失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("查询订单操作记录列表失败")
	}

	var list []*omsclient.OrderOperationLogData

	for _, item := range result {
		list = append(list, &omsclient.OrderOperationLogData{
			Id:            item.ID,                                       // 主键ID
			OrderId:       item.OrderID,                                  // 订单ID
			OrderNo:       item.OrderNo,                                  // 订单编号
			OperationType: item.OperationType,                            // 操作类型：1-创建订单，2-支付订单，3-发货，4-确认收货，5-取消订单，6-退款
			OperatorId:    item.OperatorID,                               // 操作人ID
			OperatorType:  item.OperatorType,                             // 操作人类型：1-用户，2-系统，3-管理员
			OperatorNote:  item.OperatorNote,                             // 操作备注
			CreateTime:    item.CreateTime.Format("2006-01-02 15:04:05"), // 操作时间

		})
	}

	return &omsclient.QueryOrderOperationLogListResp{
		Total: count,
		List:  list,
	}, nil
}
