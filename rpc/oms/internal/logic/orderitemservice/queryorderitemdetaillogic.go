package orderitemservicelogic

import (
	"context"

	"github.com/feihua/zero-admin/rpc/oms/internal/svc"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryOrderItemDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryOrderItemDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryOrderItemDetailLogic {
	return &QueryOrderItemDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询订单中所包含的商品详情
func (l *QueryOrderItemDetailLogic) QueryOrderItemDetail(in *omsclient.QueryOrderItemDetailReq) (*omsclient.QueryOrderItemDetailResp, error) {
	// todo: add your logic here and delete this line

	return &omsclient.QueryOrderItemDetailResp{}, nil
}
