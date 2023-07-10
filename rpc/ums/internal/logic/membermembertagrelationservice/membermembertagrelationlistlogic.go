package membermembertagrelationservicelogic

import (
	"context"

	"zero-admin/rpc/ums/internal/svc"
	"zero-admin/rpc/ums/umsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type MemberMemberTagRelationListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberMemberTagRelationListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberMemberTagRelationListLogic {
	return &MemberMemberTagRelationListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MemberMemberTagRelationListLogic) MemberMemberTagRelationList(in *umsclient.MemberMemberTagRelationListReq) (*umsclient.MemberMemberTagRelationListResp, error) {
	// todo: add your logic here and delete this line

	return &umsclient.MemberMemberTagRelationListResp{}, nil
}
