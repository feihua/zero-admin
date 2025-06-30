package orderpaymentservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/pkg/time_util"
	"github.com/feihua/zero-admin/rpc/oms/gen/query"
	"github.com/feihua/zero-admin/rpc/oms/internal/svc"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

// QueryOrderPaymentDetailLogic 查询订单支付详情
/*
Author: LiuFeiHua
Date: 2025/06/30 15:49:27
*/
type QueryOrderPaymentDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryOrderPaymentDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryOrderPaymentDetailLogic {
	return &QueryOrderPaymentDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryOrderPaymentDetail 查询订单支付详情
func (l *QueryOrderPaymentDetailLogic) QueryOrderPaymentDetail(in *omsclient.QueryOrderPaymentDetailReq) (*omsclient.OrderPaymentData, error) {
	item, err := query.OmsOrderPayment.WithContext(l.ctx).Where(query.OmsOrderPayment.ID.Eq(in.Id)).First()

	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		logc.Errorf(l.ctx, "订单支付不存在, 请求参数：%+v, 异常信息: %s", in, err.Error())
		return nil, errors.New("订单支付不存在")
	case err != nil:
		logc.Errorf(l.ctx, "查询订单支付异常, 请求参数：%+v, 异常信息: %s", in, err.Error())
		return nil, errors.New("查询订单支付异常")
	}

	data := &omsclient.OrderPaymentData{
		Id:            item.ID,                                 // 主键ID
		OrderId:       item.OrderID,                            // 订单ID
		OrderNo:       item.OrderNo,                            // 订单编号
		PayType:       item.PayType,                            // 支付方式：1-支付宝，2-微信，3-银联
		TransactionId: item.TransactionID,                      // 支付流水号
		TotalAmount:   float32(item.TotalAmount),               // 订单金额
		PayAmount:     float32(item.PayAmount),                 // 支付金额
		PayStatus:     item.PayStatus,                          // 支付状态：0-待支付，1-支付成功，2-支付失败
		PayTime:       time_util.TimeToString(item.PayTime),    // 支付时间
		CreateTime:    time_util.TimeToStr(item.CreateTime),    // 创建时间
		UpdateTime:    time_util.TimeToString(item.UpdateTime), // 更新时间
		IsDeleted:     item.IsDeleted,                          // 是否删除
	}

	return data, nil
}
