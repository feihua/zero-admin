package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"zero-admin/rpc/model/omsmodel"
	"zero-admin/rpc/oms/internal/svc"
	"zero-admin/rpc/oms/oms"
	"zero-admin/rpc/oms/omsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrderListByMemberIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOrderListByMemberIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderListByMemberIdLogic {
	return &OrderListByMemberIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *OrderListByMemberIdLogic) OrderListByMemberId(in *omsclient.OrderListByMemberIdReq) (*omsclient.OrderListByMemberIdResp, error) {
	var (
		orderList *[]omsmodel.OmsOrder
		resulterr error
	)
	if in.Status == 0 {
		orderList, resulterr = l.svcCtx.OmsOrderModel.FindListByMemberId(in.MemberId, in.Current, in.PageSize)
	} else {
		orderList, resulterr = l.svcCtx.OmsOrderModel.FindListByMemberIdAndStatus(in.MemberId, in.Status, in.Current, in.PageSize)
	}

	if resulterr != nil {
		reqStr, _ := json.Marshal(in)
		logx.WithContext(l.ctx).Errorf("查询订单列表信息失败,参数:%s,异常:%s", reqStr, resulterr.Error())
		return nil, resulterr
	}

	//todo 字段调整
	var list []*oms.OrderListByMemberIdData
	for _, order := range *orderList {
		statusText := strconv.FormatInt(order.Status, 10)
		productListByOrderId, _ := l.svcCtx.OmsOrderItemModel.FindProductListByOrderId(order.Id)

		var goodList []*oms.GoodsListByMemberIdData
		for _, orderItem := range *productListByOrderId {
			goodList = append(goodList, &oms.GoodsListByMemberIdData{
				Id:             int32(orderItem.ProductId),
				GoodsName:      orderItem.ProductName,
				Number:         orderItem.ProductSn,
				PicUrl:         orderItem.ProductPic,
				Specifications: orderItem.ProductName,
				Price:          fmt.Sprintf("%.2f", orderItem.ProductPrice),
				Quantity:       orderItem.ProductQuantity,
			})
		}

		//todo 字段调整
		list = append(list, &oms.OrderListByMemberIdData{
			Id:              int32(order.Id),
			OrderSn:         order.OrderSn,
			ActualPrice:     fmt.Sprintf("%.2f", order.PayAmount),
			OrderStatusText: statusText,
			HandleOption:    statusText,
			CreateTime:      order.CreateTime.Format("2006-01-02 15:04:05"),
			GoodsList:       goodList,
			IsDelivery:      order.IsDelivery,
		})
	}

	reqStr, _ := json.Marshal(in)
	listStr, _ := json.Marshal(list)
	logx.WithContext(l.ctx).Infof("查询订单列表信息,参数：%s,响应：%s", reqStr, listStr)

	return &omsclient.OrderListByMemberIdResp{
		Total: 0,
		List:  list,
	}, nil
}
