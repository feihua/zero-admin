package task

import (
	"context"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

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

func (l *UpdateMemberTaskStatusLogic) UpdateMemberTaskStatus(req *types.UpdateMemberTaskStatusReq) (resp *types.UpdateMemberTaskStatusResp, err error) {
	// todo: add your logic here and delete this line

	return
}
