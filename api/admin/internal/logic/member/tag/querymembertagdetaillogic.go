package tag

import (
	"context"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryMemberTagDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryMemberTagDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryMemberTagDetailLogic {
	return &QueryMemberTagDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryMemberTagDetailLogic) QueryMemberTagDetail(req *types.QueryMemberTagDetailReq) (resp *types.QueryMemberTagDetailResp, err error) {
	// todo: add your logic here and delete this line

	return
}
