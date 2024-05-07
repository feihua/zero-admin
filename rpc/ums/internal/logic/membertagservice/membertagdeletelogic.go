package membertagservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/ums/gen/query"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"

	"github.com/feihua/zero-admin/rpc/ums/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

// MemberTagDeleteLogic 会员标签
/*
Author: LiuFeiHua
Date: 2024/5/7 9:21
*/
type MemberTagDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberTagDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberTagDeleteLogic {
	return &MemberTagDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// MemberTagDelete 删除会员标签
func (l *MemberTagDeleteLogic) MemberTagDelete(in *umsclient.MemberTagDeleteReq) (*umsclient.MemberTagDeleteResp, error) {
	q := query.UmsMemberTag
	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Delete()
	if err != nil {
		return nil, err
	}

	return &umsclient.MemberTagDeleteResp{}, nil
}
