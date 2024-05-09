package orderreturnreasonservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/oms/gen/query"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"

	"github.com/feihua/zero-admin/rpc/oms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

// OrderReturnReasonDeleteLogic 退货原因
/*
Author: LiuFeiHua
Date: 2024/5/8 9:33
*/
type OrderReturnReasonDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOrderReturnReasonDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderReturnReasonDeleteLogic {
	return &OrderReturnReasonDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// OrderReturnReasonDelete 删除退货原因
func (l *OrderReturnReasonDeleteLogic) OrderReturnReasonDelete(in *omsclient.OrderReturnReasonDeleteReq) (*omsclient.OrderReturnReasonDeleteResp, error) {
	q := query.OmsOrderReturnReason
	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Delete()

	if err != nil {
		return nil, err
	}

	return &omsclient.OrderReturnReasonDeleteResp{}, nil
}
