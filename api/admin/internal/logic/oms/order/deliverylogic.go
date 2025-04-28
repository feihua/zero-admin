package order

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/api/admin/internal/common/res"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// DeliveryLogic 订单信息
/*
Author: LiuFeiHua
Date: 2024/5/14 17:04
*/
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

// Delivery 批量发货
func (l *DeliveryLogic) Delivery(req *types.DeliveryReq) (resp *types.BaseResp, err error) {
	_, err = l.svcCtx.OrderService.Delivery(l.ctx, &omsclient.DeliveryReq{
		OrderId:         req.OrderId,
		DeliveryCompany: req.DeliveryCompany,
		DeliverySn:      req.DeliverySn,
		OperateMan:      l.ctx.Value("userName").(string),
	})

	if err != nil {
		logc.Errorf(l.ctx, "批量发货失败,参数：%+v,响应：%s", req, err.Error())
		return nil, errorx.NewDefaultError("批量发货失败")
	}

	return res.Success()
}
