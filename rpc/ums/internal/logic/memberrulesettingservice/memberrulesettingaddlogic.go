package memberrulesettingservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/ums/gen/model"
	"github.com/feihua/zero-admin/rpc/ums/gen/query"
	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"
	"github.com/zeromicro/go-zero/core/logx"
)

// MemberRuleSettingAddLogic 会员积分成长规则
/*
Author: LiuFeiHua
Date: 2024/5/7 10:40
*/
type MemberRuleSettingAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberRuleSettingAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberRuleSettingAddLogic {
	return &MemberRuleSettingAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// MemberRuleSettingAdd 添加会员积分成长规则
func (l *MemberRuleSettingAddLogic) MemberRuleSettingAdd(in *umsclient.MemberRuleSettingAddReq) (*umsclient.MemberRuleSettingAddResp, error) {
	err := query.UmsMemberRuleSetting.WithContext(l.ctx).Create(&model.UmsMemberRuleSetting{
		ContinueSignDay:   in.ContinueSignDay,
		ContinueSignPoint: in.ContinueSignPoint,
		ConsumePerPoint:   in.ConsumePerPoint,
		LowOrderAmount:    in.LowOrderAmount,
		MaxPointPerOrder:  in.MaxPointPerOrder,
		RuleType:          in.RuleType,
		CreateBy:          in.CreateBy,
	})

	if err != nil {
		return nil, err
	}

	return &umsclient.MemberRuleSettingAddResp{}, nil
}
