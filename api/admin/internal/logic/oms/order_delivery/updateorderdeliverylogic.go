// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package order_delivery

import (
	"context"

	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc/status"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateOrderDeliveryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateOrderDeliveryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateOrderDeliveryLogic {
	return &UpdateOrderDeliveryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateOrderDeliveryLogic) UpdateOrderDelivery(req *types.UpdateOrderDeliveryReq) (resp *types.UpdateOrderDeliveryResp, err error) {
	_, err = l.svcCtx.OrderDeliveryService.UpdateOrderDelivery(l.ctx, &omsclient.UpdateOrderDeliveryReq{
		Id:               req.Id,               //
		OrderId:          req.OrderId,          // 订单ID
		OrderNo:          req.OrderNo,          // 订单编号
		ReceiverName:     req.ReceiverName,     // 收货人姓名
		ReceiverPhone:    req.ReceiverPhone,    // 收货人电话
		ReceiverProvince: req.ReceiverProvince, // 省份
		ReceiverCity:     req.ReceiverCity,     // 城市
		ReceiverDistrict: req.ReceiverDistrict, // 区县
		ReceiverAddress:  req.ReceiverAddress,  // 详细地址
		DeliveryCompany:  req.DeliveryCompany,  // 物流公司
		DeliveryNo:       req.DeliveryNo,       // 物流单号
	})

	if err != nil {
		logc.Errorf(l.ctx, "更新订单收货地址失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	return &types.UpdateOrderDeliveryResp{
		Code:    "000000",
		Message: "更新订单收货地址成功",
	}, nil
}
