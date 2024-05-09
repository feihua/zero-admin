package orderservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/oms/gen/query"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/rpc/oms/internal/svc"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"

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
	q := query.OmsOrder
	//1.查询订单是否存在
	order, err := q.WithContext(l.ctx).Where(q.ID.Eq(in.OrderId), q.MemberID.Eq(in.UserId)).First()

	if err != nil {
		return nil, err
	}

	//检测是否能够退款
	//订单状态：0->待付款；1->待发货；2->已发货；3->已完成；4->已关闭；5->无效订单
	//如果订单已付款，没有发货，则可退款
	if order.Status != 1 {
		logc.Errorf(l.ctx, "退款失败,参数：%s,订单状态：%s", in, order.Status)
		return nil, errors.New("退款失败")
	}

	//设置订单申请退款状态
	//if l.svcCtx.OmsOrderModel.UpdateOrderStatus(4, order.Id) != nil {
	//	return nil, errors.New("设置订单申请退款状态失败")
	//}

	//todo 发送邮件和短信通知，这里采用异步发送，有用户申请退款，邮件通知运营人员

	return &omsclient.OrderRefundResp{}, nil

}
