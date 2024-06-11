package memberloginlogservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/ums/gen/query"

	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// DeleteMemberLoginLogLogic 删除会员登录记录
/*
Author: LiuFeiHua
Date: 2024/6/11 14:14
*/
type DeleteMemberLoginLogLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteMemberLoginLogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteMemberLoginLogLogic {
	return &DeleteMemberLoginLogLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// DeleteMemberLoginLog 删除会员登录记录
func (l *DeleteMemberLoginLogLogic) DeleteMemberLoginLog(in *umsclient.DeleteMemberLoginLogReq) (*umsclient.DeleteMemberLoginLogResp, error) {
	q := query.UmsMemberLoginLog
	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Delete()
	if err != nil {
		return nil, err
	}

	return &umsclient.DeleteMemberLoginLogResp{}, nil
}
