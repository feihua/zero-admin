// Code generated by goctl. DO NOT EDIT.
// goctl 1.8.1
// Source: pms.proto

package server

import (
	"context"

	"github.com/feihua/zero-admin/rpc/pms/internal/logic/productattributecategoryservice"
	"github.com/feihua/zero-admin/rpc/pms/internal/svc"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"
)

type ProductAttributeCategoryServiceServer struct {
	svcCtx *svc.ServiceContext
	pmsclient.UnimplementedProductAttributeCategoryServiceServer
}

func NewProductAttributeCategoryServiceServer(svcCtx *svc.ServiceContext) *ProductAttributeCategoryServiceServer {
	return &ProductAttributeCategoryServiceServer{
		svcCtx: svcCtx,
	}
}

// 添加产品属性分类表
func (s *ProductAttributeCategoryServiceServer) AddProductAttributeCategory(ctx context.Context, in *pmsclient.AddProductAttributeCategoryReq) (*pmsclient.AddProductAttributeCategoryResp, error) {
	l := productattributecategoryservicelogic.NewAddProductAttributeCategoryLogic(ctx, s.svcCtx)
	return l.AddProductAttributeCategory(in)
}

// 删除产品属性分类表
func (s *ProductAttributeCategoryServiceServer) DeleteProductAttributeCategory(ctx context.Context, in *pmsclient.DeleteProductAttributeCategoryReq) (*pmsclient.DeleteProductAttributeCategoryResp, error) {
	l := productattributecategoryservicelogic.NewDeleteProductAttributeCategoryLogic(ctx, s.svcCtx)
	return l.DeleteProductAttributeCategory(in)
}

// 更新产品属性分类表
func (s *ProductAttributeCategoryServiceServer) UpdateProductAttributeCategory(ctx context.Context, in *pmsclient.UpdateProductAttributeCategoryReq) (*pmsclient.UpdateProductAttributeCategoryResp, error) {
	l := productattributecategoryservicelogic.NewUpdateProductAttributeCategoryLogic(ctx, s.svcCtx)
	return l.UpdateProductAttributeCategory(in)
}

// 查询产品属性分类表列表
func (s *ProductAttributeCategoryServiceServer) QueryProductAttributeCategoryList(ctx context.Context, in *pmsclient.QueryProductAttributeCategoryListReq) (*pmsclient.QueryProductAttributeCategoryListResp, error) {
	l := productattributecategoryservicelogic.NewQueryProductAttributeCategoryListLogic(ctx, s.svcCtx)
	return l.QueryProductAttributeCategoryList(in)
}

// 获取所有商品属性分类及其下属性
func (s *ProductAttributeCategoryServiceServer) QueryCategoryWithAttrList(ctx context.Context, in *pmsclient.QueryProductAttributeCategoryListReq) (*pmsclient.QueryProductAttributeCategoryListResp, error) {
	l := productattributecategoryservicelogic.NewQueryCategoryWithAttrListLogic(ctx, s.svcCtx)
	return l.QueryCategoryWithAttrList(in)
}
