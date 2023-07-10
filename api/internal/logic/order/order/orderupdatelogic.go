package logic

import (
	"context"
	"encoding/json"
	"zero-admin/api/internal/common/errorx"
	"zero-admin/rpc/oms/omsclient"

	"zero-admin/api/internal/svc"
	"zero-admin/api/internal/types"

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
		reqStr, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("更新订单信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, errorx.NewDefaultError("更新订单信息失败")
	}

	return &types.UpdateOrderResp{
		Code:    "000000",
		Message: "更新订单信息成功",
	}, nil
}
