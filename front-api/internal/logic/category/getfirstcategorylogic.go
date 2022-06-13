package category

import (
	"context"

	"zero-admin/front-api/internal/svc"
	"zero-admin/front-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFirstCategoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetFirstCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetFirstCategoryLogic {
	return GetFirstCategoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetFirstCategoryLogic) GetFirstCategory() (resp *types.CategoryResp, err error) {
	// todo: add your logic here and delete this line

	return
}
