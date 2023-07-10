package history

import (
	"context"
	"zero-admin/rpc/ums/umsclient"

	"zero-admin/front-api/internal/svc"
	"zero-admin/front-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ReadHistoryDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewReadHistoryDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ReadHistoryDeleteLogic {
	return &ReadHistoryDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ReadHistoryDeleteLogic) ReadHistoryDelete(req *types.ReadHistoryDeleteReq) (resp *types.ReadHistoryDeleteResp, err error) {
	_, _ = l.svcCtx.MemberReadHistoryService.MemberReadHistoryDelete(l.ctx, &umsclient.MemberReadHistoryDeleteReq{
		Ids:      req.Ids,
		MemberId: l.ctx.Value("memberId").(int64),
	})

	return &types.ReadHistoryDeleteResp{
		Code:    0,
		Message: "操作成功",
	}, nil
}
