package orderservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/oms/gen/model"
	"github.com/feihua/zero-admin/rpc/oms/gen/query"
	"github.com/zeromicro/go-zero/core/logc"
	"gorm.io/gorm"
	"time"

	"github.com/feihua/zero-admin/rpc/oms/internal/svc"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateOrderLogic {
	return &UpdateOrderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateOrder 更新订单
func (l *UpdateOrderLogic) UpdateOrder(in *omsclient.UpdateOrderReq) (*omsclient.UpdateOrderResp, error) {
	q := query.OmsOrderMain.WithContext(l.ctx)

	// 1.根据订单id查询订单是否已存在
	_, err := q.Where(query.OmsOrderMain.ID.Eq(in.Id)).First()

	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		logc.Errorf(l.ctx, "订单不存在, 请求参数：%+v, 异常信息: %s", in, err.Error())
		return nil, errors.New("订单不存在")
	case err != nil:
		logc.Errorf(l.ctx, "查询订单异常, 请求参数：%+v, 异常信息: %s", in, err.Error())
		return nil, errors.New("查询订单异常")
	}

	item := &model.OmsOrderMain{
		ID:                 in.Id,                     // 订单ID
		Remark:             in.Remark,                 // 订单备注
		ExpressOrderNumber: in.ExpressOrderNumber,     // 快递单号
		OrderStatus:        in.OrderStatus,            // 订单状态：1-待支付,2-已支付,3-已发货,4-已完成,5-已取消,6-已退款,7-售后中
		FreightAmount:      float64(in.FreightAmount), // 运费金额
	}

	if in.OrderStatus != 0 {
		item.OrderStatus = in.OrderStatus
	}
	if in.ExpressOrderNumber != "" { // 快递单号
		now := time.Now()
		item.DeliveryTime = &now
	}

	// 2.订单存在时,则直接更新订单
	_, err = q.Updates(item)

	if err != nil {
		logc.Errorf(l.ctx, "更新订单失败,参数:%+v,异常:%s", item, err.Error())
		return nil, errors.New("更新订单失败")
	}

	return &omsclient.UpdateOrderResp{}, nil
}
