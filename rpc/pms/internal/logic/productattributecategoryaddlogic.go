package logic

import (
	"context"

	"go-zero-admin/rpc/pms/internal/svc"
	"go-zero-admin/rpc/pms/pms"

	"github.com/tal-tech/go-zero/core/logx"
)

type ProductAttributeCategoryAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProductAttributeCategoryAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductAttributeCategoryAddLogic {
	return &ProductAttributeCategoryAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ProductAttributeCategoryAddLogic) ProductAttributeCategoryAdd(in *pms.ProductAttributeCategoryAddReq) (*pms.ProductAttributeCategoryAddResp, error) {
	// todo: add your logic here and delete this line

	return &pms.ProductAttributeCategoryAddResp{}, nil
}
