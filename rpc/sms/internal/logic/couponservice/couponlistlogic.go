package couponservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/sms/gen/query"
	"github.com/feihua/zero-admin/rpc/sms/internal/svc"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
	"time"
)

// CouponListLogic 优惠券
/*
Author: LiuFeiHua
Date: 2024/5/15 9:22
*/
type CouponListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCouponListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CouponListLogic {
	return &CouponListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// CouponList 查询优惠券列表
func (l *CouponListLogic) CouponList(in *smsclient.CouponListReq) (*smsclient.CouponListResp, error) {

	q := query.SmsCoupon.WithContext(l.ctx)
	if len(in.Name) > 0 {
		q = q.Where(query.SmsCoupon.Name.Like("%" + in.Name + "%"))
	}

	if in.Type != 4 {
		q = q.Where(query.SmsCoupon.Type.Eq(in.Type))
	}
	if in.Platform != 3 {
		q = q.Where(query.SmsCoupon.Platform.Eq(in.Platform))
	}
	if in.UseType != 3 {
		q = q.Where(query.SmsCoupon.UseType.Eq(in.UseType))
	}
	if len(in.StartTime) > 0 {
		startTime, _ := time.Parse("2006-01-02 15:04:05", in.StartTime)
		q = q.Where(query.SmsCoupon.StartTime.Gte(startTime))
	}
	if len(in.EndTime) > 0 {
		endTime, _ := time.Parse("2006-01-02 15:04:05", in.EndTime)
		q = q.Where(query.SmsCoupon.EndTime.Lte(endTime))
	}

	offset := (in.Current - 1) * in.PageSize
	result, err := q.Offset(int(offset)).Limit(int(in.PageSize)).Find()
	count, err := q.Count()

	if err != nil {
		logc.Errorf(l.ctx, "查询优惠券列表信息失败,参数：%+v,异常:%s", in, err.Error())
		return nil, err
	}

	var list []*smsclient.CouponListData
	for _, coupon := range result {

		list = append(list, &smsclient.CouponListData{
			Id:           coupon.ID,
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
		})

		logc.Infof(l.ctx, "查询优惠券列表信息,参数：%+v,响应：%+v", in, list)
	}

	return &smsclient.CouponListResp{
		Total: count,
		List:  list,
	}, nil
}
