package memberattentionservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/ums/gen/query"

	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// MemberBrandAttentionClearLogic
/*
Author: LiuFeiHua
Date: 2023/12/4 16:53
*/
type MemberBrandAttentionClearLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberBrandAttentionClearLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberBrandAttentionClearLogic {
	return &MemberBrandAttentionClearLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// MemberBrandAttentionClear 清空会员的所有关注的品牌
func (l *MemberBrandAttentionClearLogic) MemberBrandAttentionClear(in *umsclient.MemberBrandAttentionClearReq) (*umsclient.MemberBrandAttentionClearResp, error) {
	q := query.UmsMemberBrandAttention
	_, err := q.WithContext(l.ctx).Where(q.MemberID.Eq(in.MemberId)).Delete()

	if err != nil {
		return nil, err
	}

	return &umsclient.MemberBrandAttentionClearResp{}, nil
}
