package couponhistoryservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/sms/gen/query"

	"github.com/feihua/zero-admin/rpc/sms/internal/svc"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// DeleteCouponHistoryLogic 删除优惠券使用、领取历史表
/*
Author: LiuFeiHua
Date: 2024/6/12 17:28
*/
type DeleteCouponHistoryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteCouponHistoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteCouponHistoryLogic {
	return &DeleteCouponHistoryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// DeleteCouponHistory 删除优惠券使用、领取历史表
func (l *DeleteCouponHistoryLogic) DeleteCouponHistory(in *smsclient.DeleteCouponHistoryReq) (*smsclient.DeleteCouponHistoryResp, error) {
	q := query.SmsCouponHistory
	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Delete()
	if err != nil {
		return nil, err
	}

	return &smsclient.DeleteCouponHistoryResp{}, nil
}
