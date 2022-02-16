package logic

import (
	"context"
	"encoding/json"
	"zero-admin/api/internal/common/errorx"
	"zero-admin/rpc/sys/sysclient"

	"zero-admin/api/internal/svc"
	"zero-admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
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
		data, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("参数: %s,查询岗位列表异常:%s", string(data), err.Error())
		return nil, errorx.NewDefaultError("查询岗位失败")
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
		Message:  "查询岗位成功",
		Current:  req.Current,
		Data:     list,
		PageSize: req.PageSize,
		Success:  true,
		Total:    resp.Total,
	}, nil

}
