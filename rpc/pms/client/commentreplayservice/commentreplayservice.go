// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.6
// Source: pms.proto

package commentreplayservice

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
	AddProductAttributeGroupReq                   = pmsclient.AddProductAttributeGroupReq
	AddProductAttributeGroupResp                  = pmsclient.AddProductAttributeGroupResp
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
	AddProductSkuReq                              = pmsclient.AddProductSkuReq
	AddProductSkuResp                             = pmsclient.AddProductSkuResp
	AddProductSpecReq                             = pmsclient.AddProductSpecReq
	AddProductSpecResp                            = pmsclient.AddProductSpecResp
	AddProductSpecValueReq                        = pmsclient.AddProductSpecValueReq
	AddProductSpecValueResp                       = pmsclient.AddProductSpecValueResp
	AddProductVertifyRecordReq                    = pmsclient.AddProductVertifyRecordReq
	AddProductVertifyRecordResp                   = pmsclient.AddProductVertifyRecordResp
	BrandData                                     = pmsclient.BrandData
	CommentListData                               = pmsclient.CommentListData
	CommentReplayListData                         = pmsclient.CommentReplayListData
	DeleteCommentReplayReq                        = pmsclient.DeleteCommentReplayReq
	DeleteCommentReplayResp                       = pmsclient.DeleteCommentReplayResp
	DeleteCommentReq                              = pmsclient.DeleteCommentReq
	DeleteCommentResp                             = pmsclient.DeleteCommentResp
	DeleteFeightTemplateReq                       = pmsclient.DeleteFeightTemplateReq
	DeleteFeightTemplateResp                      = pmsclient.DeleteFeightTemplateResp
	DeleteProductAttributeGroupReq                = pmsclient.DeleteProductAttributeGroupReq
	DeleteProductAttributeGroupResp               = pmsclient.DeleteProductAttributeGroupResp
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
	DeleteProductSkuReq                           = pmsclient.DeleteProductSkuReq
	DeleteProductSkuResp                          = pmsclient.DeleteProductSkuResp
	DeleteProductSpecReq                          = pmsclient.DeleteProductSpecReq
	DeleteProductSpecResp                         = pmsclient.DeleteProductSpecResp
	DeleteProductSpecValueReq                     = pmsclient.DeleteProductSpecValueReq
	DeleteProductSpecValueResp                    = pmsclient.DeleteProductSpecValueResp
	DeleteProductSpuReq                           = pmsclient.DeleteProductSpuReq
	DeleteProductSpuResp                          = pmsclient.DeleteProductSpuResp
	FeightTemplateListData                        = pmsclient.FeightTemplateListData
	MemberPriceList                               = pmsclient.MemberPriceList
	MemberPriceListData                           = pmsclient.MemberPriceListData
	ProductAttributeDataList                      = pmsclient.ProductAttributeDataList
	ProductAttributeGroupListData                 = pmsclient.ProductAttributeGroupListData
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
	ProductOperateLogListData                     = pmsclient.ProductOperateLogListData
	ProductSkuListData                            = pmsclient.ProductSkuListData
	ProductSpecListData                           = pmsclient.ProductSpecListData
	ProductSpecValueListData                      = pmsclient.ProductSpecValueListData
	ProductSpuListData                            = pmsclient.ProductSpuListData
	ProductSpuReq                                 = pmsclient.ProductSpuReq
	ProductSpuResp                                = pmsclient.ProductSpuResp
	ProductVertifyRecordListData                  = pmsclient.ProductVertifyRecordListData
	QueryBrandListByIdsReq                        = pmsclient.QueryBrandListByIdsReq
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
	QueryProductAttributeDetailReq                = pmsclient.QueryProductAttributeDetailReq
	QueryProductAttributeDetailResp               = pmsclient.QueryProductAttributeDetailResp
	QueryProductAttributeGroupDetailReq           = pmsclient.QueryProductAttributeGroupDetailReq
	QueryProductAttributeGroupDetailResp          = pmsclient.QueryProductAttributeGroupDetailResp
	QueryProductAttributeGroupListReq             = pmsclient.QueryProductAttributeGroupListReq
	QueryProductAttributeGroupListResp            = pmsclient.QueryProductAttributeGroupListResp
	QueryProductAttributeListReq                  = pmsclient.QueryProductAttributeListReq
	QueryProductAttributeListResp                 = pmsclient.QueryProductAttributeListResp
	QueryProductAttributeValueDetailReq           = pmsclient.QueryProductAttributeValueDetailReq
	QueryProductAttributeValueDetailResp          = pmsclient.QueryProductAttributeValueDetailResp
	QueryProductAttributeValueListReq             = pmsclient.QueryProductAttributeValueListReq
	QueryProductAttributeValueListResp            = pmsclient.QueryProductAttributeValueListResp
	QueryProductBrandDetailReq                    = pmsclient.QueryProductBrandDetailReq
	QueryProductBrandDetailResp                   = pmsclient.QueryProductBrandDetailResp
	QueryProductBrandListReq                      = pmsclient.QueryProductBrandListReq
	QueryProductBrandListResp                     = pmsclient.QueryProductBrandListResp
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
	QueryProductFullReductionListReq              = pmsclient.QueryProductFullReductionListReq
	QueryProductFullReductionListResp             = pmsclient.QueryProductFullReductionListResp
	QueryProductLadderDetailReq                   = pmsclient.QueryProductLadderDetailReq
	QueryProductLadderDetailResp                  = pmsclient.QueryProductLadderDetailResp
	QueryProductLadderListReq                     = pmsclient.QueryProductLadderListReq
	QueryProductLadderListResp                    = pmsclient.QueryProductLadderListResp
	QueryProductOperateLogDetailReq               = pmsclient.QueryProductOperateLogDetailReq
	QueryProductOperateLogDetailResp              = pmsclient.QueryProductOperateLogDetailResp
	QueryProductOperateLogListReq                 = pmsclient.QueryProductOperateLogListReq
	QueryProductOperateLogListResp                = pmsclient.QueryProductOperateLogListResp
	QueryProductSkuDetailReq                      = pmsclient.QueryProductSkuDetailReq
	QueryProductSkuDetailResp                     = pmsclient.QueryProductSkuDetailResp
	QueryProductSkuListReq                        = pmsclient.QueryProductSkuListReq
	QueryProductSkuListResp                       = pmsclient.QueryProductSkuListResp
	QueryProductSpecDetailReq                     = pmsclient.QueryProductSpecDetailReq
	QueryProductSpecDetailResp                    = pmsclient.QueryProductSpecDetailResp
	QueryProductSpecListReq                       = pmsclient.QueryProductSpecListReq
	QueryProductSpecListResp                      = pmsclient.QueryProductSpecListResp
	QueryProductSpecValueDetailReq                = pmsclient.QueryProductSpecValueDetailReq
	QueryProductSpecValueDetailResp               = pmsclient.QueryProductSpecValueDetailResp
	QueryProductSpecValueListReq                  = pmsclient.QueryProductSpecValueListReq
	QueryProductSpecValueListResp                 = pmsclient.QueryProductSpecValueListResp
	QueryProductSpuByIdsReq                       = pmsclient.QueryProductSpuByIdsReq
	QueryProductSpuDetailReq                      = pmsclient.QueryProductSpuDetailReq
	QueryProductSpuDetailResp                     = pmsclient.QueryProductSpuDetailResp
	QueryProductSpuListReq                        = pmsclient.QueryProductSpuListReq
	QueryProductSpuListResp                       = pmsclient.QueryProductSpuListResp
	QueryProductVertifyRecordDetailReq            = pmsclient.QueryProductVertifyRecordDetailReq
	QueryProductVertifyRecordDetailResp           = pmsclient.QueryProductVertifyRecordDetailResp
	QueryProductVertifyRecordListReq              = pmsclient.QueryProductVertifyRecordListReq
	QueryProductVertifyRecordListResp             = pmsclient.QueryProductVertifyRecordListResp
	SkuStockData                                  = pmsclient.SkuStockData
	SkuStockList                                  = pmsclient.SkuStockList
	UpdateCommentReplayReq                        = pmsclient.UpdateCommentReplayReq
	UpdateCommentReplayResp                       = pmsclient.UpdateCommentReplayResp
	UpdateCommentReq                              = pmsclient.UpdateCommentReq
	UpdateCommentResp                             = pmsclient.UpdateCommentResp
	UpdateFeightTemplateReq                       = pmsclient.UpdateFeightTemplateReq
	UpdateFeightTemplateResp                      = pmsclient.UpdateFeightTemplateResp
	UpdateProductAttributeGroupReq                = pmsclient.UpdateProductAttributeGroupReq
	UpdateProductAttributeGroupResp               = pmsclient.UpdateProductAttributeGroupResp
	UpdateProductAttributeGroupStatusReq          = pmsclient.UpdateProductAttributeGroupStatusReq
	UpdateProductAttributeGroupStatusResp         = pmsclient.UpdateProductAttributeGroupStatusResp
	UpdateProductAttributeReq                     = pmsclient.UpdateProductAttributeReq
	UpdateProductAttributeResp                    = pmsclient.UpdateProductAttributeResp
	UpdateProductAttributeStatusReq               = pmsclient.UpdateProductAttributeStatusReq
	UpdateProductAttributeStatusResp              = pmsclient.UpdateProductAttributeStatusResp
	UpdateProductAttributeValueReq                = pmsclient.UpdateProductAttributeValueReq
	UpdateProductAttributeValueResp               = pmsclient.UpdateProductAttributeValueResp
	UpdateProductAttributeValueStatusReq          = pmsclient.UpdateProductAttributeValueStatusReq
	UpdateProductAttributeValueStatusResp         = pmsclient.UpdateProductAttributeValueStatusResp
	UpdateProductBrandReq                         = pmsclient.UpdateProductBrandReq
	UpdateProductBrandResp                        = pmsclient.UpdateProductBrandResp
	UpdateProductBrandSortReq                     = pmsclient.UpdateProductBrandSortReq
	UpdateProductBrandStatusReq                   = pmsclient.UpdateProductBrandStatusReq
	UpdateProductBrandStatusResp                  = pmsclient.UpdateProductBrandStatusResp
	UpdateProductCategoryReq                      = pmsclient.UpdateProductCategoryReq
	UpdateProductCategoryResp                     = pmsclient.UpdateProductCategoryResp
	UpdateProductCategoryStatusReq                = pmsclient.UpdateProductCategoryStatusReq
	UpdateProductCategoryStatusResp               = pmsclient.UpdateProductCategoryStatusResp
	UpdateProductLadderReq                        = pmsclient.UpdateProductLadderReq
	UpdateProductLadderResp                       = pmsclient.UpdateProductLadderResp
	UpdateProductSkuData                          = pmsclient.UpdateProductSkuData
	UpdateProductSkuReq                           = pmsclient.UpdateProductSkuReq
	UpdateProductSkuResp                          = pmsclient.UpdateProductSkuResp
	UpdateProductSortReq                          = pmsclient.UpdateProductSortReq
	UpdateProductSpecReq                          = pmsclient.UpdateProductSpecReq
	UpdateProductSpecResp                         = pmsclient.UpdateProductSpecResp
	UpdateProductSpecStatusReq                    = pmsclient.UpdateProductSpecStatusReq
	UpdateProductSpecStatusResp                   = pmsclient.UpdateProductSpecStatusResp
	UpdateProductSpecValueReq                     = pmsclient.UpdateProductSpecValueReq
	UpdateProductSpecValueResp                    = pmsclient.UpdateProductSpecValueResp
	UpdateProductSpecValueStatusReq               = pmsclient.UpdateProductSpecValueStatusReq
	UpdateProductSpecValueStatusResp              = pmsclient.UpdateProductSpecValueStatusResp
	UpdateProductSpuStatusReq                     = pmsclient.UpdateProductSpuStatusReq
	UpdateProductSpuStatusResp                    = pmsclient.UpdateProductSpuStatusResp
	UpdateSkuStockData                            = pmsclient.UpdateSkuStockData
	UpdateSkuStockLockResp                        = pmsclient.UpdateSkuStockLockResp
	UpdateSkuStockReq                             = pmsclient.UpdateSkuStockReq

	CommentReplayService interface {
		// 添加产品评价回复
		AddCommentReplay(ctx context.Context, in *AddCommentReplayReq, opts ...grpc.CallOption) (*AddCommentReplayResp, error)
		// 删除产品评价回复
		DeleteCommentReplay(ctx context.Context, in *DeleteCommentReplayReq, opts ...grpc.CallOption) (*DeleteCommentReplayResp, error)
		// 更新产品评价回复
		UpdateCommentReplay(ctx context.Context, in *UpdateCommentReplayReq, opts ...grpc.CallOption) (*UpdateCommentReplayResp, error)
		// 查询产品评价回复详情
		QueryCommentReplayDetail(ctx context.Context, in *QueryCommentReplayDetailReq, opts ...grpc.CallOption) (*QueryCommentReplayDetailResp, error)
		// 查询产品评价回复列表
		QueryCommentReplayList(ctx context.Context, in *QueryCommentReplayListReq, opts ...grpc.CallOption) (*QueryCommentReplayListResp, error)
	}

	defaultCommentReplayService struct {
		cli zrpc.Client
	}
)

func NewCommentReplayService(cli zrpc.Client) CommentReplayService {
	return &defaultCommentReplayService{
		cli: cli,
	}
}

// 添加产品评价回复
func (m *defaultCommentReplayService) AddCommentReplay(ctx context.Context, in *AddCommentReplayReq, opts ...grpc.CallOption) (*AddCommentReplayResp, error) {
	client := pmsclient.NewCommentReplayServiceClient(m.cli.Conn())
	return client.AddCommentReplay(ctx, in, opts...)
}

// 删除产品评价回复
func (m *defaultCommentReplayService) DeleteCommentReplay(ctx context.Context, in *DeleteCommentReplayReq, opts ...grpc.CallOption) (*DeleteCommentReplayResp, error) {
	client := pmsclient.NewCommentReplayServiceClient(m.cli.Conn())
	return client.DeleteCommentReplay(ctx, in, opts...)
}

// 更新产品评价回复
func (m *defaultCommentReplayService) UpdateCommentReplay(ctx context.Context, in *UpdateCommentReplayReq, opts ...grpc.CallOption) (*UpdateCommentReplayResp, error) {
	client := pmsclient.NewCommentReplayServiceClient(m.cli.Conn())
	return client.UpdateCommentReplay(ctx, in, opts...)
}

// 查询产品评价回复详情
func (m *defaultCommentReplayService) QueryCommentReplayDetail(ctx context.Context, in *QueryCommentReplayDetailReq, opts ...grpc.CallOption) (*QueryCommentReplayDetailResp, error) {
	client := pmsclient.NewCommentReplayServiceClient(m.cli.Conn())
	return client.QueryCommentReplayDetail(ctx, in, opts...)
}

// 查询产品评价回复列表
func (m *defaultCommentReplayService) QueryCommentReplayList(ctx context.Context, in *QueryCommentReplayListReq, opts ...grpc.CallOption) (*QueryCommentReplayListResp, error) {
	client := pmsclient.NewCommentReplayServiceClient(m.cli.Conn())
	return client.QueryCommentReplayList(ctx, in, opts...)
}
