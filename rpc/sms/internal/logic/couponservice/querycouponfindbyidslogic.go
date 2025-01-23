package couponservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/pkg/time_util"
	"github.com/feihua/zero-admin/rpc/sms/gen/query"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/rpc/sms/internal/svc"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryCouponFindByIdsLogic 根据优惠券ids查询优惠券
/*
Author: LiuFeiHua
Date: 2024/6/12 17:36
*/
type QueryCouponFindByIdsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryCouponFindByIdsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryCouponFindByIdsLogic {
	return &QueryCouponFindByIdsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryCouponFindByIds 根据优惠券ids查询优惠券
func (l *QueryCouponFindByIdsLogic) QueryCouponFindByIds(in *smsclient.QueryCouponFindByIdsReq) (*smsclient.QueryCouponFindByIdsResp, error) {
	q := query.SmsCoupon
	result, err := q.WithContext(l.ctx).Where(q.ID.In(in.CouponIds...)).Find()

	if err != nil {
		logc.Errorf(l.ctx, "查询优惠券列表信息失败,参数：%+v,异常:%s", in, err.Error())
		return nil, err
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

		logc.Infof(l.ctx, "查询优惠券列表信息,参数：%+v,响应：%+v", in, list)
	}

	return &smsclient.QueryCouponFindByIdsResp{
		Total: 0,
		List:  list,
	}, nil

}
