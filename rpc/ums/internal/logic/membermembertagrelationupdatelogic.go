package logic

import (
	"context"
	"go-zero-admin/rpc/model/umsmodel"

	"go-zero-admin/rpc/ums/internal/svc"
	"go-zero-admin/rpc/ums/ums"

	"github.com/zeromicro/go-zero/core/logx"
)

type MemberMemberTagRelationUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberMemberTagRelationUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberMemberTagRelationUpdateLogic {
	return &MemberMemberTagRelationUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MemberMemberTagRelationUpdateLogic) MemberMemberTagRelationUpdate(in *ums.MemberMemberTagRelationUpdateReq) (*ums.MemberMemberTagRelationUpdateResp, error) {
	err := l.svcCtx.UmsMemberMemberTagRelationModel.Update(umsmodel.UmsMemberMemberTagRelation{
		Id:       in.Id,
		MemberId: in.MemberId,
		TagId:    in.TagId,
	})
	if err != nil {
		return nil, err
	}

	return &ums.MemberMemberTagRelationUpdateResp{}, nil
}
