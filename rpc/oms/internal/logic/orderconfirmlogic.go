package logic

import (
	"context"
	"encoding/json"
	"errors"

	"zero-admin/rpc/oms/internal/svc"
	"zero-admin/rpc/oms/omsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrderConfirmLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOrderConfirmLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderConfirmLogic {
	return &OrderConfirmLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *OrderConfirmLogic) OrderConfirm(in *omsclient.OrderConfirmReq) (*omsclient.OrderConfirmResp, error) {
	order, err := l.svcCtx.OmsOrderModel.FindOne(in.OrderId)

	if err != nil {
		return nil, err
	}

	if order.MemberId != in.UserId {
		return nil, errors.New("用户订单不存在,确认收货失败")
	}

	//检测是否能够确认收货
	//订单状态：0->待付款；1->待发货；2->已发货；3->已完成；4->已关闭；5->无效订单
	//如果订单已经发货，没有收货，则可收货操作,
	if order.Status != 2 {
		reqStr, _ := json.Marshal(in)
		logx.WithContext(l.ctx).Errorf("确认收货失败,参数：%s,订单状态：%s", reqStr, order.Status)
		return nil, errors.New("确认收货失败")
	}

	//设置订单已确认状态
	if l.svcCtx.OmsOrderModel.UpdateOrderStatus(3, order.Id) != nil {
		return nil, errors.New("设置订单已确认状态失败")
	}
	return &omsclient.OrderConfirmResp{}, nil

}
