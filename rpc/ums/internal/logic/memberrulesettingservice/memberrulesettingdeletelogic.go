package memberrulesettingservicelogic

import (
	"context"
	"zero-admin/rpc/ums/umsclient"

	"zero-admin/rpc/ums/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

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

func (l *MemberRuleSettingDeleteLogic) MemberRuleSettingDelete(in *umsclient.MemberRuleSettingDeleteReq) (*umsclient.MemberRuleSettingDeleteResp, error) {
	err := l.svcCtx.UmsMemberRuleSettingModel.DeleteByIds(l.ctx, in.Ids)

	if err != nil {
		return nil, err
	}

	return &umsclient.MemberRuleSettingDeleteResp{}, nil
}
