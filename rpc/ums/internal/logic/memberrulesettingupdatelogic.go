package logic

import (
	"context"
	"go-zero-admin/rpc/model/umsmodel"

	"go-zero-admin/rpc/ums/internal/svc"
	"go-zero-admin/rpc/ums/ums"

	"github.com/tal-tech/go-zero/core/logx"
)

type MemberRuleSettingUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberRuleSettingUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberRuleSettingUpdateLogic {
	return &MemberRuleSettingUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MemberRuleSettingUpdateLogic) MemberRuleSettingUpdate(in *ums.MemberRuleSettingUpdateReq) (*ums.MemberRuleSettingUpdateResp, error) {
	err := l.svcCtx.UmsMemberRuleSettingModel.Update(umsmodel.UmsMemberRuleSetting{
		Id:                in.Id,
		ContinueSignDay:   in.ContinueSignDay,
		ContinueSignPoint: in.ContinueSignPoint,
		ConsumePerPoint:   float64(in.ConsumePerPoint),
		LowOrderAmount:    float64(in.LowOrderAmount),
		MaxPointPerOrder:  in.MaxPointPerOrder,
		Type:              in.Type,
	})
	if err != nil {
		return nil, err
	}

	return &ums.MemberRuleSettingUpdateResp{}, nil
}
