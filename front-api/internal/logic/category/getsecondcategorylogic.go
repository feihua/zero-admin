package category

import (
	"context"

	"zero-admin/front-api/internal/svc"
	"zero-admin/front-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSecondCategoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetSecondCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetSecondCategoryLogic {
	return GetSecondCategoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetSecondCategoryLogic) GetSecondCategory(req types.SecondCategoryReq) (resp *types.CategoryResp, err error) {
	// todo: add your logic here and delete this line

	return
}
