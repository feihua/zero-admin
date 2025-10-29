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

// DeleteMemberRuleSettingLogic 删除会员积分成长规则
/*
Author: LiuFeiHua
Date: 2025/05/23 11:19:51
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

// DeleteMemberRuleSetting 删除会员积分成长规则
func (l *DeleteMemberRuleSettingLogic) DeleteMemberRuleSetting(in *umsclient.DeleteMemberRuleSettingReq) (*umsclient.DeleteMemberRuleSettingResp, error) {
	q := query.UmsMemberRuleSetting

	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Delete()

	if err != nil {
		logc.Errorf(l.ctx, "删除会员积分成长规则失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("删除会员积分成长规则失败")
	}

	return &umsclient.DeleteMemberRuleSettingResp{}, nil
}
