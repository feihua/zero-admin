package membertagrelationservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/ums/gen/model"
	"github.com/feihua/zero-admin/rpc/ums/gen/query"
	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
)

// AddMemberTagRelationLogic 添加会员标签关联
/*
Author: LiuFeiHua
Date: 2025/05/22 10:44:59
*/
type AddMemberTagRelationLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddMemberTagRelationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddMemberTagRelationLogic {
	return &AddMemberTagRelationLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// AddMemberTagRelation 添加会员标签关联
func (l *AddMemberTagRelationLogic) AddMemberTagRelation(in *umsclient.AddMemberTagRelationReq) (*umsclient.AddMemberTagRelationResp, error) {
	q := query.UmsMemberTagRelation

	item := &model.UmsMemberTagRelation{
		MemberID: in.MemberId, // 会员ID
		TagID:    in.TagId,    // 标签ID
	}

	err := q.WithContext(l.ctx).Create(item)
	if err != nil {
		logc.Errorf(l.ctx, "添加会员标签关联失败,参数:%+v,异常:%s", item, err.Error())
		return nil, errors.New("添加会员标签关联失败")
	}

	return &umsclient.AddMemberTagRelationResp{}, nil
}
