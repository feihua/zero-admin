package member_rule_setting

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/api/admin/internal/common/res"
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
func (l *UpdateMemberRuleSettingLogic) UpdateMemberRuleSetting(req *types.UpdateMemberRuleSettingReq) (resp *types.BaseResp, err error) {
	userId, err := common.GetUserId(l.ctx)
	if err != nil {
		return nil, err
	}
	_, err = l.svcCtx.MemberRuleSettingService.UpdateMemberRuleSetting(l.ctx, &umsclient.UpdateMemberRuleSettingReq{
		Id:               req.Id,               //
		ConsumePerPoint:  req.ConsumePerPoint,  // 每消费多少元获取1个点
		LowOrderAmount:   req.LowOrderAmount,   // 最低获取点数的订单金额
		MaxPointPerOrder: req.MaxPointPerOrder, // 每笔订单最高获取点数
		RuleType:         req.RuleType,         // 类型：0->积分规则；1->成长值规则
		Status:           req.Status,           // 状态：0->禁用；1->启用
		UpdateBy:         userId,
	})

	if err != nil {
		logc.Errorf(l.ctx, "更新会员积分规则信息失败,参数：%+v,响应：%s", req, err.Error())
		return nil, errorx.NewDefaultError("更新会员积分规则成功")
	}

	return res.Success()
}
