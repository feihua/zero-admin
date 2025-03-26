package membermembertagrelationservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/ums/gen/query"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// DeleteMemberMemberTagRelationLogic 删除用户和标签关系
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

// DeleteMemberMemberTagRelation 删除用户和标签关系
func (l *DeleteMemberMemberTagRelationLogic) DeleteMemberMemberTagRelation(in *umsclient.DeleteMemberMemberTagRelationReq) (*umsclient.DeleteMemberMemberTagRelationResp, error) {
	q := query.UmsMemberMemberTagRelation
	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Delete()
	if err != nil {
		logc.Errorf(l.ctx, "删除用户和标签关系失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("删除用户和标签关系失败")
	}

	return &umsclient.DeleteMemberMemberTagRelationResp{}, nil
}
