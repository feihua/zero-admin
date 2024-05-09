package ordersettingservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/oms/gen/query"
	"github.com/feihua/zero-admin/rpc/oms/internal/svc"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/zeromicro/go-zero/core/logx"
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

func (l *OrderSettingListLogic) OrderSettingList(in *omsclient.OrderSettingListReq) (*omsclient.OrderSettingListResp, error) {
	q := query.OmsOrderSetting.WithContext(l.ctx)

	offset := (in.Current - 1) * in.PageSize
	result, err := q.Offset(int(offset)).Limit(int(in.PageSize)).Find()
	count, err := q.Count()

	if err != nil {
		logc.Errorf(l.ctx, "查询订单设置列表信息失败,参数：%+v,异常:%s", in, err.Error())
		return nil, err
	}

	var list []*omsclient.OrderSettingListData
	for _, item := range result {

		list = append(list, &omsclient.OrderSettingListData{
			Id:                  item.ID,
			FlashOrderOvertime:  item.FinishOvertime,
			NormalOrderOvertime: item.NormalOrderOvertime,
			ConfirmOvertime:     item.ConfirmOvertime,
			FinishOvertime:      item.FinishOvertime,
			CommentOvertime:     item.CommentOvertime,
		})
	}

	logc.Infof(l.ctx, "查询订单设置列表信息,参数：%+v,响应：%+v", in, list)
	return &omsclient.OrderSettingListResp{
		Total: count,
		List:  list,
	}, nil
}
