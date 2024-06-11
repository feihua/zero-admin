package attention

import (
	"context"
	"encoding/json"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"

	"github.com/feihua/zero-admin/api/front/internal/svc"
	"github.com/feihua/zero-admin/api/front/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// DeleteAttentionLogic //取消品牌关注/清空当前用户品牌关注列表
/*
Author: LiuFeiHua
Date: 2024/5/16 11:03
*/
type DeleteAttentionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteAttentionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteAttentionLogic {
	return &DeleteAttentionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// DeleteAttention //取消品牌关注/清空当前用户品牌关注列表
func (l *DeleteAttentionLogic) DeleteAttention(req *types.DeleteAttentionReq) (resp *types.DeleteAttentionResp, err error) {
	memberId, _ := l.ctx.Value("memberId").(json.Number).Int64()
	_, _ = l.svcCtx.MemberBrandAttentionService.DeleteMemberBrandAttention(l.ctx, &umsclient.DeleteMemberBrandAttentionReq{
		Ids:      req.BrandIds,
		MemberId: memberId,
	})

	return &types.DeleteAttentionResp{
		Code:    0,
		Message: "操作成功",
	}, nil
}
