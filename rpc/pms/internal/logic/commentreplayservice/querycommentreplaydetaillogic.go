package commentreplayservicelogic

import (
	"context"

	"github.com/feihua/zero-admin/rpc/pms/internal/svc"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryCommentReplayDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryCommentReplayDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryCommentReplayDetailLogic {
	return &QueryCommentReplayDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询产品评价回复表详情
func (l *QueryCommentReplayDetailLogic) QueryCommentReplayDetail(in *pmsclient.QueryCommentReplayDetailReq) (*pmsclient.QueryCommentReplayDetailResp, error) {
	// todo: add your logic here and delete this line

	return &pmsclient.QueryCommentReplayDetailResp{}, nil
}
