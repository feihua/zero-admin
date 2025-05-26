package ordersettingservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/oms/gen/query"
	"github.com/feihua/zero-admin/rpc/oms/internal/svc"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateOrderSettingStatusLogic 更新订单设置
/*
Author: LiuFeiHua
Date: 2025/05/26 15:21:44
*/
type UpdateOrderSettingStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateOrderSettingStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateOrderSettingStatusLogic {
	return &UpdateOrderSettingStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateOrderSettingStatus 更新订单设置状态
func (l *UpdateOrderSettingStatusLogic) UpdateOrderSettingStatus(in *omsclient.UpdateOrderSettingStatusReq) (*omsclient.UpdateOrderSettingStatusResp, error) {
	q := query.OmsOrderSetting

	_, err := q.WithContext(l.ctx).Where(q.ID.Eq(in.Id)).UpdateSimple(q.Status.Value(in.Status), q.UpdateBy.Value(in.UpdateBy))

	if err != nil {
		logc.Errorf(l.ctx, "更新订单设置状态失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("更新订单设置状态失败")
	}

	return &omsclient.UpdateOrderSettingStatusResp{}, nil
}
