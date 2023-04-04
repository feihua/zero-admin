package logic

import (
	"context"

	"zero-admin/rpc/ums/internal/svc"
	"zero-admin/rpc/ums/umsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type MemberDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberDeleteLogic {
	return &MemberDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MemberDeleteLogic) MemberDelete(in *umsclient.MemberDeleteReq) (*umsclient.MemberDeleteResp, error) {
	err := l.svcCtx.UmsMemberModel.Delete(l.ctx, in.Id)

	if err != nil {
		return nil, err
	}
	return &umsclient.MemberDeleteResp{}, nil
}
