package membertagservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/model/umsmodel"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"

	"github.com/feihua/zero-admin/rpc/ums/internal/svc"

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

func (l *MemberMemberTagRelationUpdateLogic) MemberMemberTagRelationUpdate(in *umsclient.MemberMemberTagRelationUpdateReq) (*umsclient.MemberMemberTagRelationUpdateResp, error) {
	err := l.svcCtx.UmsMemberMemberTagRelationModel.Update(l.ctx, &umsmodel.UmsMemberMemberTagRelation{
		Id:       in.Id,
		MemberId: in.MemberId,
		TagId:    in.TagId,
	})
	if err != nil {
		return nil, err
	}

	return &umsclient.MemberMemberTagRelationUpdateResp{}, nil
}
