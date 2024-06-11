package topiccommentservicelogic

import (
	"context"

	"github.com/feihua/zero-admin/rpc/cms/cmsclient"
	"github.com/feihua/zero-admin/rpc/cms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddTopicCommentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddTopicCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddTopicCommentLogic {
	return &AddTopicCommentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 添加专题评论表
func (l *AddTopicCommentLogic) AddTopicComment(in *cmsclient.AddTopicCommentReq) (*cmsclient.AddTopicCommentResp, error) {
	// todo: add your logic here and delete this line

	return &cmsclient.AddTopicCommentResp{}, nil
}
