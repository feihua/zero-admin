package jobservicelogic

import (
	"context"
	"database/sql"
	"zero-admin/rpc/model/sysmodel"
	"zero-admin/rpc/sys/sysclient"

	"zero-admin/rpc/sys/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

// JobAddLogic
/*
Author: LiuFeiHua
Date: 2023/12/18 17:04
*/
type JobAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewJobAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *JobAddLogic {
	return &JobAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// JobAdd 添加岗位
func (l *JobAddLogic) JobAdd(in *sysclient.JobAddReq) (*sysclient.JobAddResp, error) {
	job := &sysmodel.SysJob{
		JobName:  in.JobName,
		OrderNum: in.OrderNum,
		CreateBy: in.CreateBy,
		Remarks:  sql.NullString{String: in.Remarks, Valid: true},
		DelFlag:  in.DelFlag,
	}
	if _, err := l.svcCtx.JobModel.Insert(l.ctx, job); err != nil {
		return nil, err
	}

	return &sysclient.JobAddResp{}, nil
}
