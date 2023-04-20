package logic

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	"zero-admin/rpc/oms/internal/svc"
	"zero-admin/rpc/oms/oms"
	"zero-admin/rpc/oms/omsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrderByOrderIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOrderByOrderIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderByOrderIdLogic {
	return &OrderByOrderIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *OrderByOrderIdLogic) OrderByOrderId(in *omsclient.OrderByOrderIdReq) (*omsclient.OrderByOrderIdResp, error) {
	order, err := l.svcCtx.OmsOrderModel.FindOne(l.ctx, in.OrderId)

	if err != nil {
		reqStr, _ := json.Marshal(in)
		logx.WithContext(l.ctx).Errorf("查询订单列表信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, err
	}
	if order.MemberId != in.MemberId {
		return nil, errors.New("用户订单不存在,取消订单失败")
	}
	//todo 字段调整

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
	list := &oms.OrderByOrderIdData{
		Id:              int32(order.Id),
		OrderSn:         order.OrderSn,
		ActualPrice:     fmt.Sprintf("%.2f", order.PayAmount),
		OrderStatusText: statusText,
		HandleOption:    statusText,
		CreateTime:      order.CreateTime.Format("2006-01-02 15:04:05"),
		GoodsList:       goodList,
		IsDelivery:      order.IsDelivery,
	}

	reqStr, _ := json.Marshal(in)
	listStr, _ := json.Marshal(list)
	logx.WithContext(l.ctx).Infof("查询订单列表信息,参数：%s,响应：%s", reqStr, listStr)

	return &omsclient.OrderByOrderIdResp{
		Total: 0,
		List:  list,
	}, nil
}
