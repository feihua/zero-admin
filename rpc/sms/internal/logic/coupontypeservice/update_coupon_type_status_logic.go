package coupontypeservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/sms/gen/query"
	"github.com/feihua/zero-admin/rpc/sms/internal/svc"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateCouponTypeStatusLogic 更新优惠券类型
/*
Author: LiuFeiHua
Date: 2025/06/11 10:44:50
*/
type UpdateCouponTypeStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateCouponTypeStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCouponTypeStatusLogic {
	return &UpdateCouponTypeStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateCouponTypeStatus 更新优惠券类型状态
func (l *UpdateCouponTypeStatusLogic) UpdateCouponTypeStatus(in *smsclient.UpdateCouponTypeStatusReq) (*smsclient.UpdateCouponTypeStatusResp, error) {
	q := query.SmsCouponType

	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Update(q.Status, in.Status)

	if err != nil {
		logc.Errorf(l.ctx, "更新优惠券类型状态失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("更新优惠券类型状态失败")
	}

	return &smsclient.UpdateCouponTypeStatusResp{}, nil
}
