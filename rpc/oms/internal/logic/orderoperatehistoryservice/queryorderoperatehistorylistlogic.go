package orderoperatehistoryservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/pkg/time_util"
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
		logc.Errorf(l.ctx, "查询订单操作历史记录列表失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("查询订单操作历史记录列表失败")
	}

	var list []*omsclient.OrderOperateHistoryListData
	for _, item := range result {

		list = append(list, &omsclient.OrderOperateHistoryListData{
			Id:          item.ID,                              //
			OrderId:     item.OrderID,                         // 订单id
			OperateMan:  item.OperateMan,                      // 操作人：用户；系统；后台管理员
			CreateTime:  time_util.TimeToStr(item.CreateTime), // 创建时间
			OrderStatus: item.OrderStatus,                     // 订单状态：0->待付款；1->待发货；2->已发货；3->已完成；4->已关闭；5->无效订单
			Note:        item.Note,                            // 备注
		})
	}

	return &omsclient.QueryOrderOperateHistoryListResp{
		Total: count,
		List:  list,
	}, nil

}
