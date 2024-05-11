package logic

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrderUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOrderUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) OrderUpdateLogic {
	return OrderUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OrderUpdateLogic) OrderUpdate(req types.UpdateOrderReq) (*types.UpdateOrderResp, error) {
	_, err := l.svcCtx.OrderService.OrderUpdate(l.ctx, &omsclient.OrderUpdateReq{
		Id:              req.Id,
		Status:          req.Status,
		DeliveryCompany: req.DeliveryCompany,
		DeliverySn:      req.DeliverySn,
		Note:            req.Note,
	})

	if err != nil {
		logc.Errorf(l.ctx, "更新订单信息失败,参数：%+v,响应：%s", req, err.Error())
		return nil, errorx.NewDefaultError("更新订单信息失败")
	}

	return &types.UpdateOrderResp{
		Code:    "000000",
		Message: "更新订单信息成功",
	}, nil
}
