package topiccommentservicelogic

import (
	"context"

	"github.com/feihua/zero-admin/rpc/cms/cmsclient"
	"github.com/feihua/zero-admin/rpc/cms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteTopicCommentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteTopicCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteTopicCommentLogic {
	return &DeleteTopicCommentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 删除专题评论表
func (l *DeleteTopicCommentLogic) DeleteTopicComment(in *cmsclient.DeleteTopicCommentReq) (*cmsclient.DeleteTopicCommentResp, error) {
	// todo: add your logic here and delete this line

	return &cmsclient.DeleteTopicCommentResp{}, nil
}
