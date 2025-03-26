package couponhistoryservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/sms/gen/model"
	"github.com/feihua/zero-admin/rpc/sms/gen/query"
	"github.com/feihua/zero-admin/rpc/sms/internal/svc"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddCouponHistoryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddCouponHistoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddCouponHistoryLogic {
	return &AddCouponHistoryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// AddCouponHistory 添加优惠券使用、领取历史表
// 1.查询优惠券是否存在
// 2.查询是否已经领取过优惠券了
// 3.添加领取优惠券记录
// 4.更新优惠券数量
func (l *AddCouponHistoryLogic) AddCouponHistory(in *smsclient.AddCouponHistoryReq) (*smsclient.AddCouponHistoryResp, error) {
	// 1.查询优惠券是否存在
	coupon := query.SmsCoupon
	smsCoupon, err := coupon.WithContext(l.ctx).Where(coupon.ID.Eq(in.CouponId)).First()

	if err != nil {
		return nil, errors.New("优惠券不存在")
	}

	if smsCoupon.Count <= 0 {
		return nil, errors.New("优惠券已经领完了")
	}

	// //2.查询是否已经领取过优惠券了
	history := query.SmsCouponHistory
	count, _ := history.WithContext(l.ctx).Where(history.CouponID.Eq(in.CouponId)).Where(history.MemberID.Eq(in.MemberId)).Count()

	if count >= int64(smsCoupon.PerLimit) {
		return nil, errors.New("您已经领取过该优惠券")
	}

	// 3.添加领取优惠券记录
	err = query.SmsCouponHistory.WithContext(l.ctx).Create(&model.SmsCouponHistory{
		CouponID:       in.CouponId,       // 优惠券id
		MemberID:       in.MemberId,       // 会员id
		CouponCode:     in.CouponCode,     // 优惠码
		MemberNickname: in.MemberNickname, // 领取人昵称
		GetType:        in.GetType,        // 获取类型：0->后台赠送；1->主动获取
		UseStatus:      in.UseStatus,      // 使用状态：0->未使用；1->已使用；2->已过期
	})
	if err != nil {
		logc.Errorf(l.ctx, "添加领取优惠券记录失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("添加领取优惠券记录失败")
	}

	// 4.更新优惠券数量
	smsCoupon.Count = smsCoupon.Count - 1
	smsCoupon.ReceiveCount = smsCoupon.ReceiveCount + 1
	_, err = coupon.WithContext(l.ctx).Where(coupon.ID.Eq(in.CouponId)).Updates(smsCoupon)
	if err != nil {
		logc.Errorf(l.ctx, "更新优惠券数量失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("更新优惠券数量失败")
	}

	return &smsclient.AddCouponHistoryResp{}, nil
}
