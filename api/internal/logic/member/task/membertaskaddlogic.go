package logic

import (
	"context"

	"go-zero-admin/api/internal/svc"
	"go-zero-admin/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type MemberTaskAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMemberTaskAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) MemberTaskAddLogic {
	return MemberTaskAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MemberTaskAddLogic) MemberTaskAdd(req types.AddMemberTaskReq) (*types.AddMemberTaskResp, error) {
	// todo: add your logic here and delete this line

	return &types.AddMemberTaskResp{}, nil
}
