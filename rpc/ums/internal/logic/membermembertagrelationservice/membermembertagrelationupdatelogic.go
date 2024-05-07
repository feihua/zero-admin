package membermembertagrelationservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/ums/gen/model"
	"github.com/feihua/zero-admin/rpc/ums/gen/query"

	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// MemberMemberTagRelationUpdateLogic 用户和标签关系
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

// MemberMemberTagRelationUpdate 更新用户和标签关系
func (l *MemberMemberTagRelationUpdateLogic) MemberMemberTagRelationUpdate(in *umsclient.MemberMemberTagRelationUpdateReq) (*umsclient.MemberMemberTagRelationUpdateResp, error) {
	_, err := query.UmsMemberMemberTagRelation.WithContext(l.ctx).Updates(&model.UmsMemberMemberTagRelation{
		ID:       in.Id,
		MemberID: in.MemberId,
		TagID:    in.TagId,
	})

	if err != nil {
		return nil, err
	}

	return &umsclient.MemberMemberTagRelationUpdateResp{}, nil
}
