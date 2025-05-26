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

// UpdateOrderSettingDefaultStatusLogic 更新订单设置默认状态
/*
Author: LiuFeiHua
Date: 2025/5/26 15:53
*/
type UpdateOrderSettingDefaultStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateOrderSettingDefaultStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateOrderSettingDefaultStatusLogic {
	return &UpdateOrderSettingDefaultStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateOrderSettingDefaultStatus 更新订单设置默认状态
func (l *UpdateOrderSettingDefaultStatusLogic) UpdateOrderSettingDefaultStatus(in *omsclient.UpdateOrderSettingStatusReq) (*omsclient.UpdateOrderSettingStatusResp, error) {
	q := query.OmsOrderSetting
	if in.Status == 1 {
		if _, err := q.WithContext(l.ctx).Where(q.IsDefault.Eq(1)).Update(q.IsDefault, 0); err != nil {
			return nil, err
		}
	}
	_, err := q.WithContext(l.ctx).Where(q.ID.Eq(in.Id)).UpdateSimple(q.IsDefault.Value(in.Status), q.UpdateBy.Value(in.UpdateBy))

	if err != nil {
		logc.Errorf(l.ctx, "更新订单设置是否为默认失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("更新订单设置是否为默认失败")
	}

	return &omsclient.UpdateOrderSettingStatusResp{}, nil
}
