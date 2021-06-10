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

	//todo 请求微信统一下单

	_ = l.svcCtx.WxRecordModel.Update(paymodel.PayWxRecord{
		Id:         result.Id,
		BusinessId: in.BusinessId,
		Amount:     result.Amount,
		// 支付类型(1:app支付 2:小程序支付 3:h5支付 4:公众号支付)
		PayType:    1,
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
