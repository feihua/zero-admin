package couponhistoryservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/sms/gen/query"

	"github.com/feihua/zero-admin/rpc/sms/internal/svc"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateCouponHistoryStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateCouponHistoryStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCouponHistoryStatusLogic {
	return &UpdateCouponHistoryStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateCouponHistoryStatus 更新优惠券使用、领取历史表状态
// 1.更新用户优惠券状态
// 2.更新优惠券数量
func (l *UpdateCouponHistoryStatusLogic) UpdateCouponHistoryStatus(in *smsclient.UpdateCouponHistoryStatusReq) (*smsclient.UpdateCouponHistoryStatusResp, error) {
	// 1.更新用户优惠券状态
	q := query.SmsCouponHistory
	_, err := q.WithContext(l.ctx).Where(q.MemberID.Eq(in.MemberId), q.CouponID.Eq(in.CouponId)).Update(q.UseStatus, in.UseStatus)
	if err != nil {
		return nil, err
	}

	// 2.更新优惠券数量
	var sql string
	if in.UseStatus == 0 {
		sql = "update sms_coupon set use_count=use_count+1 where `id` = ? and use_count<count"
	} else {
		sql = "update sms_coupon set use_count=use_count-1 where `id` = ? and use_count>0"
	}

	db := l.svcCtx.DB
	err = db.Exec(sql, in.CouponId).Error
	if err != nil {
		return nil, err
	}

	return &smsclient.UpdateCouponHistoryStatusResp{}, nil
}
