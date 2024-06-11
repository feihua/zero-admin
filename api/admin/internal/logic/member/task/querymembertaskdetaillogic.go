package task

import (
	"context"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryMemberTaskDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryMemberTaskDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryMemberTaskDetailLogic {
	return &QueryMemberTaskDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryMemberTaskDetailLogic) QueryMemberTaskDetail(req *types.QueryMemberTaskDetailReq) (resp *types.QueryMemberTaskDetailResp, err error) {
	// todo: add your logic here and delete this line

	return
}
