package membertaskservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/ums/gen/model"
	"github.com/feihua/zero-admin/rpc/ums/gen/query"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"

	"github.com/feihua/zero-admin/rpc/ums/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

// MemberTaskUpdateLogic 会员任务
/*
Author: LiuFeiHua
Date: 2024/5/7 9:45
*/
type MemberTaskUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberTaskUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberTaskUpdateLogic {
	return &MemberTaskUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// MemberTaskUpdate 更新会员任务
func (l *MemberTaskUpdateLogic) MemberTaskUpdate(in *umsclient.MemberTaskUpdateReq) (*umsclient.MemberTaskUpdateResp, error) {
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

	return &umsclient.MemberTaskUpdateResp{}, nil
}
