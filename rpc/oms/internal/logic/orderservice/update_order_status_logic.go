package orderservicelogic

import (
	"context"

	"github.com/feihua/zero-admin/rpc/oms/internal/svc"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateOrderStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateOrderStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateOrderStatusLogic {
	return &UpdateOrderStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateOrderStatus 更新订单状态
func (l *UpdateOrderStatusLogic) UpdateOrderStatus(in *omsclient.UpdateOrderStatusReq) (*omsclient.UpdateOrderStatusResp, error) {
	// todo: add your logic here and delete this line

	return &omsclient.UpdateOrderStatusResp{}, nil
}
