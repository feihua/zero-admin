package logic

import (
	"context"
	"go-zero-admin/api/internal/common/errorx"
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
		logx.WithContext(l.ctx).Errorf("根据jobId: %d,删除岗位异常:%s", req.Id, err.Error())
		return nil, errorx.NewDefaultError("删除岗位失败")
	}

	return &types.DeleteJobResp{
		Code:    "000000",
		Message: "删除岗位成功",
	}, nil
}
