package logic

import (
	"bytes"
	"context"
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
	"zero-admin/rpc/model/paymodel"

	"zero-admin/rpc/pay/internal/svc"
	"zero-admin/rpc/pay/pay"

	"github.com/zeromicro/go-zero/core/logx"
)

type H5UnifiedOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewH5UnifiedOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *H5UnifiedOrderLogic {
	return &H5UnifiedOrderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *H5UnifiedOrderLogic) H5UnifiedOrder(in *pay.UnifiedOrderReq) (*pay.H5UnifiedOrderResp, error) {
	result, err := l.svcCtx.WxRecordModel.Insert(paymodel.PayWxRecord{
		BusinessId: in.BusinessId,
		Amount:     in.Amount,
		PayType:    in.PayType,
		Remarks:    in.Remarks,
		CreateTime: time.Now(),
		// 0：初始化 1：已发送 2：成功 3：失败
		PayStatus: 0,
	})
	if err != nil {
		return nil, err
	}
	id, _ := result.LastInsertId()

	merchants, _ := l.svcCtx.WxMerchantsModel.FindOneByMerId(in.MerId, in.PayType)

	res, _ := h5Pay(merchants)

	_ = l.svcCtx.WxRecordModel.Update(paymodel.PayWxRecord{
		Id:         id,
		BusinessId: in.BusinessId,
		Amount:     in.Amount,
		PayType:    in.PayType,
		Remarks:    in.Remarks,
		UpdateTime: time.Now(),
		ReturnCode: res.ReturnCode,
		ReturnMsg:  res.ReturnMsg,
		ResultCode: res.ResultCode,
		ResultMsg:  "",
		// 0：初始化 1：已发送 2：成功 3：失败
		PayStatus: 1,
	})

	return &pay.H5UnifiedOrderResp{}, nil
}

/**
构建h5支付统一下单参数
*/
func h5Pay(merchants *paymodel.PayWxMerchants) (commonPayRes CommonPayResponse, err error) {
	amount := 1
	payParam := make(map[string]string)
	payParam["appid"] = merchants.AppId
	payParam["mch_id"] = merchants.MchId
	payParam["nonce_str"] = getRandomString(32)
	payParam["body"] = fmt.Sprintf("微信充值:￥%d", amount)
	payParam["notify_url"] = "https://hy.life23.cn/order/notify"
	payParam["out_trade_no"] = fmt.Sprintf("test%s%s", time.Now().Format("20060102150405"), randNumber())
	payParam["spbill_create_ip"] = "127.0.0.1"
	payParam["total_fee"] = fmt.Sprintf("%v", amount)
	payParam["trade_type"] = merchants.PayType
	payParam["sign_type"] = md5SignType
	payParam["sign"] = getMd5Sign(payParam, merchants.MchKey)
	commonPayParam := CommonPayParam{
		AppId:          payParam["appid"],
		MchId:          payParam["mch_id"],
		NonceStr:       payParam["nonce_str"],
		Body:           payParam["body"],
		NotifyUrl:      payParam["notify_url"],
		OutTradeNo:     payParam["out_trade_no"],
		SpBillCreateIp: payParam["spbill_create_ip"],
		TotalFee:       payParam["total_fee"],
		TradeType:      payParam["trade_type"],
		Sign:           payParam["sign"],
		SignType:       payParam["sign_type"],
	}
	payXmlBytes, err := xml.Marshal(commonPayParam)
	if err != nil {
		return commonPayRes, err
	}
	//fmt.Println(string(payXmlBytes))
	request, err := http.NewRequest(http.MethodPost, CommonPayUrl, bytes.NewReader(payXmlBytes))
	if err != nil {
		return commonPayRes, err
	}
	client := http.DefaultClient
	response, err := client.Do(request)
	if err != nil {
		return commonPayRes, err
	}
	defer response.Body.Close()
	bodyBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return commonPayRes, err
	}
	if err = xml.Unmarshal(bodyBytes, &commonPayRes); err != nil {
		return commonPayRes, err
	}
	commonPayResParam := make(map[string]string)
	commonPayResParam["return_code"] = commonPayRes.ReturnCode
	commonPayResParam["return_msg"] = commonPayRes.ReturnMsg
	commonPayResParam["appid"] = commonPayRes.AppId
	commonPayResParam["mch_id"] = commonPayRes.MchId
	commonPayResParam["nonce_str"] = commonPayRes.NonceStr
	commonPayResParam["result_code"] = commonPayRes.ResultCode
	commonPayResParam["prepay_id"] = commonPayRes.PrepayId
	commonPayResParam["trade_type"] = commonPayRes.TradeType
	if !checkMd5Sign(commonPayResParam, merchants.MchKey, commonPayRes.Sign) {
		return commonPayRes, errors.New("common pay response sign verify error")
	}
	return commonPayRes, nil
}
