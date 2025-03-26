package ordersettingservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/oms/gen/query"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/rpc/oms/internal/svc"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateOrderSettingIsDefaultLogic 更新订单设置是否为默认
/*
Author: LiuFeiHua
Date: 2024/6/15 11:20
*/
type UpdateOrderSettingIsDefaultLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateOrderSettingIsDefaultLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateOrderSettingIsDefaultLogic {
	return &UpdateOrderSettingIsDefaultLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateOrderSettingIsDefault 更新订单设置是否为默认
func (l *UpdateOrderSettingIsDefaultLogic) UpdateOrderSettingIsDefault(in *omsclient.UpdateOrderSettingIsDefaultReq) (*omsclient.UpdateOrderSettingStatusResp, error) {
	q := query.OmsOrderSetting
	if in.IsDefault == 1 {
		if _, err := q.WithContext(l.ctx).Where(q.IsDefault.Eq(1)).Update(q.IsDefault, 0); err != nil {
			return nil, err
		}
	}
	_, err := q.WithContext(l.ctx).Where(q.ID.Eq(in.Id)).Update(q.IsDefault, in.IsDefault)

	if err != nil {
		logc.Errorf(l.ctx, "更新订单设置是否为默认失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("更新订单设置是否为默认失败")
	}

	return &omsclient.UpdateOrderSettingStatusResp{}, nil
}
