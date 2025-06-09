package couponscopeservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/sms/gen/model"
	"github.com/feihua/zero-admin/rpc/sms/gen/query"
	"github.com/feihua/zero-admin/rpc/sms/internal/svc"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

// UpdateCouponScopeLogic 更新优惠券使用范围
/*
Author: LiuFeiHua
Date: 2025/06/11 10:44:50
*/
type UpdateCouponScopeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateCouponScopeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCouponScopeLogic {
	return &UpdateCouponScopeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateCouponScope 更新优惠券使用范围
func (l *UpdateCouponScopeLogic) UpdateCouponScope(in *smsclient.UpdateCouponScopeReq) (*smsclient.UpdateCouponScopeResp, error) {
	q := query.SmsCouponScope.WithContext(l.ctx)

	// 1.根据优惠券使用范围id查询优惠券使用范围是否已存在
	_, err := q.Where(query.SmsCouponScope.ID.Eq(in.Id)).First()

	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		logc.Errorf(l.ctx, "优惠券使用范围不存在, 请求参数：%+v, 异常信息: %s", in, err.Error())
		return nil, errors.New("优惠券使用范围不存在")
	case err != nil:
		logc.Errorf(l.ctx, "查询优惠券使用范围异常, 请求参数：%+v, 异常信息: %s", in, err.Error())
		return nil, errors.New("查询优惠券使用范围异常")
	}

	item := &model.SmsCouponScope{
		ID:        in.Id,        //
		CouponID:  in.CouponId,  // 优惠券ID
		ScopeType: in.ScopeType, // 范围类型：0-全场，1-分类，2-商品
		ScopeID:   in.ScopeId,   // 范围ID（分类ID或商品ID）
	}

	// 2.优惠券使用范围存在时,则直接更新优惠券使用范围
	_, err = q.Where(query.SmsCouponScope.ID.Eq(in.Id)).Updates(item)

	if err != nil {
		logc.Errorf(l.ctx, "更新优惠券使用范围失败,参数:%+v,异常:%s", item, err.Error())
		return nil, errors.New("更新优惠券使用范围失败")
	}

	return &smsclient.UpdateCouponScopeResp{}, nil
}
