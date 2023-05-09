package logic

import (
	"context"
	"database/sql"
	"time"
	"zero-admin/rpc/model/sysmodel"

	"zero-admin/rpc/sys/internal/svc"
	"zero-admin/rpc/sys/sys"

	"github.com/zeromicro/go-zero/core/logx"
)

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

func (l *JobAddLogic) JobAdd(in *sys.JobAddReq) (*sys.JobAddResp, error) {
	_, err := l.svcCtx.JobModel.Insert(l.ctx, &sysmodel.SysJob{
		JobName:    in.JobName,
		OrderNum:   in.OrderNum,
		CreateBy:   in.CreateBy,
		CreateTime: time.Now(),
		Remarks:    sql.NullString{String: in.Remarks},
		DelFlag:    0,
	})

	if err != nil {
		return nil, err
	}

	return &sys.JobAddResp{}, nil
}
