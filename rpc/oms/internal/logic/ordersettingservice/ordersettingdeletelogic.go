package ordersettingservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/oms/gen/query"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"

	"github.com/feihua/zero-admin/rpc/oms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

// OrderSettingDeleteLogic 订单设置
/*
Author: LiuFeiHua
Date: 2024/5/8 9:34
*/
type OrderSettingDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOrderSettingDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderSettingDeleteLogic {
	return &OrderSettingDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// OrderSettingDelete 删除订单设置
func (l *OrderSettingDeleteLogic) OrderSettingDelete(in *omsclient.OrderSettingDeleteReq) (*omsclient.OrderSettingDeleteResp, error) {
	q := query.OmsOrderSetting
	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Delete()

	if err != nil {
		return nil, err
	}

	return &omsclient.OrderSettingDeleteResp{}, nil
}
