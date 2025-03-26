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

// UpdateOrderSettingStatusLogic 更新订单设置状态
/*
Author: LiuFeiHua
Date: 2024/6/15 11:17
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
	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Update(q.Status, in.Status)

	if err != nil {
		logc.Errorf(l.ctx, "更新订单设置状态失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("更新订单设置状态失败")
	}

	return &omsclient.UpdateOrderSettingStatusResp{}, nil
}
