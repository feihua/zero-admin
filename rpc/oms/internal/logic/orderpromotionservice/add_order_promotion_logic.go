package orderpromotionservicelogic

import (
	"context"
	"errors"

	"github.com/feihua/zero-admin/rpc/oms/gen/model"
	"github.com/feihua/zero-admin/rpc/oms/gen/query"
	"github.com/feihua/zero-admin/rpc/oms/internal/svc"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
)

// AddOrderPromotionLogic 添加订单优惠信息
/*
Author: LiuFeiHua
Date: 2025/07/03 09:32:56
*/
type AddOrderPromotionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddOrderPromotionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddOrderPromotionLogic {
	return &AddOrderPromotionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// AddOrderPromotion 添加订单优惠信息
func (l *AddOrderPromotionLogic) AddOrderPromotion(in *omsclient.AddOrderPromotionReq) (*omsclient.AddOrderPromotionResp, error) {
	q := query.OmsOrderPromotion

	item := &model.OmsOrderPromotion{
		OrderID:        in.OrderId,                 // 订单ID
		OrderNo:        in.OrderNo,                 // 订单编号
		PromotionType:  in.PromotionType,           // 优惠类型：1-优惠券，2-积分抵扣，3-会员折扣，4-促销活动
		PromotionID:    &in.PromotionId,            // 优惠ID
		PromotionName:  in.PromotionName,           // 优惠名称
		DiscountAmount: float64(in.DiscountAmount), // 优惠金额
	}

	err := q.WithContext(l.ctx).Create(item)
	if err != nil {
		logc.Errorf(l.ctx, "添加订单优惠信息失败,参数:%+v,异常:%s", item, err.Error())
		return nil, errors.New("添加订单优惠信息失败")
	}

	return &omsclient.AddOrderPromotionResp{}, nil
}
