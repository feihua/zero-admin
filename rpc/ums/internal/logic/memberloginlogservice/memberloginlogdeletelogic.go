package memberloginlogservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/ums/gen/query"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"

	"github.com/feihua/zero-admin/rpc/ums/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

// MemberLoginLogDeleteLogic 会员登录日志
/*
Author: LiuFeiHua
Date: 2024/5/7 9:34
*/
type MemberLoginLogDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberLoginLogDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberLoginLogDeleteLogic {
	return &MemberLoginLogDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// MemberLoginLogDelete 删除会员登录日志
func (l *MemberLoginLogDeleteLogic) MemberLoginLogDelete(in *umsclient.MemberLoginLogDeleteReq) (*umsclient.MemberLoginLogDeleteResp, error) {
	q := query.UmsMemberLoginLog
	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Delete()
	if err != nil {
		return nil, err
	}

	return &umsclient.MemberLoginLogDeleteResp{}, nil
}
