package coupontypeservicelogic

import (
	"context"
	"errors"
	"fmt"
	"github.com/feihua/zero-admin/rpc/sms/gen/model"
	"github.com/feihua/zero-admin/rpc/sms/gen/query"
	"github.com/feihua/zero-admin/rpc/sms/internal/svc"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
	"time"
)

// UpdateCouponTypeLogic 更新优惠券类型
/*
Author: LiuFeiHua
Date: 2025/06/11 10:44:50
*/
type UpdateCouponTypeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateCouponTypeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCouponTypeLogic {
	return &UpdateCouponTypeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateCouponType 更新优惠券类型
func (l *UpdateCouponTypeLogic) UpdateCouponType(in *smsclient.UpdateCouponTypeReq) (*smsclient.UpdateCouponTypeResp, error) {
	couponType := query.SmsCouponType
	q := couponType.WithContext(l.ctx)

	// 1.根据优惠券类型id查询优惠券类型是否已存在
	detail, err := q.Where(couponType.ID.Eq(in.Id)).First()

	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		logc.Errorf(l.ctx, "优惠券类型不存在, 请求参数：%+v, 异常信息: %s", in, err.Error())
		return nil, errors.New("优惠券类型不存在")
	case err != nil:
		logc.Errorf(l.ctx, "查询优惠券类型异常, 请求参数：%+v, 异常信息: %s", in, err.Error())
		return nil, errors.New("查询优惠券类型异常")
	}

	count, err := couponType.WithContext(l.ctx).Where(couponType.Name.Eq(in.Name), couponType.ID.Neq(in.Id)).Count()

	if err != nil {
		logc.Errorf(l.ctx, "更新优惠券类型失败,参数：%+v, 异常:%s", in, err.Error())
		return nil, errors.New(fmt.Sprintf("更新优惠券类型失败"))
	}

	if count > 0 {
		return nil, errors.New(fmt.Sprintf("优惠券类型：%s,已存在", in.Name))
	}

	now := time.Now()
	item := &model.SmsCouponType{
		ID:           in.Id,             // 主键ID
		Name:         in.Name,           // 类型名称
		Code:         in.Code,           // 类型编码
		Description:  in.Description,    // 描述
		DiscountType: in.DiscountType,   // 优惠方式：1-固定金额，2-折扣率，3-第N件特惠，4-买赠，5-特价，6-套装优惠，7-搭配优惠，8-积分抵现，9-积分倍率，10-免运费，11-运费减免，12-限时特权，13-会员特权
		Status:       in.Status,         // 是否启用
		Sort:         in.Sort,           // 排序
		CreateBy:     detail.CreateBy,   // 创建人ID
		CreateTime:   detail.CreateTime, // 创建时间
		UpdateBy:     &in.UpdateBy,      // 更新人ID
		UpdateTime:   &now,              // 更新时间

	}

	// 2.优惠券类型存在时,则直接更新优惠券类型
	err = l.svcCtx.DB.Model(&model.SmsCouponType{}).WithContext(l.ctx).Where(couponType.ID.Eq(in.Id)).Save(item).Error

	if err != nil {
		logc.Errorf(l.ctx, "更新优惠券类型失败,参数:%+v,异常:%s", item, err.Error())
		return nil, errors.New("更新优惠券类型失败")
	}

	return &smsclient.UpdateCouponTypeResp{}, nil
}
