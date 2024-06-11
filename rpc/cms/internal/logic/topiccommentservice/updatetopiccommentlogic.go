package topiccommentservicelogic

import (
	"context"

	"github.com/feihua/zero-admin/rpc/cms/cmsclient"
	"github.com/feihua/zero-admin/rpc/cms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateTopicCommentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateTopicCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateTopicCommentLogic {
	return &UpdateTopicCommentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新专题评论表
func (l *UpdateTopicCommentLogic) UpdateTopicComment(in *cmsclient.UpdateTopicCommentReq) (*cmsclient.UpdateTopicCommentResp, error) {
	// todo: add your logic here and delete this line

	return &cmsclient.UpdateTopicCommentResp{}, nil
}
