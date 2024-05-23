package job

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateJobLogic
/*
Author: LiuFeiHua
Date: 2023/12/18 17:19
*/
type UpdateJobLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateJobLogic(ctx context.Context, svcCtx *svc.ServiceContext) UpdateJobLogic {
	return UpdateJobLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// UpdateJob 更新岗位信息
func (l *UpdateJobLogic) UpdateJob(req *types.UpdateJobReq) (*types.UpdateJobResp, error) {
	_, err := l.svcCtx.JobService.JobUpdate(l.ctx, &sysclient.JobUpdateReq{
		Id:       req.Id,
		JobName:  req.JobName,
		OrderNum: req.OrderNum,
		Remarks:  req.Remarks,
		UpdateBy: l.ctx.Value("userName").(string),
		DelFlag:  req.DelFlag,
	})

	if err != nil {
		logc.Errorf(l.ctx, "更新岗位信息失败,参数：%+v,响应：%s", req, err.Error())
		return nil, errorx.NewDefaultError("更新岗位失败")
	}

	return &types.UpdateJobResp{
		Code:    "000000",
		Message: "更新岗位信息成功",
	}, nil
}
