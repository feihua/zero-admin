package orderpromotionservicelogic

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

// QueryOrderPromotionDetailLogic 查询订单优惠信息详情
/*
Author: LiuFeiHua
Date: 2025/07/03 09:32:56
*/
type QueryOrderPromotionDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryOrderPromotionDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryOrderPromotionDetailLogic {
	return &QueryOrderPromotionDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryOrderPromotionDetail 查询订单优惠信息详情
func (l *QueryOrderPromotionDetailLogic) QueryOrderPromotionDetail(in *omsclient.QueryOrderPromotionDetailReq) (*omsclient.QueryOrderPromotionDetailResp, error) {
	item, err := query.OmsOrderPromotion.WithContext(l.ctx).Where(query.OmsOrderPromotion.ID.Eq(in.Id)).First()

	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		logc.Errorf(l.ctx, "订单优惠信息不存在, 请求参数：%+v, 异常信息: %s", in, err.Error())
		return nil, errors.New("订单优惠信息不存在")
	case err != nil:
		logc.Errorf(l.ctx, "查询订单优惠信息异常, 请求参数：%+v, 异常信息: %s", in, err.Error())
		return nil, errors.New("查询订单优惠信息异常")
	}

	data := &omsclient.QueryOrderPromotionDetailResp{
		Id:             item.ID,                              // 主键ID
		OrderId:        item.OrderID,                         // 订单ID
		OrderNo:        item.OrderNo,                         // 订单编号
		PromotionType:  item.PromotionType,                   // 优惠类型：1-优惠券，2-积分抵扣，3-会员折扣，4-促销活动
		PromotionName:  item.PromotionName,                   // 优惠名称
		DiscountAmount: float32(item.DiscountAmount),         // 优惠金额
		CreateTime:     time_util.TimeToStr(item.CreateTime), //
	}

	if item.PromotionID != nil {
		data.PromotionId = *item.PromotionID
	}
	return data, nil
}
