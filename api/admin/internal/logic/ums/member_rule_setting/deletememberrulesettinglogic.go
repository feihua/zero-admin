package member_rule_setting

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// DeleteMemberRuleSettingLogic 删除积分规则
/*
Author: LiuFeiHua
Date: 2024/5/23 9:24
*/
type DeleteMemberRuleSettingLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteMemberRuleSettingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteMemberRuleSettingLogic {
	return &DeleteMemberRuleSettingLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// DeleteMemberRuleSetting 删除积分规则
func (l *DeleteMemberRuleSettingLogic) DeleteMemberRuleSetting(req *types.DeleteMemberRuleSettingReq) (resp *types.DeleteMemberRuleSettingResp, err error) {
	_, err = l.svcCtx.MemberRuleSettingService.DeleteMemberRuleSetting(l.ctx, &umsclient.DeleteMemberRuleSettingReq{
		Ids: req.Ids,
	})

	if err != nil {
		logc.Errorf(l.ctx, "根据Id: %+v,删除积分规则异常:%s", req, err.Error())
		return nil, errorx.NewDefaultError("删除积分规则失败")
	}
	return &types.DeleteMemberRuleSettingResp{
		Code:    "000000",
		Message: "删除积分规则成功",
	}, nil
}
