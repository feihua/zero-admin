package pay

import (
	"context"
	"encoding/json"
	"zero-admin/front-api/internal/svc"
	"zero-admin/front-api/internal/types"
	"zero-admin/rpc/oms/omsclient"
	"zero-admin/rpc/pay/payclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// OrderPayLogic
/*
Author: LiuFeiHua
Date: 2023/12/14 14:42
*/
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
//1.判断订单是否存在
//2.调用支付rpc进行预下单
//4.返回唤起客户端的信息
func (l *OrderPayLogic) OrderPay(req *types.OrderPayReq) (*types.OrderPayResp, error) {
	memberId, _ := l.ctx.Value("memberId").(json.Number).Int64()

	//1.判断订单是否存在
	orderInfo, err := l.svcCtx.OrderService.OrderListByMemberId(l.ctx, &omsclient.OrderListByMemberIdReq{
		Id:       req.OrderId,
		MemberId: memberId,
	})
	if err != nil {
		return payInfo(1, "查询订单异常", "")
	}

	//2.调用支付rpc进行预下单
	order := orderInfo.Data
	orderPay, err := l.svcCtx.OrderPayService.OrderPay(l.ctx, &payclient.OrderPayReq{
		BusinessId: req.OrderId,
		MemberId:   order.MemberId,
		MemberName: order.MemberUsername,
		Amount:     int64(order.PayAmount * 100), //订单系统保存的金额单位为元,支付rpc统一采用分为单位(因为微信和支付宝支付的单位不一样)
		Channel:    req.Channel,
		PayType:    req.PayType,
		Remarks:    req.Remark,
	})
	if err != nil {
		return payInfo(1, "预下单异常", "")
	}
	//{"申请编号"：用户个信息,”银行卡“：[第一张银行卡信息，第二张银行卡信息],“申请记录”：[第一条记录，第二条]，“人脸识别”：[]}
	//3.返回唤起客户端的信息
	return payInfo(0, "预下单成功", orderPay.PayInfo)

}

func payInfo(code int64, message, data string) (*types.OrderPayResp, error) {
	return &types.OrderPayResp{
		Code:    code,
		Message: message,
		Data:    data,
	}, nil
}
