package integrationconsumesettingservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/ums/gen/query"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryIntegrationConsumeSettingByIdLogic
/*
Author: LiuFeiHua
Date: 2023/12/11 18:10
*/
type QueryIntegrationConsumeSettingByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryIntegrationConsumeSettingByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryIntegrationConsumeSettingByIdLogic {
	return &QueryIntegrationConsumeSettingByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryIntegrationConsumeSettingById 获取积分使用规则
func (l *QueryIntegrationConsumeSettingByIdLogic) QueryIntegrationConsumeSettingById(in *umsclient.QueryIntegrationConsumeSettingByIdReq) (*umsclient.IntegrationConsumeSettingListData, error) {
	item, err := query.UmsIntegrationConsumeSetting.WithContext(l.ctx).Where(query.UmsIntegrationConsumeSetting.ID.Eq(in.Id)).First()

	if err != nil {
		logc.Errorf(l.ctx, "查询积分消费设置列表信息失败,参数：%+v,异常:%s", in, err.Error())
		return nil, err
	}

	logc.Infof(l.ctx, "查询积分消费设置列表信息,参数：%+v,响应：%+v", in, item)
	return &umsclient.IntegrationConsumeSettingListData{
		Id:                 item.ID,
		DeductionPerAmount: item.DeductionPerAmount,
		MaxPercentPerOrder: item.MaxPercentPerOrder,
		UseUnit:            item.UseUnit,
		CouponStatus:       item.CouponStatus,
	}, nil

}
