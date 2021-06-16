package logic

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/xml"
	"fmt"
	"math/rand"
	"sort"
	"strings"
	"time"
)

const md5SignType = "MD5"
const CommonPayUrl = "https://api.mch.weixin.qq.com/pay/unifiedorder"
const searchTradeUrl = "https://api.mch.weixin.qq.com/pay/orderquery"

func getRandomString(length int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	byts := []byte(str)
	bytesLen := len(byts)
	var result []byte
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result = append(result, byts[r.Intn(bytesLen)])
	}
	return string(result)
}
func randNumber() string {
	return fmt.Sprintf("%05v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(100000))
}

// MD5加密
func hashMd5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	cipherStr := h.Sum(nil)
	return hex.EncodeToString(cipherStr)
}
func getMd5Sign(paramMap map[string]string, signKey string) string {
	sortString := getSortString(paramMap)
	//fmt.Println("sortString", sortString)
	sign := hashMd5(sortString + "&key=" + signKey)
	return strings.ToUpper(sign)
}

// 验证签名
func checkMd5Sign(rspMap map[string]string, md5Key, sign string) bool {
	calculateSign := getMd5Sign(rspMap, md5Key)
	//fmt.Println(calculateSign, sign, md5Key)
	return calculateSign == sign
}

// 支付字符串拼接
func getSortString(m map[string]string) string {
	var buf bytes.Buffer
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		vs := m[k]
		if vs == "" {
			continue
		}
		if buf.Len() > 0 {
			buf.WriteByte('&')
		}
		buf.WriteString(k)
		buf.WriteByte('=')
		buf.WriteString(vs)
	}
	return buf.String()
}

type CommonPayParam struct {
	XMLName        xml.Name `xml:"xml" json:"-"`
	AppId          string   `json:"appid" xml:"appid"`
	MchId          string   `json:"mch_id" xml:"mch_id"`
	NonceStr       string   `json:"nonce_str" xml:"nonce_str"`
	Body           string   `json:"body" xml:"body"`
	NotifyUrl      string   `json:"notify_url" xml:"notify_url"`
	Openid         string   `json:"openid" xml:"openid"`
	OutTradeNo     string   `json:"out_trade_no" xml:"out_trade_no"`
	SpBillCreateIp string   `json:"spbill_create_ip" xml:"spbill_create_ip"`
	TotalFee       string   `json:"total_fee" xml:"total_fee"`
	TradeType      string   `json:"trade_type" xml:"trade_type"`
	Sign           string   `json:"sign" xml:"sign"`
	SignType       string   `json:"sign_type" xml:"sign_type"`
}

type CommonPayResponse struct {
	XMLName    xml.Name `xml:"xml" json:"-"`
	ReturnCode string   `json:"return_code" xml:"return_code"`
	ReturnMsg  string   `json:"return_msg" xml:"return_msg"`
	AppId      string   `json:"appid" xml:"appid"`
	MchId      string   `json:"mch_id" xml:"mch_id"`
	NonceStr   string   `json:"nonce_str" xml:"nonce_str"`
	Sign       string   `json:"sign" xml:"sign"`
	ResultCode string   `json:"result_code" xml:"result_code"`
	PrepayId   string   `json:"prepay_id" xml:"prepay_id"`
	TradeType  string   `json:"trade_type" xml:"trade_type"`
}
