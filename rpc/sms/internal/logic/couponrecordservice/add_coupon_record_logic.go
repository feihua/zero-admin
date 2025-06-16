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
	"time"
)

// AddCouponRecordLogic 添加优惠券领取记录
/*
Author: LiuFeiHua
Date: 2025/06/11 10:44:50
*/
type AddCouponRecordLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddCouponRecordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddCouponRecordLogic {
	return &AddCouponRecordLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// AddCouponRecord 添加优惠券领取记录
// 1.查询优惠券是否存在
// 2.查询是否已经领取过优惠券了
// 3.添加领取优惠券记录
// 4.更新优惠券数量
func (l *AddCouponRecordLogic) AddCouponRecord(in *smsclient.AddCouponRecordReq) (*smsclient.AddCouponRecordResp, error) {
	// 1.查询优惠券是否存在
	coupon := query.SmsCoupon
	smsCoupon, err := coupon.WithContext(l.ctx).Where(coupon.ID.Eq(in.CouponId)).First()

	if err != nil {
		return nil, errors.New("优惠券不存在")
	}

	if smsCoupon.TotalCount-smsCoupon.ReceivedCount <= 0 {
		return nil, errors.New("优惠券已经领完了")
	}

	// //2.查询是否已经领取过优惠券了
	history := query.SmsCouponRecord
	count, _ := history.WithContext(l.ctx).Where(history.CouponID.Eq(in.CouponId)).Where(history.MemberID.Eq(in.MemberId)).Count()

	if count >= int64(smsCoupon.PerLimit) {
		return nil, errors.New("您已经领取过该优惠券")
	}

	// 3.添加领取优惠券记录
	item := &model.SmsCouponRecord{
		CouponID: in.CouponId, // 优惠券ID
		MemberID: in.MemberId, // 用户ID
		GetTime:  time.Now(),  // 领取时间
		GetType:  in.GetType,  // 获取类型：0->后台赠送；1->主动获取
		Status:   0,           // 状态：0-未使用，1-已使用，2-已过期，3-已失效
	}

	q := query.SmsCouponRecord
	err = q.WithContext(l.ctx).Create(item)
	if err != nil {
		logc.Errorf(l.ctx, "添加优惠券领取记录失败,参数:%+v,异常:%s", item, err.Error())
		return nil, errors.New("添加优惠券领取记录失败")
	}

	// 4.更新优惠券数量
	smsCoupon.ReceivedCount = smsCoupon.ReceivedCount + 1
	_, err = coupon.WithContext(l.ctx).Where(coupon.ID.Eq(in.CouponId)).Updates(smsCoupon)
	if err != nil {
		logc.Errorf(l.ctx, "更新优惠券数量失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("更新优惠券数量失败")
	}

	return &smsclient.AddCouponRecordResp{}, nil
}
