package rulesetting

import (
	"context"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryMemberRuleSettingDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryMemberRuleSettingDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryMemberRuleSettingDetailLogic {
	return &QueryMemberRuleSettingDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryMemberRuleSettingDetailLogic) QueryMemberRuleSettingDetail(req *types.QueryMemberRuleSettingDetailReq) (resp *types.QueryMemberRuleSettingDetailResp, err error) {
	// todo: add your logic here and delete this line

	return
}
