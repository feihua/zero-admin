package membermembertagrelationservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/ums/gen/model"
	"github.com/feihua/zero-admin/rpc/ums/gen/query"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// AddMemberMemberTagRelationLogic 添加用户和标签关系
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

// AddMemberMemberTagRelation 添加用户和标签关系
func (l *AddMemberMemberTagRelationLogic) AddMemberMemberTagRelation(in *umsclient.AddMemberMemberTagRelationReq) (*umsclient.AddMemberMemberTagRelationResp, error) {
	err := query.UmsMemberMemberTagRelation.WithContext(l.ctx).Create(&model.UmsMemberMemberTagRelation{
		MemberID: in.MemberId, // 会员id
		TagID:    in.TagId,    // 标签id
	})

	if err != nil {
		logc.Errorf(l.ctx, "添加用户和标签关系失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("添加用户和标签关系失败")
	}

	return &umsclient.AddMemberMemberTagRelationResp{}, nil
}
