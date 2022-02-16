package logic

import (
	"context"
	"encoding/json"
	"zero-admin/rpc/oms/internal/svc"
	"zero-admin/rpc/oms/oms"

	"github.com/zeromicro/go-zero/core/logx"
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
	all, err := l.svcCtx.OmsOrderReturnReasonModel.FindAll(in.Current, in.PageSize)
	count, _ := l.svcCtx.OmsOrderReturnReasonModel.Count()

	if err != nil {
		reqStr, _ := json.Marshal(in)
		logx.WithContext(l.ctx).Errorf("查询退货原因列表信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, err
	}

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

	reqStr, _ := json.Marshal(in)
	listStr, _ := json.Marshal(list)
	logx.WithContext(l.ctx).Infof("查询退货原因列表信息,参数：%s,响应：%s", reqStr, listStr)
	return &oms.OrderReturnReasonListResp{
		Total: count,
		List:  list,
	}, nil
}
