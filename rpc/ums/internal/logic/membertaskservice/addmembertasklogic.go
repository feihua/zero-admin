package membertaskservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/ums/gen/model"
	"github.com/feihua/zero-admin/rpc/ums/gen/query"
	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// AddMemberTaskLogic 会员任务
/*
Author: LiuFeiHua
Date: 2024/5/7 9:44
*/
type AddMemberTaskLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddMemberTaskLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddMemberTaskLogic {
	return &AddMemberTaskLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// AddMemberTask 添加会员任务表
func (l *AddMemberTaskLogic) AddMemberTask(in *umsclient.AddMemberTaskReq) (*umsclient.AddMemberTaskResp, error) {
	err := query.UmsMemberTask.WithContext(l.ctx).Create(&model.UmsMemberTask{
		TaskName:     in.TaskName,
		TaskGrowth:   in.TaskGrowth,
		TaskIntegral: in.TaskIntegral,
		TaskType:     in.TaskType,
		Status:       1,
		CreateBy:     in.CreateBy,
	})

	if err != nil {
		return nil, err
	}

	return &umsclient.AddMemberTaskResp{}, nil
}
