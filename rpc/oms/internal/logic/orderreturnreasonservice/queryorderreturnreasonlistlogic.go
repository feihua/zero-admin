package orderreturnreasonservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/oms/gen/query"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/rpc/oms/internal/svc"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryOrderReturnReasonListLogic 查询退货原因表列表
/*
Author: LiuFeiHua
Date: 2024/6/12 10:19
*/
type QueryOrderReturnReasonListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryOrderReturnReasonListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryOrderReturnReasonListLogic {
	return &QueryOrderReturnReasonListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryOrderReturnReasonList 查询退货原因表列表
func (l *QueryOrderReturnReasonListLogic) QueryOrderReturnReasonList(in *omsclient.QueryOrderReturnReasonListReq) (*omsclient.QueryOrderReturnReasonListResp, error) {
	q := query.OmsOrderReturnReason.WithContext(l.ctx)
	if len(in.Name) > 0 {
		q = q.Where(query.OmsOrderReturnReason.Name.Like("%" + in.Name + "%"))
	}

	if in.Status != 2 {
		q = q.Where(query.OmsOrderReturnReason.Status.Eq(in.Status))
	}

	result, count, err := q.FindByPage(int((in.PageNum-1)*in.PageSize), int(in.PageSize))

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

	return &omsclient.QueryOrderReturnReasonListResp{
		Total: count,
		List:  list,
	}, nil

}
