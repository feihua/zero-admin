package memberrulesettingservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/ums/gen/model"
	"github.com/feihua/zero-admin/rpc/ums/gen/query"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"

	"github.com/feihua/zero-admin/rpc/ums/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

// MemberRuleSettingUpdateLogic 会员积分成长规则
/*
Author: LiuFeiHua
Date: 2024/5/7 10:34
*/
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

// MemberRuleSettingUpdate 更新会员积分成长规则
func (l *MemberRuleSettingUpdateLogic) MemberRuleSettingUpdate(in *umsclient.MemberRuleSettingUpdateReq) (*umsclient.MemberRuleSettingUpdateResp, error) {
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

	return &umsclient.MemberRuleSettingUpdateResp{}, nil
}
