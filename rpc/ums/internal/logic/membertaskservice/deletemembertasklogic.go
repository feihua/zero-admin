package membertaskservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/ums/gen/query"

	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// DeleteMemberTaskLogic 会员任务
/*
Author: LiuFeiHua
Date: 2024/5/7 9:20
*/
type DeleteMemberTaskLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteMemberTaskLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteMemberTaskLogic {
	return &DeleteMemberTaskLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// DeleteMemberTask 删除会员任务表
func (l *DeleteMemberTaskLogic) DeleteMemberTask(in *umsclient.DeleteMemberTaskReq) (*umsclient.DeleteMemberTaskResp, error) {
	q := query.UmsMemberTask
	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Delete()
	if err != nil {
		return nil, err
	}

	return &umsclient.DeleteMemberTaskResp{}, nil
}
