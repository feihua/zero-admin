package couponhistoryservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/sms/gen/model"
	"github.com/feihua/zero-admin/rpc/sms/gen/query"
	"github.com/feihua/zero-admin/rpc/sms/internal/svc"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

// CouponHistoryAddLogic
/*
Author: LiuFeiHua
Date: 2024/5/6 11:34
*/
type CouponHistoryAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCouponHistoryAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CouponHistoryAddLogic {
	return &CouponHistoryAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// CouponHistoryAdd 添加优惠券使用记录
// 1.查询优惠券是否存在
// 2.查询是否已经领取过优惠券了
// 3.添加领取优惠券记录
// 4.更新优惠券数量
func (l *CouponHistoryAddLogic) CouponHistoryAdd(in *smsclient.CouponHistoryAddReq) (*smsclient.CouponHistoryAddResp, error) {
	//1.查询优惠券是否存在
	coupon := query.SmsCoupon
	smsCoupon, err := coupon.WithContext(l.ctx).Where(coupon.ID.Eq(in.CouponId)).First()

	if err != nil {
		return nil, errors.New("优惠券不存在")
	}

	if smsCoupon.Count <= 0 {
		return nil, errors.New("优惠券已经领完了")
	}

	////2.查询是否已经领取过优惠券了
	history := query.SmsCouponHistory
	count, _ := history.WithContext(l.ctx).Where(history.CouponID.Eq(in.CouponId)).Where(history.MemberID.Eq(in.MemberId)).Count()

	if count >= int64(smsCoupon.PerLimit) {
		return nil, errors.New("您已经领取过该优惠券")
	}

	//3.添加领取优惠券记录
	err = query.SmsCouponHistory.WithContext(l.ctx).Create(&model.SmsCouponHistory{
		CouponID:       in.CouponId,
		MemberID:       in.MemberId,
		CouponCode:     in.CouponCode,
		MemberNickname: in.MemberNickname,
		GetType:        in.GetType,
		CreateTime:     time.Now(),
		UseStatus:      in.UseStatus,
	})
	if err != nil {
		return nil, err
	}

	//4.更新优惠券数量
	smsCoupon.Count = smsCoupon.Count - 1
	smsCoupon.ReceiveCount = smsCoupon.ReceiveCount + 1
	_, err = coupon.WithContext(l.ctx).Where(coupon.ID.Eq(in.CouponId)).Updates(smsCoupon)
	if err != nil {
		return nil, err
	}

	return &smsclient.CouponHistoryAddResp{}, nil
}
