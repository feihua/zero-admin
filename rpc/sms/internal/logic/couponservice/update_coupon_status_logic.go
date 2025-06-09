package couponservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/sms/gen/query"
	"github.com/feihua/zero-admin/rpc/sms/internal/svc"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateCouponStatusLogic 更新优惠券
/*
Author: LiuFeiHua
Date: 2025/06/11 10:44:50
*/
type UpdateCouponStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateCouponStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCouponStatusLogic {
	return &UpdateCouponStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateCouponStatus 更新优惠券状态
func (l *UpdateCouponStatusLogic) UpdateCouponStatus(in *smsclient.UpdateCouponStatusReq) (*smsclient.UpdateCouponStatusResp, error) {
	q := query.SmsCoupon

	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Update(q.Status, in.Status)

	if err != nil {
		logc.Errorf(l.ctx, "更新优惠券状态失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("更新优惠券状态失败")
	}

	return &smsclient.UpdateCouponStatusResp{}, nil
}
