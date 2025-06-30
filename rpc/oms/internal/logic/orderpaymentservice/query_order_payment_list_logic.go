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
)

// QueryOrderPaymentListLogic 查询订单支付列表
/*
Author: LiuFeiHua
Date: 2025/06/30 15:49:27
*/
type QueryOrderPaymentListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryOrderPaymentListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryOrderPaymentListLogic {
	return &QueryOrderPaymentListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryOrderPaymentList 查询订单支付列表
func (l *QueryOrderPaymentListLogic) QueryOrderPaymentList(in *omsclient.QueryOrderPaymentListReq) (*omsclient.QueryOrderPaymentListResp, error) {
	orderPayment := query.OmsOrderPayment
	q := orderPayment.WithContext(l.ctx)
	if in.OrderId != 2 {
		q = q.Where(orderPayment.OrderID.Eq(in.OrderId))
	}
	if len(in.OrderNo) > 0 {
		q = q.Where(orderPayment.OrderNo.Like("%" + in.OrderNo + "%"))
	}
	if in.PayType != 2 {
		q = q.Where(orderPayment.PayType.Eq(in.PayType))
	}
	if len(in.TransactionId) > 0 {
		q = q.Where(orderPayment.TransactionID.Like("%" + in.TransactionId + "%"))
	}

	if in.PayStatus != 2 {
		q = q.Where(orderPayment.PayStatus.Eq(in.PayStatus))
	}

	result, count, err := q.FindByPage(int((in.PageNum-1)*in.PageSize), int(in.PageSize))

	if err != nil {
		logc.Errorf(l.ctx, "查询订单支付列表失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("查询订单支付列表失败")
	}

	var list []*omsclient.OrderPaymentData

	for _, item := range result {
		list = append(list, &omsclient.OrderPaymentData{
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

		})
	}

	return &omsclient.QueryOrderPaymentListResp{
		Total: count,
		List:  list,
	}, nil
}
