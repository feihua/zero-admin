package integrationconsumesetting

import (
	"context"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryIntegrationConsumeSettingDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryIntegrationConsumeSettingDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryIntegrationConsumeSettingDetailLogic {
	return &QueryIntegrationConsumeSettingDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryIntegrationConsumeSettingDetailLogic) QueryIntegrationConsumeSettingDetail(req *types.QueryIntegrationConsumeSettingDetailReq) (resp *types.QueryIntegrationConsumeSettingDetailResp, err error) {
	// todo: add your logic here and delete this line

	return
}
