package membermembertagrelationservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/ums/gen/query"

	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// DeleteMemberMemberTagRelationLogic 删除用户和标签关系表
/*
Author: LiuFeiHua
Date: 2024/6/11 14:12
*/
type DeleteMemberMemberTagRelationLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteMemberMemberTagRelationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteMemberMemberTagRelationLogic {
	return &DeleteMemberMemberTagRelationLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// DeleteMemberMemberTagRelation 删除用户和标签关系表
func (l *DeleteMemberMemberTagRelationLogic) DeleteMemberMemberTagRelation(in *umsclient.DeleteMemberMemberTagRelationReq) (*umsclient.DeleteMemberMemberTagRelationResp, error) {
	q := query.UmsMemberMemberTagRelation
	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Delete()
	if err != nil {
		return nil, err
	}

	return &umsclient.DeleteMemberMemberTagRelationResp{}, nil
}
