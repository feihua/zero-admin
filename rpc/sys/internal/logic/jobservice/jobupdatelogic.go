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

// JobUpdateLogic 更新岗位信息
/*
Author: LiuFeiHua
Date: 2023/12/18 17:06
*/
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

// JobUpdate 更新岗位信息
// 1.根据岗位id查询岗位是否已存在
// 3.岗位存在时,则直接更新岗位
func (l *JobUpdateLogic) JobUpdate(in *sysclient.JobUpdateReq) (*sysclient.JobUpdateResp, error) {
	q := query.SysJob.WithContext(l.ctx)

	//1.根据岗位id查询岗位是否已存在
	_, err := q.Where(query.SysJob.ID.Eq(in.Id)).First()

	if err != nil {
		logc.Errorf(l.ctx, "根据岗位id：%d,查询岗位信息失败,异常:%s", in.Id, err.Error())
		return nil, errors.New(fmt.Sprintf("查询岗位信息失败"))
	}

	// 3.岗位存在时,则直接更新岗位
	var job = &model.SysJob{
		ID:       in.Id,
		JobName:  in.JobName,
		OrderNum: in.OrderNum,
		UpdateBy: in.UpdateBy,
		Remarks:  in.Remarks,
		DelFlag:  in.DelFlag,
	}
	_, err = q.Updates(job)

	if err != nil {
		logc.Errorf(l.ctx, "更新岗位信息失败,参数:%+v,异常:%s", job, err.Error())
		return nil, errors.New("更新岗位信息失败")
	}

	return &sysclient.JobUpdateResp{}, nil
}
