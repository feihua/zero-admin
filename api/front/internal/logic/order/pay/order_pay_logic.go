package pay

import (
	"context"
	"github.com/feihua/zero-admin/api/front/internal/logic/common"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"

	"github.com/feihua/zero-admin/api/front/internal/svc"
	"github.com/feihua/zero-admin/api/front/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// OrderPayLogic 预下单
type OrderPayLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOrderPayLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderPayLogic {
	return &OrderPayLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// OrderPay 预下单
func (l *OrderPayLogic) OrderPay(req *types.OrderPayReq) (resp *types.OrderPayResp, err error) {
	memberId, err := common.GetMemberId(l.ctx)
	if err != nil {
		return nil, err
	}

	// 1.判断订单是否存在
	orderInfo, err := l.svcCtx.OrderService.QueryOrderDetail(l.ctx, &omsclient.QueryOrderDetailReq{
		Id:     req.OrderId,
		UserId: memberId,
	})
	if err != nil {
		return orderPayResp(1, "查询订单异常", "")
	}

	// 2.调用支付rpc进行预下单
	operationsUtils := NewPaymentOperationsUtils(l.ctx, l.svcCtx)
	outTradeNo := orderInfo.Data.OrderNo
	message, err := operationsUtils.TradeAppPay(outTradeNo, "0.01", "支付测试")
	if err != nil {
		return orderPayResp(1, message, "")
	}

	// 3.返回唤起客户端的信息
	return orderPayResp(0, "预下单成功", message)

}

func orderPayResp(code int64, message, data string) (*types.OrderPayResp, error) {
	return &types.OrderPayResp{
		Code:    code,
		Message: message,
		Data:    data,
	}, nil
}
