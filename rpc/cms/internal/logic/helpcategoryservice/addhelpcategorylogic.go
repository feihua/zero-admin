package helpcategoryservicelogic

import (
	"context"

	"github.com/feihua/zero-admin/rpc/cms/cmsclient"
	"github.com/feihua/zero-admin/rpc/cms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddHelpCategoryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddHelpCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddHelpCategoryLogic {
	return &AddHelpCategoryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 添加帮助分类表
func (l *AddHelpCategoryLogic) AddHelpCategory(in *cmsclient.AddHelpCategoryReq) (*cmsclient.AddHelpCategoryResp, error) {
	// todo: add your logic here and delete this line

	return &cmsclient.AddHelpCategoryResp{}, nil
}
