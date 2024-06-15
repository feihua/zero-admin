package memberrulesettingservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/ums/gen/query"

	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateMemberRuleSettingStatusLogic 更新会员积分成长规则表状态
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

// UpdateMemberRuleSettingStatus 更新会员积分成长规则表状态
func (l *UpdateMemberRuleSettingStatusLogic) UpdateMemberRuleSettingStatus(in *umsclient.UpdateMemberRuleSettingStatusReq) (*umsclient.UpdateMemberRuleSettingStatusResp, error) {
	q := query.UmsMemberRuleSetting
	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Update(q.Status, in.Status)
	if err != nil {
		return nil, err
	}

	return &umsclient.UpdateMemberRuleSettingStatusResp{}, nil
}
