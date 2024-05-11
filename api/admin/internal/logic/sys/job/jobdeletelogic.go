package logic

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// JobDeleteLogic
/*
Author: LiuFeiHua
Date: 2023/12/18 17:19
*/
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

// JobDelete 删除岗位信息
func (l *JobDeleteLogic) JobDelete(req types.DeleteJobReq) (*types.DeleteJobResp, error) {
	_, err := l.svcCtx.JobService.JobDelete(l.ctx, &sysclient.JobDeleteReq{
		Ids: req.Ids,
	})

	if err != nil {
		logc.Errorf(l.ctx, "根据jobId: %d,删除岗位异常:%s", req, err.Error())
		return nil, errorx.NewDefaultError("删除岗位失败")
	}

	return &types.DeleteJobResp{
		Code:    "000000",
		Message: "删除岗位成功",
	}, nil
}
