package logic

import (
	"context"
	"encoding/json"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// JobAddLogic
/*
Author: LiuFeiHua
Date: 2023/12/18 17:18
*/
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

// JobAdd 添加岗位信息
func (l *JobAddLogic) JobAdd(req types.AddJobReq) (*types.AddJobResp, error) {
	_, err := l.svcCtx.JobService.JobAdd(l.ctx, &sysclient.JobAddReq{
		JobName:  req.JobName,
		OrderNum: req.OrderNum,
		CreateBy: l.ctx.Value("userName").(string),
		Remarks:  req.Remarks,
		DelFlag:  req.DelFlag,
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
