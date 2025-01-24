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
		ContinueSignDay:   in.ContinueSignDay,   // 连续签到天数
		ContinueSignPoint: in.ContinueSignPoint, // 连续签到赠送数量
		ConsumePerPoint:   in.ConsumePerPoint,   // 每消费多少元获取1个点
		LowOrderAmount:    in.LowOrderAmount,    // 最低获取点数的订单金额
		MaxPointPerOrder:  in.MaxPointPerOrder,  // 每笔订单最高获取点数
		RuleType:          in.RuleType,          // 类型：0->积分规则；1->成长值规则
		Status:            in.Status,            // 状态：0->禁用；1->启用
		UpdateBy:          in.UpdateBy,          // 创建者
	})

	if err != nil {
		return nil, err
	}

	return &umsclient.UpdateMemberRuleSettingResp{}, nil
}
