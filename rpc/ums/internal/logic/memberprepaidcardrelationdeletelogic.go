package logic

import (
	"context"

	"zero-admin/rpc/ums/internal/svc"
	"zero-admin/rpc/ums/umsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type MemberPrepaidCardRelationDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberPrepaidCardRelationDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberPrepaidCardRelationDeleteLogic {
	return &MemberPrepaidCardRelationDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MemberPrepaidCardRelationDeleteLogic) MemberPrepaidCardRelationDelete(in *umsclient.MemberPrepaidCardRelationDeleteReq) (*umsclient.MemberPrepaidCardRelationDeleteResp, error) {
	// todo: add your logic here and delete this line

	return &umsclient.MemberPrepaidCardRelationDeleteResp{}, nil
}
