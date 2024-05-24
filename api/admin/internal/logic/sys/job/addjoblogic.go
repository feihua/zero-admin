package job

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc/status"

	"github.com/zeromicro/go-zero/core/logx"
)

// AddJobLogic
/*
Author: LiuFeiHua
Date: 2023/12/18 17:18
*/
type AddJobLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddJobLogic(ctx context.Context, svcCtx *svc.ServiceContext) AddJobLogic {
	return AddJobLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// AddJob 添加岗位信息
func (l *AddJobLogic) AddJob(req *types.AddJobReq) (*types.AddJobResp, error) {
	_, err := l.svcCtx.JobService.JobAdd(l.ctx, &sysclient.JobAddReq{
		JobName:  req.JobName,
		OrderNum: req.OrderNum,
		CreateBy: l.ctx.Value("userName").(string),
		Remarks:  req.Remarks,
		DelFlag:  req.DelFlag,
	})

	if err != nil {
		logc.Errorf(l.ctx, "添加岗位信息失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	return &types.AddJobResp{
		Code:    "000000",
		Message: "添加岗位成功",
	}, nil
}
