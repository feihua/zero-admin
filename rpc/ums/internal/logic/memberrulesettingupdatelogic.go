package logic

import (
	"context"

	"zero-admin/rpc/ums/internal/svc"
	"zero-admin/rpc/ums/umsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type MemberRuleSettingUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberRuleSettingUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberRuleSettingUpdateLogic {
	return &MemberRuleSettingUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MemberRuleSettingUpdateLogic) MemberRuleSettingUpdate(in *umsclient.MemberRuleSettingUpdateReq) (*umsclient.MemberRuleSettingUpdateResp, error) {
	// todo: add your logic here and delete this line

	return &umsclient.MemberRuleSettingUpdateResp{}, nil
}
