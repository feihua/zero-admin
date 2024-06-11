package membermembertagrelationservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/ums/gen/model"
	"github.com/feihua/zero-admin/rpc/ums/gen/query"

	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// AddMemberMemberTagRelationLogic 添加用户和标签关系表
/*
Author: LiuFeiHua
Date: 2024/6/11 14:12
*/
type AddMemberMemberTagRelationLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddMemberMemberTagRelationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddMemberMemberTagRelationLogic {
	return &AddMemberMemberTagRelationLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// AddMemberMemberTagRelation 添加用户和标签关系表
func (l *AddMemberMemberTagRelationLogic) AddMemberMemberTagRelation(in *umsclient.AddMemberMemberTagRelationReq) (*umsclient.AddMemberMemberTagRelationResp, error) {
	err := query.UmsMemberMemberTagRelation.WithContext(l.ctx).Create(&model.UmsMemberMemberTagRelation{
		MemberID: in.MemberId,
		TagID:    in.TagId,
	})

	if err != nil {
		return nil, err
	}

	return &umsclient.AddMemberMemberTagRelationResp{}, nil
}
