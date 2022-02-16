package logic

import (
	"context"
	"encoding/json"
	"fmt"

	"zero-admin/rpc/sys/internal/svc"
	"zero-admin/rpc/sys/sys"

	"github.com/zeromicro/go-zero/core/logx"
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
		reqStr, _ := json.Marshal(in)
		logx.WithContext(l.ctx).Errorf("查询岗位列表信息失败,参数:%s,异常:%s", reqStr, err.Error())
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

	reqStr, _ := json.Marshal(in)
	listStr, _ := json.Marshal(list)
	logx.WithContext(l.ctx).Infof("查询岗位列表信息,参数：%s,响应：%s", reqStr, listStr)
	return &sys.JobListResp{
		Total: count,
		List:  list,
	}, nil
}
