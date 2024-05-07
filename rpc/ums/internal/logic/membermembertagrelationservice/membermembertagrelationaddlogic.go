package membermembertagrelationservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/ums/gen/model"
	"github.com/feihua/zero-admin/rpc/ums/gen/query"
	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// MemberMemberTagRelationAddLogic 用户和标签关系
/*
Author: LiuFeiHua
Date: 2024/5/7 10:12
*/
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

// MemberMemberTagRelationAdd 添加用户和标签关系
func (l *MemberMemberTagRelationAddLogic) MemberMemberTagRelationAdd(in *umsclient.MemberMemberTagRelationAddReq) (*umsclient.MemberMemberTagRelationAddResp, error) {
	err := query.UmsMemberMemberTagRelation.WithContext(l.ctx).Create(&model.UmsMemberMemberTagRelation{
		MemberID: in.MemberId,
		TagID:    in.TagId,
	})

	if err != nil {
		return nil, err
	}

	return &umsclient.MemberMemberTagRelationAddResp{}, nil
}
