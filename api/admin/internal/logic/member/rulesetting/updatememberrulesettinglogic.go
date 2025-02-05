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
	_, err = l.svcCtx.MemberRuleSettingService.UpdateMemberRuleSetting(l.ctx, &umsclient.UpdateMemberRuleSettingReq{
		Id:                req.Id,                //
		ContinueSignDay:   req.ContinueSignDay,   // 连续签到天数
		ContinueSignPoint: req.ContinueSignPoint, // 连续签到赠送数量
		ConsumePerPoint:   req.ConsumePerPoint,   // 每消费多少元获取1个点
		LowOrderAmount:    req.LowOrderAmount,    // 最低获取点数的订单金额
		MaxPointPerOrder:  req.MaxPointPerOrder,  // 每笔订单最高获取点数
		RuleType:          req.RuleType,          // 类型：0->积分规则；1->成长值规则
		Status:            req.Status,            // 状态：0->禁用；1->启用
		UpdateBy:          l.ctx.Value("userName").(string),
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
