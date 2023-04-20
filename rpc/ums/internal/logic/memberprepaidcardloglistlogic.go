package logic

import (
	"context"

	"zero-admin/rpc/ums/internal/svc"
	"zero-admin/rpc/ums/umsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type MemberPrepaidCardLogListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberPrepaidCardLogListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberPrepaidCardLogListLogic {
	return &MemberPrepaidCardLogListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MemberPrepaidCardLogListLogic) MemberPrepaidCardLogList(in *umsclient.MemberPrepaidCardLogListReq) (*umsclient.MemberPrepaidCardLogListResp, error) {
	// todo: add your logic here and delete this line

	return &umsclient.MemberPrepaidCardLogListResp{}, nil
}
