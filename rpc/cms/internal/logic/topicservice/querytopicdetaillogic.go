package topicservicelogic

import (
	"context"

	"github.com/feihua/zero-admin/rpc/cms/cmsclient"
	"github.com/feihua/zero-admin/rpc/cms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryTopicDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryTopicDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryTopicDetailLogic {
	return &QueryTopicDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询话题表详情
func (l *QueryTopicDetailLogic) QueryTopicDetail(in *cmsclient.QueryTopicDetailReq) (*cmsclient.QueryTopicDetailResp, error) {
	// todo: add your logic here and delete this line

	return &cmsclient.QueryTopicDetailResp{}, nil
}
