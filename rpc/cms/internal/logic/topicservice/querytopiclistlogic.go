package topicservicelogic

import (
	"context"

	"github.com/feihua/zero-admin/rpc/cms/cmsclient"
	"github.com/feihua/zero-admin/rpc/cms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryTopicListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryTopicListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryTopicListLogic {
	return &QueryTopicListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询话题表列表
func (l *QueryTopicListLogic) QueryTopicList(in *cmsclient.QueryTopicListReq) (*cmsclient.QueryTopicListResp, error) {
	// todo: add your logic here and delete this line

	return &cmsclient.QueryTopicListResp{}, nil
}
