package orderdeliveryservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/oms/gen/model"
	"github.com/feihua/zero-admin/rpc/oms/gen/query"
	"github.com/feihua/zero-admin/rpc/oms/internal/svc"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
)

// AddOrderDeliveryLogic 添加订单收货地址
/*
Author: LiuFeiHua
Date: 2025/05/26 15:21:43
*/
type AddOrderDeliveryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddOrderDeliveryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddOrderDeliveryLogic {
	return &AddOrderDeliveryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// AddOrderDelivery 添加订单收货地址
func (l *AddOrderDeliveryLogic) AddOrderDelivery(in *omsclient.AddOrderDeliveryReq) (*omsclient.AddOrderDeliveryResp, error) {
	q := query.OmsOrderDelivery

	item := &model.OmsOrderDelivery{
		OrderID:          in.OrderId,          // 订单ID
		OrderNo:          in.OrderNo,          // 订单编号
		ReceiverName:     in.ReceiverName,     // 收货人姓名
		ReceiverPhone:    in.ReceiverPhone,    // 收货人电话
		ReceiverProvince: in.ReceiverProvince, // 省份
		ReceiverCity:     in.ReceiverCity,     // 城市
		ReceiverDistrict: in.ReceiverDistrict, // 区县
		ReceiverAddress:  in.ReceiverAddress,  // 详细地址
		DeliveryCompany:  in.DeliveryCompany,  // 物流公司
		DeliveryNo:       in.DeliveryNo,       // 物流单号
	}

	err := q.WithContext(l.ctx).Create(item)
	if err != nil {
		logc.Errorf(l.ctx, "添加订单收货地址失败,参数:%+v,异常:%s", item, err.Error())
		return nil, errors.New("添加订单收货地址失败")
	}

	return &omsclient.AddOrderDeliveryResp{}, nil
}
