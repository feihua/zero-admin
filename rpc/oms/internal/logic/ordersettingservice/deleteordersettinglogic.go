package ordersettingservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/oms/gen/query"

	"github.com/feihua/zero-admin/rpc/oms/internal/svc"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// DeleteOrderSettingLogic 删除订单设置表
/*
Author: LiuFeiHua
Date: 2024/6/12 9:39
*/
type DeleteOrderSettingLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteOrderSettingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteOrderSettingLogic {
	return &DeleteOrderSettingLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// DeleteOrderSetting 删除订单设置表
func (l *DeleteOrderSettingLogic) DeleteOrderSetting(in *omsclient.DeleteOrderSettingReq) (*omsclient.DeleteOrderSettingResp, error) {
	q := query.OmsOrderSetting
	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Delete()

	if err != nil {
		return nil, err
	}

	return &omsclient.DeleteOrderSettingResp{}, nil
}
