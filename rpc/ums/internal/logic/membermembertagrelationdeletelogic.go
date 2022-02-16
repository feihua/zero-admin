package logic

import (
	"context"

	"zero-admin/rpc/ums/internal/svc"
	"zero-admin/rpc/ums/ums"

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

func (l *MemberMemberTagRelationDeleteLogic) MemberMemberTagRelationDelete(in *ums.MemberMemberTagRelationDeleteReq) (*ums.MemberMemberTagRelationDeleteResp, error) {
	err := l.svcCtx.UmsMemberMemberTagRelationModel.Delete(in.Id)

	if err != nil {
		return nil, err
	}

	return &ums.MemberMemberTagRelationDeleteResp{}, nil
}
