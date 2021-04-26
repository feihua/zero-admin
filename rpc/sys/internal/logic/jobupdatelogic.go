package logic

import (
	"context"
	"go-zero-admin/rpc/model/sysmodel"
	"time"

	"go-zero-admin/rpc/sys/internal/svc"
	"go-zero-admin/rpc/sys/sys"

	"github.com/tal-tech/go-zero/core/logx"
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
	err := l.svcCtx.JobModel.Update(sysmodel.SysJob{
		Id:             in.Id,
		JobName:        in.JobName,
		OrderNum:       in.OrderNum,
		LastUpdateBy:   in.LastUpdateBy,
		LastUpdateTime: time.Now(),
		Remarks:        in.Remarks,
	})

	if err != nil {
		return nil, err
	}

	return &sys.JobUpdateResp{}, nil
}
