package couponhistoryservicelogic

import (
	"context"

	"zero-admin/rpc/sms/internal/svc"
	"zero-admin/rpc/sms/smsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

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
	count, _ := l.svcCtx.SmsCouponHistoryModel.Count(l.ctx, &smsclient.CouponHistoryListReq{
		Current:   1,
		PageSize:  100,
		CouponId:  0,
		MemberId:  in.MemberId,
		UseStatus: 0, //'使用状态：0->未使用；1->已使用；2->已过期',
	})

	return &smsclient.CouponCountResp{
		Total: count,
	}, nil
}
