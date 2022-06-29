package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"zero-admin/rpc/oms/oms"

	"zero-admin/rpc/oms/internal/svc"
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
	orderList, err := l.svcCtx.OmsOrderModel.FindListByMemberId(in.MemberId)

	if err != nil {
		reqStr, _ := json.Marshal(in)
		logx.WithContext(l.ctx).Errorf("查询订单列表信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, err
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
			})
		}

		//todo 字段调整
		list = append(list, &oms.OrderListByMemberIdData{
			Id:              int32(order.Id),
			OrderSn:         order.OrderSn,
			ActualPrice:     fmt.Sprintf("%.2f", order.PayAmount),
			OrderStatusText: statusText,
			HandleOption:    statusText,
			GoodsList:       goodList,
		})
	}

	reqStr, _ := json.Marshal(in)
	listStr, _ := json.Marshal(list)
	logx.WithContext(l.ctx).Infof("查询订单列表信息,参数：%s,响应：%s", reqStr, listStr)
	return &oms.OrderListByMemberIdResp{
		Total: 0,
		List:  list,
	}, nil
}
