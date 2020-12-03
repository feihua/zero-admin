package logic

import (
	"context"

	"go-zero-admin/api/internal/svc"
	"go-zero-admin/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type MemberRuleSettingAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMemberRuleSettingAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) MemberRuleSettingAddLogic {
	return MemberRuleSettingAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MemberRuleSettingAddLogic) MemberRuleSettingAdd(req types.AddMemberRuleSettingReq) (*types.AddMemberRuleSettingResp, error) {
	// todo: add your logic here and delete this line

	return &types.AddMemberRuleSettingResp{}, nil
}
