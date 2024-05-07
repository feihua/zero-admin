package membermembertagrelationservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/ums/gen/query"

	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// MemberMemberTagRelationDeleteLogic 用户和标签关系
/*
Author: LiuFeiHua
Date: 2024/5/7 9:33
*/
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

// MemberMemberTagRelationDelete 删除用户和标签关系
func (l *MemberMemberTagRelationDeleteLogic) MemberMemberTagRelationDelete(in *umsclient.MemberMemberTagRelationDeleteReq) (*umsclient.MemberMemberTagRelationDeleteResp, error) {
	q := query.UmsMemberMemberTagRelation
	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Delete()
	if err != nil {
		return nil, err
	}

	return &umsclient.MemberMemberTagRelationDeleteResp{}, nil
}
