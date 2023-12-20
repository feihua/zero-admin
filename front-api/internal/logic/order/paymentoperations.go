package order

import (
	"context"
	"errors"
	"github.com/smartwalle/alipay/v3"
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"
	"zero-admin/front-api/internal/svc"
	"zero-admin/rpc/oms/omsclient"
)

// PaymentOperationsUtils 支付相关工具
/*
Author: LiuFeiHua
Date: 2023/12/15 10:05
*/
type PaymentOperationsUtils struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPaymentOperationsUtils(ctx context.Context, svcCtx *svc.ServiceContext) *PaymentOperationsUtils {
	return &PaymentOperationsUtils{Logger: logx.WithContext(ctx), ctx: ctx, svcCtx: svcCtx}
}

// TradeAppPay 下单支付
func (l *PaymentOperationsUtils) TradeAppPay(outTradeNo, totalAmount, subject string) (string, error) {
	var p = alipay.TradeAppPay{}
	p.NotifyURL = l.svcCtx.Config.Alipay.NotifyURL
	p.Subject = subject         //订单标题
	p.OutTradeNo = outTradeNo   //商户订单号，64个字符以内、可包含字母、数字、下划线；需保证在商户端不重复
	p.TotalAmount = totalAmount // 订单总金额，单位为元，精确到小数点后两位，取值范围[0.01,100000000]
	return l.svcCtx.AlipayClient.TradeAppPay(p)
}

// AliPayNotify 回调通知
func (l *PaymentOperationsUtils) AliPayNotify(writer http.ResponseWriter, request *http.Request) {

	if err := request.ParseForm(); err != nil {
		_, _ = writer.Write([]byte("error"))
		return
	}

	var notification, err = l.svcCtx.AlipayClient.DecodeNotification(request.Form)
	if err != nil {
		_, _ = writer.Write([]byte("error"))
		return
	}

	l.WithContext(l.ctx).Infof("支付宝支付回调参数：%+v", notification)
	var outTradeNo = notification.OutTradeNo
	var tradeStatus = notification.TradeStatus

	//订单状态修改可以为分二种方式实现
	//1.直接更新订单表
	//2.发送到mq,后再处理
	//目前实现的是第一种,直接更新订单状态
	if alipay.TradeStatusSuccess == tradeStatus {
		_, err = l.svcCtx.OrderService.UpdateOrderStatusByOutTradeNo(l.ctx, &omsclient.UpdateOrderStatusByOutTradeNoReq{
			OutTradeNo:  outTradeNo,
			OrderStatus: 1, //支付成功后,订单状态修改为待发货
		})

		if err == nil {
			l.svcCtx.AlipayClient.ACKNotification(writer)
		} else {
			_, _ = writer.Write([]byte("error"))
		}
	}
	_, _ = writer.Write([]byte("error"))
}

//TradeQuery 查询订单
func (l *PaymentOperationsUtils) TradeQuery(outTradeNo string) (string, int64, error) {
	var p = alipay.TradeQuery{}
	p.OutTradeNo = outTradeNo
	tradeQuery, err := l.svcCtx.AlipayClient.TradeQuery(p)
	if err != nil {
		return outTradeNo, 0, err
	}
	if tradeQuery.IsFailure() {
		return outTradeNo, 0, errors.New("接口调用失败")
	}
	if tradeQuery.TradeStatus == alipay.TradeStatusSuccess {
		return outTradeNo, 1, nil
	}
	return outTradeNo, 0, nil

}
