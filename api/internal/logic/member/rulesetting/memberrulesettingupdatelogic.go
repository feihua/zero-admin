package logic

import (
	"context"
	"encoding/json"
	"go-zero-admin/api/internal/common/errorx"
	"go-zero-admin/rpc/ums/umsclient"

	"go-zero-admin/api/internal/svc"
	"go-zero-admin/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
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
	_, err := l.svcCtx.Ums.MemberRuleSettingUpdate(l.ctx, &umsclient.MemberRuleSettingUpdateReq{
		Id:                req.Id,
		ContinueSignDay:   req.ContinueSignDay,
		ContinueSignPoint: req.ContinueSignPoint,
		ConsumePerPoint:   int64(req.ConsumePerPoint),
		LowOrderAmount:    int64(req.LowOrderAmount),
		MaxPointPerOrder:  req.MaxPointPerOrder,
		Type:              req.Type,
	})

	if err != nil {
		reqStr, _ := json.Marshal(req)
		logx.Errorf("更新会员积分规则参数:%s,异常:%s", reqStr, err.Error())
		return nil, errorx.NewDefaultError("更新会员积分规则成功")
	}

	return &types.UpdateMemberRuleSettingResp{
		Code:    "000000",
		Message: "更新会员积分规则成功",
	}, nil
}
