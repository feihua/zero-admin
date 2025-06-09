package couponservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/pkg/pointerprocess"
	"github.com/feihua/zero-admin/pkg/time_util"
	"github.com/feihua/zero-admin/rpc/sms/gen/query"
	"github.com/feihua/zero-admin/rpc/sms/internal/svc"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

// QueryCouponDetailLogic 查询优惠券详情
/*
Author: LiuFeiHua
Date: 2025/06/11 10:44:50
*/
type QueryCouponDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryCouponDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryCouponDetailLogic {
	return &QueryCouponDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryCouponDetail 查询优惠券详情
func (l *QueryCouponDetailLogic) QueryCouponDetail(in *smsclient.QueryCouponDetailReq) (*smsclient.QueryCouponDetailResp, error) {
	item, err := query.SmsCoupon.WithContext(l.ctx).Where(query.SmsCoupon.ID.Eq(in.Id)).First()

	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		logc.Errorf(l.ctx, "优惠券不存在, 请求参数：%+v, 异常信息: %s", in, err.Error())
		return nil, errors.New("优惠券不存在")
	case err != nil:
		logc.Errorf(l.ctx, "查询优惠券异常, 请求参数：%+v, 异常信息: %s", in, err.Error())
		return nil, errors.New("查询优惠券异常")
	}

	data := &smsclient.QueryCouponDetailResp{
		Id:            item.ID,                                          // 优惠券ID
		TypeId:        item.TypeID,                                      // 优惠券类型ID
		Name:          item.Name,                                        // 优惠券名称
		Code:          item.Code,                                        // 优惠券码
		Amount:        float32(item.Amount),                             // 优惠金额/折扣率
		MinAmount:     float32(item.MinAmount),                          // 最低使用金额
		StartTime:     time_util.TimeToStr(item.StartTime),              // 开始时间
		EndTime:       time_util.TimeToStr(item.EndTime),                // 结束时间
		TotalCount:    item.TotalCount,                                  // 发放总量
		ReceivedCount: item.ReceivedCount,                               // 已领取数量
		UsedCount:     item.UsedCount,                                   // 已使用数量
		PerLimit:      item.PerLimit,                                    // 每人限领数量
		Status:        item.Status,                                      // 状态：0-未开始，1-进行中，2-已结束，3-已取消
		IsEnabled:     item.IsEnabled,                                   // 是否启用
		Description:   item.Description,                                 // 使用说明
		CreateBy:      item.CreateBy,                                    // 创建人ID
		CreateTime:    time_util.TimeToStr(item.CreateTime),             // 创建时间
		UpdateBy:      pointerprocess.DefaltData(item.UpdateBy).(int64), // 更新人ID
		UpdateTime:    time_util.TimeToString(item.UpdateTime),          // 更新时间
	}

	return data, nil
}
