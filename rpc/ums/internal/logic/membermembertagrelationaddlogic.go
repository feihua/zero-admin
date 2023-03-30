package logic

import (
	"context"

	"zero-admin/rpc/model/umsmodel"
	"zero-admin/rpc/ums/internal/svc"
	"zero-admin/rpc/ums/umsclient"

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

func (l *MemberMemberTagRelationAddLogic) MemberMemberTagRelationAdd(in *umsclient.MemberMemberTagRelationAddReq) (*umsclient.MemberMemberTagRelationAddResp, error) {
	// todo: add your logic here and delete this line
	_, err := l.svcCtx.UmsMemberMemberTagRelationModel.Insert(l.ctx, &umsmodel.UmsMemberMemberTagRelation{
		MemberId: in.MemberId,
		TagId:    in.TagId,
	})
	if err != nil {
		return nil, err
	}
	return &umsclient.MemberMemberTagRelationAddResp{}, nil
}
