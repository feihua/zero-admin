package memberrulesettingservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/ums/gen/query"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"

	"github.com/feihua/zero-admin/rpc/ums/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

// MemberRuleSettingDeleteLogic 会员积分成长规则
/*
Author: LiuFeiHua
Date: 2024/5/7 9:23
*/
type MemberRuleSettingDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberRuleSettingDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberRuleSettingDeleteLogic {
	return &MemberRuleSettingDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// MemberRuleSettingDelete 删除会员积分成长规则
func (l *MemberRuleSettingDeleteLogic) MemberRuleSettingDelete(in *umsclient.MemberRuleSettingDeleteReq) (*umsclient.MemberRuleSettingDeleteResp, error) {
	q := query.UmsMemberRuleSetting
	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Delete()
	if err != nil {
		return nil, err
	}

	return &umsclient.MemberRuleSettingDeleteResp{}, nil
}
