package member_task

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/api/admin/internal/common/res"
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
func (l *UpdateMemberTaskLogic) UpdateMemberTask(req *types.UpdateMemberTaskReq) (resp *types.BaseResp, err error) {
	_, err = l.svcCtx.MemberTaskService.UpdateMemberTask(l.ctx, &umsclient.UpdateMemberTaskReq{
		Id:           req.Id,
		TaskName:     req.TaskName,     // 任务名称
		TaskGrowth:   req.TaskGrowth,   // 赠送成长值
		TaskIntegral: req.TaskIntegral, // 赠送积分
		TaskType:     req.TaskType,     // 任务类型：0->新手任务；1->日常任务
		Status:       req.Status,       // 状态：0->禁用；1->启用
		UpdateBy:     l.ctx.Value("userName").(string),
	})

	if err != nil {
		logc.Errorf(l.ctx, "更新会员任务信息失败,参数：%+v,响应：%s", req, err.Error())
		return nil, errorx.NewDefaultError("更新会员任务失败")
	}

	return res.Success()
}
