package logic

import (
	"context"
	"go-zero-admin/rpc/model/paymodel"
	"time"

	"go-zero-admin/rpc/pay/internal/svc"
	"go-zero-admin/rpc/pay/pay"

	"github.com/tal-tech/go-zero/core/logx"
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
		// 支付类型(APP:APP支付 JSAPI:小程序,公众号 MWEB:H5支付)
		PayType:    "MWEB",
		Remarks:    in.Remarks,
		CreateTime: time.Now(),
		// 0：初始化 1：已发送 2：成功 3：失败
		PayStatus: 0,
	})
	if err != nil {
		return nil, err
	}
	id, _ := result.LastInsertId()

	buildH5OrderReqVo(in, l)

	_ = l.svcCtx.WxRecordModel.Update(paymodel.PayWxRecord{
		Id:         id,
		BusinessId: in.BusinessId,
		Amount:     in.Amount,
		// 支付类型(APP:APP支付 JSAPI:小程序,公众号 MWEB:H5支付)
		PayType:    h5TradeType,
		Remarks:    in.Remarks,
		UpdateTime: time.Now(),
		ReturnCode: "",
		ReturnMsg:  "",
		ResultCode: "",
		ResultMsg:  "",
		// 0：初始化 1：已发送 2：成功 3：失败
		PayStatus: 1,
	})

	return &pay.H5UnifiedOrderResp{}, nil
}

/**
构建h5支付统一下单参数
*/
func buildH5OrderReqVo(in *pay.UnifiedOrderReq, l *H5UnifiedOrderLogic) map[string]string {

	merchants, _ := l.svcCtx.WxMerchantsModel.FindOneByMerId(in.MerId, "APP")

	reqVo := make(map[string]string)
	reqVo["appid"] = merchants.AppId
	reqVo["mch_id"] = merchants.MchId
	reqVo["nonce_str"] = ""
	reqVo["sign_type"] = md5SignType
	reqVo["body"] = in.Remarks
	reqVo["out_trade_no"] = in.BusinessId
	reqVo["fee_type"] = "CNY"
	reqVo["total_fee"] = in.Amount
	reqVo["spbill_create_ip"] = ""
	reqVo["notify_url"] = merchants.NotifyUrl
	reqVo["trade_type"] = h5TradeType
	reqVo["scene_info"] = ""
	//reqVo["attach"] = ""
	reqVo["sign"] = ""

	return reqVo
}
