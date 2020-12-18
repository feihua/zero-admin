package logic

import (
	"context"

	"github.com/tal-tech/go-zero/core/logx"
	"go-zero-admin/front-api/internal/svc"
)

type ProductCategoryAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProductCategoryAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) ProductCategoryAddLogic {
	return ProductCategoryAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ProductCategoryAddLogic) ProductCategoryAdd() error {
	// todo: add your logic here and delete this line

	return nil
}
