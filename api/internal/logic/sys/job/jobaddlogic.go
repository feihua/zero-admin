package logic

import (
	"context"
	"go-zero-admin/api/internal/svc"
	"go-zero-admin/api/internal/types"
	"go-zero-admin/rpc/sys/sysclient"

	"github.com/tal-tech/go-zero/core/logx"
)

type JobAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewJobAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) JobAddLogic {
	return JobAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *JobAddLogic) JobAdd(req types.AddJobReq) (*types.AddJobResp, error) {
	_, err := l.svcCtx.Sys.JobAdd(l.ctx, &sysclient.JobAddReq{
		JobName:  req.JobName,
		OrderNum: req.OrderNum,
		CreateBy: "admin",
		Remarks:  req.Remarks,
	})

	if err != nil {
		return nil, err
	}

	return &types.AddJobResp{
		Code:    "000000",
		Message: "",
	}, nil
}
