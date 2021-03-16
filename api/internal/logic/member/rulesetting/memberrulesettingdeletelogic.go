package logic

import (
	"context"
	"go-zero-admin/rpc/ums/umsclient"

	"go-zero-admin/api/internal/svc"
	"go-zero-admin/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type MemberRuleSettingDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMemberRuleSettingDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) MemberRuleSettingDeleteLogic {
	return MemberRuleSettingDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MemberRuleSettingDeleteLogic) MemberRuleSettingDelete(req types.DeleteMemberRuleSettingReq) (*types.DeleteMemberRuleSettingResp, error) {
	_, _ = l.svcCtx.Ums.MemberRuleSettingDelete(l.ctx, &umsclient.MemberRuleSettingDeleteReq{
		Id: req.Id,
	})

	return &types.DeleteMemberRuleSettingResp{
		Code:    "000000",
		Message: "",
	}, nil
}
