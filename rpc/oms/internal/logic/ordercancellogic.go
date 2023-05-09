package logic

import (
	"context"
	"encoding/json"
	"errors"

	"zero-admin/rpc/oms/internal/svc"
	"zero-admin/rpc/oms/omsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

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

func (l *OrderCancelLogic) OrderCancel(in *omsclient.OrderCancelReq) (*omsclient.OrderCancelResp, error) {
	order, err := l.svcCtx.OmsOrderModel.FindOne(l.ctx, in.OrderId)

	if err != nil {
		return nil, err
	}

	if order.MemberId != in.UserId {
		return nil, errors.New("用户订单不存在,取消订单失败")
	}

	//检测是否能够取消
	//订单状态：0->待付款；1->待发货；2->已发货；3->已完成；4->已关闭；5->无效订单
	//例如订单没有被取消，且没有支付，则可支付，可取消
	if order.Status != 0 {
		reqStr, _ := json.Marshal(in)
		logx.WithContext(l.ctx).Errorf("取消订单失败,参数：%s,订单状态：%s", reqStr, order.Status)
		return nil, errors.New("取消订单失败")
	}

	//设置订单已取消状态
	//if l.svcCtx.OmsOrderModel.UpdateOrderStatus(5, order.Id) != nil {
	//	return nil, errors.New("设置订单已取消状态失败")
	//}

	//todo 商品货品数量增加

	return &omsclient.OrderCancelResp{}, nil

}
