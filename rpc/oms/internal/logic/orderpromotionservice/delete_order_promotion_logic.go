package orderpromotionservicelogic

import (
	"context"
	"errors"

	"github.com/feihua/zero-admin/rpc/oms/gen/query"
	"github.com/feihua/zero-admin/rpc/oms/internal/svc"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
)

// DeleteOrderPromotionLogic 删除订单优惠信息
/*
Author: LiuFeiHua
Date: 2025/07/03 09:32:56
*/
type DeleteOrderPromotionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteOrderPromotionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteOrderPromotionLogic {
	return &DeleteOrderPromotionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// DeleteOrderPromotion 删除订单优惠信息
func (l *DeleteOrderPromotionLogic) DeleteOrderPromotion(in *omsclient.DeleteOrderPromotionReq) (*omsclient.DeleteOrderPromotionResp, error) {
	q := query.OmsOrderPromotion

	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Delete()

	if err != nil {
		logc.Errorf(l.ctx, "删除订单优惠信息失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("删除订单优惠信息失败")
	}

	return &omsclient.DeleteOrderPromotionResp{}, nil
}
