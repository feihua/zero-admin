package membertaskservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/ums/gen/model"
	"github.com/feihua/zero-admin/rpc/ums/gen/query"
	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"
	"github.com/zeromicro/go-zero/core/logx"
)

// MemberTaskAddLogic 会员任务
/*
Author: LiuFeiHua
Date: 2024/5/7 9:44
*/
type MemberTaskAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberTaskAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberTaskAddLogic {
	return &MemberTaskAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// MemberTaskAdd 添加会员任务
func (l *MemberTaskAddLogic) MemberTaskAdd(in *umsclient.MemberTaskAddReq) (*umsclient.MemberTaskAddResp, error) {
	err := query.UmsMemberTask.WithContext(l.ctx).Create(&model.UmsMemberTask{
		TaskName:     in.TaskName,
		TaskGrowth:   in.TaskGrowth,
		TaskIntegral: in.TaskIntegral,
		TaskType:     in.TaskType,
		CreateBy:     in.CreateBy,
	})

	if err != nil {
		return nil, err
	}

	return &umsclient.MemberTaskAddResp{}, nil
}
