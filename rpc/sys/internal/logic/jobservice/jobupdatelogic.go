package jobservicelogic

import (
	"context"
	"database/sql"
	"zero-admin/rpc/sys/internal/svc"
	"zero-admin/rpc/sys/sysclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type JobUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewJobUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *JobUpdateLogic {
	return &JobUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *JobUpdateLogic) JobUpdate(in *sysclient.JobUpdateReq) (*sysclient.JobUpdateResp, error) {
	//更新之前查询记录是否存在
	job, err := l.svcCtx.JobModel.FindOne(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}

	job.JobName = in.JobName
	job.OrderNum = in.OrderNum
	job.UpdateBy = sql.NullString{String: in.LastUpdateBy, Valid: true}
	job.Remarks = sql.NullString{String: in.Remarks, Valid: true}
	job.DelFlag = in.DelFlag

	err = l.svcCtx.JobModel.Update(l.ctx, job)

	if err != nil {
		return nil, err
	}

	return &sysclient.JobUpdateResp{}, nil
}
