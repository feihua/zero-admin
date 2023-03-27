package logic

import (
	"context"

	"zero-admin/rpc/ums/internal/svc"
	"zero-admin/rpc/ums/umsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type MemberTagUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberTagUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberTagUpdateLogic {
	return &MemberTagUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MemberTagUpdateLogic) MemberTagUpdate(in *umsclient.MemberTagUpdateReq) (*umsclient.MemberTagUpdateResp, error) {
	// todo: add your logic here and delete this line

	return &umsclient.MemberTagUpdateResp{}, nil
}
