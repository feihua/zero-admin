package orderreturnreasonservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/oms/gen/query"
	"github.com/feihua/zero-admin/rpc/oms/internal/svc"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"
	"github.com/zeromicro/go-zero/core/logc"

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
	q := query.OmsOrderReturnReason.WithContext(l.ctx)
	if len(in.Name) > 0 {
		q = q.Where(query.OmsOrderReturnReason.Name.Like("%" + in.Name + "%"))
	}

	if in.Status != 2 {
		q = q.Where(query.OmsOrderReturnReason.Status.Eq(in.Status))
	}

	offset := (in.Current - 1) * in.PageSize
	result, err := q.Offset(int(offset)).Limit(int(in.PageSize)).Find()
	count, err := q.Count()

	if err != nil {
		logc.Errorf(l.ctx, "查询退货原因列表信息失败,参数：%+v,异常:%s", in, err.Error())
		return nil, err
	}

	var list []*omsclient.OrderReturnReasonListData
	for _, item := range result {
		list = append(list, &omsclient.OrderReturnReasonListData{
			Id:         item.ID,
			Name:       item.Name,
			Sort:       item.Sort,
			Status:     item.Status,
			CreateTime: item.CreateTime.Format("2006-01-02 15:04:05"),
		})
	}

	logc.Infof(l.ctx, "查询退货原因列表信息,参数：%+v,响应：%+v", in, list)
	return &omsclient.OrderReturnReasonListResp{
		Total: count,
		List:  list,
	}, nil
}
