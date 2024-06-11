package topiccategoryservicelogic

import (
	"context"

	"github.com/feihua/zero-admin/rpc/cms/cmsclient"
	"github.com/feihua/zero-admin/rpc/cms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteTopicCategoryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteTopicCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteTopicCategoryLogic {
	return &DeleteTopicCategoryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 删除话题分类表
func (l *DeleteTopicCategoryLogic) DeleteTopicCategory(in *cmsclient.DeleteTopicCategoryReq) (*cmsclient.DeleteTopicCategoryResp, error) {
	// todo: add your logic here and delete this line

	return &cmsclient.DeleteTopicCategoryResp{}, nil
}
