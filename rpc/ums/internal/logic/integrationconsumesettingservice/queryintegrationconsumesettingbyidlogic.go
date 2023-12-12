package integrationconsumesettingservicelogic

import (
	"context"
	"encoding/json"

	"zero-admin/rpc/ums/internal/svc"
	"zero-admin/rpc/ums/umsclient"

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
	item, err := l.svcCtx.UmsIntegrationConsumeSettingModel.FindOne(l.ctx, in.Id)

	if err != nil {
		reqStr, _ := json.Marshal(in)
		logx.WithContext(l.ctx).Errorf("查询积分消费设置列表信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, err
	}

	reqStr, _ := json.Marshal(in)
	listStr, _ := json.Marshal(item)
	logx.WithContext(l.ctx).Infof("查询积分消费设置列表信息,参数：%s,响应：%s", reqStr, listStr)
	return &umsclient.IntegrationConsumeSettingListData{
		Id:                 item.Id,
		DeductionPerAmount: item.DeductionPerAmount,
		MaxPercentPerOrder: item.MaxPercentPerOrder,
		UseUnit:            item.UseUnit,
		CouponStatus:       item.CouponStatus,
	}, nil

}
