package memberlevelservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/ums/gen/query"

	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// DeleteMemberLevelLogic 删除会员等级表
/*
Author: LiuFeiHua
Date: 2024/6/11 14:16
*/
type DeleteMemberLevelLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteMemberLevelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteMemberLevelLogic {
	return &DeleteMemberLevelLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// DeleteMemberLevel 删除会员等级表
func (l *DeleteMemberLevelLogic) DeleteMemberLevel(in *umsclient.DeleteMemberLevelReq) (*umsclient.DeleteMemberLevelResp, error) {
	q := query.UmsMemberLevel
	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Delete()
	if err != nil {
		return nil, err
	}

	return &umsclient.DeleteMemberLevelResp{}, nil
}
