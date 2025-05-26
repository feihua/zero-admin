package orderreturnreasonservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/oms/gen/query"
	"github.com/feihua/zero-admin/rpc/oms/internal/svc"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
)

// DeleteOrderReturnReasonLogic 删除退货原因
/*
Author: LiuFeiHua
Date: 2025/05/26 15:21:44
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

// DeleteOrderReturnReason 删除退货原因
func (l *DeleteOrderReturnReasonLogic) DeleteOrderReturnReason(in *omsclient.DeleteOrderReturnReasonReq) (*omsclient.DeleteOrderReturnReasonResp, error) {
	q := query.OmsOrderReturnReason

	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Delete()

	if err != nil {
		logc.Errorf(l.ctx, "删除退货原因失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("删除退货原因失败")
	}

	return &omsclient.DeleteOrderReturnReasonResp{}, nil
}
