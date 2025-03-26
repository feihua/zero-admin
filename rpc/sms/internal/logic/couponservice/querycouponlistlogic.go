package couponservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/pkg/time_util"
	"github.com/feihua/zero-admin/rpc/sms/gen/query"
	"github.com/zeromicro/go-zero/core/logc"
	"time"

	"github.com/feihua/zero-admin/rpc/sms/internal/svc"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryCouponListLogic 查询优惠券
/*
Author: LiuFeiHua
Date: 2024/6/12 17:37
*/
type QueryCouponListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryCouponListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryCouponListLogic {
	return &QueryCouponListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryCouponList 查询优惠券
func (l *QueryCouponListLogic) QueryCouponList(in *smsclient.QueryCouponListReq) (*smsclient.QueryCouponListResp, error) {
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

	result, count, err := q.FindByPage(int((in.Current-1)*in.PageSize), int(in.PageSize))

	if err != nil {
		logc.Errorf(l.ctx, "查询优惠券失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("查询优惠券失败")
	}

	var list []*smsclient.QueryCouponListData
	for _, coupon := range result {

		list = append(list, &smsclient.QueryCouponListData{
			Id:           coupon.ID,                              //
			Type:         coupon.Type,                            // 优惠券类型；0->全场赠券；1->会员赠券；2->购物赠券；3->注册赠券
			Name:         coupon.Name,                            // 名称
			Platform:     coupon.Platform,                        // 使用平台：0->全部；1->移动；2->PC
			Count:        coupon.Count,                           // 数量
			Amount:       coupon.Amount,                          // 金额
			PerLimit:     coupon.PerLimit,                        // 每人限领张数
			MinPoint:     coupon.MinPoint,                        // 使用门槛；0表示无门槛
			StartTime:    time_util.TimeToStr(coupon.StartTime),  // 开始时间
			EndTime:      time_util.TimeToStr(coupon.EndTime),    // 结束时间
			UseType:      coupon.UseType,                         // 使用类型：0->全场通用；1->指定分类；2->指定商品
			Note:         coupon.Note,                            // 备注
			PublishCount: coupon.PublishCount,                    // 发行数量
			UseCount:     coupon.UseCount,                        // 已使用数量
			ReceiveCount: coupon.ReceiveCount,                    // 领取数量
			EnableTime:   time_util.TimeToStr(coupon.EnableTime), // 可以领取的日期
			Code:         coupon.Code,                            // 优惠码
			MemberLevel:  coupon.MemberLevel,                     // 可领取的会员类型：0->无限时
		})

	}

	return &smsclient.QueryCouponListResp{
		Total: count,
		List:  list,
	}, nil

}
