package memberrulesettingservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/ums/gen/model"
	"github.com/feihua/zero-admin/rpc/ums/gen/query"

	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateMemberRuleSettingLogic 更新会员积分成长规则表
/*
Author: LiuFeiHua
Date: 2024/6/11 14:02
*/
type UpdateMemberRuleSettingLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateMemberRuleSettingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateMemberRuleSettingLogic {
	return &UpdateMemberRuleSettingLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateMemberRuleSetting 更新会员积分成长规则表
func (l *UpdateMemberRuleSettingLogic) UpdateMemberRuleSetting(in *umsclient.UpdateMemberRuleSettingReq) (*umsclient.UpdateMemberRuleSettingResp, error) {
	_, err := query.UmsMemberRuleSetting.WithContext(l.ctx).Updates(&model.UmsMemberRuleSetting{
		ID:                in.Id,
		ContinueSignDay:   in.ContinueSignDay,
		ContinueSignPoint: in.ContinueSignPoint,
		ConsumePerPoint:   in.ConsumePerPoint,
		LowOrderAmount:    in.LowOrderAmount,
		MaxPointPerOrder:  in.MaxPointPerOrder,
		RuleType:          in.RuleType,
		UpdateBy:          in.UpdateBy,
	})

	if err != nil {
		return nil, err
	}

	return &umsclient.UpdateMemberRuleSettingResp{}, nil
}
