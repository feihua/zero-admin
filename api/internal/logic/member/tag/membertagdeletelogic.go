package logic

import (
	"context"

	"go-zero-admin/api/internal/svc"
	"go-zero-admin/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type MemberTagDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMemberTagDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) MemberTagDeleteLogic {
	return MemberTagDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MemberTagDeleteLogic) MemberTagDelete(req types.DeleteMemberTagReq) (*types.DeleteMemberTagResp, error) {
	// todo: add your logic here and delete this line

	return &types.DeleteMemberTagResp{}, nil
}
