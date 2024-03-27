package orderreturnreasonservicelogic

import (
	"context"
	"encoding/json"
	"github.com/feihua/zero-admin/rpc/oms/internal/svc"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"

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

func (l *OrderReturnReasonListLogic) OrderReturnReasonList(in *omsclient.OrderReturnReasonListReq) (*omsclient.OrderReturnReasonListResp, error) {
	all, err := l.svcCtx.OmsOrderReturnReasonModel.FindAll(l.ctx, in)
	count, _ := l.svcCtx.OmsOrderReturnReasonModel.Count(l.ctx, in)

	if err != nil {
		reqStr, _ := json.Marshal(in)
		logx.WithContext(l.ctx).Errorf("查询退货原因列表信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, err
	}

	var list []*omsclient.OrderReturnReasonListData
	for _, item := range *all {
		list = append(list, &omsclient.OrderReturnReasonListData{
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
	return &omsclient.OrderReturnReasonListResp{
		Total: count,
		List:  list,
	}, nil
}
