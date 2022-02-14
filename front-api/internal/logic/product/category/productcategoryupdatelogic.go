package logic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-admin/front-api/internal/svc"
)

type ProductCategoryUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProductCategoryUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) ProductCategoryUpdateLogic {
	return ProductCategoryUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ProductCategoryUpdateLogic) ProductCategoryUpdate() error {
	// todo: add your logic here and delete this line

	return nil
}
