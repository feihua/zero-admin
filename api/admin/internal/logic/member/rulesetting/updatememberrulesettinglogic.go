package rulesetting

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateMemberRuleSettingLogic 更新会员积分规则
/*
Author: LiuFeiHua
Date: 2024/5/23 9:25
*/
type UpdateMemberRuleSettingLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateMemberRuleSettingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateMemberRuleSettingLogic {
	return &UpdateMemberRuleSettingLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// UpdateMemberRuleSetting 更新会员积分规则
func (l *UpdateMemberRuleSettingLogic) UpdateMemberRuleSetting(req *types.UpdateMemberRuleSettingReq) (resp *types.UpdateMemberRuleSettingResp, err error) {
	_, err = l.svcCtx.MemberRuleSettingService.MemberRuleSettingUpdate(l.ctx, &umsclient.MemberRuleSettingUpdateReq{
		Id:                req.Id,
		ContinueSignDay:   req.ContinueSignDay,
		ContinueSignPoint: req.ContinueSignPoint,
		ConsumePerPoint:   req.ConsumePerPoint,
		LowOrderAmount:    req.LowOrderAmount,
		MaxPointPerOrder:  req.MaxPointPerOrder,
		RuleType:          req.Type,
	})

	if err != nil {
		logc.Errorf(l.ctx, "更新会员积分规则信息失败,参数：%+v,响应：%s", req, err.Error())
		return nil, errorx.NewDefaultError("更新会员积分规则成功")
	}

	return &types.UpdateMemberRuleSettingResp{
		Code:    "000000",
		Message: "更新会员积分规则成功",
	}, nil
}
