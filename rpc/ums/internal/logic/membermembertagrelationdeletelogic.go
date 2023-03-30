package logic

import (
	"context"

	"zero-admin/rpc/ums/internal/svc"
	"zero-admin/rpc/ums/umsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type MemberMemberTagRelationDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberMemberTagRelationDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberMemberTagRelationDeleteLogic {
	return &MemberMemberTagRelationDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MemberMemberTagRelationDeleteLogic) MemberMemberTagRelationDelete(in *umsclient.MemberMemberTagRelationDeleteReq) (*umsclient.MemberMemberTagRelationDeleteResp, error) {
	// todo: add your logic here and delete this line
	err := l.svcCtx.UmsMemberMemberTagRelationModel.Delete(l.ctx, in.Id)

	if err != nil {
		return nil, err
	}
	return &umsclient.MemberMemberTagRelationDeleteResp{}, nil
}
