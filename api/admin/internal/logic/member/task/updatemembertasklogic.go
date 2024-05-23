package task

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateMemberTaskLogic 更新会员任务
/*
Author: LiuFeiHua
Date: 2024/5/23 9:13
*/
type UpdateMemberTaskLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateMemberTaskLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateMemberTaskLogic {
	return &UpdateMemberTaskLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// UpdateMemberTask 更新会员任务
func (l *UpdateMemberTaskLogic) UpdateMemberTask(req *types.UpdateMemberTaskReq) (resp *types.UpdateMemberTaskResp, err error) {
	_, err = l.svcCtx.MemberTaskService.MemberTaskUpdate(l.ctx, &umsclient.MemberTaskUpdateReq{
		Id:           req.Id,
		TaskName:     req.Name,
		TaskGrowth:   req.Growth,
		TaskIntegral: req.Intergration,
		TaskType:     req.Type,
	})

	if err != nil {
		logc.Errorf(l.ctx, "更新会员任务信息失败,参数：%+v,响应：%s", req, err.Error())
		return nil, errorx.NewDefaultError("更新会员任务失败")
	}

	return &types.UpdateMemberTaskResp{
		Code:    "000000",
		Message: "更新会员任务成功",
	}, nil
}
