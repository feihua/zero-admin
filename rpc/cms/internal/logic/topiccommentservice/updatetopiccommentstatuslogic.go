package topiccommentservicelogic

import (
	"context"

	"github.com/feihua/zero-admin/rpc/cms/cmsclient"
	"github.com/feihua/zero-admin/rpc/cms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateTopicCommentStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateTopicCommentStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateTopicCommentStatusLogic {
	return &UpdateTopicCommentStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新专题评论表状态
func (l *UpdateTopicCommentStatusLogic) UpdateTopicCommentStatus(in *cmsclient.UpdateTopicCommentStatusReq) (*cmsclient.UpdateTopicCommentStatusResp, error) {
	// todo: add your logic here and delete this line

	return &cmsclient.UpdateTopicCommentStatusResp{}, nil
}
