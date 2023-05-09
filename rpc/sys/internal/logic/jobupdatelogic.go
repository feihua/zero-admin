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

func (l *JobUpdateLogic) JobUpdate(in *sys.JobUpdateReq) (*sys.JobUpdateResp, error) {
	err := l.svcCtx.JobModel.Update(l.ctx, &sysmodel.SysJob{
		Id:         in.Id,
		JobName:    in.JobName,
		OrderNum:   in.OrderNum,
		UpdateBy:   sql.NullString{String: in.LastUpdateBy, Valid: true},
		UpdateTime: sql.NullTime{Time: time.Now()},
		Remarks:    sql.NullString{String: in.Remarks},
	})

	if err != nil {
		return nil, err
	}

	return &sys.JobUpdateResp{}, nil
}
