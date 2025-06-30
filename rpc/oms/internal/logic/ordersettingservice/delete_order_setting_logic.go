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

// DeleteOrderSettingLogic 删除订单设置
/*
Author: LiuFeiHua
Date: 2025/05/26 15:21:44
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

// DeleteOrderSetting 删除订单设置
func (l *DeleteOrderSettingLogic) DeleteOrderSetting(in *omsclient.DeleteOrderSettingReq) (*omsclient.DeleteOrderSettingResp, error) {
	q := query.OmsOrderSetting

	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Delete()

	if err != nil {
		logc.Errorf(l.ctx, "删除订单设置失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("删除订单设置失败")
	}

	return &omsclient.DeleteOrderSettingResp{}, nil
}
