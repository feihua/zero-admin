package couponrecordservicelogic

import (
	"context"
	"time"

	"github.com/feihua/zero-admin/pkg/time_util"
	"github.com/feihua/zero-admin/rpc/sms/gen/model"
	"github.com/feihua/zero-admin/rpc/sms/gen/query"
	"github.com/feihua/zero-admin/rpc/sms/internal/svc"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryMemberCouponListLogic 获取会员优惠券
/*
Author: LiuFeiHua
Date: 2025/6/11 14:48
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
	sql := `select t2.*
				from sms_coupon_record t1
						 left join sms_coupon t2 on t1.coupon_id = t2.id
				where t1.member_id = ?
				  and t1.status = ?;`
	db := l.svcCtx.DB
	err := db.Where(l.ctx).Raw(sql, in.MemberId, in.Status).Find(&result).Error

	if err != nil {
		logc.Errorf(l.ctx, "查询优惠券领取记录列表失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("查询优惠券领取记录列表失败")
	}

	var list []*smsclient.QueryCouponData

	for _, item := range result {
		x := &smsclient.QueryCouponData{
			Id:          item.ID,                             // 优惠券ID
			TypeId:      item.TypeID,                         // 优惠券类型ID
			Name:        item.Name,                           // 优惠券名称
			Code:        item.Code,                           // 优惠券码
			Amount:      float32(item.Amount),                // 优惠金额/折扣率
			MinAmount:   float32(item.MinAmount),             // 最低使用金额
			StartTime:   time_util.TimeToStr(item.StartTime), // 生效时间
			EndTime:     time_util.TimeToStr(item.EndTime),   // 失效时间
			PerLimit:    item.PerLimit,                       // 每人限领数量
			Status:      item.Status,                         // 状态：0-未开始，1-进行中，2-已结束，3-已取消
			Description: item.Description,                    // 使用说明
		}

		isExpired := item.EndTime.Before(time.Now())
		if in.Status == 0 && isExpired {
			_, _ = query.SmsCouponRecord.WithContext(l.ctx).Where(query.SmsCouponRecord.CouponID.Eq(item.ID)).Update(query.SmsCouponRecord.Status, 2)
			continue
		}

		scope, _ := query.SmsCouponScope.WithContext(l.ctx).Where(query.SmsCouponScope.CouponID.Eq(item.ID)).First()
		x.ScopeType = scope.ScopeType
		list = append(list, x)

	}
	return &smsclient.QueryMemberCouponListResp{
		List: list,
	}, nil
}
