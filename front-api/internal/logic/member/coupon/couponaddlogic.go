package coupon

import (
	"context"
	"errors"
	"time"
	"zero-admin/rpc/sms/smsclient"

	"zero-admin/front-api/internal/svc"
	"zero-admin/front-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// CouponAddLogic
/*
Author: LiuFeiHua
Date: 2023/11/30 16:32
*/
type CouponAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCouponAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CouponAddLogic {
	return &CouponAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// CouponAdd 领取指定优惠券
func (l *CouponAddLogic) CouponAdd(req *types.AddCouponReq) (resp *types.AddCouponResp, err error) {
	coupon, err := l.svcCtx.CouponService.CouponFindById(l.ctx, &smsclient.CouponFindByIdReq{CouponId: req.CouponId})

	if err != nil {
		return nil, errors.New("优惠券不存在")
	}

	if coupon.Count <= 0 {
		return nil, errors.New("优惠券已经领完了")
	}

	couponHistoryList, _ := l.svcCtx.CouponHistoryService.CouponHistoryList(l.ctx, &smsclient.CouponHistoryListReq{
		Current:  1,
		PageSize: 100,
		CouponId: req.CouponId,
		MemberId: l.ctx.Value("memberId").(int64),
	})

	if couponHistoryList.Total >= coupon.PerLimit {
		return nil, errors.New("您已经领取过该优惠券")
	}

	_, err = l.svcCtx.CouponHistoryService.CouponHistoryAdd(l.ctx, &smsclient.CouponHistoryAddReq{
		CouponId:       req.CouponId,
		MemberId:       l.ctx.Value("memberId").(int64),
		CouponCode:     "fsdf",
		MemberNickname: l.ctx.Value("memberName").(string),
		GetType:        1,
		CreateTime:     time.Now().Format("2006-01-02 15:04:05"),
		UseStatus:      0,
	})

	if err != nil {
		return nil, errors.New("添加优惠券异常")
	}

	coupon.Count = coupon.Count - 1
	coupon.ReceiveCount = coupon.ReceiveCount + 1

	_, err = l.svcCtx.CouponService.CouponUpdate(l.ctx, &smsclient.CouponUpdateReq{
		Id:           coupon.Id,
		Type:         coupon.Type,
		Name:         coupon.Name,
		Platform:     coupon.Platform,
		Count:        coupon.Count,
		Amount:       coupon.Amount,
		PerLimit:     coupon.PerLimit,
		MinPoint:     coupon.MinPoint,
		StartTime:    coupon.StartTime,
		EndTime:      coupon.EndTime,
		UseType:      coupon.UseType,
		Note:         coupon.Note,
		PublishCount: coupon.PublishCount,
		UseCount:     coupon.UseCount,
		ReceiveCount: coupon.ReceiveCount,
		EnableTime:   coupon.EnableTime,
		Code:         coupon.Code,
		MemberLevel:  coupon.MemberLevel,
	})
	if err != nil {
		return nil, errors.New("更新优惠券异常")
	}

	return &types.AddCouponResp{
		Code:    0,
		Message: "领取优惠券成功",
	}, nil
}
