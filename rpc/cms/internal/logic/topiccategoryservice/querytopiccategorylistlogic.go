package topiccategoryservicelogic

import (
	"context"

	"github.com/feihua/zero-admin/rpc/cms/cmsclient"
	"github.com/feihua/zero-admin/rpc/cms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryTopicCategoryListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryTopicCategoryListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryTopicCategoryListLogic {
	return &QueryTopicCategoryListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询话题分类表列表
func (l *QueryTopicCategoryListLogic) QueryTopicCategoryList(in *cmsclient.QueryTopicCategoryListReq) (*cmsclient.QueryTopicCategoryListResp, error) {
	// todo: add your logic here and delete this line

	return &cmsclient.QueryTopicCategoryListResp{}, nil
}
