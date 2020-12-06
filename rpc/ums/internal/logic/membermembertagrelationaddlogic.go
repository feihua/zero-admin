package logic

import (
	"context"

	"go-zero-admin/rpc/ums/internal/svc"
	"go-zero-admin/rpc/ums/ums"

	"github.com/tal-tech/go-zero/core/logx"
)

type MemberMemberTagRelationAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberMemberTagRelationAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberMemberTagRelationAddLogic {
	return &MemberMemberTagRelationAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MemberMemberTagRelationAddLogic) MemberMemberTagRelationAdd(in *ums.MemberMemberTagRelationAddReq) (*ums.MemberMemberTagRelationAddResp, error) {
	// todo: add your logic here and delete this line

	return &ums.MemberMemberTagRelationAddResp{}, nil
}
