package member_task

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc/status"

	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateMemberTaskStatusLogic 更新会员任务状态状态
/*
Author: LiuFeiHua
Date: 2025/05/22 10:44:59
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
func (l *UpdateMemberTaskStatusLogic) UpdateMemberTaskStatus(req *types.UpdateMemberTaskStatusReq) (resp *types.BaseResp, err error) {
	userId, err := common.GetUserId(l.ctx)
	if err != nil {
		return nil, err
	}
	_, err = l.svcCtx.MemberTaskService.UpdateMemberTaskStatus(l.ctx, &umsclient.UpdateMemberTaskStatusReq{
		Ids:      req.Ids,    // 主键ID
		Status:   req.Status, // 状态：0-禁用，1-启用
		UpdateBy: userId,
	})

	if err != nil {
		logc.Errorf(l.ctx, "更新会员任务状态失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	return &types.BaseResp{
		Code:    "000000",
		Message: "更新会员任务状态成功",
	}, nil
}
