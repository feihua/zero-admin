package setting

import (
	"context"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryOrderSettingDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryOrderSettingDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryOrderSettingDetailLogic {
	return &QueryOrderSettingDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryOrderSettingDetailLogic) QueryOrderSettingDetail(req *types.QueryOrderSettingDetailReq) (resp *types.QueryOrderSettingDetailResp, err error) {
	// todo: add your logic here and delete this line

	return
}
