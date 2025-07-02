package couponrecordservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/sms/gen/model"
	"github.com/feihua/zero-admin/rpc/sms/gen/query"
	"github.com/feihua/zero-admin/rpc/sms/internal/svc"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
	"time"
)

// UpdateCouponRecordLogic 更新优惠券领取记录
/*
Author: LiuFeiHua
Date: 2025/06/11 10:44:50
*/
type UpdateCouponRecordLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateCouponRecordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCouponRecordLogic {
	return &UpdateCouponRecordLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateCouponRecord 更新优惠券领取记录
func (l *UpdateCouponRecordLogic) UpdateCouponRecord(in *smsclient.UpdateCouponRecordReq) (*smsclient.UpdateCouponRecordResp, error) {
	q := query.SmsCouponRecord

	for _, couponId := range in.CouponIds {
		// 1.查询优惠券领取记录是否已存在
		coupon, err := q.WithContext(l.ctx).Where(q.MemberID.Eq(in.MemberId), q.CouponID.Eq(couponId)).First()

		switch {
		case errors.Is(err, gorm.ErrRecordNotFound):
			logc.Errorf(l.ctx, "优惠券领取记录不存在, 请求参数：%+v, 异常信息: %s", in, err.Error())
			return nil, errors.New("优惠券领取记录不存在")
		case err != nil:
			logc.Errorf(l.ctx, "查询优惠券领取记录异常, 请求参数：%+v, 异常信息: %s", in, err.Error())
			return nil, errors.New("查询优惠券领取记录异常")
		}

		now := time.Now()
		item := &model.SmsCouponRecord{
			ID:       coupon.ID,   //
			CouponID: couponId,    // 优惠券ID
			MemberID: in.MemberId, // 用户ID
			Status:   in.Status,   // 状态：0-未使用，1-已使用，2-已过期，3-已失效
		}

		if in.Status == 1 {
			item.UseTime = &now
			item.OrderID = in.OrderId
			item.OrderAmount = float64(in.OrderAmount)
			item.DiscountAmount = float64(in.DiscountAmount)
		}

		if in.Status == 3 {
			invalidTime, _ := time.Parse("2006-01-02 15:04:05", in.InvalidTime)
			item.InvalidTime = &invalidTime
			item.InvalidReason = in.InvalidReason
		}

		// 2.优惠券领取记录存在时,则直接更新优惠券领取记录
		_, err = q.WithContext(l.ctx).Updates(item)

		if err != nil {
			logc.Errorf(l.ctx, "更新优惠券领取记录失败,参数:%+v,异常:%s", item, err.Error())
			return nil, errors.New("更新优惠券领取记录失败")
		}

		if in.Status == 2 || in.Status == 3 {
			return &smsclient.UpdateCouponRecordResp{}, nil
		}

		c := query.SmsCoupon
		smsCoupon, err := c.WithContext(l.ctx).Where(c.ID.Eq(couponId)).First()
		if err != nil {
			logc.Errorf(l.ctx, "更新优惠券使用失败,参数:%+v,异常:%s", in, err.Error())
			return nil, errors.New("更新优惠券使用失败")
		}

		count := smsCoupon.UsedCount
		if in.Status == 0 {
			count = count - 1
		} else {
			count = count + 1
		}
		_, err = c.WithContext(l.ctx).Where(c.ID.Eq(couponId)).Update(c.UsedCount, count)

		if err != nil {
			logc.Errorf(l.ctx, "更新优惠券使用失败,参数:%+v,异常:%s", in, err.Error())
			return nil, errors.New("更新优惠券使用失败")
		}

	}
	return &smsclient.UpdateCouponRecordResp{}, nil
}
