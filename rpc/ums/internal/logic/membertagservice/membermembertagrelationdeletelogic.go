package membertagservicelogic

import (
	"context"
	"zero-admin/rpc/ums/umsclient"

	"zero-admin/rpc/ums/internal/svc"

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
	err := l.svcCtx.UmsMemberMemberTagRelationModel.DeleteByIds(l.ctx, in.Ids)

	if err != nil {
		return nil, err
	}

	return &umsclient.MemberMemberTagRelationDeleteResp{}, nil
}
