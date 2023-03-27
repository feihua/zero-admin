package logic

import (
	"context"

	"zero-admin/rpc/ums/internal/svc"
	"zero-admin/rpc/ums/umsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type MemberLoginLogUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberLoginLogUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberLoginLogUpdateLogic {
	return &MemberLoginLogUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MemberLoginLogUpdateLogic) MemberLoginLogUpdate(in *umsclient.MemberLoginLogUpdateReq) (*umsclient.MemberLoginLogUpdateResp, error) {
	// todo: add your logic here and delete this line

	return &umsclient.MemberLoginLogUpdateResp{}, nil
}
