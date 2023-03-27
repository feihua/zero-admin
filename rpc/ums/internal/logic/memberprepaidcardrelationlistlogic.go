package logic

import (
	"context"

	"zero-admin/rpc/ums/internal/svc"
	"zero-admin/rpc/ums/umsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type MemberPrepaidCardRelationListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberPrepaidCardRelationListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberPrepaidCardRelationListLogic {
	return &MemberPrepaidCardRelationListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MemberPrepaidCardRelationListLogic) MemberPrepaidCardRelationList(in *umsclient.MemberPrepaidCardRelationListReq) (*umsclient.MemberPrepaidCardRelationListResp, error) {
	// todo: add your logic here and delete this line

	return &umsclient.MemberPrepaidCardRelationListResp{}, nil
}
