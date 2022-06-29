package logic

import (
	"context"
	"encoding/json"
	"errors"

	"zero-admin/rpc/oms/internal/svc"
	"zero-admin/rpc/oms/omsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrderRefundLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOrderRefundLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderRefundLogic {
	return &OrderRefundLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *OrderRefundLogic) OrderRefund(in *omsclient.OrderRefundReq) (*omsclient.OrderRefundResp, error) {
	order, err := l.svcCtx.OmsOrderModel.FindOne(in.OrderId)

	if err != nil {
		return nil, err
	}

	if order.MemberId != in.UserId {
		return nil, errors.New("用户订单不存在,退款失败")
	}

	//检测是否能够退款
	//订单状态：0->待付款；1->待发货；2->已发货；3->已完成；4->已关闭；5->无效订单
	//如果订单已付款，没有发货，则可退款
	if order.Status != 1 {
		reqStr, _ := json.Marshal(in)
		logx.WithContext(l.ctx).Errorf("退款失败,参数：%s,订单状态：%s", reqStr, order.Status)
		return nil, errors.New("退款失败")
	}

	//设置订单申请退款状态
	if l.svcCtx.OmsOrderModel.UpdateOrderStatus(4, order.Id) != nil {
		return nil, errors.New("设置订单申请退款状态失败")
	}

	//todo 发送邮件和短信通知，这里采用异步发送，有用户申请退款，邮件通知运营人员

	return &omsclient.OrderRefundResp{}, nil

}
