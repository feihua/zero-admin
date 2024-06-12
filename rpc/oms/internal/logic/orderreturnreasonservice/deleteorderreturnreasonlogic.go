package orderreturnreasonservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/oms/gen/query"

	"github.com/feihua/zero-admin/rpc/oms/internal/svc"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// DeleteOrderReturnReasonLogic 删除退货原因表
/*
Author: LiuFeiHua
Date: 2024/6/12 10:18
*/
type DeleteOrderReturnReasonLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteOrderReturnReasonLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteOrderReturnReasonLogic {
	return &DeleteOrderReturnReasonLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// DeleteOrderReturnReason 删除退货原因表
func (l *DeleteOrderReturnReasonLogic) DeleteOrderReturnReason(in *omsclient.DeleteOrderReturnReasonReq) (*omsclient.DeleteOrderReturnReasonResp, error) {
	q := query.OmsOrderReturnReason
	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Delete()

	if err != nil {
		return nil, err
	}

	return &omsclient.DeleteOrderReturnReasonResp{}, nil
}
