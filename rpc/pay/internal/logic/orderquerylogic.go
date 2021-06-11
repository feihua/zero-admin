package logic

import (
	"context"
	"go-zero-admin/rpc/model/paymodel"
	"time"

	"go-zero-admin/rpc/pay/internal/svc"
	"go-zero-admin/rpc/pay/pay"

	"github.com/tal-tech/go-zero/core/logx"
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

	buildOderQueryVo(in, l)

	_ = l.svcCtx.WxRecordModel.Update(paymodel.PayWxRecord{
		Id:         result.Id,
		BusinessId: in.BusinessId,
		Amount:     result.Amount,
		Remarks:    result.Remarks,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
		ReturnCode: "",
		ReturnMsg:  "",
		ResultCode: "",
		ResultMsg:  "",
		// 0：初始化 1：已发送 2：成功 3：失败
		PayStatus: 2,
	})

	return &pay.OrderQueryResp{}, nil
}

/**
构建微信支付订单查询参数
*/
func buildOderQueryVo(in *pay.OrderQueryReq, l *OrderQueryLogic) map[string]string {

	merchants, _ := l.svcCtx.WxMerchantsModel.FindOneByMerId(in.MerId, "APP")

	return map[string]string{
		"appid":        merchants.AppId,
		"mch_id":       merchants.MchId,
		"nonce_str":    "",
		"out_trade_no": in.BusinessId,
		"sign":         "",
	}
}
