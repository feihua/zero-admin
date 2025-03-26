package memberbrandattentionservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/ums/gen/query"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// DeleteMemberBrandAttentionLogic 取消品牌关注/清空当前用户品牌关注列表
/*
Author: LiuFeiHua
Date: 2024/6/11 14:19
*/
type DeleteMemberBrandAttentionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteMemberBrandAttentionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteMemberBrandAttentionLogic {
	return &DeleteMemberBrandAttentionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// DeleteMemberBrandAttention 取消品牌关注/清空当前用户品牌关注列表
func (l *DeleteMemberBrandAttentionLogic) DeleteMemberBrandAttention(in *umsclient.DeleteMemberBrandAttentionReq) (*umsclient.DeleteMemberBrandAttentionResp, error) {
	q := query.UmsMemberBrandAttention
	attentionDo := q.WithContext(l.ctx).Where(q.MemberID.Eq(in.MemberId))
	if len(in.Ids) > 0 {
		attentionDo = attentionDo.Where(q.ID.In(in.Ids...))
	}
	_, err := attentionDo.Delete()
	if err != nil {
		logc.Errorf(l.ctx, "取消品牌关注失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("取消品牌关注失败")
	}

	return &umsclient.DeleteMemberBrandAttentionResp{}, nil
}
