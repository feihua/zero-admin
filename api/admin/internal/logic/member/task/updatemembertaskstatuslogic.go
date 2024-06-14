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

// UpdateMemberTaskStatusLogic 更新会员任务状态
/*
Author: LiuFeiHua
Date: 2024/5/23 9:13
*/
type UpdateMemberTaskStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateMemberTaskStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateMemberTaskStatusLogic {
	return &UpdateMemberTaskStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// UpdateMemberTaskStatus 更新会员任务状态
func (l *UpdateMemberTaskStatusLogic) UpdateMemberTaskStatus(req *types.UpdateMemberTaskStatusReq) (resp *types.UpdateMemberTaskStatusResp, err error) {
	_, err = l.svcCtx.MemberTaskService.UpdateMemberTaskStatus(l.ctx, &umsclient.UpdateMemberTaskStatusReq{
		UpdateBy: l.ctx.Value("userName").(string),
		Ids:      req.Ids,
		Status:   req.Status,
	})

	if err != nil {
		logc.Errorf(l.ctx, "更新会员任务状态信息失败,参数：%+v,响应：%s", req, err.Error())
		return nil, errorx.NewDefaultError("更新会员任务状态失败")
	}

	return &types.UpdateMemberTaskStatusResp{
		Code:    "000000",
		Message: "更新会员任务状态成功",
	}, nil

}
