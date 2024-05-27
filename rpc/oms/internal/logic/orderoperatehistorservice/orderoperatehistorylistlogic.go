package orderoperatehistorservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/oms/gen/query"
	"github.com/feihua/zero-admin/rpc/oms/internal/svc"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrderOperateHistoryListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOrderOperateHistoryListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderOperateHistoryListLogic {
	return &OrderOperateHistoryListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *OrderOperateHistoryListLogic) OrderOperateHistoryList(in *omsclient.OrderOperateHistoryListReq) (*omsclient.OrderOperateHistoryListResp, error) {
	q := query.OmsOrderOperateHistory.WithContext(l.ctx)

	result, count, err := q.FindByPage(int((in.Current-1)*in.PageSize), int(in.PageSize))

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

	logc.Infof(l.ctx, "查询订单操作历史列表信息,参数：%+v,响应：%+v", in, list)
	return &omsclient.OrderOperateHistoryListResp{
		Total: count,
		List:  list,
	}, nil
}
