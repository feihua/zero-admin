package logic

import (
	"context"
	"encoding/json"
	"go-zero-admin/api/internal/common/errorx"
	"go-zero-admin/rpc/sys/sysclient"

	"go-zero-admin/api/internal/svc"
	"go-zero-admin/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type JobUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewJobUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) JobUpdateLogic {
	return JobUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *JobUpdateLogic) JobUpdate(req types.UpdateJobReq) (*types.UpdateJobResp, error) {
	_, err := l.svcCtx.Sys.JobUpdate(l.ctx, &sysclient.JobUpdateReq{
		Id:       req.Id,
		JobName:  req.JobName,
		OrderNum: req.OrderNum,
		Remarks:  req.Remarks,
		//todo 从token里面拿
		LastUpdateBy: "admin",
	})

	if err != nil {
		reqStr, _ := json.Marshal(req)
		logx.Errorf("更新岗位信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, errorx.NewDefaultError("删除岗位失败")
	}

	return &types.UpdateJobResp{
		Code:    "000000",
		Message: "删除岗位信息成功",
	}, nil
}
