package returnreason

import (
	"context"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryOrderReturnReasonDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryOrderReturnReasonDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryOrderReturnReasonDetailLogic {
	return &QueryOrderReturnReasonDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryOrderReturnReasonDetailLogic) QueryOrderReturnReasonDetail(req *types.QueryOrderReturnReasonDetailReq) (resp *types.QueryOrderReturnReasonDetailResp, err error) {
	// todo: add your logic here and delete this line

	return
}
