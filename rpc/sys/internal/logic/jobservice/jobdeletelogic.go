package jobservicelogic

import (
	"context"
	"zero-admin/rpc/sys/sysclient"

	"zero-admin/rpc/sys/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type JobDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewJobDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *JobDeleteLogic {
	return &JobDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *JobDeleteLogic) JobDelete(in *sysclient.JobDeleteReq) (*sysclient.JobDeleteResp, error) {
	err := l.svcCtx.JobModel.DeleteByIds(l.ctx, in.Ids)

	if err != nil {
		return nil, err
	}

	return &sysclient.JobDeleteResp{}, nil
}
