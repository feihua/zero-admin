package memberservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/ums/gen/query"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"

	"github.com/feihua/zero-admin/rpc/ums/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

// MemberDeleteLogic 会员信息
/*
Author: LiuFeiHua
Date: 2024/5/7 9:22
*/
type MemberDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberDeleteLogic {
	return &MemberDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// MemberDelete 删除会员信息
func (l *MemberDeleteLogic) MemberDelete(in *umsclient.MemberDeleteReq) (*umsclient.MemberDeleteResp, error) {
	q := query.UmsMember
	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Delete()
	if err != nil {
		return nil, err
	}

	return &umsclient.MemberDeleteResp{}, nil
}
