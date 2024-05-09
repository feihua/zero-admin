package orderservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/oms/gen/model"
	"github.com/feihua/zero-admin/rpc/oms/gen/query"
	"github.com/zeromicro/go-zero/core/logc"
	"time"

	"github.com/feihua/zero-admin/rpc/oms/internal/svc"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"

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
	q := query.OmsOrder
	//1.查询订单是否存在
	order, err := q.WithContext(l.ctx).Where(q.ID.Eq(in.OrderId), q.MemberID.Eq(in.MemberId)).First()

	if err != nil {
		return nil, errors.New("用户订单不存在,取消用户订单失败")
	}

	//2.判断是否可被取消
	if order.Status == 0 && order.DeleteStatus == 0 {
		item := model.OmsOrder{}
		item.ID = in.OrderId
		item.ConfirmStatus = 0
		item.Status = 4
		item.ReceiveTime = time.Now()
		_, err = q.WithContext(l.ctx).Where(q.ID.Eq(in.OrderId)).Updates(item)
		if err != nil {
			return nil, err
		}

		//3.解除订单商品库存锁定ReleaseSkuStockLockData
		result, _ := query.OmsOrderItem.WithContext(l.ctx).Where(query.OmsOrderItem.OrderID.Eq(order.ID)).Find()

		var list []*omsclient.ReleaseSkuStockLockData
		for _, x := range result {

			list = append(list, &omsclient.ReleaseSkuStockLockData{
				ProductQuantity: x.ProductQuantity,
				ProductSkuId:    x.ProductSkuID,
			})
		}

		//4和5返回api中操作
		return &omsclient.OrderCancelResp{
			CouponId:    order.CouponID,    //4.修改优惠券使用状态
			Integration: order.Integration, //5.返还使用积分
			Data:        list,
		}, nil
	}

	logc.Errorf(l.ctx, "取消订单失败,参数：%+v,订单状态：%d", in, order.Status)
	return nil, errors.New("取消订单失败")

}
