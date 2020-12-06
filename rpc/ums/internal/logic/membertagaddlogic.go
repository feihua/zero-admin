package logic

import (
	"context"

	"go-zero-admin/rpc/ums/internal/svc"
	"go-zero-admin/rpc/ums/ums"

	"github.com/tal-tech/go-zero/core/logx"
)

type MemberTagAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberTagAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberTagAddLogic {
	return &MemberTagAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MemberTagAddLogic) MemberTagAdd(in *ums.MemberTagAddReq) (*ums.MemberTagAddResp, error) {
	// todo: add your logic here and delete this line

	return &ums.MemberTagAddResp{}, nil
}
