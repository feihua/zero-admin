package logic

import (
	"context"

	"go-zero-admin/api/internal/svc"
	"go-zero-admin/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type MemberTaskUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMemberTaskUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) MemberTaskUpdateLogic {
	return MemberTaskUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MemberTaskUpdateLogic) MemberTaskUpdate(req types.UpdateMemberTaskReq) (*types.UpdateMemberTaskResp, error) {
	// todo: add your logic here and delete this line

	return &types.UpdateMemberTaskResp{}, nil
}
