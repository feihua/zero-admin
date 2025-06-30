package orderreturnservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/oms/gen/query"
	"github.com/feihua/zero-admin/rpc/oms/internal/svc"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
)

// DeleteOrderReturnLogic 删除退货/售后
/*
Author: LiuFeiHua
Date: 2025/06/30 15:49:28
*/
type DeleteOrderReturnLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteOrderReturnLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteOrderReturnLogic {
	return &DeleteOrderReturnLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// DeleteOrderReturn 删除退货/售后
func (l *DeleteOrderReturnLogic) DeleteOrderReturn(in *omsclient.DeleteOrderReturnReq) (*omsclient.OrderReturnResp, error) {
	q := query.OmsOrderReturn

	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Delete()

	if err != nil {
		logc.Errorf(l.ctx, "删除退货/售后失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("删除退货/售后失败")
	}

	return &omsclient.OrderReturnResp{}, nil
}
