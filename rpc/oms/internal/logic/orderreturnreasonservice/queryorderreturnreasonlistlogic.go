package orderreturnreasonservicelogic

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

// QueryOrderReturnReasonListLogic 查询退货原因列表
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

// QueryOrderReturnReasonList 查询退货原因列表
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
		logc.Errorf(l.ctx, "查询退货原因详情失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("查询退货原因详情失败")
	}

	var list []*omsclient.OrderReturnReasonListData
	for _, item := range result {
		list = append(list, &omsclient.OrderReturnReasonListData{
			Id:         item.ID,                              //
			Name:       item.Name,                            // 退货类型
			Sort:       item.Sort,                            // 排序
			Status:     item.Status,                          // 状态：0->不启用；1->启用
			CreateTime: time_util.TimeToStr(item.CreateTime), // 创建时间
		})
	}

	return &omsclient.QueryOrderReturnReasonListResp{
		Total: count,
		List:  list,
	}, nil

}
