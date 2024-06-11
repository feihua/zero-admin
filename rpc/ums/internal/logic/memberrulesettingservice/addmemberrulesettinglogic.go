package memberrulesettingservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/ums/gen/model"
	"github.com/feihua/zero-admin/rpc/ums/gen/query"

	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// AddMemberRuleSettingLogic 添加会员积分成长规则表
/*
Author: LiuFeiHua
Date: 2024/6/11 14:00
*/
type AddMemberRuleSettingLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddMemberRuleSettingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddMemberRuleSettingLogic {
	return &AddMemberRuleSettingLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// AddMemberRuleSetting 添加会员积分成长规则表
func (l *AddMemberRuleSettingLogic) AddMemberRuleSetting(in *umsclient.AddMemberRuleSettingReq) (*umsclient.AddMemberRuleSettingResp, error) {
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

	return &umsclient.AddMemberRuleSettingResp{}, nil
}
