package logic

import (
	"context"

	"go-zero-admin/rpc/pms/internal/svc"
	"go-zero-admin/rpc/pms/pms"

	"github.com/tal-tech/go-zero/core/logx"
)

type ProductCategoryListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProductCategoryListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductCategoryListLogic {
	return &ProductCategoryListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ProductCategoryListLogic) ProductCategoryList(in *pms.ProductCategoryListReq) (*pms.ProductCategoryListResp, error) {
	// todo: add your logic here and delete this line

	return &pms.ProductCategoryListResp{}, nil
}
