package couponservicelogic

import (
	"context"
	"errors"

	"zero-admin/rpc/sms/internal/svc"
	"zero-admin/rpc/sms/smsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type CouponFindByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCouponFindByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CouponFindByIdLogic {
	return &CouponFindByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// CouponFindById 根据优惠券id查询优惠券
func (l *CouponFindByIdLogic) CouponFindById(in *smsclient.CouponFindByIdReq) (*smsclient.CouponFindByIdResp, error) {
	coupon, err := l.svcCtx.SmsCouponModel.FindOne(l.ctx, in.CouponId)

	if err != nil {
		errors.New("根据优惠券id查询优惠券异常")
	}

	return &smsclient.CouponFindByIdResp{
		Id:           coupon.Id,
		Type:         coupon.Type,
		Name:         coupon.Name,
		Platform:     coupon.Platform,
		Count:        coupon.Count,
		Amount:       coupon.Amount,
		PerLimit:     coupon.PerLimit,
		MinPoint:     coupon.MinPoint,
		StartTime:    coupon.StartTime.Format("2006-01-02 15:04:05"),
		EndTime:      coupon.EndTime.Format("2006-01-02 15:04:05"),
		UseType:      coupon.UseType,
		Note:         coupon.Note,
		PublishCount: coupon.PublishCount,
		UseCount:     coupon.UseCount,
		ReceiveCount: coupon.ReceiveCount,
		EnableTime:   coupon.EnableTime.Format("2006-01-02 15:04:05"),
		Code:         coupon.Code,
		MemberLevel:  coupon.MemberLevel,
	}, nil
}
