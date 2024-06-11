package rulesetting

import (
	"context"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateMemberRuleSettingStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateMemberRuleSettingStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateMemberRuleSettingStatusLogic {
	return &UpdateMemberRuleSettingStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateMemberRuleSettingStatusLogic) UpdateMemberRuleSettingStatus(req *types.UpdateMemberRuleSettingStatusReq) (resp *types.UpdateMemberRuleSettingStatusResp, err error) {
	// todo: add your logic here and delete this line

	return
}
