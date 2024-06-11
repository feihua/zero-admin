package memberrulesettingservicelogic

import (
	"context"

	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateMemberRuleSettingStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateMemberRuleSettingStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateMemberRuleSettingStatusLogic {
	return &UpdateMemberRuleSettingStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新会员积分成长规则表状态
func (l *UpdateMemberRuleSettingStatusLogic) UpdateMemberRuleSettingStatus(in *umsclient.UpdateMemberRuleSettingStatusReq) (*umsclient.UpdateMemberRuleSettingStatusResp, error) {
	// todo: add your logic here and delete this line

	return &umsclient.UpdateMemberRuleSettingStatusResp{}, nil
}
