package job

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc/status"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryJobListLogic
/*
Author: LiuFeiHua
Date: 2023/12/18 17:19
*/
type QueryJobListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryJobListLogic(ctx context.Context, svcCtx *svc.ServiceContext) QueryJobListLogic {
	return QueryJobListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// QueryJobList 岗位信息列表
func (l *QueryJobListLogic) QueryJobList(req *types.ListJobReq) (*types.ListJobResp, error) {
	resp, err := l.svcCtx.JobService.JobList(l.ctx, &sysclient.JobListReq{
		Current:  req.Current,
		PageSize: req.PageSize,
		JobName:  req.JobName,
		DelFlag:  req.DelFlag,
	})

	if err != nil {
		logc.Errorf(l.ctx, "参数: %+v,查询岗位列表异常:%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	var list []*types.ListJobData

	for _, job := range resp.List {
		list = append(list, &types.ListJobData{
			Id:         job.Id,
			JobName:    job.JobName,
			OrderNum:   job.OrderNum,
			CreateBy:   job.CreateBy,
			CreateTime: job.CreateTime,
			UpdateBy:   job.UpdateBy,
			UpdateTime: job.UpdateTime,
			DelFlag:    job.DelFlag,
			Remarks:    job.Remarks,
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
