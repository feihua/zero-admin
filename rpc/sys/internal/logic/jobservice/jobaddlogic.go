package jobservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/sys/gen/model"
	"github.com/feihua/zero-admin/rpc/sys/gen/query"
	"github.com/feihua/zero-admin/rpc/sys/internal/svc"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"

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
// 1.根据部门名称查询部门是否已存在
// 2.如果部门已存在,则直接返回
// 3.部门不存在时,则直接添加部门
func (l *JobAddLogic) JobAdd(in *sysclient.JobAddReq) (*sysclient.JobAddResp, error) {
	err := query.SysJob.WithContext(l.ctx).Create(&model.SysJob{
		JobName:  in.JobName,
		OrderNum: in.OrderNum,
		CreateBy: in.CreateBy,
		Remarks:  in.Remarks,
		DelFlag:  in.DelFlag,
	})
	if err != nil {
		return nil, err
	}

	return &sysclient.JobAddResp{}, nil
}
