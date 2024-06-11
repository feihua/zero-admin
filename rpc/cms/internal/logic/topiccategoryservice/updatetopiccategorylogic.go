package topiccategoryservicelogic

import (
	"context"

	"github.com/feihua/zero-admin/rpc/cms/cmsclient"
	"github.com/feihua/zero-admin/rpc/cms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateTopicCategoryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateTopicCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateTopicCategoryLogic {
	return &UpdateTopicCategoryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新话题分类表
func (l *UpdateTopicCategoryLogic) UpdateTopicCategory(in *cmsclient.UpdateTopicCategoryReq) (*cmsclient.UpdateTopicCategoryResp, error) {
	// todo: add your logic here and delete this line

	return &cmsclient.UpdateTopicCategoryResp{}, nil
}
