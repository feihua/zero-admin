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
)

// QueryOrderDeliveryListLogic 查询订单收货地址列表
/*
Author: LiuFeiHua
Date: 2025/05/26 15:21:44
*/
type QueryOrderDeliveryListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryOrderDeliveryListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryOrderDeliveryListLogic {
	return &QueryOrderDeliveryListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryOrderDeliveryList 查询订单收货地址列表
func (l *QueryOrderDeliveryListLogic) QueryOrderDeliveryList(in *omsclient.QueryOrderDeliveryListReq) (*omsclient.QueryOrderDeliveryListResp, error) {
	orderDelivery := query.OmsOrderDelivery
	q := orderDelivery.WithContext(l.ctx)
	if in.OrderId != 2 {
		q = q.Where(orderDelivery.OrderID.Eq(in.OrderId))
	}
	if len(in.OrderNo) > 0 {
		q = q.Where(orderDelivery.OrderNo.Like("%" + in.OrderNo + "%"))
	}
	if len(in.ReceiverName) > 0 {
		q = q.Where(orderDelivery.ReceiverName.Like("%" + in.ReceiverName + "%"))
	}
	if len(in.ReceiverPhone) > 0 {
		q = q.Where(orderDelivery.ReceiverPhone.Like("%" + in.ReceiverPhone + "%"))
	}

	if len(in.ReceiverAddress) > 0 {
		q = q.Where(orderDelivery.ReceiverAddress.Like("%" + in.ReceiverAddress + "%"))
	}
	if len(in.DeliveryCompany) > 0 {
		q = q.Where(orderDelivery.DeliveryCompany.Like("%" + in.DeliveryCompany + "%"))
	}
	if len(in.DeliveryNo) > 0 {
		q = q.Where(orderDelivery.DeliveryNo.Like("%" + in.DeliveryNo + "%"))
	}

	result, count, err := q.FindByPage(int((in.PageNum-1)*in.PageSize), int(in.PageSize))

	if err != nil {
		logc.Errorf(l.ctx, "查询订单收货地址列表失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("查询订单收货地址列表失败")
	}

	var list []*omsclient.OrderDeliveryListData

	for _, item := range result {
		list = append(list, &omsclient.OrderDeliveryListData{
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

		})
	}

	return &omsclient.QueryOrderDeliveryListResp{
		Total: count,
		List:  list,
	}, nil
}
