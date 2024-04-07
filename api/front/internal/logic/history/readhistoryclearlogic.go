package history

import (
	"context"
	"encoding/json"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"

	"github.com/feihua/zero-admin/api/front/internal/svc"
	"github.com/feihua/zero-admin/api/front/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// ReadHistoryClearLogic
/*
Author: LiuFeiHua
Date: 2023/11/30 16:29
*/
type ReadHistoryClearLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewReadHistoryClearLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ReadHistoryClearLogic {
	return &ReadHistoryClearLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// ReadHistoryClear 清空浏览记录
func (l *ReadHistoryClearLogic) ReadHistoryClear() (resp *types.ReadHistoryClearResp, err error) {
	memberId, _ := l.ctx.Value("memberId").(json.Number).Int64()
	_, _ = l.svcCtx.MemberReadHistoryService.MemberReadHistoryDelete(l.ctx, &umsclient.MemberReadHistoryDeleteReq{
		MemberId: memberId,
	})

	return &types.ReadHistoryClearResp{
		Code:    0,
		Message: "操作成功",
	}, nil
}
