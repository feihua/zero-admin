package prefrenceArea

import (
	"context"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryPreferredAreaDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryPreferredAreaDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryPreferredAreaDetailLogic {
	return &QueryPreferredAreaDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryPreferredAreaDetailLogic) QueryPreferredAreaDetail(req *types.QueryPreferredAreaDetailReq) (resp *types.QueryPreferredAreaDetailResp, err error) {
	// todo: add your logic here and delete this line

	return
}
