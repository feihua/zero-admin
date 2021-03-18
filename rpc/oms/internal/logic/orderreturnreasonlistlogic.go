package logic

import (
	"context"
	"fmt"

	"go-zero-admin/rpc/oms/internal/svc"
	"go-zero-admin/rpc/oms/oms"

	"github.com/tal-tech/go-zero/core/logx"
)

type OrderReturnReasonListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOrderReturnReasonListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderReturnReasonListLogic {
	return &OrderReturnReasonListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *OrderReturnReasonListLogic) OrderReturnReasonList(in *oms.OrderReturnReasonListReq) (*oms.OrderReturnReasonListResp, error) {
	all, _ := l.svcCtx.OmsOrderReturnReasonModel.FindAll(in.Current, in.PageSize)
	count, _ := l.svcCtx.OmsOrderReturnReasonModel.Count()

	var list []*oms.OrderReturnReasonListData
	for _, item := range *all {
		list = append(list, &oms.OrderReturnReasonListData{
			Id:         item.Id,
			Name:       item.Name,
			Sort:       item.Sort,
			Status:     item.Status,
			CreateTime: item.CreateTime.Format("2006-01-02 15:04:05"),
		})
	}

	fmt.Println(list)
	return &oms.OrderReturnReasonListResp{
		Total: count,
		List:  list,
	}, nil
}
