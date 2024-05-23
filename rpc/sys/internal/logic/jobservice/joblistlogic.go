package jobservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/sys/gen/query"
	"github.com/feihua/zero-admin/rpc/sys/internal/logic/common"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/rpc/sys/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

// JobListLogic
/*
Author: LiuFeiHua
Date: 2023/12/18 17:05
*/
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

// JobList 岗位列表
func (l *JobListLogic) JobList(in *sysclient.JobListReq) (*sysclient.JobListResp, error) {
	q := query.SysJob.WithContext(l.ctx)
	if len(in.JobName) > 0 {
		q = q.Where(query.SysJob.JobName.Like("%" + in.JobName + "%"))
	}

	if in.DelFlag != 2 {
		q = q.Where(query.SysJob.DelFlag.Eq(in.DelFlag))
	}

	offset := (in.Current - 1) * in.PageSize
	result, err := q.Offset(int(offset)).Limit(int(in.PageSize)).Find()
	count, err := q.Count()

	if err != nil {
		logc.Errorf(l.ctx, "查询岗位列表信息失败,参数：%+v,异常:%s", in, err.Error())
		return nil, err
	}

	var list []*sysclient.JobListData
	for _, job := range result {
		list = append(list, &sysclient.JobListData{
			Id:         job.ID,
			JobName:    job.JobName,
			OrderNum:   job.OrderNum,
			CreateBy:   job.CreateBy,
			CreateTime: job.CreateTime.Format("2006-01-02 15:04:05"),
			UpdateBy:   job.UpdateBy,
			UpdateTime: common.TimeToString(job.UpdateTime),
			DelFlag:    job.DelFlag,
			Remarks:    job.Remarks,
		})
	}

	logc.Infof(l.ctx, "查询岗位列表信息,参数：%+v,响应：%+v", in, list)
	return &sysclient.JobListResp{
		Total: count,
		List:  list,
	}, nil
}
