package topicservicelogic

import (
	"context"

	"github.com/feihua/zero-admin/rpc/cms/cmsclient"
	"github.com/feihua/zero-admin/rpc/cms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteTopicLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteTopicLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteTopicLogic {
	return &DeleteTopicLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 删除话题表
func (l *DeleteTopicLogic) DeleteTopic(in *cmsclient.DeleteTopicReq) (*cmsclient.DeleteTopicResp, error) {
	// todo: add your logic here and delete this line

	return &cmsclient.DeleteTopicResp{}, nil
}
