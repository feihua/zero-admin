// Code generated by goctl. DO NOT EDIT.
// goctl 1.8.2
// Source: pms.proto

package productbrandservice

import (
	"context"

	"github.com/feihua/zero-admin/rpc/pms/pmsclient"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	AddCommentReplayReq                           = pmsclient.AddCommentReplayReq
	AddCommentReplayResp                          = pmsclient.AddCommentReplayResp
	AddCommentReq                                 = pmsclient.AddCommentReq
	AddCommentResp                                = pmsclient.AddCommentResp
	AddFeightTemplateReq                          = pmsclient.AddFeightTemplateReq
	AddFeightTemplateResp                         = pmsclient.AddFeightTemplateResp
	AddProductAttributeCategoryReq                = pmsclient.AddProductAttributeCategoryReq
	AddProductAttributeCategoryResp               = pmsclient.AddProductAttributeCategoryResp
	AddProductAttributeReq                        = pmsclient.AddProductAttributeReq
	AddProductAttributeResp                       = pmsclient.AddProductAttributeResp
	AddProductAttributeValueReq                   = pmsclient.AddProductAttributeValueReq
	AddProductAttributeValueResp                  = pmsclient.AddProductAttributeValueResp
	AddProductBrandReq                            = pmsclient.AddProductBrandReq
	AddProductBrandResp                           = pmsclient.AddProductBrandResp
	AddProductCategoryAttributeRelationReq        = pmsclient.AddProductCategoryAttributeRelationReq
	AddProductCategoryAttributeRelationResp       = pmsclient.AddProductCategoryAttributeRelationResp
	AddProductCategoryReq                         = pmsclient.AddProductCategoryReq
	AddProductCategoryResp                        = pmsclient.AddProductCategoryResp
	AddProductCollectReq                          = pmsclient.AddProductCollectReq
	AddProductCollectResp                         = pmsclient.AddProductCollectResp
	AddProductFullReductionReq                    = pmsclient.AddProductFullReductionReq
	AddProductFullReductionResp                   = pmsclient.AddProductFullReductionResp
	AddProductLadderReq                           = pmsclient.AddProductLadderReq
	AddProductLadderResp                          = pmsclient.AddProductLadderResp
	AddProductOperateLogReq                       = pmsclient.AddProductOperateLogReq
	AddProductOperateLogResp                      = pmsclient.AddProductOperateLogResp
	AddProductReq                                 = pmsclient.AddProductReq
	AddProductResp                                = pmsclient.AddProductResp
	AddProductVertifyRecordReq                    = pmsclient.AddProductVertifyRecordReq
	AddProductVertifyRecordResp                   = pmsclient.AddProductVertifyRecordResp
	AddSkuStockReq                                = pmsclient.AddSkuStockReq
	AddSkuStockResp                               = pmsclient.AddSkuStockResp
	BrandData                                     = pmsclient.BrandData
	CommentListData                               = pmsclient.CommentListData
	CommentReplayListData                         = pmsclient.CommentReplayListData
	DeleteCommentReplayReq                        = pmsclient.DeleteCommentReplayReq
	DeleteCommentReplayResp                       = pmsclient.DeleteCommentReplayResp
	DeleteCommentReq                              = pmsclient.DeleteCommentReq
	DeleteCommentResp                             = pmsclient.DeleteCommentResp
	DeleteFeightTemplateReq                       = pmsclient.DeleteFeightTemplateReq
	DeleteFeightTemplateResp                      = pmsclient.DeleteFeightTemplateResp
	DeleteProductAttributeCategoryReq             = pmsclient.DeleteProductAttributeCategoryReq
	DeleteProductAttributeCategoryResp            = pmsclient.DeleteProductAttributeCategoryResp
	DeleteProductAttributeReq                     = pmsclient.DeleteProductAttributeReq
	DeleteProductAttributeResp                    = pmsclient.DeleteProductAttributeResp
	DeleteProductAttributeValueReq                = pmsclient.DeleteProductAttributeValueReq
	DeleteProductAttributeValueResp               = pmsclient.DeleteProductAttributeValueResp
	DeleteProductBrandReq                         = pmsclient.DeleteProductBrandReq
	DeleteProductBrandResp                        = pmsclient.DeleteProductBrandResp
	DeleteProductCategoryReq                      = pmsclient.DeleteProductCategoryReq
	DeleteProductCategoryResp                     = pmsclient.DeleteProductCategoryResp
	DeleteProductCollectReq                       = pmsclient.DeleteProductCollectReq
	DeleteProductCollectResp                      = pmsclient.DeleteProductCollectResp
	DeleteProductFullReductionReq                 = pmsclient.DeleteProductFullReductionReq
	DeleteProductFullReductionResp                = pmsclient.DeleteProductFullReductionResp
	DeleteProductLadderReq                        = pmsclient.DeleteProductLadderReq
	DeleteProductLadderResp                       = pmsclient.DeleteProductLadderResp
	DeleteProductOperateLogReq                    = pmsclient.DeleteProductOperateLogReq
	DeleteProductOperateLogResp                   = pmsclient.DeleteProductOperateLogResp
	DeleteProductReq                              = pmsclient.DeleteProductReq
	DeleteProductResp                             = pmsclient.DeleteProductResp
	DeleteProductVertifyRecordReq                 = pmsclient.DeleteProductVertifyRecordReq
	DeleteProductVertifyRecordResp                = pmsclient.DeleteProductVertifyRecordResp
	DeleteSkuStockReq                             = pmsclient.DeleteSkuStockReq
	DeleteSkuStockResp                            = pmsclient.DeleteSkuStockResp
	FeightTemplateListData                        = pmsclient.FeightTemplateListData
	LockSkuStockLockData                          = pmsclient.LockSkuStockLockData
	LockSkuStockLockReq                           = pmsclient.LockSkuStockLockReq
	LockSkuStockLockResp                          = pmsclient.LockSkuStockLockResp
	MemberPriceList                               = pmsclient.MemberPriceList
	MemberPriceListData                           = pmsclient.MemberPriceListData
	ProductAttributeCategoryListData              = pmsclient.ProductAttributeCategoryListData
	ProductAttributeData                          = pmsclient.ProductAttributeData
	ProductAttributeDataList                      = pmsclient.ProductAttributeDataList
	ProductAttributeListData                      = pmsclient.ProductAttributeListData
	ProductAttributeValueData                     = pmsclient.ProductAttributeValueData
	ProductAttributeValueList                     = pmsclient.ProductAttributeValueList
	ProductAttributeValueListData                 = pmsclient.ProductAttributeValueListData
	ProductBrandListData                          = pmsclient.ProductBrandListData
	ProductCategoryAttributeRelationListData      = pmsclient.ProductCategoryAttributeRelationListData
	ProductCategoryListData                       = pmsclient.ProductCategoryListData
	ProductCollectListData                        = pmsclient.ProductCollectListData
	ProductFullReductionData                      = pmsclient.ProductFullReductionData
	ProductFullReductionList                      = pmsclient.ProductFullReductionList
	ProductFullReductionListData                  = pmsclient.ProductFullReductionListData
	ProductLadderData                             = pmsclient.ProductLadderData
	ProductLadderList                             = pmsclient.ProductLadderList
	ProductLadderListData                         = pmsclient.ProductLadderListData
	ProductListData                               = pmsclient.ProductListData
	ProductOperateLogListData                     = pmsclient.ProductOperateLogListData
	ProductVertifyRecordListData                  = pmsclient.ProductVertifyRecordListData
	QueryBrandListByIdsReq                        = pmsclient.QueryBrandListByIdsReq
	QueryByproductCategoryIdData                  = pmsclient.QueryByproductCategoryIdData
	QueryByproductCategoryIdReq                   = pmsclient.QueryByproductCategoryIdReq
	QueryByproductCategoryIdResp                  = pmsclient.QueryByproductCategoryIdResp
	QueryCommentDetailReq                         = pmsclient.QueryCommentDetailReq
	QueryCommentDetailResp                        = pmsclient.QueryCommentDetailResp
	QueryCommentListReq                           = pmsclient.QueryCommentListReq
	QueryCommentListResp                          = pmsclient.QueryCommentListResp
	QueryCommentReplayDetailReq                   = pmsclient.QueryCommentReplayDetailReq
	QueryCommentReplayDetailResp                  = pmsclient.QueryCommentReplayDetailResp
	QueryCommentReplayListReq                     = pmsclient.QueryCommentReplayListReq
	QueryCommentReplayListResp                    = pmsclient.QueryCommentReplayListResp
	QueryFeightTemplateDetailReq                  = pmsclient.QueryFeightTemplateDetailReq
	QueryFeightTemplateDetailResp                 = pmsclient.QueryFeightTemplateDetailResp
	QueryFeightTemplateListReq                    = pmsclient.QueryFeightTemplateListReq
	QueryFeightTemplateListResp                   = pmsclient.QueryFeightTemplateListResp
	QueryProductAttributeCategoryListReq          = pmsclient.QueryProductAttributeCategoryListReq
	QueryProductAttributeCategoryListResp         = pmsclient.QueryProductAttributeCategoryListResp
	QueryProductAttributeListReq                  = pmsclient.QueryProductAttributeListReq
	QueryProductAttributeListResp                 = pmsclient.QueryProductAttributeListResp
	QueryProductAttributeValueListReq             = pmsclient.QueryProductAttributeValueListReq
	QueryProductAttributeValueListResp            = pmsclient.QueryProductAttributeValueListResp
	QueryProductBrandDetailReq                    = pmsclient.QueryProductBrandDetailReq
	QueryProductBrandDetailResp                   = pmsclient.QueryProductBrandDetailResp
	QueryProductBrandListReq                      = pmsclient.QueryProductBrandListReq
	QueryProductBrandListResp                     = pmsclient.QueryProductBrandListResp
	QueryProductByIdsReq                          = pmsclient.QueryProductByIdsReq
	QueryProductCategoryAttributeRelationListReq  = pmsclient.QueryProductCategoryAttributeRelationListReq
	QueryProductCategoryAttributeRelationListResp = pmsclient.QueryProductCategoryAttributeRelationListResp
	QueryProductCategoryDetailReq                 = pmsclient.QueryProductCategoryDetailReq
	QueryProductCategoryDetailResp                = pmsclient.QueryProductCategoryDetailResp
	QueryProductCategoryListReq                   = pmsclient.QueryProductCategoryListReq
	QueryProductCategoryListResp                  = pmsclient.QueryProductCategoryListResp
	QueryProductCategoryListTreeData              = pmsclient.QueryProductCategoryListTreeData
	QueryProductCategoryListTreeResp              = pmsclient.QueryProductCategoryListTreeResp
	QueryProductCategoryTreeListReq               = pmsclient.QueryProductCategoryTreeListReq
	QueryProductCollectDetailReq                  = pmsclient.QueryProductCollectDetailReq
	QueryProductCollectDetailResp                 = pmsclient.QueryProductCollectDetailResp
	QueryProductCollectListReq                    = pmsclient.QueryProductCollectListReq
	QueryProductCollectListResp                   = pmsclient.QueryProductCollectListResp
	QueryProductDetailByIdReq                     = pmsclient.QueryProductDetailByIdReq
	QueryProductDetailByIdResp                    = pmsclient.QueryProductDetailByIdResp
	QueryProductFullReductionListReq              = pmsclient.QueryProductFullReductionListReq
	QueryProductFullReductionListResp             = pmsclient.QueryProductFullReductionListResp
	QueryProductLadderDetailReq                   = pmsclient.QueryProductLadderDetailReq
	QueryProductLadderDetailResp                  = pmsclient.QueryProductLadderDetailResp
	QueryProductLadderListReq                     = pmsclient.QueryProductLadderListReq
	QueryProductLadderListResp                    = pmsclient.QueryProductLadderListResp
	QueryProductListReq                           = pmsclient.QueryProductListReq
	QueryProductListResp                          = pmsclient.QueryProductListResp
	QueryProductOperateLogDetailReq               = pmsclient.QueryProductOperateLogDetailReq
	QueryProductOperateLogDetailResp              = pmsclient.QueryProductOperateLogDetailResp
	QueryProductOperateLogListReq                 = pmsclient.QueryProductOperateLogListReq
	QueryProductOperateLogListResp                = pmsclient.QueryProductOperateLogListResp
	QueryProductVertifyRecordDetailReq            = pmsclient.QueryProductVertifyRecordDetailReq
	QueryProductVertifyRecordDetailResp           = pmsclient.QueryProductVertifyRecordDetailResp
	QueryProductVertifyRecordListReq              = pmsclient.QueryProductVertifyRecordListReq
	QueryProductVertifyRecordListResp             = pmsclient.QueryProductVertifyRecordListResp
	QuerySkuStockByProductSkuIdReq                = pmsclient.QuerySkuStockByProductSkuIdReq
	QuerySkuStockDetailReq                        = pmsclient.QuerySkuStockDetailReq
	QuerySkuStockDetailResp                       = pmsclient.QuerySkuStockDetailResp
	QuerySkuStockListReq                          = pmsclient.QuerySkuStockListReq
	QuerySkuStockListResp                         = pmsclient.QuerySkuStockListResp
	ReleaseSkuStockLockData                       = pmsclient.ReleaseSkuStockLockData
	ReleaseSkuStockLockReq                        = pmsclient.ReleaseSkuStockLockReq
	ReleaseSkuStockLockResp                       = pmsclient.ReleaseSkuStockLockResp
	SkuStockData                                  = pmsclient.SkuStockData
	SkuStockList                                  = pmsclient.SkuStockList
	SkuStockListData                              = pmsclient.SkuStockListData
	UpdateCommentReplayReq                        = pmsclient.UpdateCommentReplayReq
	UpdateCommentReplayResp                       = pmsclient.UpdateCommentReplayResp
	UpdateCommentReq                              = pmsclient.UpdateCommentReq
	UpdateCommentResp                             = pmsclient.UpdateCommentResp
	UpdateFeightTemplateReq                       = pmsclient.UpdateFeightTemplateReq
	UpdateFeightTemplateResp                      = pmsclient.UpdateFeightTemplateResp
	UpdateProductAttributeCategoryReq             = pmsclient.UpdateProductAttributeCategoryReq
	UpdateProductAttributeCategoryResp            = pmsclient.UpdateProductAttributeCategoryResp
	UpdateProductAttributeValueReq                = pmsclient.UpdateProductAttributeValueReq
	UpdateProductAttributeValueResp               = pmsclient.UpdateProductAttributeValueResp
	UpdateProductBrandReq                         = pmsclient.UpdateProductBrandReq
	UpdateProductBrandResp                        = pmsclient.UpdateProductBrandResp
	UpdateProductBrandStatusReq                   = pmsclient.UpdateProductBrandStatusReq
	UpdateProductBrandStatusResp                  = pmsclient.UpdateProductBrandStatusResp
	UpdateProductCategoryReq                      = pmsclient.UpdateProductCategoryReq
	UpdateProductCategoryResp                     = pmsclient.UpdateProductCategoryResp
	UpdateProductCategoryStatusReq                = pmsclient.UpdateProductCategoryStatusReq
	UpdateProductCategoryStatusResp               = pmsclient.UpdateProductCategoryStatusResp
	UpdateProductLadderReq                        = pmsclient.UpdateProductLadderReq
	UpdateProductLadderResp                       = pmsclient.UpdateProductLadderResp
	UpdateProductReq                              = pmsclient.UpdateProductReq
	UpdateProductResp                             = pmsclient.UpdateProductResp
	UpdateProductStatusReq                        = pmsclient.UpdateProductStatusReq
	UpdateProductStatusResp                       = pmsclient.UpdateProductStatusResp
	UpdateProductVertifyRecordReq                 = pmsclient.UpdateProductVertifyRecordReq
	UpdateProductVertifyRecordResp                = pmsclient.UpdateProductVertifyRecordResp
	UpdateProductVertifyRecordStatusReq           = pmsclient.UpdateProductVertifyRecordStatusReq
	UpdateProductVertifyRecordStatusResp          = pmsclient.UpdateProductVertifyRecordStatusResp
	UpdateSkuStockData                            = pmsclient.UpdateSkuStockData
	UpdateSkuStockReq                             = pmsclient.UpdateSkuStockReq
	UpdateSkuStockResp                            = pmsclient.UpdateSkuStockResp

	ProductBrandService interface {
		// 添加商品品牌
		AddProductBrand(ctx context.Context, in *AddProductBrandReq, opts ...grpc.CallOption) (*AddProductBrandResp, error)
		// 删除商品品牌
		DeleteProductBrand(ctx context.Context, in *DeleteProductBrandReq, opts ...grpc.CallOption) (*DeleteProductBrandResp, error)
		// 更新商品品牌
		UpdateProductBrand(ctx context.Context, in *UpdateProductBrandReq, opts ...grpc.CallOption) (*UpdateProductBrandResp, error)
		// 更新商品品牌状态
		UpdateProductBrandStatus(ctx context.Context, in *UpdateProductBrandStatusReq, opts ...grpc.CallOption) (*UpdateProductBrandStatusResp, error)
		// 查询商品品牌详情
		QueryProductBrandDetail(ctx context.Context, in *QueryProductBrandDetailReq, opts ...grpc.CallOption) (*QueryProductBrandDetailResp, error)
		// 查询商品品牌列表
		QueryProductBrandList(ctx context.Context, in *QueryProductBrandListReq, opts ...grpc.CallOption) (*QueryProductBrandListResp, error)
		QueryBrandListByIds(ctx context.Context, in *QueryBrandListByIdsReq, opts ...grpc.CallOption) (*QueryProductBrandListResp, error)
		// 更新品牌的推荐状态
		UpdateBrandRecommendStatus(ctx context.Context, in *UpdateProductBrandStatusReq, opts ...grpc.CallOption) (*UpdateProductBrandStatusResp, error)
	}

	defaultProductBrandService struct {
		cli zrpc.Client
	}
)

func NewProductBrandService(cli zrpc.Client) ProductBrandService {
	return &defaultProductBrandService{
		cli: cli,
	}
}

// 添加商品品牌
func (m *defaultProductBrandService) AddProductBrand(ctx context.Context, in *AddProductBrandReq, opts ...grpc.CallOption) (*AddProductBrandResp, error) {
	client := pmsclient.NewProductBrandServiceClient(m.cli.Conn())
	return client.AddProductBrand(ctx, in, opts...)
}

// 删除商品品牌
func (m *defaultProductBrandService) DeleteProductBrand(ctx context.Context, in *DeleteProductBrandReq, opts ...grpc.CallOption) (*DeleteProductBrandResp, error) {
	client := pmsclient.NewProductBrandServiceClient(m.cli.Conn())
	return client.DeleteProductBrand(ctx, in, opts...)
}

// 更新商品品牌
func (m *defaultProductBrandService) UpdateProductBrand(ctx context.Context, in *UpdateProductBrandReq, opts ...grpc.CallOption) (*UpdateProductBrandResp, error) {
	client := pmsclient.NewProductBrandServiceClient(m.cli.Conn())
	return client.UpdateProductBrand(ctx, in, opts...)
}

// 更新商品品牌状态
func (m *defaultProductBrandService) UpdateProductBrandStatus(ctx context.Context, in *UpdateProductBrandStatusReq, opts ...grpc.CallOption) (*UpdateProductBrandStatusResp, error) {
	client := pmsclient.NewProductBrandServiceClient(m.cli.Conn())
	return client.UpdateProductBrandStatus(ctx, in, opts...)
}

// 查询商品品牌详情
func (m *defaultProductBrandService) QueryProductBrandDetail(ctx context.Context, in *QueryProductBrandDetailReq, opts ...grpc.CallOption) (*QueryProductBrandDetailResp, error) {
	client := pmsclient.NewProductBrandServiceClient(m.cli.Conn())
	return client.QueryProductBrandDetail(ctx, in, opts...)
}

// 查询商品品牌列表
func (m *defaultProductBrandService) QueryProductBrandList(ctx context.Context, in *QueryProductBrandListReq, opts ...grpc.CallOption) (*QueryProductBrandListResp, error) {
	client := pmsclient.NewProductBrandServiceClient(m.cli.Conn())
	return client.QueryProductBrandList(ctx, in, opts...)
}

func (m *defaultProductBrandService) QueryBrandListByIds(ctx context.Context, in *QueryBrandListByIdsReq, opts ...grpc.CallOption) (*QueryProductBrandListResp, error) {
	client := pmsclient.NewProductBrandServiceClient(m.cli.Conn())
	return client.QueryBrandListByIds(ctx, in, opts...)
}

// 更新品牌的推荐状态
func (m *defaultProductBrandService) UpdateBrandRecommendStatus(ctx context.Context, in *UpdateProductBrandStatusReq, opts ...grpc.CallOption) (*UpdateProductBrandStatusResp, error) {
	client := pmsclient.NewProductBrandServiceClient(m.cli.Conn())
	return client.UpdateBrandRecommendStatus(ctx, in, opts...)
}
