package membertaskrelationservicelogic

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

// AddMemberTaskRelationLogic 添加会员任务关联
/*
Author: LiuFeiHua
Date: 2025/05/22 10:44:59
*/
type AddMemberTaskRelationLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddMemberTaskRelationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddMemberTaskRelationLogic {
	return &AddMemberTaskRelationLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// AddMemberTaskRelation 添加会员任务关联
func (l *AddMemberTaskRelationLogic) AddMemberTaskRelation(in *umsclient.AddMemberTaskRelationReq) (*umsclient.AddMemberTaskRelationResp, error) {
	q := query.UmsMemberTaskRelation

	item := &model.UmsMemberTaskRelation{
		MemberID: in.MemberId, // 会员ID
		TaskID:   in.TaskId,   // 任务ID
	}

	err := q.WithContext(l.ctx).Create(item)
	if err != nil {
		logc.Errorf(l.ctx, "添加会员任务关联失败,参数:%+v,异常:%s", item, err.Error())
		return nil, errors.New("添加会员任务关联失败")
	}

	return &umsclient.AddMemberTaskRelationResp{}, nil
}
