package memberrulesettingservicelogic

import (
	"context"

	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryMemberRuleSettingDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryMemberRuleSettingDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryMemberRuleSettingDetailLogic {
	return &QueryMemberRuleSettingDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询会员积分成长规则表详情
func (l *QueryMemberRuleSettingDetailLogic) QueryMemberRuleSettingDetail(in *umsclient.QueryMemberRuleSettingDetailReq) (*umsclient.QueryMemberRuleSettingDetailResp, error) {
	// todo: add your logic here and delete this line

	return &umsclient.QueryMemberRuleSettingDetailResp{}, nil
}
