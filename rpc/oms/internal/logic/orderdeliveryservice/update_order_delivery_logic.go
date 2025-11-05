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
	"gorm.io/gorm"
)

// UpdateOrderDeliveryLogic 更新订单收货地址
/*
Author: LiuFeiHua
Date: 2025/05/26 15:21:44
*/
type UpdateOrderDeliveryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateOrderDeliveryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateOrderDeliveryLogic {
	return &UpdateOrderDeliveryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateOrderDelivery 更新订单收货地址
func (l *UpdateOrderDeliveryLogic) UpdateOrderDelivery(in *omsclient.UpdateOrderDeliveryReq) (*omsclient.UpdateOrderDeliveryResp, error) {
	q := query.OmsOrderDelivery.WithContext(l.ctx)

	// 1.根据订单收货地址id查询订单收货地址是否已存在
	_, err := q.Where(query.OmsOrderDelivery.ID.Eq(in.Id)).First()

	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		logc.Errorf(l.ctx, "订单收货地址不存在, 请求参数：%+v, 异常信息: %s", in, err.Error())
		return nil, errors.New("订单收货地址不存在")
	case err != nil:
		logc.Errorf(l.ctx, "查询订单收货地址异常, 请求参数：%+v, 异常信息: %s", in, err.Error())
		return nil, errors.New("查询订单收货地址异常")
	}

	item := &model.OmsOrderDelivery{
		ID:               in.Id,               //
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

	// 2.订单收货地址存在时,则直接更新订单收货地址
	_, err = q.Where(query.OmsOrderDelivery.ID.Eq(in.Id)).Updates(item)

	if err != nil {
		logc.Errorf(l.ctx, "更新订单收货地址失败,参数:%+v,异常:%s", item, err.Error())
		return nil, errors.New("更新订单收货地址失败")
	}

	return &omsclient.UpdateOrderDeliveryResp{}, nil
}
