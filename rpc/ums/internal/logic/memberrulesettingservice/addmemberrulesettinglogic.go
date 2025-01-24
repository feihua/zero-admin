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
		ContinueSignDay:   in.ContinueSignDay,   // 连续签到天数
		ContinueSignPoint: in.ContinueSignPoint, // 连续签到赠送数量
		ConsumePerPoint:   in.ConsumePerPoint,   // 每消费多少元获取1个点
		LowOrderAmount:    in.LowOrderAmount,    // 最低获取点数的订单金额
		MaxPointPerOrder:  in.MaxPointPerOrder,  // 每笔订单最高获取点数
		RuleType:          in.RuleType,          // 类型：0->积分规则；1->成长值规则
		Status:            in.Status,            // 状态：0->禁用；1->启用
		CreateBy:          in.CreateBy,          // 创建者
	})

	if err != nil {
		return nil, err
	}

	return &umsclient.AddMemberRuleSettingResp{}, nil
}
