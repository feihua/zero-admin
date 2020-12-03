package logic

import (
	"context"

	"go-zero-admin/api/internal/svc"
	"go-zero-admin/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type MemberLevelAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMemberLevelAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) MemberLevelAddLogic {
	return MemberLevelAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MemberLevelAddLogic) MemberLevelAdd(req types.AddMemberLevelReq) (*types.AddMemberLevelResp, error) {
	// todo: add your logic here and delete this line

	return &types.AddMemberLevelResp{}, nil
}
