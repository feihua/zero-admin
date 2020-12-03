package logic

import (
	"context"

	"go-zero-admin/api/internal/svc"
	"go-zero-admin/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type MemberTaskDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMemberTaskDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) MemberTaskDeleteLogic {
	return MemberTaskDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MemberTaskDeleteLogic) MemberTaskDelete(req types.DeleteMemberTaskReq) (*types.DeleteMemberTaskResp, error) {
	// todo: add your logic here and delete this line

	return &types.DeleteMemberTaskResp{}, nil
}
