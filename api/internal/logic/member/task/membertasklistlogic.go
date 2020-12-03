package logic

import (
	"context"

	"go-zero-admin/api/internal/svc"
	"go-zero-admin/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type MemberTaskListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMemberTaskListLogic(ctx context.Context, svcCtx *svc.ServiceContext) MemberTaskListLogic {
	return MemberTaskListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MemberTaskListLogic) MemberTaskList(req types.ListMemberTaskReq) (*types.ListMemberTaskResp, error) {
	// todo: add your logic here and delete this line

	return &types.ListMemberTaskResp{}, nil
}
