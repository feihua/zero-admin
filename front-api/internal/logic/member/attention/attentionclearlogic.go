package attention

import (
	"context"
	"encoding/json"
	"zero-admin/rpc/ums/umsclient"

	"zero-admin/front-api/internal/svc"
	"zero-admin/front-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// AttentionClearLogic
/*
Author: LiuFeiHua
Date: 2023/12/4 17:11
*/
type AttentionClearLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAttentionClearLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AttentionClearLogic {
	return &AttentionClearLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// AttentionClear 清空所有关注
func (l *AttentionClearLogic) AttentionClear() (resp *types.ClearAttentionResp, err error) {
	memberId, _ := l.ctx.Value("memberId").(json.Number).Int64()
	_, _ = l.svcCtx.MemberAttentionService.MemberBrandAttentionClear(l.ctx, &umsclient.MemberBrandAttentionClearReq{
		MemberId: memberId,
	})

	return &types.ClearAttentionResp{
		Code:    0,
		Message: "操作成功",
	}, nil
}
