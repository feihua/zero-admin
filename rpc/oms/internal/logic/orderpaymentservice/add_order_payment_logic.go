package orderpaymentservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/oms/gen/model"
	"github.com/feihua/zero-admin/rpc/oms/gen/query"
	"github.com/feihua/zero-admin/rpc/oms/internal/svc"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
	"time"
)

// AddOrderPaymentLogic 添加订单支付
/*
Author: LiuFeiHua
Date: 2025/06/30 15:49:27
*/
type AddOrderPaymentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddOrderPaymentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddOrderPaymentLogic {
	return &AddOrderPaymentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// AddOrderPayment 添加订单支付
func (l *AddOrderPaymentLogic) AddOrderPayment(in *omsclient.OrderPaymentReq) (*omsclient.OrderPaymentResp, error) {
	q := query.OmsOrderPayment

	now := time.Now()
	item := &model.OmsOrderPayment{
		OrderID:       in.OrderId,              // 订单ID
		OrderNo:       in.OrderNo,              // 订单编号
		PayType:       in.PayType,              // 支付方式：1-支付宝，2-微信，3-银联
		TransactionID: in.TransactionId,        // 支付流水号
		TotalAmount:   float64(in.TotalAmount), // 订单金额
		PayAmount:     float64(in.PayAmount),   // 支付金额
		PayStatus:     in.PayStatus,            // 支付状态：0-待支付，1-支付成功，2-支付失败
		PayTime:       &now,                    // 支付时间
		IsDeleted:     in.IsDeleted,            // 是否删除
	}

	err := q.WithContext(l.ctx).Create(item)
	if err != nil {
		logc.Errorf(l.ctx, "添加订单支付失败,参数:%+v,异常:%s", item, err.Error())
		return nil, errors.New("添加订单支付失败")
	}

	return &omsclient.OrderPaymentResp{}, nil
}
