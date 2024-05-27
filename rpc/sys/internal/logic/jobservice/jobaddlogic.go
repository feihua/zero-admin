package jobservicelogic

import (
	"context"
	"errors"
	"fmt"
	"github.com/feihua/zero-admin/rpc/sys/gen/model"
	"github.com/feihua/zero-admin/rpc/sys/gen/query"
	"github.com/feihua/zero-admin/rpc/sys/internal/svc"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/zeromicro/go-zero/core/logx"
)

// JobAddLogic 添加岗位
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
// 1.根据岗位名称查询岗位是否已存在
// 2.如果岗位已存在,则直接返回
// 3.岗位不存在时,则直接添加岗位
func (l *JobAddLogic) JobAdd(in *sysclient.JobAddReq) (*sysclient.JobAddResp, error) {

	q := query.SysJob.WithContext(l.ctx)

	//1.根据岗位名称查询岗位是否已存在
	jobName := in.JobName
	count, err := q.Where(query.SysJob.JobName.Eq(jobName)).Count()

	if err != nil {
		logc.Errorf(l.ctx, "根据岗位名称：%s,查询岗位信息失败,异常:%s", jobName, err.Error())
		return nil, errors.New(fmt.Sprintf("查询岗位信息失败"))
	}

	//2.如果岗位已存在,则直接返回
	if count > 0 {
		logc.Errorf(l.ctx, "岗位信息已存在：%+v", in)
		return nil, errors.New(fmt.Sprintf("岗位：%s,已存在", jobName))
	}

	//3.岗位不存在时,则直接添加岗位
	job := &model.SysJob{
		JobName:  in.JobName,
		OrderNum: in.OrderNum,
		CreateBy: in.CreateBy,
		Remarks:  in.Remarks,
		DelFlag:  in.DelFlag,
	}

	err = q.Create(job)

	if err != nil {
		logc.Errorf(l.ctx, "添加岗位信息失败,参数:%+v,异常:%s", job, err.Error())
		return nil, errors.New("添加岗位信息失败")
	}

	return &sysclient.JobAddResp{}, nil
}
