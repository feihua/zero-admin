package level

import (
	"context"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryMemberLevelDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryMemberLevelDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryMemberLevelDetailLogic {
	return &QueryMemberLevelDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryMemberLevelDetailLogic) QueryMemberLevelDetail(req *types.QueryMemberLevelDetailReq) (resp *types.QueryMemberLevelDetailResp, err error) {
	// todo: add your logic here and delete this line

	return
}
