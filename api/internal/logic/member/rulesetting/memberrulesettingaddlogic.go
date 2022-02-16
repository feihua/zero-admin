package logic

import (
	"context"
	"encoding/json"
	"zero-admin/api/internal/common/errorx"
	"zero-admin/rpc/ums/umsclient"

	"zero-admin/api/internal/svc"
	"zero-admin/api/internal/types"

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
	_, err := l.svcCtx.Ums.MemberRuleSettingAdd(l.ctx, &umsclient.MemberRuleSettingAddReq{
		ContinueSignDay:   req.ContinueSignDay,
		ContinueSignPoint: req.ContinueSignPoint,
		ConsumePerPoint:   int64(req.ConsumePerPoint),
		LowOrderAmount:    int64(req.LowOrderAmount),
		MaxPointPerOrder:  req.MaxPointPerOrder,
		Type:              req.Type,
	})

	if err != nil {
		reqStr, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("添加会员积分规则信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, errorx.NewDefaultError("添加会员积分规则失败")
	}

	return &types.AddMemberRuleSettingResp{
		Code:    "000000",
		Message: "添加会员积分规则成功",
	}, nil
}
