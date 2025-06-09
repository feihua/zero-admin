package couponscopeservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/sms/gen/query"
	"github.com/feihua/zero-admin/rpc/sms/internal/svc"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
)

// DeleteCouponScopeLogic 删除优惠券使用范围
/*
Author: LiuFeiHua
Date: 2025/06/11 10:44:50
*/
type DeleteCouponScopeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteCouponScopeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteCouponScopeLogic {
	return &DeleteCouponScopeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// DeleteCouponScope 删除优惠券使用范围
func (l *DeleteCouponScopeLogic) DeleteCouponScope(in *smsclient.DeleteCouponScopeReq) (*smsclient.DeleteCouponScopeResp, error) {
	q := query.SmsCouponScope

	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Delete()

	if err != nil {
		logc.Errorf(l.ctx, "删除优惠券使用范围失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("删除优惠券使用范围失败")
	}

	return &smsclient.DeleteCouponScopeResp{}, nil
}
