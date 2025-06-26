package logic

import (
	"context"

	"github.com/feihua/zero-admin/job/internal/svc"
	"github.com/feihua/zero-admin/job/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type JobLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewJobLogic(ctx context.Context, svcCtx *svc.ServiceContext) *JobLogic {
	return &JobLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *JobLogic) Job(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
