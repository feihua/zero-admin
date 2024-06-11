package topiccategoryservicelogic

import (
	"context"

	"github.com/feihua/zero-admin/rpc/cms/cmsclient"
	"github.com/feihua/zero-admin/rpc/cms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryTopicCategoryDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryTopicCategoryDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryTopicCategoryDetailLogic {
	return &QueryTopicCategoryDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询话题分类表详情
func (l *QueryTopicCategoryDetailLogic) QueryTopicCategoryDetail(in *cmsclient.QueryTopicCategoryDetailReq) (*cmsclient.QueryTopicCategoryDetailResp, error) {
	// todo: add your logic here and delete this line

	return &cmsclient.QueryTopicCategoryDetailResp{}, nil
}
