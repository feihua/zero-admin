package topiccommentservicelogic

import (
	"context"

	"github.com/feihua/zero-admin/rpc/cms/cmsclient"
	"github.com/feihua/zero-admin/rpc/cms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryTopicCommentListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryTopicCommentListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryTopicCommentListLogic {
	return &QueryTopicCommentListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询专题评论表列表
func (l *QueryTopicCommentListLogic) QueryTopicCommentList(in *cmsclient.QueryTopicCommentListReq) (*cmsclient.QueryTopicCommentListResp, error) {
	// todo: add your logic here and delete this line

	return &cmsclient.QueryTopicCommentListResp{}, nil
}
