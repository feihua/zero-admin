package history

import (
	"context"
	"encoding/json"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"

	"github.com/feihua/zero-admin/api/front/internal/svc"
	"github.com/feihua/zero-admin/api/front/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// DeleteReadHistoryLogic
/*
Author: LiuFeiHua
Date: 2024/5/16 10:47
*/
type DeleteReadHistoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteReadHistoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteReadHistoryLogic {
	return &DeleteReadHistoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// DeleteReadHistory 删除浏览记录
func (l *DeleteReadHistoryLogic) DeleteReadHistory(req *types.ReadHistoryDeleteReq) (resp *types.ReadHistoryDeleteResp, err error) {
	memberId, _ := l.ctx.Value("memberId").(json.Number).Int64()
	_, _ = l.svcCtx.MemberReadHistoryService.MemberReadHistoryDelete(l.ctx, &umsclient.MemberReadHistoryDeleteReq{
		Ids:      req.Ids,
		MemberId: memberId,
	})

	return &types.ReadHistoryDeleteResp{
		Code:    0,
		Message: "删除浏览记录",
	}, nil
}
