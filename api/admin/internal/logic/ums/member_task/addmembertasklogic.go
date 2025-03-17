package member_task

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// AddMemberTaskLogic 添加会员任务
/*
Author: LiuFeiHua
Date: 2024/5/23 9:12
*/
type AddMemberTaskLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddMemberTaskLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddMemberTaskLogic {
	return &AddMemberTaskLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// AddMemberTask 添加会员任务
func (l *AddMemberTaskLogic) AddMemberTask(req *types.AddMemberTaskReq) (resp *types.AddMemberTaskResp, err error) {
	_, err = l.svcCtx.MemberTaskService.AddMemberTask(l.ctx, &umsclient.AddMemberTaskReq{
		TaskName:     req.TaskName,     // 任务名称
		TaskGrowth:   req.TaskGrowth,   // 赠送成长值
		TaskIntegral: req.TaskIntegral, // 赠送积分
		TaskType:     req.TaskType,     // 任务类型：0->新手任务；1->日常任务
		Status:       req.Status,       // 状态：0->禁用；1->启用
		CreateBy:     l.ctx.Value("userName").(string),
	})

	if err != nil {
		logc.Errorf(l.ctx, "添加会员任务信息失败,参数：%+v,响应：%s", req, err.Error())
		return nil, errorx.NewDefaultError("添加会员任务失败")
	}

	return &types.AddMemberTaskResp{
		Code:    "000000",
		Message: "添加会员任务成功",
	}, nil
}
