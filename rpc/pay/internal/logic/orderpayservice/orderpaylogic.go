package orderpayservicelogic

import (
	"context"
	"fmt"
	"github.com/smartwalle/alipay/v3"
	"github.com/zeromicro/go-zero/core/logx"
	"zero-admin/rpc/pay/internal/svc"
	"zero-admin/rpc/pay/payclient"
)

// OrderPayLogic
/*
Author: LiuFeiHua
Date: 2023/12/14 14:13
*/
type OrderPayLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOrderPayLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderPayLogic {
	return &OrderPayLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

const (
	// TODO 设置回调地址域名
	kServerDomain = ""
)

// OrderPay 预下单,返回信息给app唤起客户端进行支付
func (l *OrderPayLogic) OrderPay(in *payclient.OrderPayReq) (*payclient.OrderPayResp, error) {

	var tradeNo = fmt.Sprintf("%d", Next())

	println(tradeNo)
	var p = alipay.TradeAppPay{}
	p.NotifyURL = kServerDomain + "/alipay/notify"
	p.ReturnURL = kServerDomain + "/alipay/callback"
	p.Subject = "支付测试:" + tradeNo
	p.OutTradeNo = tradeNo
	p.TotalAmount = "0.01"
	//p.ProductCode = "FAST_INSTANT_TRADE_PAY"

	result, _ := l.svcCtx.AlipayClient.TradeAppPay(p)
	return &payclient.OrderPayResp{
		PayInfo: result,
	}, nil
}
