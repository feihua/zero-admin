package topiccategoryservicelogic

import (
	"context"

	"github.com/feihua/zero-admin/rpc/cms/cmsclient"
	"github.com/feihua/zero-admin/rpc/cms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateTopicCategoryStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateTopicCategoryStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateTopicCategoryStatusLogic {
	return &UpdateTopicCategoryStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新话题分类表状态
func (l *UpdateTopicCategoryStatusLogic) UpdateTopicCategoryStatus(in *cmsclient.UpdateTopicCategoryStatusReq) (*cmsclient.UpdateTopicCategoryStatusResp, error) {
	// todo: add your logic here and delete this line

	return &cmsclient.UpdateTopicCategoryStatusResp{}, nil
}
