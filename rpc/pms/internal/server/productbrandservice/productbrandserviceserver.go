// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.6
// Source: pms.proto

package server

import (
	"context"

	"github.com/feihua/zero-admin/rpc/pms/internal/logic/productbrandservice"
	"github.com/feihua/zero-admin/rpc/pms/internal/svc"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"
)

type ProductBrandServiceServer struct {
	svcCtx *svc.ServiceContext
	pmsclient.UnimplementedProductBrandServiceServer
}

func NewProductBrandServiceServer(svcCtx *svc.ServiceContext) *ProductBrandServiceServer {
	return &ProductBrandServiceServer{
		svcCtx: svcCtx,
	}
}

// 添加商品品牌
func (s *ProductBrandServiceServer) AddProductBrand(ctx context.Context, in *pmsclient.AddProductBrandReq) (*pmsclient.AddProductBrandResp, error) {
	l := productbrandservicelogic.NewAddProductBrandLogic(ctx, s.svcCtx)
	return l.AddProductBrand(in)
}

// 删除商品品牌
func (s *ProductBrandServiceServer) DeleteProductBrand(ctx context.Context, in *pmsclient.DeleteProductBrandReq) (*pmsclient.DeleteProductBrandResp, error) {
	l := productbrandservicelogic.NewDeleteProductBrandLogic(ctx, s.svcCtx)
	return l.DeleteProductBrand(in)
}

// 更新商品品牌
func (s *ProductBrandServiceServer) UpdateProductBrand(ctx context.Context, in *pmsclient.UpdateProductBrandReq) (*pmsclient.UpdateProductBrandResp, error) {
	l := productbrandservicelogic.NewUpdateProductBrandLogic(ctx, s.svcCtx)
	return l.UpdateProductBrand(in)
}

// 更新商品品牌状态
func (s *ProductBrandServiceServer) UpdateProductBrandStatus(ctx context.Context, in *pmsclient.UpdateProductBrandStatusReq) (*pmsclient.UpdateProductBrandStatusResp, error) {
	l := productbrandservicelogic.NewUpdateProductBrandStatusLogic(ctx, s.svcCtx)
	return l.UpdateProductBrandStatus(in)
}

// 查询商品品牌详情
func (s *ProductBrandServiceServer) QueryProductBrandDetail(ctx context.Context, in *pmsclient.QueryProductBrandDetailReq) (*pmsclient.QueryProductBrandDetailResp, error) {
	l := productbrandservicelogic.NewQueryProductBrandDetailLogic(ctx, s.svcCtx)
	return l.QueryProductBrandDetail(in)
}

// 查询商品品牌列表
func (s *ProductBrandServiceServer) QueryProductBrandList(ctx context.Context, in *pmsclient.QueryProductBrandListReq) (*pmsclient.QueryProductBrandListResp, error) {
	l := productbrandservicelogic.NewQueryProductBrandListLogic(ctx, s.svcCtx)
	return l.QueryProductBrandList(in)
}

func (s *ProductBrandServiceServer) QueryBrandListByIds(ctx context.Context, in *pmsclient.QueryBrandListByIdsReq) (*pmsclient.QueryProductBrandListResp, error) {
	l := productbrandservicelogic.NewQueryBrandListByIdsLogic(ctx, s.svcCtx)
	return l.QueryBrandListByIds(in)
}

// 更新品牌的推荐状态
func (s *ProductBrandServiceServer) UpdateBrandRecommendStatus(ctx context.Context, in *pmsclient.UpdateProductBrandStatusReq) (*pmsclient.UpdateProductBrandStatusResp, error) {
	l := productbrandservicelogic.NewUpdateBrandRecommendStatusLogic(ctx, s.svcCtx)
	return l.UpdateBrandRecommendStatus(in)
}

// 更新品牌的排序
func (s *ProductBrandServiceServer) UpdateBrandSort(ctx context.Context, in *pmsclient.UpdateProductBrandSortReq) (*pmsclient.UpdateProductBrandStatusResp, error) {
	l := productbrandservicelogic.NewUpdateBrandSortLogic(ctx, s.svcCtx)
	return l.UpdateBrandSort(in)
}
