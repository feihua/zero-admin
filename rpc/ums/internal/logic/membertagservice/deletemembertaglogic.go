package membertagservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/ums/gen/query"

	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// DeleteMemberTagLogic 删除用户标签表
/*
Author: LiuFeiHua
Date: 2024/6/11 13:46
*/
type DeleteMemberTagLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteMemberTagLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteMemberTagLogic {
	return &DeleteMemberTagLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// DeleteMemberTag 删除用户标签表
func (l *DeleteMemberTagLogic) DeleteMemberTag(in *umsclient.DeleteMemberTagReq) (*umsclient.DeleteMemberTagResp, error) {
	q := query.UmsMemberTag
	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Delete()
	if err != nil {
		return nil, err
	}

	return &umsclient.DeleteMemberTagResp{}, nil
}
