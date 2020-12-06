package logic

import (
	"context"

	"go-zero-admin/rpc/pms/internal/svc"
	"go-zero-admin/rpc/pms/pms"

	"github.com/tal-tech/go-zero/core/logx"
)

type ProductCategoryAttributeRelationListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProductCategoryAttributeRelationListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductCategoryAttributeRelationListLogic {
	return &ProductCategoryAttributeRelationListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ProductCategoryAttributeRelationListLogic) ProductCategoryAttributeRelationList(in *pms.ProductCategoryAttributeRelationListReq) (*pms.ProductCategoryAttributeRelationListResp, error) {
	// todo: add your logic here and delete this line

	return &pms.ProductCategoryAttributeRelationListResp{}, nil
}
