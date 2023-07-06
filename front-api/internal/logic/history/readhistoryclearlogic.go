package history

import (
	"context"
	"zero-admin/rpc/ums/umsclient"

	"zero-admin/front-api/internal/svc"
	"zero-admin/front-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

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

func (l *ReadHistoryClearLogic) ReadHistoryClear() (resp *types.ReadHistoryClearResp, err error) {
	_, _ = l.svcCtx.Ums.MemberReadHistoryDelete(l.ctx, &umsclient.MemberReadHistoryDeleteReq{
		Ids:      nil,
		MemberId: l.ctx.Value("memberId").(int64),
	})

	return &types.ReadHistoryClearResp{
		Code:    0,
		Message: "操作成功",
	}, nil
}
