package integrationconsumesettingservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/ums/gen/query"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"

	"github.com/feihua/zero-admin/rpc/ums/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

// IntegrationConsumeSettingDeleteLogic 积分消费设置
/*
Author: LiuFeiHua
Date: 2024/5/7 9:37
*/
type IntegrationConsumeSettingDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewIntegrationConsumeSettingDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IntegrationConsumeSettingDeleteLogic {
	return &IntegrationConsumeSettingDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// IntegrationConsumeSettingDelete 删除积分消费设置
func (l *IntegrationConsumeSettingDeleteLogic) IntegrationConsumeSettingDelete(in *umsclient.IntegrationConsumeSettingDeleteReq) (*umsclient.IntegrationConsumeSettingDeleteResp, error) {
	q := query.UmsIntegrationConsumeSetting
	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Delete()
	if err != nil {
		return nil, err
	}

	return &umsclient.IntegrationConsumeSettingDeleteResp{}, nil
}
