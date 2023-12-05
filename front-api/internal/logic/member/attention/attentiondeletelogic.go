package attention

import (
	"context"
	"encoding/json"
	"zero-admin/rpc/ums/umsclient"

	"zero-admin/front-api/internal/svc"
	"zero-admin/front-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// AttentionDeleteLogic
/*
Author: LiuFeiHua
Date: 2023/12/4 17:12
*/
type AttentionDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAttentionDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AttentionDeleteLogic {
	return &AttentionDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// AttentionDelete 取消关注
func (l *AttentionDeleteLogic) AttentionDelete(req *types.DeleteAttentionReq) (resp *types.DeleteAttentionResp, err error) {
	memberId, _ := l.ctx.Value("memberId").(json.Number).Int64()
	_, _ = l.svcCtx.MemberAttentionService.MemberBrandAttentionDelete(l.ctx, &umsclient.MemberBrandAttentionDeleteReq{
		Id:       req.Id,
		MemberId: memberId,
	})

	return &types.DeleteAttentionResp{
		Code:    0,
		Message: "操作成功",
	}, nil
}
