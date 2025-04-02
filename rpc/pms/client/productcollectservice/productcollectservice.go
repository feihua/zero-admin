// Code generated by goctl. DO NOT EDIT.
// goctl 1.8.1
// Source: pms.proto

package productcollectservice

import (
	"context"

	"github.com/feihua/zero-admin/rpc/pms/pmsclient"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	AddBrandReq                                   = pmsclient.AddBrandReq
	AddBrandResp                                  = pmsclient.AddBrandResp
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
	BrandListData                                 = pmsclient.BrandListData
	CommentListData                               = pmsclient.CommentListData
	CommentReplayListData                         = pmsclient.CommentReplayListData
	DeleteBrandReq                                = pmsclient.DeleteBrandReq
	DeleteBrandResp                               = pmsclient.DeleteBrandResp
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
	QueryBrandDetailReq                           = pmsclient.QueryBrandDetailReq
	QueryBrandDetailResp                          = pmsclient.QueryBrandDetailResp
	QueryBrandListByIdsReq                        = pmsclient.QueryBrandListByIdsReq
	QueryBrandListReq                             = pmsclient.QueryBrandListReq
	QueryBrandListResp                            = pmsclient.QueryBrandListResp
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
	UpdateBrandFactoryStatusReq                   = pmsclient.UpdateBrandFactoryStatusReq
	UpdateBrandRecommendStatusReq                 = pmsclient.UpdateBrandRecommendStatusReq
	UpdateBrandReq                                = pmsclient.UpdateBrandReq
	UpdateBrandResp                               = pmsclient.UpdateBrandResp
	UpdateBrandShowStatusReq                      = pmsclient.UpdateBrandShowStatusReq
	UpdateBrandStatusResp                         = pmsclient.UpdateBrandStatusResp
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

	ProductCollectService interface {
		// 添加收藏表
		AddProductCollect(ctx context.Context, in *AddProductCollectReq, opts ...grpc.CallOption) (*AddProductCollectResp, error)
		// 删除收藏表
		DeleteProductCollect(ctx context.Context, in *DeleteProductCollectReq, opts ...grpc.CallOption) (*DeleteProductCollectResp, error)
		// 查询收藏表详情
		QueryProductCollectDetail(ctx context.Context, in *QueryProductCollectDetailReq, opts ...grpc.CallOption) (*QueryProductCollectDetailResp, error)
		// 查询收藏表列表
		QueryProductCollectList(ctx context.Context, in *QueryProductCollectListReq, opts ...grpc.CallOption) (*QueryProductCollectListResp, error)
	}

	defaultProductCollectService struct {
		cli zrpc.Client
	}
)

func NewProductCollectService(cli zrpc.Client) ProductCollectService {
	return &defaultProductCollectService{
		cli: cli,
	}
}

// 添加收藏表
func (m *defaultProductCollectService) AddProductCollect(ctx context.Context, in *AddProductCollectReq, opts ...grpc.CallOption) (*AddProductCollectResp, error) {
	client := pmsclient.NewProductCollectServiceClient(m.cli.Conn())
	return client.AddProductCollect(ctx, in, opts...)
}

// 删除收藏表
func (m *defaultProductCollectService) DeleteProductCollect(ctx context.Context, in *DeleteProductCollectReq, opts ...grpc.CallOption) (*DeleteProductCollectResp, error) {
	client := pmsclient.NewProductCollectServiceClient(m.cli.Conn())
	return client.DeleteProductCollect(ctx, in, opts...)
}

// 查询收藏表详情
func (m *defaultProductCollectService) QueryProductCollectDetail(ctx context.Context, in *QueryProductCollectDetailReq, opts ...grpc.CallOption) (*QueryProductCollectDetailResp, error) {
	client := pmsclient.NewProductCollectServiceClient(m.cli.Conn())
	return client.QueryProductCollectDetail(ctx, in, opts...)
}

// 查询收藏表列表
func (m *defaultProductCollectService) QueryProductCollectList(ctx context.Context, in *QueryProductCollectListReq, opts ...grpc.CallOption) (*QueryProductCollectListResp, error) {
	client := pmsclient.NewProductCollectServiceClient(m.cli.Conn())
	return client.QueryProductCollectList(ctx, in, opts...)
}
