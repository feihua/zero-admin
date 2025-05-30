// Code generated by goctl. DO NOT EDIT.
// goctl 1.8.2
// Source: oms.proto

package companyaddressservice

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

	CompanyAddressService interface {
		// 添加公司收发货地址
		AddCompanyAddress(ctx context.Context, in *AddCompanyAddressReq, opts ...grpc.CallOption) (*AddCompanyAddressResp, error)
		// 删除公司收发货地址
		DeleteCompanyAddress(ctx context.Context, in *DeleteCompanyAddressReq, opts ...grpc.CallOption) (*DeleteCompanyAddressResp, error)
		// 更新公司收发货地址
		UpdateCompanyAddress(ctx context.Context, in *UpdateCompanyAddressReq, opts ...grpc.CallOption) (*UpdateCompanyAddressResp, error)
		// 查询公司收发货地址详情
		QueryCompanyAddressDetail(ctx context.Context, in *QueryCompanyAddressDetailReq, opts ...grpc.CallOption) (*QueryCompanyAddressDetailResp, error)
		// 查询公司收发货地址列表
		QueryCompanyAddressList(ctx context.Context, in *QueryCompanyAddressListReq, opts ...grpc.CallOption) (*QueryCompanyAddressListResp, error)
		// 更新公司默认发货地址
		UpdateCompanyAddressSendStatus(ctx context.Context, in *UpdateCompanyAddressStatusReq, opts ...grpc.CallOption) (*UpdateCompanyAddressStatusResp, error)
		// 更新公司默认收货地址
		UpdateCompanyAddressReceiveStatus(ctx context.Context, in *UpdateCompanyAddressStatusReq, opts ...grpc.CallOption) (*UpdateCompanyAddressStatusResp, error)
	}

	defaultCompanyAddressService struct {
		cli zrpc.Client
	}
)

func NewCompanyAddressService(cli zrpc.Client) CompanyAddressService {
	return &defaultCompanyAddressService{
		cli: cli,
	}
}

// 添加公司收发货地址
func (m *defaultCompanyAddressService) AddCompanyAddress(ctx context.Context, in *AddCompanyAddressReq, opts ...grpc.CallOption) (*AddCompanyAddressResp, error) {
	client := omsclient.NewCompanyAddressServiceClient(m.cli.Conn())
	return client.AddCompanyAddress(ctx, in, opts...)
}

// 删除公司收发货地址
func (m *defaultCompanyAddressService) DeleteCompanyAddress(ctx context.Context, in *DeleteCompanyAddressReq, opts ...grpc.CallOption) (*DeleteCompanyAddressResp, error) {
	client := omsclient.NewCompanyAddressServiceClient(m.cli.Conn())
	return client.DeleteCompanyAddress(ctx, in, opts...)
}

// 更新公司收发货地址
func (m *defaultCompanyAddressService) UpdateCompanyAddress(ctx context.Context, in *UpdateCompanyAddressReq, opts ...grpc.CallOption) (*UpdateCompanyAddressResp, error) {
	client := omsclient.NewCompanyAddressServiceClient(m.cli.Conn())
	return client.UpdateCompanyAddress(ctx, in, opts...)
}

// 查询公司收发货地址详情
func (m *defaultCompanyAddressService) QueryCompanyAddressDetail(ctx context.Context, in *QueryCompanyAddressDetailReq, opts ...grpc.CallOption) (*QueryCompanyAddressDetailResp, error) {
	client := omsclient.NewCompanyAddressServiceClient(m.cli.Conn())
	return client.QueryCompanyAddressDetail(ctx, in, opts...)
}

// 查询公司收发货地址列表
func (m *defaultCompanyAddressService) QueryCompanyAddressList(ctx context.Context, in *QueryCompanyAddressListReq, opts ...grpc.CallOption) (*QueryCompanyAddressListResp, error) {
	client := omsclient.NewCompanyAddressServiceClient(m.cli.Conn())
	return client.QueryCompanyAddressList(ctx, in, opts...)
}

// 更新公司默认发货地址
func (m *defaultCompanyAddressService) UpdateCompanyAddressSendStatus(ctx context.Context, in *UpdateCompanyAddressStatusReq, opts ...grpc.CallOption) (*UpdateCompanyAddressStatusResp, error) {
	client := omsclient.NewCompanyAddressServiceClient(m.cli.Conn())
	return client.UpdateCompanyAddressSendStatus(ctx, in, opts...)
}

// 更新公司默认收货地址
func (m *defaultCompanyAddressService) UpdateCompanyAddressReceiveStatus(ctx context.Context, in *UpdateCompanyAddressStatusReq, opts ...grpc.CallOption) (*UpdateCompanyAddressStatusResp, error) {
	client := omsclient.NewCompanyAddressServiceClient(m.cli.Conn())
	return client.UpdateCompanyAddressReceiveStatus(ctx, in, opts...)
}
