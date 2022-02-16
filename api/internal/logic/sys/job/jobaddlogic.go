package logic

import (
	"context"
	"encoding/json"
	"zero-admin/api/internal/common/errorx"
	"zero-admin/api/internal/svc"
	"zero-admin/api/internal/types"
	"zero-admin/rpc/sys/sysclient"

	"github.com/zeromicro/go-zero/core/logx"
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
		reqStr, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("添加岗位信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, errorx.NewDefaultError("添加岗位失败")
	}

	return &types.AddJobResp{
		Code:    "000000",
		Message: "添加岗位成功",
	}, nil
}
