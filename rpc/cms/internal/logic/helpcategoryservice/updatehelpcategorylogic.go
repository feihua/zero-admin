package helpcategoryservicelogic

import (
	"context"

	"github.com/feihua/zero-admin/rpc/cms/cmsclient"
	"github.com/feihua/zero-admin/rpc/cms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateHelpCategoryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateHelpCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateHelpCategoryLogic {
	return &UpdateHelpCategoryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新帮助分类表
func (l *UpdateHelpCategoryLogic) UpdateHelpCategory(in *cmsclient.UpdateHelpCategoryReq) (*cmsclient.UpdateHelpCategoryResp, error) {
	// todo: add your logic here and delete this line

	return &cmsclient.UpdateHelpCategoryResp{}, nil
}
