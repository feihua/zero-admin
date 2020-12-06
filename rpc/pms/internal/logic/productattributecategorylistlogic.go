package logic

import (
	"context"

	"go-zero-admin/rpc/pms/internal/svc"
	"go-zero-admin/rpc/pms/pms"

	"github.com/tal-tech/go-zero/core/logx"
)

type ProductAttributeCategoryListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProductAttributeCategoryListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductAttributeCategoryListLogic {
	return &ProductAttributeCategoryListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ProductAttributeCategoryListLogic) ProductAttributeCategoryList(in *pms.ProductAttributeCategoryListReq) (*pms.ProductAttributeCategoryListResp, error) {
	// todo: add your logic here and delete this line

	return &pms.ProductAttributeCategoryListResp{}, nil
}
