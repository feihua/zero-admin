package logic

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MemberRuleSettingAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMemberRuleSettingAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) MemberRuleSettingAddLogic {
	return MemberRuleSettingAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MemberRuleSettingAddLogic) MemberRuleSettingAdd(req types.AddMemberRuleSettingReq) (*types.AddMemberRuleSettingResp, error) {
	_, err := l.svcCtx.MemberRuleSettingService.MemberRuleSettingAdd(l.ctx, &umsclient.MemberRuleSettingAddReq{
		ContinueSignDay:   req.ContinueSignDay,
		ContinueSignPoint: req.ContinueSignPoint,
		ConsumePerPoint:   req.ConsumePerPoint,
		LowOrderAmount:    req.LowOrderAmount,
		MaxPointPerOrder:  req.MaxPointPerOrder,
		RuleType:          req.Type,
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
