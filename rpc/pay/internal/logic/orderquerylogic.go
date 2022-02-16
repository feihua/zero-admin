package logic

import (
	"bytes"
	"context"
	"encoding/xml"
	"errors"
	"io/ioutil"
	"net/http"
	"time"
	"zero-admin/rpc/model/paymodel"

	"zero-admin/rpc/pay/internal/svc"
	"zero-admin/rpc/pay/pay"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrderQueryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOrderQueryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderQueryLogic {
	return &OrderQueryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *OrderQueryLogic) OrderQuery(in *pay.OrderQueryReq) (*pay.OrderQueryResp, error) {
	// 查询本地订单是否已经支付
	result, err := l.svcCtx.WxRecordModel.FindOneByBusinessId(in.BusinessId)
	if err != nil {
		return nil, err
	}

	if result.PayStatus == 2 {
		return &pay.OrderQueryResp{}, nil
	}

	merchants, _ := l.svcCtx.WxMerchantsModel.FindOneByMerId(in.MerId, in.PayType)

	res, _ := searchTrade(merchants, in.BusinessId)

	_ = l.svcCtx.WxRecordModel.Update(paymodel.PayWxRecord{
		Id:         result.Id,
		BusinessId: in.BusinessId,
		Amount:     result.Amount,
		Remarks:    result.Remarks,
		UpdateTime: time.Now(),
		ReturnCode: res.ReturnCode,
		ReturnMsg:  res.ReturnMsg,
		ResultCode: res.ResultCode,
		ResultMsg:  "",
		// 0：初始化 1：已发送 2：成功 3：失败
		PayStatus: 2,
	})

	return &pay.OrderQueryResp{}, nil
}

type SearchTradeParam struct {
	XMLName    xml.Name `xml:"xml" json:"-"`
	AppId      string   `json:"appid" xml:"appid"`
	MchId      string   `json:"mch_id" xml:"mch_id"`
	OutTradeNo string   `json:"out_trade_no" xml:"out_trade_no"`
	NonceStr   string   `json:"nonce_str" xml:"nonce_str"`
	SignType   string   `json:"sign_type" xml:"sign_type"`
	Sign       string   `json:"sign" xml:"sign"`
}
type SearchTradeResponse struct {
	ReturnCode     string `json:"return_code" xml:"return_code"`
	ReturnMsg      string `json:"return_msg" xml:"return_msg"`
	AppId          string `json:"appid" xml:"appid"`
	MchId          string `json:"mch_id" xml:"mch_id"`
	DeviceInfo     string `json:"device_info" xml:"device_info"`
	NonceStr       string `json:"nonce_str" xml:"nonce_str"`
	Sign           string `json:"sign" xml:"sign"`
	ResultCode     string `json:"result_code" xml:"result_code"`
	Openid         string `json:"openid" xml:"openid"`
	IsSubscribe    string `json:"is_subscribe" xml:"is_subscribe"`
	TradeType      string `json:"trade_type" xml:"trade_type"`
	BankType       string `json:"bank_type" xml:"bank_type"`
	TotalFee       string `json:"total_fee" xml:"total_fee"`
	FeeType        string `json:"fee_type" xml:"fee_type"`
	TransactionId  string `json:"transaction_id" xml:"transaction_id"`
	OutTradeNo     string `json:"out_trade_no" xml:"out_trade_no"`
	Attach         string `json:"attach" xml:"attach"`
	TimeEnd        string `json:"time_end" xml:"time_end"`
	TradeState     string `json:"trade_state" xml:"trade_state"`
	CashFee        string `json:"cash_fee" xml:"cash_fee"`
	TradeStateDesc string `json:"trade_state_desc" xml:"trade_state_desc"`
	CashFeeType    string `json:"cash_fee_type" xml:"cash_fee_type"`
}

// 查询交易
func searchTrade(merchants *paymodel.PayWxMerchants, businessId string) (trade SearchTradeResponse, err error) {
	paramMap := make(map[string]string)
	paramMap["appid"] = merchants.AppId
	paramMap["mch_id"] = merchants.MchId
	paramMap["out_trade_no"] = businessId
	paramMap["nonce_str"] = getRandomString(32)
	paramMap["sign_type"] = md5SignType
	paramMap["sign"] = getMd5Sign(paramMap, merchants.MchKey)
	searchTradeParam := SearchTradeParam{
		AppId:      paramMap["appid"],
		MchId:      paramMap["mch_id"],
		OutTradeNo: paramMap["out_trade_no"],
		NonceStr:   paramMap["nonce_str"],
		SignType:   paramMap["sign_type"],
		Sign:       paramMap["sign"],
	}
	searchTradeParamBytes, err := xml.Marshal(searchTradeParam)
	if err != nil {
		return trade, err
	}
	request, err := http.NewRequest(http.MethodPost, searchTradeUrl, bytes.NewReader(searchTradeParamBytes))
	if err != nil {
		return trade, err
	}
	client := http.DefaultClient
	response, err := client.Do(request)
	if err != nil {
		return trade, err
	}
	defer response.Body.Close()
	bodyBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return trade, err
	}
	if err := xml.Unmarshal(bodyBytes, &trade); err != nil {
		return trade, err
	}
	searchTradeMap := make(map[string]string)
	searchTradeMap["return_code"] = trade.ReturnCode
	searchTradeMap["return_msg"] = trade.ReturnMsg
	searchTradeMap["appid"] = trade.AppId
	searchTradeMap["mch_id"] = trade.MchId
	searchTradeMap["device_info"] = trade.DeviceInfo
	searchTradeMap["nonce_str"] = trade.NonceStr
	searchTradeMap["result_code"] = trade.ResultCode
	searchTradeMap["openid"] = trade.Openid
	searchTradeMap["is_subscribe"] = trade.IsSubscribe
	searchTradeMap["trade_type"] = trade.TradeType
	searchTradeMap["bank_type"] = trade.BankType
	searchTradeMap["total_fee"] = trade.TotalFee
	searchTradeMap["fee_type"] = trade.FeeType
	searchTradeMap["transaction_id"] = trade.TransactionId
	searchTradeMap["out_trade_no"] = trade.OutTradeNo
	searchTradeMap["attach"] = trade.Attach
	searchTradeMap["time_end"] = trade.TimeEnd
	searchTradeMap["trade_state"] = trade.TradeState
	searchTradeMap["cash_fee"] = trade.CashFee
	searchTradeMap["trade_state_desc"] = trade.TradeStateDesc
	searchTradeMap["cash_fee_type"] = trade.CashFeeType

	// 签名校验
	if !checkMd5Sign(searchTradeMap, merchants.MchKey, trade.Sign) {
		return trade, errors.New("search trade result sign verify error")
	}

	return trade, nil
}
