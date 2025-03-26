package integrationconsumesettingservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/ums/gen/query"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// DeleteIntegrationConsumeSettingLogic 删除积分消费设置
/*
Author: LiuFeiHua
Date: 2024/6/11 14:21
*/
type DeleteIntegrationConsumeSettingLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteIntegrationConsumeSettingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteIntegrationConsumeSettingLogic {
	return &DeleteIntegrationConsumeSettingLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// DeleteIntegrationConsumeSetting 删除积分消费设置
func (l *DeleteIntegrationConsumeSettingLogic) DeleteIntegrationConsumeSetting(in *umsclient.DeleteIntegrationConsumeSettingReq) (*umsclient.DeleteIntegrationConsumeSettingResp, error) {
	q := query.UmsIntegrationConsumeSetting
	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Delete()
	if err != nil {
		logc.Errorf(l.ctx, "删除积分消费设置失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("删除积分消费设置失败")
	}

	return &umsclient.DeleteIntegrationConsumeSettingResp{}, nil
}
