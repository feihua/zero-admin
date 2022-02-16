package logic

import (
	"context"
	"zero-admin/front-api/internal/svc"
	"zero-admin/front-api/internal/types"
	"zero-admin/rpc/pay/payclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type UnifiedOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUnifiedOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) UnifiedOrderLogic {
	return UnifiedOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UnifiedOrderLogic) UnifiedOrder(req types.UnifiedOrderReq) (*types.UnifiedOrderResp, error) {

	orderReq := &payclient.UnifiedOrderReq{
		BusinessId: req.BusinessId,
		Amount:     req.Amount,
		Remarks:    req.Remarks,
		MerId:      req.MerId,
		Code:       req.Code,
		PayType:    req.PayType,
	}

	resp := &types.UnifiedOrderResp{}
	resp.Code = "000000"

	// 支付类型(APP:APP支付 JSAPI:小程序,公众号 MWEB:H5支付)
	switch req.PayType {
	case "APP":
		appUnifiedOrder(orderReq, l, resp)
	case "JSAPI":
		jsApiUnifiedOrder(orderReq, l, resp)
	case "MWEB":
		h5UnifiedOrder(orderReq, l, resp)
	}

	return resp, nil
}

/**
微信app统一下单
*/
func appUnifiedOrder(req *payclient.UnifiedOrderReq, l *UnifiedOrderLogic, resp *types.UnifiedOrderResp) {

	orderResp, _ := l.svcCtx.Pay.AppUnifiedOrder(l.ctx, req)

	var data = make(map[string]string)
	data["appId"] = orderResp.AppId
	data["partnerId"] = orderResp.PartnerId
	data["prepayId"] = orderResp.PrepayId
	data["package"] = orderResp.PackageStr
	data["nonceStr"] = orderResp.NonceStr
	data["timestamp"] = orderResp.Timestamp
	data["sign"] = orderResp.Sign

	resp.Message = "微信app统一下单成功"
	resp.Data = data
}

/**
微信jsapi统一下单(小程序统,公从号)
*/
func jsApiUnifiedOrder(req *payclient.UnifiedOrderReq, l *UnifiedOrderLogic, resp *types.UnifiedOrderResp) {

	orderResp, _ := l.svcCtx.Pay.JsUnifiedOrder(l.ctx, req)

	var data = make(map[string]string)
	data["appId"] = orderResp.AppId
	data["timeStamp"] = orderResp.PartnerId
	data["nonceStr"] = orderResp.PrepayId
	data["package"] = orderResp.PackageStr
	data["signType"] = "MD5"
	data["sign"] = orderResp.Sign

	resp.Message = "微信jsapi统一下单成功"
	resp.Data = data
}

/**
微信h5统一下单
*/
func h5UnifiedOrder(req *payclient.UnifiedOrderReq, l *UnifiedOrderLogic, resp *types.UnifiedOrderResp) {

	orderResp, _ := l.svcCtx.Pay.H5UnifiedOrder(l.ctx, req)

	var data = make(map[string]string)
	data["nWebUrl"] = orderResp.MWebUrl

	resp.Message = "微信h5统一下单成功"
	resp.Data = data
}
