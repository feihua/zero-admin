package returnapply

import (
	"context"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryOrderReturnApplyDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryOrderReturnApplyDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryOrderReturnApplyDetailLogic {
	return &QueryOrderReturnApplyDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryOrderReturnApplyDetailLogic) QueryOrderReturnApplyDetail(req *types.QueryOrderReturnApplyDetailReq) (resp *types.QueryOrderReturnApplyDetailResp, err error) {
	// todo: add your logic here and delete this line

	return
}
