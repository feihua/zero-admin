package memberrulesettingservicelogic

import (
	"context"
	"zero-admin/rpc/model/umsmodel"
	"zero-admin/rpc/ums/umsclient"

	"zero-admin/rpc/ums/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type MemberRuleSettingAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberRuleSettingAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberRuleSettingAddLogic {
	return &MemberRuleSettingAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MemberRuleSettingAddLogic) MemberRuleSettingAdd(in *umsclient.MemberRuleSettingAddReq) (*umsclient.MemberRuleSettingAddResp, error) {
	_, err := l.svcCtx.UmsMemberRuleSettingModel.Insert(l.ctx, &umsmodel.UmsMemberRuleSetting{
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

	return &umsclient.MemberRuleSettingAddResp{}, nil
}
