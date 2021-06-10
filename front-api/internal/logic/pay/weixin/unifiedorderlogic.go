package logic

import (
	"context"
	"go-zero-admin/front-api/internal/svc"
	"go-zero-admin/front-api/internal/types"
	"go-zero-admin/rpc/pay/payclient"

	"github.com/tal-tech/go-zero/core/logx"
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

	resp := &types.UnifiedOrderResp{}
	resp.Code = "000000"

	// 支付类型(1:app支付 2:小程序支付 3:h5支付 4:公众号支付)
	switch req.PayType {
	case 1:
		appUnifiedOrder(req, l, resp)
	case 2:
		smallUnisafiedOrder(req, l, resp)
	case 3:
		h5UnifiedOrder(req, l, resp)
	case 4:
		jsApiUnifiedOrder(req, l, resp)
	}

	return resp, nil
}

/**
微信app统一下单
*/
func appUnifiedOrder(req types.UnifiedOrderReq, l *UnifiedOrderLogic, resp *types.UnifiedOrderResp) {
	orderResp, _ := l.svcCtx.Pay.AppUnifiedOrder(l.ctx, &payclient.UnifiedOrderReq{
		BusinessId: req.BusinessId,
		Amount:     req.Amount,
		Remarks:    req.Remarks,
	})

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
微信小程序统一下单
*/
func smallUnisafiedOrder(req types.UnifiedOrderReq, l *UnifiedOrderLogic, resp *types.UnifiedOrderResp) {
	orderResp, _ := l.svcCtx.Pay.SmallUnifiedOrder(l.ctx, &payclient.UnifiedOrderReq{
		BusinessId: req.BusinessId,
		Amount:     req.Amount,
		Remarks:    req.Remarks,
	})

	var data = make(map[string]string)
	data["appId"] = orderResp.AppId
	data["timeStamp"] = orderResp.PartnerId
	data["nonceStr"] = orderResp.PrepayId
	data["package"] = orderResp.PackageStr
	data["signType"] = "MD5"
	data["sign"] = orderResp.Sign

	resp.Message = "微信小程序统一下单成功"
	resp.Data = data
}

/**
微信h5统一下单
*/
func h5UnifiedOrder(req types.UnifiedOrderReq, l *UnifiedOrderLogic, resp *types.UnifiedOrderResp) {
	orderResp, _ := l.svcCtx.Pay.H5UnifiedOrder(l.ctx, &payclient.UnifiedOrderReq{
		BusinessId: req.BusinessId,
		Amount:     req.Amount,
		Remarks:    req.Remarks,
	})

	var data = make(map[string]string)
	data[""] = orderResp.MWebUrl

	resp.Message = "微信h5统一下单成功"
	resp.Data = data
}

/**
微信jsapi统一下单
*/
func jsApiUnifiedOrder(req types.UnifiedOrderReq, l *UnifiedOrderLogic, resp *types.UnifiedOrderResp) {

	orderResp, _ := l.svcCtx.Pay.JsUnifiedOrder(l.ctx, &payclient.UnifiedOrderReq{
		BusinessId: req.BusinessId,
		Amount:     req.Amount,
		Remarks:    req.Remarks,
		Code:       req.Code,
	})

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
