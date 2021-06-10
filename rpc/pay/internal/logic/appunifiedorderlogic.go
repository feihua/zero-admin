package logic

import (
	"context"
	"go-zero-admin/rpc/model/paymodel"
	"go-zero-admin/rpc/pay/internal/svc"
	"go-zero-admin/rpc/pay/pay"
	"time"

	"github.com/tal-tech/go-zero/core/logx"
)

type AppUnifiedOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAppUnifiedOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AppUnifiedOrderLogic {
	return &AppUnifiedOrderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AppUnifiedOrderLogic) AppUnifiedOrder(in *pay.UnifiedOrderReq) (*pay.UnifiedOrderResp, error) {
	result, err := l.svcCtx.WxRecordModel.Insert(paymodel.PayWxRecord{
		BusinessId: in.BusinessId,
		Amount:     in.Amount,
		// 支付类型(1:app支付 2:小程序支付 3:h5支付 4:公众号支付)
		PayType:    1,
		Remarks:    in.Remarks,
		CreateTime: time.Now(),
		// 0：初始化 1：已发送 2：成功 3：失败
		PayStatus: 0,
	})
	if err != nil {
		return nil, err
	}
	id, _ := result.LastInsertId()

	//todo 请求微信统一下单

	_ = l.svcCtx.WxRecordModel.Update(paymodel.PayWxRecord{
		Id:         id,
		BusinessId: in.BusinessId,
		Amount:     in.Amount,
		// 支付类型(1:app支付 2:小程序支付 3:h5支付 4:公众号支付)
		PayType:    1,
		Remarks:    "",
		UpdateTime: time.Now(),
		ReturnCode: "",
		ReturnMsg:  "",
		ResultCode: "",
		ResultMsg:  "",
		// 0：初始化 1：已发送 2：成功 3：失败
		PayStatus: 1,
	})

	return &pay.UnifiedOrderResp{}, nil
}
