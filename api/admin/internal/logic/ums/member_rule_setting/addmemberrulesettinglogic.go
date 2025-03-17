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

// AddMemberRuleSettingLogic 添加会员积分规则
/*
Author: LiuFeiHua
Date: 2024/5/23 9:23
*/
type AddMemberRuleSettingLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddMemberRuleSettingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddMemberRuleSettingLogic {
	return &AddMemberRuleSettingLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// AddMemberRuleSetting 添加会员积分规则
func (l *AddMemberRuleSettingLogic) AddMemberRuleSetting(req *types.AddMemberRuleSettingReq) (resp *types.AddMemberRuleSettingResp, err error) {
	_, err = l.svcCtx.MemberRuleSettingService.AddMemberRuleSetting(l.ctx, &umsclient.AddMemberRuleSettingReq{
		ContinueSignDay:   req.ContinueSignDay,   // 连续签到天数
		ContinueSignPoint: req.ContinueSignPoint, // 连续签到赠送数量
		ConsumePerPoint:   req.ConsumePerPoint,   // 每消费多少元获取1个点
		LowOrderAmount:    req.LowOrderAmount,    // 最低获取点数的订单金额
		MaxPointPerOrder:  req.MaxPointPerOrder,  // 每笔订单最高获取点数
		RuleType:          req.RuleType,          // 类型：0->积分规则；1->成长值规则
		Status:            req.Status,            // 状态：0->禁用；1->启用
		CreateBy:          l.ctx.Value("userName").(string),
	})

	if err != nil {
		logc.Errorf(l.ctx, "添加会员积分规则信息失败,参数：%+v,响应：%s", req, err.Error())
		return nil, errorx.NewDefaultError("添加会员积分规则失败")
	}

	return &types.AddMemberRuleSettingResp{
		Code:    "000000",
		Message: "添加会员积分规则成功",
	}, nil
}
