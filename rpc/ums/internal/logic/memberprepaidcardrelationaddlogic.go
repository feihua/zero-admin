package logic

import (
	"context"

	"zero-admin/rpc/ums/internal/svc"
	"zero-admin/rpc/ums/umsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type MemberPrepaidCardRelationAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberPrepaidCardRelationAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberPrepaidCardRelationAddLogic {
	return &MemberPrepaidCardRelationAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MemberPrepaidCardRelationAddLogic) MemberPrepaidCardRelationAdd(in *umsclient.MemberPrepaidCardRelationAddReq) (*umsclient.MemberPrepaidCardRelationAddResp, error) {
	// todo: add your logic here and delete this line

	return &umsclient.MemberPrepaidCardRelationAddResp{}, nil
}
