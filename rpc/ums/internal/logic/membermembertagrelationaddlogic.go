package logic

import (
	"context"
	"zero-admin/rpc/model/umsmodel"

	"zero-admin/rpc/ums/internal/svc"
	"zero-admin/rpc/ums/ums"

	"github.com/zeromicro/go-zero/core/logx"
)

type MemberMemberTagRelationAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberMemberTagRelationAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberMemberTagRelationAddLogic {
	return &MemberMemberTagRelationAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MemberMemberTagRelationAddLogic) MemberMemberTagRelationAdd(in *ums.MemberMemberTagRelationAddReq) (*ums.MemberMemberTagRelationAddResp, error) {
	_, err := l.svcCtx.UmsMemberMemberTagRelationModel.Insert(umsmodel.UmsMemberMemberTagRelation{
		MemberId: in.MemberId,
		TagId:    in.TagId,
	})
	if err != nil {
		return nil, err
	}

	return &ums.MemberMemberTagRelationAddResp{}, nil
}
