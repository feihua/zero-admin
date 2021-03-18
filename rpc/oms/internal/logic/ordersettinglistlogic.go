package logic

import (
	"context"
	"fmt"

	"go-zero-admin/rpc/oms/internal/svc"
	"go-zero-admin/rpc/oms/oms"

	"github.com/tal-tech/go-zero/core/logx"
)

type OrderSettingListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOrderSettingListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderSettingListLogic {
	return &OrderSettingListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *OrderSettingListLogic) OrderSettingList(in *oms.OrderSettingListReq) (*oms.OrderSettingListResp, error) {
	all, _ := l.svcCtx.OmsOrderSettingModel.FindAll(in.Current, in.PageSize)
	count, _ := l.svcCtx.OmsOrderSettingModel.Count()

	var list []*oms.OrderSettingListData
	for _, item := range *all {

		list = append(list, &oms.OrderSettingListData{
			Id:                  item.Id,
			FlashOrderOvertime:  item.FinishOvertime,
			NormalOrderOvertime: item.NormalOrderOvertime,
			ConfirmOvertime:     item.ConfirmOvertime,
			FinishOvertime:      item.FinishOvertime,
			CommentOvertime:     item.CommentOvertime,
		})
	}

	fmt.Println(list)
	return &oms.OrderSettingListResp{
		Total: count,
		List:  list,
	}, nil
}
