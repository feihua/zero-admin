package logic

import (
	"context"

	"zero-admin/rpc/ums/internal/svc"
	"zero-admin/rpc/ums/umsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type MemberPrepaidCardRelationUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberPrepaidCardRelationUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberPrepaidCardRelationUpdateLogic {
	return &MemberPrepaidCardRelationUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MemberPrepaidCardRelationUpdateLogic) MemberPrepaidCardRelationUpdate(in *umsclient.MemberPrepaidCardRelationUpdateReq) (*umsclient.MemberPrepaidCardRelationUpdateResp, error) {
	// todo: add your logic here and delete this line

	return &umsclient.MemberPrepaidCardRelationUpdateResp{}, nil
}
