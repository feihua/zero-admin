package integrationconsumesettingservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/ums/gen/query"

	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateIntegrationConsumeSettingStatusLogic 更新积分消费设置状态
type UpdateIntegrationConsumeSettingStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateIntegrationConsumeSettingStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateIntegrationConsumeSettingStatusLogic {
	return &UpdateIntegrationConsumeSettingStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateIntegrationConsumeSettingStatus 更新积分消费设置状态
func (l *UpdateIntegrationConsumeSettingStatusLogic) UpdateIntegrationConsumeSettingStatus(in *umsclient.UpdateIntegrationConsumeSettingStatusReq) (*umsclient.UpdateIntegrationConsumeSettingStatusResp, error) {
	err := query.Q.Transaction(func(tx *query.Query) error {
		q := tx.UmsIntegrationConsumeSetting
		if in.IsDefault == 1 {
			_, err := q.WithContext(l.ctx).Where(q.IsDefault.Eq(1)).Update(q.IsDefault, 0)
			if err != nil {
				return err
			}
		}

		_, err := q.WithContext(l.ctx).Where(q.ID.Eq(in.Id)).Update(q.IsDefault, in.IsDefault)
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return &umsclient.UpdateIntegrationConsumeSettingStatusResp{}, nil
}
