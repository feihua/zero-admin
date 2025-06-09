package couponrecordservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/pkg/time_util"
	"github.com/feihua/zero-admin/rpc/sms/gen/query"
	"github.com/feihua/zero-admin/rpc/sms/internal/svc"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

// QueryCouponRecordDetailLogic 查询优惠券领取记录详情
/*
Author: LiuFeiHua
Date: 2025/06/11 10:44:50
*/
type QueryCouponRecordDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryCouponRecordDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryCouponRecordDetailLogic {
	return &QueryCouponRecordDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryCouponRecordDetail 查询优惠券领取记录详情
func (l *QueryCouponRecordDetailLogic) QueryCouponRecordDetail(in *smsclient.QueryCouponRecordDetailReq) (*smsclient.QueryCouponRecordDetailResp, error) {
	item, err := query.SmsCouponRecord.WithContext(l.ctx).Where(query.SmsCouponRecord.ID.Eq(in.Id)).First()

	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		logc.Errorf(l.ctx, "优惠券领取记录不存在, 请求参数：%+v, 异常信息: %s", in, err.Error())
		return nil, errors.New("优惠券领取记录不存在")
	case err != nil:
		logc.Errorf(l.ctx, "查询优惠券领取记录异常, 请求参数：%+v, 异常信息: %s", in, err.Error())
		return nil, errors.New("查询优惠券领取记录异常")
	}

	data := &smsclient.QueryCouponRecordDetailResp{
		Id:             item.ID,                                  //
		CouponId:       item.CouponID,                            // 优惠券ID
		MemberId:       item.MemberID,                            // 用户ID
		GetTime:        time_util.TimeToStr(item.GetTime),        // 领取时间
		GetType:        item.GetType,                             // 获取类型：0->后台赠送；1->主动获取
		UseTime:        time_util.TimeToString(item.UseTime),     // 使用时间
		OrderId:        item.OrderID,                             // 使用订单ID
		OrderAmount:    float32(item.OrderAmount),                // 订单金额
		DiscountAmount: float32(item.DiscountAmount),             // 优惠金额
		Status:         item.Status,                              // 状态：0-未使用，1-已使用，2-已过期，3-已失效
		InvalidTime:    time_util.TimeToString(item.InvalidTime), // 失效时间
		InvalidReason:  item.InvalidReason,                       // 失效原因
		CreateTime:     time_util.TimeToStr(item.CreateTime),     // 创建时间
	}

	return data, nil
}
