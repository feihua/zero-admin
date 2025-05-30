// Code generated by goctl. DO NOT EDIT.
// goctl 1.8.2
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
	AddCartItemResp                   = omsclient.AddCartItemResp
	AddCompanyAddressReq              = omsclient.AddCompanyAddressReq
	AddCompanyAddressResp             = omsclient.AddCompanyAddressResp
	AddOrderDeliveryReq               = omsclient.AddOrderDeliveryReq
	AddOrderDeliveryResp              = omsclient.AddOrderDeliveryResp
	AddOrderItemReq                   = omsclient.AddOrderItemReq
	AddOrderItemResp                  = omsclient.AddOrderItemResp
	AddOrderOperateHistoryReq         = omsclient.AddOrderOperateHistoryReq
	AddOrderOperateHistoryResp        = omsclient.AddOrderOperateHistoryResp
	AddOrderReturnApplyReq            = omsclient.AddOrderReturnApplyReq
	AddOrderReturnApplyResp           = omsclient.AddOrderReturnApplyResp
	AddOrderReturnReasonReq           = omsclient.AddOrderReturnReasonReq
	AddOrderReturnReasonResp          = omsclient.AddOrderReturnReasonResp
	AddOrderSettingReq                = omsclient.AddOrderSettingReq
	AddOrderSettingResp               = omsclient.AddOrderSettingResp
	CartItemListData                  = omsclient.CartItemListData
	CloseOrderReq                     = omsclient.CloseOrderReq
	CloseOrderResp                    = omsclient.CloseOrderResp
	CompanyAddressListData            = omsclient.CompanyAddressListData
	DeleteCartItemReq                 = omsclient.DeleteCartItemReq
	DeleteCartItemResp                = omsclient.DeleteCartItemResp
	DeleteCompanyAddressReq           = omsclient.DeleteCompanyAddressReq
	DeleteCompanyAddressResp          = omsclient.DeleteCompanyAddressResp
	DeleteOrderDeliveryReq            = omsclient.DeleteOrderDeliveryReq
	DeleteOrderDeliveryResp           = omsclient.DeleteOrderDeliveryResp
	DeleteOrderOperateHistoryReq      = omsclient.DeleteOrderOperateHistoryReq
	DeleteOrderOperateHistoryResp     = omsclient.DeleteOrderOperateHistoryResp
	DeleteOrderReturnApplyReq         = omsclient.DeleteOrderReturnApplyReq
	DeleteOrderReturnApplyResp        = omsclient.DeleteOrderReturnApplyResp
	DeleteOrderReturnReasonReq        = omsclient.DeleteOrderReturnReasonReq
	DeleteOrderReturnReasonResp       = omsclient.DeleteOrderReturnReasonResp
	DeleteOrderSettingReq             = omsclient.DeleteOrderSettingReq
	DeleteOrderSettingResp            = omsclient.DeleteOrderSettingResp
	DeliveryReq                       = omsclient.DeliveryReq
	DeliveryResp                      = omsclient.DeliveryResp
	OrderAddReq                       = omsclient.OrderAddReq
	OrderAddResp                      = omsclient.OrderAddResp
	OrderCancelReq                    = omsclient.OrderCancelReq
	OrderCancelResp                   = omsclient.OrderCancelResp
	OrderConfirmReq                   = omsclient.OrderConfirmReq
	OrderConfirmResp                  = omsclient.OrderConfirmResp
	OrderDeleteByIdReq                = omsclient.OrderDeleteByIdReq
	OrderDeleteReq                    = omsclient.OrderDeleteReq
	OrderDeleteResp                   = omsclient.OrderDeleteResp
	OrderDeliveryListData             = omsclient.OrderDeliveryListData
	OrderDetailReq                    = omsclient.OrderDetailReq
	OrderDetailResp                   = omsclient.OrderDetailResp
	OrderItemData                     = omsclient.OrderItemData
	OrderItemListData                 = omsclient.OrderItemListData
	OrderListByMemberIdReq            = omsclient.OrderListByMemberIdReq
	OrderListByMemberIdResp           = omsclient.OrderListByMemberIdResp
	OrderListData                     = omsclient.OrderListData
	OrderListReq                      = omsclient.OrderListReq
	OrderListResp                     = omsclient.OrderListResp
	OrderOperateHistoryData           = omsclient.OrderOperateHistoryData
	OrderOperateHistoryListData       = omsclient.OrderOperateHistoryListData
	OrderRefundReq                    = omsclient.OrderRefundReq
	OrderRefundResp                   = omsclient.OrderRefundResp
	OrderReturnApplyListData          = omsclient.OrderReturnApplyListData
	OrderReturnReasonListData         = omsclient.OrderReturnReasonListData
	OrderSettingListData              = omsclient.OrderSettingListData
	OrderUpdateReq                    = omsclient.OrderUpdateReq
	OrderUpdateResp                   = omsclient.OrderUpdateResp
	QueryCartItemDetailReq            = omsclient.QueryCartItemDetailReq
	QueryCartItemDetailResp           = omsclient.QueryCartItemDetailResp
	QueryCartItemListReq              = omsclient.QueryCartItemListReq
	QueryCartItemListResp             = omsclient.QueryCartItemListResp
	QueryCompanyAddressDetailReq      = omsclient.QueryCompanyAddressDetailReq
	QueryCompanyAddressDetailResp     = omsclient.QueryCompanyAddressDetailResp
	QueryCompanyAddressListReq        = omsclient.QueryCompanyAddressListReq
	QueryCompanyAddressListResp       = omsclient.QueryCompanyAddressListResp
	QueryOrderDeliveryDetailReq       = omsclient.QueryOrderDeliveryDetailReq
	QueryOrderDeliveryDetailResp      = omsclient.QueryOrderDeliveryDetailResp
	QueryOrderDeliveryListReq         = omsclient.QueryOrderDeliveryListReq
	QueryOrderDeliveryListResp        = omsclient.QueryOrderDeliveryListResp
	QueryOrderItemDetailReq           = omsclient.QueryOrderItemDetailReq
	QueryOrderItemDetailResp          = omsclient.QueryOrderItemDetailResp
	QueryOrderItemListReq             = omsclient.QueryOrderItemListReq
	QueryOrderItemListResp            = omsclient.QueryOrderItemListResp
	QueryOrderListReq                 = omsclient.QueryOrderListReq
	QueryOrderOperateHistoryListReq   = omsclient.QueryOrderOperateHistoryListReq
	QueryOrderOperateHistoryListResp  = omsclient.QueryOrderOperateHistoryListResp
	QueryOrderReturnApplyDetailReq    = omsclient.QueryOrderReturnApplyDetailReq
	QueryOrderReturnApplyDetailResp   = omsclient.QueryOrderReturnApplyDetailResp
	QueryOrderReturnApplyListReq      = omsclient.QueryOrderReturnApplyListReq
	QueryOrderReturnApplyListResp     = omsclient.QueryOrderReturnApplyListResp
	QueryOrderReturnReasonDetailReq   = omsclient.QueryOrderReturnReasonDetailReq
	QueryOrderReturnReasonDetailResp  = omsclient.QueryOrderReturnReasonDetailResp
	QueryOrderReturnReasonListReq     = omsclient.QueryOrderReturnReasonListReq
	QueryOrderReturnReasonListResp    = omsclient.QueryOrderReturnReasonListResp
	QueryOrderSettingDetailReq        = omsclient.QueryOrderSettingDetailReq
	QueryOrderSettingDetailResp       = omsclient.QueryOrderSettingDetailResp
	QueryOrderSettingListReq          = omsclient.QueryOrderSettingListReq
	QueryOrderSettingListResp         = omsclient.QueryOrderSettingListResp
	ReleaseSkuStockLockData           = omsclient.ReleaseSkuStockLockData
	UpdateCartItemQuantityReq         = omsclient.UpdateCartItemQuantityReq
	UpdateCartItemQuantityResp        = omsclient.UpdateCartItemQuantityResp
	UpdateCartItemReq                 = omsclient.UpdateCartItemReq
	UpdateCartItemResp                = omsclient.UpdateCartItemResp
	UpdateCompanyAddressReq           = omsclient.UpdateCompanyAddressReq
	UpdateCompanyAddressResp          = omsclient.UpdateCompanyAddressResp
	UpdateCompanyAddressStatusReq     = omsclient.UpdateCompanyAddressStatusReq
	UpdateCompanyAddressStatusResp    = omsclient.UpdateCompanyAddressStatusResp
	UpdateMoneyInfoReq                = omsclient.UpdateMoneyInfoReq
	UpdateMoneyInfoResp               = omsclient.UpdateMoneyInfoResp
	UpdateNoteReq                     = omsclient.UpdateNoteReq
	UpdateNoteResp                    = omsclient.UpdateNoteResp
	UpdateOrderDeliveryReq            = omsclient.UpdateOrderDeliveryReq
	UpdateOrderDeliveryResp           = omsclient.UpdateOrderDeliveryResp
	UpdateOrderReturnApplyReq         = omsclient.UpdateOrderReturnApplyReq
	UpdateOrderReturnApplyResp        = omsclient.UpdateOrderReturnApplyResp
	UpdateOrderReturnReasonReq        = omsclient.UpdateOrderReturnReasonReq
	UpdateOrderReturnReasonResp       = omsclient.UpdateOrderReturnReasonResp
	UpdateOrderReturnReasonStatusReq  = omsclient.UpdateOrderReturnReasonStatusReq
	UpdateOrderReturnReasonStatusResp = omsclient.UpdateOrderReturnReasonStatusResp
	UpdateOrderSettingReq             = omsclient.UpdateOrderSettingReq
	UpdateOrderSettingResp            = omsclient.UpdateOrderSettingResp
	UpdateOrderSettingStatusReq       = omsclient.UpdateOrderSettingStatusReq
	UpdateOrderSettingStatusResp      = omsclient.UpdateOrderSettingStatusResp
	UpdateOrderStatusByOutTradeNoReq  = omsclient.UpdateOrderStatusByOutTradeNoReq
	UpdateOrderStatusByOutTradeNoResp = omsclient.UpdateOrderStatusByOutTradeNoResp
	UpdateReceiverInfoReq             = omsclient.UpdateReceiverInfoReq
	UpdateReceiverInfoResp            = omsclient.UpdateReceiverInfoResp

	CartItemService interface {
		// 添加购物车表
		AddCartItem(ctx context.Context, in *AddCartItemReq, opts ...grpc.CallOption) (*AddCartItemResp, error)
		// 删除购物车表
		DeleteCartItem(ctx context.Context, in *DeleteCartItemReq, opts ...grpc.CallOption) (*DeleteCartItemResp, error)
		// 更新购物车表
		UpdateCartItem(ctx context.Context, in *UpdateCartItemReq, opts ...grpc.CallOption) (*UpdateCartItemResp, error)
		// 修改购物车中某个商品的数量
		UpdateCartItemQuantity(ctx context.Context, in *UpdateCartItemQuantityReq, opts ...grpc.CallOption) (*UpdateCartItemQuantityResp, error)
		// 查询购物车表详情
		QueryCartItemDetail(ctx context.Context, in *QueryCartItemDetailReq, opts ...grpc.CallOption) (*QueryCartItemDetailResp, error)
		// 查询购物车表列表
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

// 添加购物车表
func (m *defaultCartItemService) AddCartItem(ctx context.Context, in *AddCartItemReq, opts ...grpc.CallOption) (*AddCartItemResp, error) {
	client := omsclient.NewCartItemServiceClient(m.cli.Conn())
	return client.AddCartItem(ctx, in, opts...)
}

// 删除购物车表
func (m *defaultCartItemService) DeleteCartItem(ctx context.Context, in *DeleteCartItemReq, opts ...grpc.CallOption) (*DeleteCartItemResp, error) {
	client := omsclient.NewCartItemServiceClient(m.cli.Conn())
	return client.DeleteCartItem(ctx, in, opts...)
}

// 更新购物车表
func (m *defaultCartItemService) UpdateCartItem(ctx context.Context, in *UpdateCartItemReq, opts ...grpc.CallOption) (*UpdateCartItemResp, error) {
	client := omsclient.NewCartItemServiceClient(m.cli.Conn())
	return client.UpdateCartItem(ctx, in, opts...)
}

// 修改购物车中某个商品的数量
func (m *defaultCartItemService) UpdateCartItemQuantity(ctx context.Context, in *UpdateCartItemQuantityReq, opts ...grpc.CallOption) (*UpdateCartItemQuantityResp, error) {
	client := omsclient.NewCartItemServiceClient(m.cli.Conn())
	return client.UpdateCartItemQuantity(ctx, in, opts...)
}

// 查询购物车表详情
func (m *defaultCartItemService) QueryCartItemDetail(ctx context.Context, in *QueryCartItemDetailReq, opts ...grpc.CallOption) (*QueryCartItemDetailResp, error) {
	client := omsclient.NewCartItemServiceClient(m.cli.Conn())
	return client.QueryCartItemDetail(ctx, in, opts...)
}

// 查询购物车表列表
func (m *defaultCartItemService) QueryCartItemList(ctx context.Context, in *QueryCartItemListReq, opts ...grpc.CallOption) (*QueryCartItemListResp, error) {
	client := omsclient.NewCartItemServiceClient(m.cli.Conn())
	return client.QueryCartItemList(ctx, in, opts...)
}
