package logic

import (
	"context"

	"zero-admin/rpc/oms/internal/svc"
	"zero-admin/rpc/oms/omsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrderListByMemberIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOrderListByMemberIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderListByMemberIdLogic {
	return &OrderListByMemberIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *OrderListByMemberIdLogic) OrderListByMemberId(in *omsclient.OrderListByMemberIdReq) (*omsclient.OrderListByMemberIdResp, error) {
	// todo: add your logic here and delete this line

	return &omsclient.OrderListByMemberIdResp{}, nil
}
