package memberattentionservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/ums/gen/query"

	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// MemberBrandAttentionDeleteLogic
/*
Author: LiuFeiHua
Date: 2023/12/4 16:53
*/
type MemberBrandAttentionDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberBrandAttentionDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberBrandAttentionDeleteLogic {
	return &MemberBrandAttentionDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// MemberBrandAttentionDelete 取消关注
func (l *MemberBrandAttentionDeleteLogic) MemberBrandAttentionDelete(in *umsclient.MemberBrandAttentionDeleteReq) (*umsclient.MemberBrandAttentionDeleteResp, error) {
	q := query.UmsMemberBrandAttention
	_, err := q.WithContext(l.ctx).Where(q.ID.Eq(in.Id), q.MemberID.Eq(in.MemberId)).Delete()
	if err != nil {
		return nil, err
	}

	return &umsclient.MemberBrandAttentionDeleteResp{}, nil
}
