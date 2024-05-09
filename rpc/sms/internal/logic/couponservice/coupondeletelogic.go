package couponservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/sms/gen/query"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"

	"github.com/feihua/zero-admin/rpc/sms/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

// CouponDeleteLogic 优惠券
/*
Author: LiuFeiHua
Date: 2024/5/8 10:12
*/
type CouponDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCouponDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CouponDeleteLogic {
	return &CouponDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// CouponDelete 删除优惠券
func (l *CouponDeleteLogic) CouponDelete(in *smsclient.CouponDeleteReq) (*smsclient.CouponDeleteResp, error) {
	q := query.SmsCoupon
	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Delete()

	if err != nil {
		return nil, err
	}

	return &smsclient.CouponDeleteResp{}, nil
}
