package logic

import (
	"context"

	"go-zero-admin/api/internal/svc"
	"go-zero-admin/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type MemberRuleSettingListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMemberRuleSettingListLogic(ctx context.Context, svcCtx *svc.ServiceContext) MemberRuleSettingListLogic {
	return MemberRuleSettingListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MemberRuleSettingListLogic) MemberRuleSettingList(req types.ListMemberRuleSettingReq) (*types.ListMemberRuleSettingResp, error) {
	// todo: add your logic here and delete this line

	return &types.ListMemberRuleSettingResp{}, nil
}
