package rule_setting

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateMemberRuleSettingStatusLogic 更新会员积分规则状态
type UpdateMemberRuleSettingStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateMemberRuleSettingStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateMemberRuleSettingStatusLogic {
	return &UpdateMemberRuleSettingStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// UpdateMemberRuleSettingStatus 更新会员积分规则状态
func (l *UpdateMemberRuleSettingStatusLogic) UpdateMemberRuleSettingStatus(req *types.UpdateMemberRuleSettingStatusReq) (resp *types.UpdateMemberRuleSettingStatusResp, err error) {
	_, err = l.svcCtx.MemberRuleSettingService.UpdateMemberRuleSettingStatus(l.ctx, &umsclient.UpdateMemberRuleSettingStatusReq{
		Ids:      req.Ids,
		Status:   req.Status,
		UpdateBy: l.ctx.Value("userName").(string),
	})

	if err != nil {
		logc.Errorf(l.ctx, "更新会员积分规则状态失败,参数：%+v,响应：%s", req, err.Error())
		return nil, errorx.NewDefaultError("更新会员积分规则状态失败")
	}

	return &types.UpdateMemberRuleSettingStatusResp{
		Code:    "000000",
		Message: "更新会员积分规则状态成功",
	}, nil
}
