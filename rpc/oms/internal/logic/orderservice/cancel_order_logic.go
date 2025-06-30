package orderservicelogic

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/feihua/zero-admin/rpc/oms/gen/query"
	"github.com/feihua/zero-admin/rpc/oms/internal/svc"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"gorm.io/gorm"

	"github.com/zeromicro/go-zero/core/logx"
)

type CancelOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCancelOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CancelOrderLogic {
	return &CancelOrderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// CancelOrder 取消订单(app)
func (l *CancelOrderLogic) CancelOrder(in *omsclient.CancelOrderReq) (*omsclient.CancelOrderResp, error) {
	order := query.OmsOrderMain
	item, err := order.WithContext(l.ctx).Where(order.ID.Eq(in.OrderId), order.UserID.Eq(in.MemberId), order.IsDeleted.Eq(0)).First()

	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		logc.Errorf(l.ctx, "订单不存在, 请求参数：%+v, 异常信息: %s", in, err.Error())
		return nil, errors.New("订单不存在")
	case err != nil:
		logc.Errorf(l.ctx, "查询订单异常, 请求参数：%+v, 异常信息: %s", in, err.Error())
		return nil, errors.New("查询订单异常")
	}

	if item.OrderStatus != 1 {
		logc.Errorf(l.ctx, "订单不能取消, 订单状态：%d, 订单ID：%d", item.OrderStatus, item.ID)
		return nil, errors.New("订单不能取消")
	}

	_, err = order.WithContext(l.ctx).Where(order.ID.Eq(in.OrderId)).Update(order.OrderStatus, 5)
	if err != nil {
		logc.Errorf(l.ctx, "取消订单失败,参数:%+v,异常:%s", item, err.Error())
		return nil, errors.New("取消订单失败")
	}

	result, _ := query.OmsOrderItem.WithContext(l.ctx).Where(query.OmsOrderItem.OrderID.Eq(item.ID)).Find()

	var list []*omsclient.ReleaseSkuStockLockData
	for _, x := range result {

		list = append(list, &omsclient.ReleaseSkuStockLockData{
			ProductQuantity: x.SkuQuantity,
			ProductSkuId:    x.SkuID,
		})
	}

	var couponIds []int64
	p := query.OmsOrderPromotion
	_ = p.WithContext(l.ctx).Select(p.PromotionID).Where(p.OrderID.Eq(item.ID), p.IsDeleted.Eq(0), p.PromotionType.Eq(1)).Scan(&couponIds)

	message := map[string]any{"id": in.OrderId}
	body, _ := json.Marshal(message)
	err = l.svcCtx.RabbitMQ.SendMessage("order.cancel.exchange", "order.cancel.by.user.queue", "order.cancel.key", body)

	return &omsclient.CancelOrderResp{
		CouponIds:   couponIds,
		Integration: item.UsePoints,
		Data:        list,
	}, nil
}
