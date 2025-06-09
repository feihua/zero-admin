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
)

// QueryCouponRecordListLogic 查询优惠券领取记录列表
/*
Author: LiuFeiHua
Date: 2025/06/11 10:44:50
*/
type QueryCouponRecordListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryCouponRecordListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryCouponRecordListLogic {
	return &QueryCouponRecordListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryCouponRecordList 查询优惠券领取记录列表
func (l *QueryCouponRecordListLogic) QueryCouponRecordList(in *smsclient.QueryCouponRecordListReq) (*smsclient.QueryCouponRecordListResp, error) {
	couponRecord := query.SmsCouponRecord
	q := couponRecord.WithContext(l.ctx)
	if in.CouponId != 2 {
		q = q.Where(couponRecord.CouponID.Eq(in.CouponId))
	}
	if in.MemberId != 2 {
		q = q.Where(couponRecord.MemberID.Eq(in.MemberId))
	}

	if in.GetType != 2 {
		q = q.Where(couponRecord.GetType.Eq(in.GetType))
	}

	if in.OrderId != 2 {
		q = q.Where(couponRecord.OrderID.Eq(in.OrderId))
	}

	if in.Status != 2 {
		q = q.Where(couponRecord.Status.Eq(in.Status))
	}

	result, count, err := q.FindByPage(int((in.PageNum-1)*in.PageSize), int(in.PageSize))

	if err != nil {
		logc.Errorf(l.ctx, "查询优惠券领取记录列表失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("查询优惠券领取记录列表失败")
	}

	var list []*smsclient.CouponRecordListData

	for _, item := range result {
		list = append(list, &smsclient.CouponRecordListData{
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

		})
	}

	return &smsclient.QueryCouponRecordListResp{
		Total: count,
		List:  list,
	}, nil
}
