package order_main

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeliveryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeliveryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeliveryLogic {
	return &DeliveryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeliveryLogic) Delivery(req *types.DeliveryReq) (resp *types.BaseResp, err error) {
	userId, err := common.GetUserId(l.ctx)
	if err != nil {
		return nil, err
	}
	_, err = l.svcCtx.OrderService.Delivery(l.ctx, &omsclient.DeliveryReq{
		OrderId:    req.OrderId,
		DeliverySn: req.DeliverySn,
		OperatorId: userId,
	})

	if err != nil {
		logc.Errorf(l.ctx, "批量发货失败,参数：%+v,响应：%s", req, err.Error())
		return nil, errorx.NewDefaultError("批量发货失败")
	}

	return
}
