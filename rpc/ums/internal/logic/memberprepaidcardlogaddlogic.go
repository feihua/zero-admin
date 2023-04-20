package logic

import (
	"context"

	"zero-admin/rpc/ums/internal/svc"
	"zero-admin/rpc/ums/umsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type MemberPrepaidCardLogAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberPrepaidCardLogAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberPrepaidCardLogAddLogic {
	return &MemberPrepaidCardLogAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MemberPrepaidCardLogAddLogic) MemberPrepaidCardLogAdd(in *umsclient.MemberPrepaidCardLogAddReq) (*umsclient.MemberPrepaidCardLogAddResp, error) {
	// todo: add your logic here and delete this line

	return &umsclient.MemberPrepaidCardLogAddResp{}, nil
}
