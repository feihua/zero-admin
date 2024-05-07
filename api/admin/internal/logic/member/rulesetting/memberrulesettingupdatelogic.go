package logic

import (
	"context"
	"encoding/json"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MemberRuleSettingUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMemberRuleSettingUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) MemberRuleSettingUpdateLogic {
	return MemberRuleSettingUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MemberRuleSettingUpdateLogic) MemberRuleSettingUpdate(req types.UpdateMemberRuleSettingReq) (*types.UpdateMemberRuleSettingResp, error) {
	_, err := l.svcCtx.MemberRuleSettingService.MemberRuleSettingUpdate(l.ctx, &umsclient.MemberRuleSettingUpdateReq{
		Id:                req.Id,
		ContinueSignDay:   req.ContinueSignDay,
		ContinueSignPoint: req.ContinueSignPoint,
		ConsumePerPoint:   req.ConsumePerPoint,
		LowOrderAmount:    req.LowOrderAmount,
		MaxPointPerOrder:  req.MaxPointPerOrder,
		Type:              req.Type,
	})

	if err != nil {
		reqStr, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("更新会员积分规则信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, errorx.NewDefaultError("更新会员积分规则成功")
	}

	return &types.UpdateMemberRuleSettingResp{
		Code:    "000000",
		Message: "更新会员积分规则成功",
	}, nil
}
