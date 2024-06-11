package memberrulesettingservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/ums/gen/query"

	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// DeleteMemberRuleSettingLogic 删除会员积分成长规则表
/*
Author: LiuFeiHua
Date: 2024/6/11 14:01
*/
type DeleteMemberRuleSettingLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteMemberRuleSettingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteMemberRuleSettingLogic {
	return &DeleteMemberRuleSettingLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// DeleteMemberRuleSetting 删除会员积分成长规则表
func (l *DeleteMemberRuleSettingLogic) DeleteMemberRuleSetting(in *umsclient.DeleteMemberRuleSettingReq) (*umsclient.DeleteMemberRuleSettingResp, error) {
	q := query.UmsMemberRuleSetting
	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Delete()
	if err != nil {
		return nil, err
	}

	return &umsclient.DeleteMemberRuleSettingResp{}, nil
}
