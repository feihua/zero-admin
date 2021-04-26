package logic

import (
	"context"
	"fmt"

	"go-zero-admin/rpc/sys/internal/svc"
	"go-zero-admin/rpc/sys/sys"

	"github.com/tal-tech/go-zero/core/logx"
)

type JobListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewJobListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *JobListLogic {
	return &JobListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *JobListLogic) JobList(in *sys.JobListReq) (*sys.JobListResp, error) {
	all, err := l.svcCtx.JobModel.FindAll(in.Current, in.PageSize)
	count, _ := l.svcCtx.JobModel.Count()

	if err != nil {
		return nil, err
	}
	var list []*sys.JobListData
	for _, job := range *all {
		fmt.Println(job)
		list = append(list, &sys.JobListData{
			Id:             job.Id,
			JobName:        job.JobName,
			OrderNum:       job.OrderNum,
			CreateBy:       job.CreateBy,
			CreateTime:     job.CreateTime.Format("2006-01-02 15:04:05"),
			LastUpdateBy:   job.LastUpdateBy,
			LastUpdateTime: job.LastUpdateTime.Format("2006-01-02 15:04:05"),
			DelFlag:        job.DelFlag,
			Remarks:        job.Remarks,
		})
	}

	return &sys.JobListResp{
		Total: count,
		List:  list,
	}, nil
}
