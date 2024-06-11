package membertaskservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/ums/gen/query"

	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateMemberTaskStatusLogic 更新会员任务表状态
/*
Author: LiuFeiHua
Date: 2024/6/11 13:37
*/
type UpdateMemberTaskStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateMemberTaskStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateMemberTaskStatusLogic {
	return &UpdateMemberTaskStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateMemberTaskStatus 更新会员任务表状态
func (l *UpdateMemberTaskStatusLogic) UpdateMemberTaskStatus(in *umsclient.UpdateMemberTaskStatusReq) (*umsclient.UpdateMemberTaskStatusResp, error) {
	q := query.UmsMemberTask
	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Update(q.Status, in.Status)
	if err != nil {
		return nil, err
	}
	return &umsclient.UpdateMemberTaskStatusResp{}, nil
}
