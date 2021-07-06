package logic

import (
	"context"
	"go-zero-admin/api/internal/common/errorx"
	"go-zero-admin/rpc/ums/umsclient"

	"go-zero-admin/api/internal/svc"
	"go-zero-admin/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
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
	_, err := l.svcCtx.Ums.MemberRuleSettingAdd(l.ctx, &umsclient.MemberRuleSettingAddReq{
		ContinueSignDay:   req.ContinueSignDay,
		ContinueSignPoint: req.ContinueSignPoint,
		ConsumePerPoint:   int64(req.ConsumePerPoint),
		LowOrderAmount:    int64(req.LowOrderAmount),
		MaxPointPerOrder:  req.MaxPointPerOrder,
		Type:              req.Type,
	})

	if err != nil {
		return nil, errorx.NewDefaultError("添加会员积分规则失败")
	}

	return &types.AddMemberRuleSettingResp{
		Code:    "000000",
		Message: "添加会员积分规则成功",
	}, nil
}
