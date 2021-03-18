package logic

import (
	"context"
	"fmt"

	"go-zero-admin/rpc/oms/internal/svc"
	"go-zero-admin/rpc/oms/oms"

	"github.com/tal-tech/go-zero/core/logx"
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

func (l *OrderOperateHistoryListLogic) OrderOperateHistoryList(in *oms.OrderOperateHistoryListReq) (*oms.OrderOperateHistoryListResp, error) {
	all, _ := l.svcCtx.OmsOrderOperateHistoryModel.FindAll(in.Current, in.PageSize)
	count, _ := l.svcCtx.OmsOrderOperateHistoryModel.Count()

	var list []*oms.OrderOperateHistoryListData
	for _, item := range *all {

		list = append(list, &oms.OrderOperateHistoryListData{
			Id:          item.Id,
			OrderId:     item.OrderId,
			OperateMan:  item.OperateMan,
			CreateTime:  item.CreateTime.Format("2006-01-02 15:04:05"),
			OrderStatus: item.OrderStatus,
			Note:        item.Note,
		})
	}

	fmt.Println(list)
	return &oms.OrderOperateHistoryListResp{
		Total: count,
		List:  list,
	}, nil
}
