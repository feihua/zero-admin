// Code generated by goctl. DO NOT EDIT.
// goctl 1.8.4
// Source: oms.proto

package cartitemservice

import (
	"context"

	"github.com/feihua/zero-admin/rpc/oms/omsclient"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	AddCartItemReq                    = omsclient.AddCartItemReq
	AddCompanyAddressReq              = omsclient.AddCompanyAddressReq
	AddCompanyAddressResp             = omsclient.AddCompanyAddressResp
	AddOrderDeliveryReq               = omsclient.AddOrderDeliveryReq
	AddOrderDeliveryResp              = omsclient.AddOrderDeliveryResp
	AddOrderOperationLogReq           = omsclient.AddOrderOperationLogReq
	AddOrderOperationLogResp          = omsclient.AddOrderOperationLogResp
	AddOrderReq                       = omsclient.AddOrderReq
	AddOrderResp                      = omsclient.AddOrderResp
	AddOrderReturnReasonReq           = omsclient.AddOrderReturnReasonReq
	AddOrderReturnReasonResp          = omsclient.AddOrderReturnReasonResp
	AddOrderSettingReq                = omsclient.AddOrderSettingReq
	AddOrderSettingResp               = omsclient.AddOrderSettingResp
	CancelOrderReq                    = omsclient.CancelOrderReq
	CancelOrderResp                   = omsclient.CancelOrderResp
	CartItemData                      = omsclient.CartItemData
	CartItemResp                      = omsclient.CartItemResp
	CloseOrderReq                     = omsclient.CloseOrderReq
	CloseOrderResp                    = omsclient.CloseOrderResp
	CompanyAddressListData            = omsclient.CompanyAddressListData
	ConfirmOrderReq                   = omsclient.ConfirmOrderReq
	ConfirmOrderResp                  = omsclient.ConfirmOrderResp
	DeleteCartItemReq                 = omsclient.DeleteCartItemReq
	DeleteCompanyAddressReq           = omsclient.DeleteCompanyAddressReq
	DeleteCompanyAddressResp          = omsclient.DeleteCompanyAddressResp
	DeleteOrderDeliveryReq            = omsclient.DeleteOrderDeliveryReq
	DeleteOrderDeliveryResp           = omsclient.DeleteOrderDeliveryResp
	DeleteOrderReq                    = omsclient.DeleteOrderReq
	DeleteOrderResp                   = omsclient.DeleteOrderResp
	DeleteOrderReturnReasonReq        = omsclient.DeleteOrderReturnReasonReq
	DeleteOrderReturnReasonResp       = omsclient.DeleteOrderReturnReasonResp
	DeleteOrderReturnReq              = omsclient.DeleteOrderReturnReq
	DeleteOrderSettingReq             = omsclient.DeleteOrderSettingReq
	DeleteOrderSettingResp            = omsclient.DeleteOrderSettingResp
	DeliveryReq                       = omsclient.DeliveryReq
	DeliveryResp                      = omsclient.DeliveryResp
	OrderDeliveryListData             = omsclient.OrderDeliveryListData
	OrderItemData                     = omsclient.OrderItemData
	OrderListData                     = omsclient.OrderListData
	OrderOperationLogData             = omsclient.OrderOperationLogData
	OrderOptLogData                   = omsclient.OrderOptLogData
	OrderPaymentData                  = omsclient.OrderPaymentData
	OrderPaymentReq                   = omsclient.OrderPaymentReq
	OrderPaymentResp                  = omsclient.OrderPaymentResp
	OrderReturnData                   = omsclient.OrderReturnData
	OrderReturnItemData               = omsclient.OrderReturnItemData
	OrderReturnReasonListData         = omsclient.OrderReturnReasonListData
	OrderReturnReq                    = omsclient.OrderReturnReq
	OrderReturnResp                   = omsclient.OrderReturnResp
	OrderSettingListData              = omsclient.OrderSettingListData
	QueryCartItemDetailReq            = omsclient.QueryCartItemDetailReq
	QueryCartItemListReq              = omsclient.QueryCartItemListReq
	QueryCartItemListResp             = omsclient.QueryCartItemListResp
	QueryCompanyAddressDetailReq      = omsclient.QueryCompanyAddressDetailReq
	QueryCompanyAddressDetailResp     = omsclient.QueryCompanyAddressDetailResp
	QueryCompanyAddressListReq        = omsclient.QueryCompanyAddressListReq
	QueryCompanyAddressListResp       = omsclient.QueryCompanyAddressListResp
	QueryDefaultSettingReq            = omsclient.QueryDefaultSettingReq
	QueryOrderDeliveryDetailReq       = omsclient.QueryOrderDeliveryDetailReq
	QueryOrderDeliveryDetailResp      = omsclient.QueryOrderDeliveryDetailResp
	QueryOrderDeliveryListReq         = omsclient.QueryOrderDeliveryListReq
	QueryOrderDeliveryListResp        = omsclient.QueryOrderDeliveryListResp
	QueryOrderDetailReq               = omsclient.QueryOrderDetailReq
	QueryOrderDetailResp              = omsclient.QueryOrderDetailResp
	QueryOrderListReq                 = omsclient.QueryOrderListReq
	QueryOrderListResp                = omsclient.QueryOrderListResp
	QueryOrderOperationLogDetailReq   = omsclient.QueryOrderOperationLogDetailReq
	QueryOrderOperationLogListReq     = omsclient.QueryOrderOperationLogListReq
	QueryOrderOperationLogListResp    = omsclient.QueryOrderOperationLogListResp
	QueryOrderPaymentDetailReq        = omsclient.QueryOrderPaymentDetailReq
	QueryOrderPaymentListReq          = omsclient.QueryOrderPaymentListReq
	QueryOrderPaymentListResp         = omsclient.QueryOrderPaymentListResp
	QueryOrderReturnDetailReq         = omsclient.QueryOrderReturnDetailReq
	QueryOrderReturnListReq           = omsclient.QueryOrderReturnListReq
	QueryOrderReturnListResp          = omsclient.QueryOrderReturnListResp
	QueryOrderReturnReasonDetailReq   = omsclient.QueryOrderReturnReasonDetailReq
	QueryOrderReturnReasonDetailResp  = omsclient.QueryOrderReturnReasonDetailResp
	QueryOrderReturnReasonListReq     = omsclient.QueryOrderReturnReasonListReq
	QueryOrderReturnReasonListResp    = omsclient.QueryOrderReturnReasonListResp
	QueryOrderSettingDetailReq        = omsclient.QueryOrderSettingDetailReq
	QueryOrderSettingDetailResp       = omsclient.QueryOrderSettingDetailResp
	QueryOrderSettingListReq          = omsclient.QueryOrderSettingListReq
	QueryOrderSettingListResp         = omsclient.QueryOrderSettingListResp
	QueryTimeOutOrderListReq          = omsclient.QueryTimeOutOrderListReq
	ReleaseSkuStockLockData           = omsclient.ReleaseSkuStockLockData
	UpdateCartItemQuantityReq         = omsclient.UpdateCartItemQuantityReq
	UpdateCartItemReq                 = omsclient.UpdateCartItemReq
	UpdateCompanyAddressReq           = omsclient.UpdateCompanyAddressReq
	UpdateCompanyAddressResp          = omsclient.UpdateCompanyAddressResp
	UpdateCompanyAddressStatusReq     = omsclient.UpdateCompanyAddressStatusReq
	UpdateCompanyAddressStatusResp    = omsclient.UpdateCompanyAddressStatusResp
	UpdateOrderDeliveryReq            = omsclient.UpdateOrderDeliveryReq
	UpdateOrderDeliveryResp           = omsclient.UpdateOrderDeliveryResp
	UpdateOrderPaymentStatusReq       = omsclient.UpdateOrderPaymentStatusReq
	UpdateOrderReq                    = omsclient.UpdateOrderReq
	UpdateOrderResp                   = omsclient.UpdateOrderResp
	UpdateOrderReturnReasonReq        = omsclient.UpdateOrderReturnReasonReq
	UpdateOrderReturnReasonResp       = omsclient.UpdateOrderReturnReasonResp
	UpdateOrderReturnReasonStatusReq  = omsclient.UpdateOrderReturnReasonStatusReq
	UpdateOrderReturnReasonStatusResp = omsclient.UpdateOrderReturnReasonStatusResp
	UpdateOrderReturnStatusReq        = omsclient.UpdateOrderReturnStatusReq
	UpdateOrderSettingReq             = omsclient.UpdateOrderSettingReq
	UpdateOrderSettingResp            = omsclient.UpdateOrderSettingResp
	UpdateOrderSettingStatusReq       = omsclient.UpdateOrderSettingStatusReq
	UpdateOrderSettingStatusResp      = omsclient.UpdateOrderSettingStatusResp
	UpdateOrderStatusReq              = omsclient.UpdateOrderStatusReq
	UpdateOrderStatusResp             = omsclient.UpdateOrderStatusResp
	UpdateReceiverInfoResp            = omsclient.UpdateReceiverInfoResp

	CartItemService interface {
		// 添加购物车
		AddCartItem(ctx context.Context, in *AddCartItemReq, opts ...grpc.CallOption) (*CartItemResp, error)
		// 删除购物车
		DeleteCartItem(ctx context.Context, in *DeleteCartItemReq, opts ...grpc.CallOption) (*CartItemResp, error)
		// 更新购物车
		UpdateCartItem(ctx context.Context, in *UpdateCartItemReq, opts ...grpc.CallOption) (*CartItemResp, error)
		// 修改购物车中某个商品的数量
		UpdateCartItemQuantity(ctx context.Context, in *UpdateCartItemQuantityReq, opts ...grpc.CallOption) (*CartItemResp, error)
		// 查询购物车详情
		QueryCartItemDetail(ctx context.Context, in *QueryCartItemDetailReq, opts ...grpc.CallOption) (*CartItemData, error)
		// 查询购物车列
		QueryCartItemList(ctx context.Context, in *QueryCartItemListReq, opts ...grpc.CallOption) (*QueryCartItemListResp, error)
	}

	defaultCartItemService struct {
		cli zrpc.Client
	}
)

func NewCartItemService(cli zrpc.Client) CartItemService {
	return &defaultCartItemService{
		cli: cli,
	}
}

// 添加购物车
func (m *defaultCartItemService) AddCartItem(ctx context.Context, in *AddCartItemReq, opts ...grpc.CallOption) (*CartItemResp, error) {
	client := omsclient.NewCartItemServiceClient(m.cli.Conn())
	return client.AddCartItem(ctx, in, opts...)
}

// 删除购物车
func (m *defaultCartItemService) DeleteCartItem(ctx context.Context, in *DeleteCartItemReq, opts ...grpc.CallOption) (*CartItemResp, error) {
	client := omsclient.NewCartItemServiceClient(m.cli.Conn())
	return client.DeleteCartItem(ctx, in, opts...)
}

// 更新购物车
func (m *defaultCartItemService) UpdateCartItem(ctx context.Context, in *UpdateCartItemReq, opts ...grpc.CallOption) (*CartItemResp, error) {
	client := omsclient.NewCartItemServiceClient(m.cli.Conn())
	return client.UpdateCartItem(ctx, in, opts...)
}

// 修改购物车中某个商品的数量
func (m *defaultCartItemService) UpdateCartItemQuantity(ctx context.Context, in *UpdateCartItemQuantityReq, opts ...grpc.CallOption) (*CartItemResp, error) {
	client := omsclient.NewCartItemServiceClient(m.cli.Conn())
	return client.UpdateCartItemQuantity(ctx, in, opts...)
}

// 查询购物车详情
func (m *defaultCartItemService) QueryCartItemDetail(ctx context.Context, in *QueryCartItemDetailReq, opts ...grpc.CallOption) (*CartItemData, error) {
	client := omsclient.NewCartItemServiceClient(m.cli.Conn())
	return client.QueryCartItemDetail(ctx, in, opts...)
}

// 查询购物车列
func (m *defaultCartItemService) QueryCartItemList(ctx context.Context, in *QueryCartItemListReq, opts ...grpc.CallOption) (*QueryCartItemListResp, error) {
	client := omsclient.NewCartItemServiceClient(m.cli.Conn())
	return client.QueryCartItemList(ctx, in, opts...)
}
