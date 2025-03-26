package membertaskservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/ums/gen/query"
	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateMemberTaskStatusLogic 更新会员任务
/*
Author: LiuFeiHua
Date: 2025/01/24 10:32:59
*/
type UpdateMemberTaskStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateMemberTaskStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateMemberTaskStatusLogic {
	return &UpdateMemberTaskStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateMemberTaskStatus 更新会员任务状态
func (l *UpdateMemberTaskStatusLogic) UpdateMemberTaskStatus(in *umsclient.UpdateMemberTaskStatusReq) (*umsclient.UpdateMemberTaskStatusResp, error) {
	q := query.UmsMemberTask

	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Update(q.Status, in.Status)

	if err != nil {
		logc.Errorf(l.ctx, "更新会员任务状态失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("更新会员任务状态失败")
	}

	return &umsclient.UpdateMemberTaskStatusResp{}, nil
}
