package membertaskservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/ums/gen/model"
	"github.com/feihua/zero-admin/rpc/ums/gen/query"
	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateMemberTaskLogic 会员任务
/*
Author: LiuFeiHua
Date: 2024/5/7 9:45
*/
type UpdateMemberTaskLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateMemberTaskLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateMemberTaskLogic {
	return &UpdateMemberTaskLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateMemberTask 更新会员任务表
func (l *UpdateMemberTaskLogic) UpdateMemberTask(in *umsclient.UpdateMemberTaskReq) (*umsclient.UpdateMemberTaskResp, error) {
	_, err := query.UmsMemberTask.WithContext(l.ctx).Updates(&model.UmsMemberTask{
		ID:           in.Id,
		TaskName:     in.TaskName,
		TaskGrowth:   in.TaskGrowth,
		TaskIntegral: in.TaskIntegral,
		TaskType:     in.TaskType,
		UpdateBy:     in.UpdateBy,
	})
	if err != nil {
		return nil, err
	}

	return &umsclient.UpdateMemberTaskResp{}, nil
}
