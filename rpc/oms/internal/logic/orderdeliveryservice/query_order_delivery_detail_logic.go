package orderdeliveryservicelogic

import (
	"context"
	"errors"

	"github.com/feihua/zero-admin/pkg/time_util"
	"github.com/feihua/zero-admin/rpc/oms/gen/query"
	"github.com/feihua/zero-admin/rpc/oms/internal/svc"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

// QueryOrderDeliveryDetailLogic 查询订单收货地址详情
/*
Author: LiuFeiHua
Date: 2025/05/26 15:21:44
*/
type QueryOrderDeliveryDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryOrderDeliveryDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryOrderDeliveryDetailLogic {
	return &QueryOrderDeliveryDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryOrderDeliveryDetail 查询订单收货地址详情
func (l *QueryOrderDeliveryDetailLogic) QueryOrderDeliveryDetail(in *omsclient.QueryOrderDeliveryDetailReq) (*omsclient.QueryOrderDeliveryDetailResp, error) {
	item, err := query.OmsOrderDelivery.WithContext(l.ctx).Where(query.OmsOrderDelivery.ID.Eq(in.Id)).First()

	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		logc.Errorf(l.ctx, "订单收货地址不存在, 请求参数：%+v, 异常信息: %s", in, err.Error())
		return nil, errors.New("订单收货地址不存在")
	case err != nil:
		logc.Errorf(l.ctx, "查询订单收货地址异常, 请求参数：%+v, 异常信息: %s", in, err.Error())
		return nil, errors.New("查询订单收货地址异常")
	}

	data := &omsclient.QueryOrderDeliveryDetailResp{
		Id:               item.ID,                                 //
		OrderId:          item.OrderID,                            // 订单ID
		OrderNo:          item.OrderNo,                            // 订单编号
		ReceiverName:     item.ReceiverName,                       // 收货人姓名
		ReceiverPhone:    item.ReceiverPhone,                      // 收货人电话
		ReceiverProvince: item.ReceiverProvince,                   // 省份
		ReceiverCity:     item.ReceiverCity,                       // 城市
		ReceiverDistrict: item.ReceiverDistrict,                   // 区县
		ReceiverAddress:  item.ReceiverAddress,                    // 详细地址
		DeliveryCompany:  item.DeliveryCompany,                    // 物流公司
		DeliveryNo:       item.DeliveryNo,                         // 物流单号
		CreateTime:       time_util.TimeToStr(item.CreateTime),    // 创建时间
		UpdateTime:       time_util.TimeToString(item.UpdateTime), // 更新时间
	}

	return data, nil
}
