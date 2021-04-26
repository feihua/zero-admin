package logic

import (
	"context"
	"go-zero-admin/rpc/sys/sysclient"

	"go-zero-admin/api/internal/svc"
	"go-zero-admin/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type JobDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewJobDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) JobDeleteLogic {
	return JobDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *JobDeleteLogic) JobDelete(req types.DeleteJobReq) (*types.DeleteJobResp, error) {
	_, err := l.svcCtx.Sys.JobDelete(l.ctx, &sysclient.JobDeleteReq{
		Id: req.Id,
	})

	if err != nil {
		return nil, err
	}

	return &types.DeleteJobResp{
		Code:    "000000",
		Message: "",
	}, nil
}
