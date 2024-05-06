package couponhistoryservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/sms/gen/query"

	"github.com/feihua/zero-admin/rpc/sms/internal/svc"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// CouponCountLogic
/*
Author: LiuFeiHua
Date: 2023/11/28 15:16
*/
type CouponCountLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCouponCountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CouponCountLogic {
	return &CouponCountLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// CouponCount 登录时获取用户还没有使用的获取优惠券数量
func (l *CouponCountLogic) CouponCount(in *smsclient.CouponCountReq) (*smsclient.CouponCountResp, error) {
	q := query.SmsCouponHistory
	count, _ := q.WithContext(l.ctx).Where(q.MemberID.Eq(in.MemberId)).Count()

	return &smsclient.CouponCountResp{
		Total: count,
	}, nil
}
