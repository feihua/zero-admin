package orderoperatehistoryservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/oms/gen/query"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/rpc/oms/internal/svc"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryOrderOperateHistoryListLogic 查询订单操作历史记录列表
/*
Author: LiuFeiHua
Date: 2024/6/12 10:14
*/
type QueryOrderOperateHistoryListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryOrderOperateHistoryListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryOrderOperateHistoryListLogic {
	return &QueryOrderOperateHistoryListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryOrderOperateHistoryList 查询订单操作历史记录列表
func (l *QueryOrderOperateHistoryListLogic) QueryOrderOperateHistoryList(in *omsclient.QueryOrderOperateHistoryListReq) (*omsclient.QueryOrderOperateHistoryListResp, error) {
	q := query.OmsOrderOperateHistory.WithContext(l.ctx)

	result, count, err := q.FindByPage(int((in.PageNum-1)*in.PageSize), int(in.PageSize))

	if err != nil {
		logc.Errorf(l.ctx, "查询订单操作历史列表信息失败,参数：%+v,异常:%s", in, err.Error())
		return nil, err
	}

	var list []*omsclient.OrderOperateHistoryListData
	for _, item := range result {

		list = append(list, &omsclient.OrderOperateHistoryListData{
			Id:          item.ID,
			OrderId:     item.OrderID,
			OperateMan:  item.OperateMan,
			CreateTime:  item.CreateTime.Format("2006-01-02 15:04:05"),
			OrderStatus: item.OrderStatus,
			Note:        *item.Note,
		})
	}

	return &omsclient.QueryOrderOperateHistoryListResp{
		Total: count,
		List:  list,
	}, nil

}
