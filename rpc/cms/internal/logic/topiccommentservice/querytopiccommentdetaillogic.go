package topiccommentservicelogic

import (
	"context"

	"github.com/feihua/zero-admin/rpc/cms/cmsclient"
	"github.com/feihua/zero-admin/rpc/cms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryTopicCommentDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryTopicCommentDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryTopicCommentDetailLogic {
	return &QueryTopicCommentDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询专题评论表详情
func (l *QueryTopicCommentDetailLogic) QueryTopicCommentDetail(in *cmsclient.QueryTopicCommentDetailReq) (*cmsclient.QueryTopicCommentDetailResp, error) {
	// todo: add your logic here and delete this line

	return &cmsclient.QueryTopicCommentDetailResp{}, nil
}
