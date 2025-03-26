package couponhistoryservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/pkg/time_util"
	"github.com/feihua/zero-admin/rpc/sms/gen/model"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/rpc/sms/internal/svc"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryMemberCouponListLogic
/*
Author: LiuFeiHua
Date: 2023/12/5 17:47
*/
type QueryMemberCouponListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryMemberCouponListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryMemberCouponListLogic {
	return &QueryMemberCouponListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryMemberCouponList 获取会员优惠券
func (l *QueryMemberCouponListLogic) QueryMemberCouponList(in *smsclient.QueryMemberCouponListReq) (*smsclient.QueryMemberCouponListResp, error) {
	var result []model.SmsCoupon
	query := `select t2.*
from sms_coupon_history t1
         left join sms_coupon t2 on t1.coupon_id = t2.id
where t1.member_id = ?
  and t1.use_status = ?`
	db := l.svcCtx.DB
	err := db.Where(l.ctx).Raw(query, in.MemberId, in.UseStatus).Find(&result).Error

	if err != nil {
		logc.Errorf(l.ctx, "获取会员优惠券失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("获取会员优惠券失败")
	}

	var list []*smsclient.QueryCouponData
	for _, coupon := range result {

		list = append(list, &smsclient.QueryCouponData{
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

	return &smsclient.QueryMemberCouponListResp{
		List: list,
	}, nil

}
