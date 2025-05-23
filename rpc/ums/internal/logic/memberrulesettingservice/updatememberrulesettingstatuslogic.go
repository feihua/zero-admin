package memberrulesettingservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/ums/gen/query"
	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateMemberRuleSettingStatusLogic 更新会员积分成长规则
/*
Author: LiuFeiHua
Date: 2025/05/23 11:19:51
*/
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

// UpdateMemberRuleSettingStatus 更新会员积分成长规则状态
func (l *UpdateMemberRuleSettingStatusLogic) UpdateMemberRuleSettingStatus(in *umsclient.UpdateMemberRuleSettingStatusReq) (*umsclient.UpdateMemberRuleSettingStatusResp, error) {
	q := query.UmsMemberRuleSetting

	_, err := q.WithContext(l.ctx).Where(q.ID.Eq(in.Id)).Update(q.Status, in.Status)

	if err != nil {
		logc.Errorf(l.ctx, "更新会员积分成长规则状态失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("更新会员积分成长规则状态失败")
	}

	return &umsclient.UpdateMemberRuleSettingStatusResp{}, nil
}
