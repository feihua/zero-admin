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
)

// AddCouponScopeLogic 添加优惠券使用范围
/*
Author: LiuFeiHua
Date: 2025/06/11 10:44:50
*/
type AddCouponScopeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddCouponScopeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddCouponScopeLogic {
	return &AddCouponScopeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// AddCouponScope 添加优惠券使用范围
func (l *AddCouponScopeLogic) AddCouponScope(in *smsclient.AddCouponScopeReq) (*smsclient.AddCouponScopeResp, error) {
	q := query.SmsCouponScope

	var data []*model.SmsCouponScope
	for _, x := range in.Data {
		data = append(data, &model.SmsCouponScope{
			CouponID:  x.CouponId,  // 优惠券ID
			ScopeType: x.ScopeType, // 范围类型：0-全场，1-分类，2-商品
			ScopeID:   x.ScopeId,   // 范围ID（分类ID或商品ID）
		})
	}

	err := q.WithContext(l.ctx).CreateInBatches(data, len(data))

	if err != nil {
		logc.Errorf(l.ctx, "添加优惠券使用范围失败,参数:%+v,异常:%s", data, err.Error())
		return nil, errors.New("添加优惠券使用范围失败")
	}

	return &smsclient.AddCouponScopeResp{}, nil
}
