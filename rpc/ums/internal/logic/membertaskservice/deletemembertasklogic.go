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

// DeleteMemberTaskLogic 删除会员任务
/*
Author: LiuFeiHua
Date: 2025/05/22 10:44:59
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

// DeleteMemberTask 删除会员任务
func (l *DeleteMemberTaskLogic) DeleteMemberTask(in *umsclient.DeleteMemberTaskReq) (*umsclient.DeleteMemberTaskResp, error) {
	q := query.UmsMemberTask

	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Delete()

	if err != nil {
		logc.Errorf(l.ctx, "删除会员任务失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("删除会员任务失败")
	}

	return &umsclient.DeleteMemberTaskResp{}, nil
}
