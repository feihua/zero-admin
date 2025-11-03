package couponservicelogic

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/feihua/zero-admin/rpc/sms/gen/model"
	"github.com/feihua/zero-admin/rpc/sms/gen/query"
	"github.com/feihua/zero-admin/rpc/sms/internal/svc"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

// UpdateCouponLogic 更新优惠券
/*
Author: LiuFeiHua
Date: 2025/06/11 10:44:50
*/
type UpdateCouponLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateCouponLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCouponLogic {
	return &UpdateCouponLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateCoupon 更新优惠券
func (l *UpdateCouponLogic) UpdateCoupon(in *smsclient.UpdateCouponReq) (*smsclient.UpdateCouponResp, error) {
	coupon := query.SmsCoupon
	q := coupon.WithContext(l.ctx)

	// 1.根据优惠券id查询优惠券是否已存在
	detail, err := q.Where(coupon.ID.Eq(in.Id)).First()

	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		logc.Errorf(l.ctx, "优惠券不存在, 请求参数：%+v, 异常信息: %s", in, err.Error())
		return nil, errors.New("优惠券不存在")
	case err != nil:
		logc.Errorf(l.ctx, "查询优惠券异常, 请求参数：%+v, 异常信息: %s", in, err.Error())
		return nil, errors.New("查询优惠券异常")
	}

	count, err := coupon.WithContext(l.ctx).Where(coupon.Name.Eq(in.Name), coupon.ID.Neq(in.Id)).Count()

	if err != nil {
		logc.Errorf(l.ctx, "更新优惠券失败,参数：%+v, 异常:%s", in, err.Error())
		return nil, errors.New(fmt.Sprintf("更新优惠券失败"))
	}

	if count > 0 {
		return nil, errors.New(fmt.Sprintf("优惠券：%s,已存在", in.Name))
	}

	startTime, _ := time.Parse("2006-01-02 15:04:05", in.StartTime)
	endTime, _ := time.Parse("2006-01-02 15:04:05", in.EndTime)
	now := time.Now()
	item := &model.SmsCoupon{
		ID:            in.Id,                 // 优惠券ID
		TypeID:        in.TypeId,             // 优惠券类型ID
		Name:          in.Name,               // 优惠券名称
		Code:          in.Code,               // 优惠券码
		Amount:        float64(in.Amount),    // 优惠金额/折扣率
		MinAmount:     float64(in.MinAmount), // 最低使用金额
		StartTime:     startTime,             // 生效时间
		EndTime:       endTime,               // 失效时间
		TotalCount:    in.TotalCount,         // 发放总量
		ReceivedCount: in.ReceivedCount,      // 已领取数量
		UsedCount:     in.UsedCount,          // 已使用数量
		PerLimit:      in.PerLimit,           // 每人限领数量
		Status:        in.Status,             // 状态：0-未开始，1-进行中，2-已结束，3-已取消
		IsEnabled:     in.IsEnabled,          // 是否启用
		Description:   in.Description,        // 使用说明
		CreateBy:      detail.CreateBy,       // 创建人ID
		CreateTime:    detail.CreateTime,     // 创建时间
		UpdateBy:      &in.UpdateBy,          // 更新人ID
		UpdateTime:    &now,                  // 更新时间
	}

	// 2.优惠券存在时,则直接更新优惠券
	err = l.svcCtx.DB.Model(&model.SmsCoupon{}).WithContext(l.ctx).Where(coupon.ID.Eq(in.Id)).Save(item).Error

	if err != nil {
		logc.Errorf(l.ctx, "更新优惠券失败,参数:%+v,异常:%s", item, err.Error())
		return nil, errors.New("更新优惠券失败")
	}

	var data []*model.SmsCouponScope
	if len(in.Scopes) == 0 {
		data = append(data, &model.SmsCouponScope{
			CouponID:  item.ID, // 优惠券ID
			ScopeType: 0,       // 范围类型：0-全场，1-分类，2-商品
			ScopeID:   0,       // 范围ID（分类ID或商品ID）
		})
	} else {
		for _, x := range in.Scopes {
			data = append(data, &model.SmsCouponScope{
				CouponID:  item.ID,     // 优惠券ID
				ScopeType: x.ScopeType, // 范围类型：0-全场，1-分类，2-商品
				ScopeID:   x.ScopeId,   // 范围ID（分类ID或商品ID）
			})
		}
	}
	_, _ = query.SmsCouponScope.WithContext(l.ctx).Where(query.SmsCouponScope.CouponID.Eq(in.Id)).Delete()

	err = query.SmsCouponScope.WithContext(l.ctx).CreateInBatches(data, len(data))

	if err != nil {
		logc.Errorf(l.ctx, "添加优惠券使用范围失败,参数:%+v,异常:%s", data, err.Error())
		return nil, errors.New("添加优惠券使用范围失败")
	}
	return &smsclient.UpdateCouponResp{}, nil
}
