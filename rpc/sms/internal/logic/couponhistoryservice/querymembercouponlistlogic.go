package couponhistoryservicelogic

import (
	"context"
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
		logc.Errorf(l.ctx, "查询优惠券列表信息失败,参数：%+v,异常:%s", in, err.Error())
		return nil, err
	}

	var list []*smsclient.QueryCouponData
	for _, coupon := range result {

		list = append(list, &smsclient.QueryCouponData{
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

	return &smsclient.QueryMemberCouponListResp{
		List: list,
	}, nil

}
