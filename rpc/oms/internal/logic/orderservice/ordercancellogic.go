package orderservicelogic

import (
	"context"
	"encoding/json"
	"errors"

	"zero-admin/rpc/oms/internal/svc"
	"zero-admin/rpc/oms/omsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// OrderCancelLogic
/*
Author: LiuFeiHua
Date: 2023/12/7 9:37
*/
type OrderCancelLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOrderCancelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderCancelLogic {
	return &OrderCancelLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// OrderCancel 取消订单(只有在未付款的时候可以取消订单,其它状态只能申请退货或者售后,所以这里取消订单不涉及退款)
func (l *OrderCancelLogic) OrderCancel(in *omsclient.OrderCancelReq) (*omsclient.OrderCancelResp, error) {
	//1.查询订单是否存在
	order, err := l.svcCtx.OmsOrderModel.FindOneByMemberIdAndOrderId(l.ctx, in.MemberId, in.OrderId)

	if err != nil {
		return nil, errors.New("用户订单不存在,取消用户订单失败")
	}

	//2.判断是否可被取消
	if order.Status == 0 && order.DeleteStatus == 0 {
		if l.svcCtx.OmsOrderModel.UpdateOrderStatus(l.ctx, 0, 4, order.Id) != nil {
			return nil, errors.New("修改订单状态为取消失败")
		}
		//3.解除订单商品库存锁定ReleaseSkuStockLockData
		all, _ := l.svcCtx.OmsOrderItemModel.FindAll(l.ctx, 1, 1000, order.Id)

		var list []*omsclient.ReleaseSkuStockLockData
		for _, item := range *all {

			list = append(list, &omsclient.ReleaseSkuStockLockData{
				ProductQuantity: item.ProductQuantity,
				ProductSkuId:    item.ProductSkuId,
			})
		}

		//4和5返回api中操作
		return &omsclient.OrderCancelResp{
			CouponId:    order.CouponId,    //4.修改优惠券使用状态
			Integration: order.Integration, //5.返还使用积分
			Data:        list,
		}, nil
	}

	reqStr, _ := json.Marshal(in)
	logx.WithContext(l.ctx).Errorf("取消订单失败,参数：%s,订单状态：%s", reqStr, order.Status)
	return nil, errors.New("取消订单失败")

}
