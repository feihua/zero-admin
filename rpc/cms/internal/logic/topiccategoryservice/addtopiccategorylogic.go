package topiccategoryservicelogic

import (
	"context"

	"github.com/feihua/zero-admin/rpc/cms/cmsclient"
	"github.com/feihua/zero-admin/rpc/cms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddTopicCategoryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddTopicCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddTopicCategoryLogic {
	return &AddTopicCategoryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 添加话题分类表
func (l *AddTopicCategoryLogic) AddTopicCategory(in *cmsclient.AddTopicCategoryReq) (*cmsclient.AddTopicCategoryResp, error) {
	// todo: add your logic here and delete this line

	return &cmsclient.AddTopicCategoryResp{}, nil
}
