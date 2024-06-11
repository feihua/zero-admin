package helpcategoryservicelogic

import (
	"context"

	"github.com/feihua/zero-admin/rpc/cms/cmsclient"
	"github.com/feihua/zero-admin/rpc/cms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteHelpCategoryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteHelpCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteHelpCategoryLogic {
	return &DeleteHelpCategoryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 删除帮助分类表
func (l *DeleteHelpCategoryLogic) DeleteHelpCategory(in *cmsclient.DeleteHelpCategoryReq) (*cmsclient.DeleteHelpCategoryResp, error) {
	// todo: add your logic here and delete this line

	return &cmsclient.DeleteHelpCategoryResp{}, nil
}
