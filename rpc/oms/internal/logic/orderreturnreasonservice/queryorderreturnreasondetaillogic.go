package orderreturnreasonservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/pkg/time_util"
	"github.com/feihua/zero-admin/rpc/oms/gen/query"
	"github.com/feihua/zero-admin/rpc/oms/internal/svc"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
)

// QueryOrderReturnReasonDetailLogic 查询退货原因表详情
/*
Author: 刘飞华
Date: 2025/02/07 10:05:47
*/
type QueryOrderReturnReasonDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryOrderReturnReasonDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryOrderReturnReasonDetailLogic {
	return &QueryOrderReturnReasonDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryOrderReturnReasonDetail 查询退货原因表详情
func (l *QueryOrderReturnReasonDetailLogic) QueryOrderReturnReasonDetail(in *omsclient.QueryOrderReturnReasonDetailReq) (*omsclient.QueryOrderReturnReasonDetailResp, error) {
	item, err := query.OmsOrderReturnReason.WithContext(l.ctx).Where(query.OmsOrderReturnReason.ID.Eq(in.Id)).First()

	if err != nil {
		logc.Errorf(l.ctx, "查询退货原因表详情失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("查询退货原因表详情失败")
	}

	data := &omsclient.QueryOrderReturnReasonDetailResp{
		Id:         item.ID,                              //
		Name:       item.Name,                            // 退货类型
		Sort:       item.Sort,                            // 排序
		Status:     item.Status,                          // 状态：0->不启用；1->启用
		CreateTime: time_util.TimeToStr(item.CreateTime), // 创建时间

	}

	logc.Infof(l.ctx, "查询退货原因表详情,参数：%+v,响应：%+v", in, data)
	return data, nil
}
