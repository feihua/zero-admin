package logic

import (
	"context"
	"go-zero-admin/rpc/sys/sysclient"

	"go-zero-admin/api/internal/svc"
	"go-zero-admin/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type JobListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewJobListLogic(ctx context.Context, svcCtx *svc.ServiceContext) JobListLogic {
	return JobListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *JobListLogic) JobList(req types.ListJobReq) (*types.ListJobResp, error) {
	resp, err := l.svcCtx.Sys.JobList(l.ctx, &sysclient.JobListReq{
		Current:  req.Current,
		PageSize: req.PageSize,
		JobName:  req.JobName,
	})

	if err != nil {
		return nil, err
	}

	var list []*types.ListJobData

	for _, job := range resp.List {
		list = append(list, &types.ListJobData{
			Id:             job.Id,
			JobName:        job.JobName,
			OrderNum:       job.OrderNum,
			CreateBy:       job.CreateBy,
			CreateTime:     job.CreateTime,
			LastUpdateBy:   job.LastUpdateBy,
			LastUpdateTime: job.LastUpdateTime,
			DelFlag:        job.DelFlag,
			Remarks:        job.Remarks,
		})
	}

	return &types.ListJobResp{
		Code:     "000000",
		Message:  "",
		Current:  req.Current,
		Data:     list,
		PageSize: req.PageSize,
		Success:  true,
		Total:    resp.Total,
	}, nil

}
