package logic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"
	"zero-admin/front-api/internal/svc"
)

type ProductCategoryDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProductCategoryDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) ProductCategoryDeleteLogic {
	return ProductCategoryDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ProductCategoryDeleteLogic) ProductCategoryDelete() error {
	// todo: add your logic here and delete this line

	return nil
}
