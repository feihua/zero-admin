// Code generated by goctl. DO NOT EDIT.
// Source: oms.proto

package server

import (
	"context"

	"zero-admin/rpc/oms/internal/logic"
	"zero-admin/rpc/oms/internal/svc"
	"zero-admin/rpc/oms/omsclient"
)

type OmsServer struct {
	svcCtx *svc.ServiceContext
	omsclient.UnimplementedOmsServer
}

func NewOmsServer(svcCtx *svc.ServiceContext) *OmsServer {
	return &OmsServer{
		svcCtx: svcCtx,
	}
}

func (s *OmsServer) OrderAdd(ctx context.Context, in *omsclient.OrderAddReq) (*omsclient.OrderAddResp, error) {
	l := logic.NewOrderAddLogic(ctx, s.svcCtx)
	return l.OrderAdd(in)
}

func (s *OmsServer) OrderList(ctx context.Context, in *omsclient.OrderListReq) (*omsclient.OrderListResp, error) {
	l := logic.NewOrderListLogic(ctx, s.svcCtx)
	return l.OrderList(in)
}

func (s *OmsServer) OrderUpdate(ctx context.Context, in *omsclient.OrderUpdateReq) (*omsclient.OrderUpdateResp, error) {
	l := logic.NewOrderUpdateLogic(ctx, s.svcCtx)
	return l.OrderUpdate(in)
}

func (s *OmsServer) OrderDelete(ctx context.Context, in *omsclient.OrderDeleteReq) (*omsclient.OrderDeleteResp, error) {
	l := logic.NewOrderDeleteLogic(ctx, s.svcCtx)
	return l.OrderDelete(in)
}

func (s *OmsServer) OrderByOrderId(ctx context.Context, in *omsclient.OrderByOrderIdReq) (*omsclient.OrderByOrderIdResp, error) {
	l := logic.NewOrderByOrderIdLogic(ctx, s.svcCtx)
	return l.OrderByOrderId(in)
}

func (s *OmsServer) OrderListByMemberId(ctx context.Context, in *omsclient.OrderListByMemberIdReq) (*omsclient.OrderListByMemberIdResp, error) {
	l := logic.NewOrderListByMemberIdLogic(ctx, s.svcCtx)
	return l.OrderListByMemberId(in)
}

func (s *OmsServer) OrderCancel(ctx context.Context, in *omsclient.OrderCancelReq) (*omsclient.OrderCancelResp, error) {
	l := logic.NewOrderCancelLogic(ctx, s.svcCtx)
	return l.OrderCancel(in)
}

func (s *OmsServer) OrderConfirm(ctx context.Context, in *omsclient.OrderConfirmReq) (*omsclient.OrderConfirmResp, error) {
	l := logic.NewOrderConfirmLogic(ctx, s.svcCtx)
	return l.OrderConfirm(in)
}

func (s *OmsServer) OrderRefund(ctx context.Context, in *omsclient.OrderRefundReq) (*omsclient.OrderRefundResp, error) {
	l := logic.NewOrderRefundLogic(ctx, s.svcCtx)
	return l.OrderRefund(in)
}

func (s *OmsServer) CartItemAdd(ctx context.Context, in *omsclient.CartItemAddReq) (*omsclient.CartItemAddResp, error) {
	l := logic.NewCartItemAddLogic(ctx, s.svcCtx)
	return l.CartItemAdd(in)
}

func (s *OmsServer) CartItemList(ctx context.Context, in *omsclient.CartItemListReq) (*omsclient.CartItemListResp, error) {
	l := logic.NewCartItemListLogic(ctx, s.svcCtx)
	return l.CartItemList(in)
}

func (s *OmsServer) CartItemUpdate(ctx context.Context, in *omsclient.CartItemUpdateReq) (*omsclient.CartItemUpdateResp, error) {
	l := logic.NewCartItemUpdateLogic(ctx, s.svcCtx)
	return l.CartItemUpdate(in)
}

func (s *OmsServer) CartItemDelete(ctx context.Context, in *omsclient.CartItemDeleteReq) (*omsclient.CartItemDeleteResp, error) {
	l := logic.NewCartItemDeleteLogic(ctx, s.svcCtx)
	return l.CartItemDelete(in)
}

func (s *OmsServer) CartItemChecked(ctx context.Context, in *omsclient.CartItemCheckedReq) (*omsclient.CartItemCheckedResp, error) {
	l := logic.NewCartItemCheckedLogic(ctx, s.svcCtx)
	return l.CartItemChecked(in)
}

func (s *OmsServer) CartItemCheckOut(ctx context.Context, in *omsclient.CartItemCheckOutReq) (*omsclient.CartItemCheckOutResp, error) {
	l := logic.NewCartItemCheckOutLogic(ctx, s.svcCtx)
	return l.CartItemCheckOut(in)
}

func (s *OmsServer) CartItemFastAdd(ctx context.Context, in *omsclient.CartItemFastAddReq) (*omsclient.CartItemFastAddResp, error) {
	l := logic.NewCartItemFastAddLogic(ctx, s.svcCtx)
	return l.CartItemFastAdd(in)
}

func (s *OmsServer) CompanyAddressAdd(ctx context.Context, in *omsclient.CompanyAddressAddReq) (*omsclient.CompanyAddressAddResp, error) {
	l := logic.NewCompanyAddressAddLogic(ctx, s.svcCtx)
	return l.CompanyAddressAdd(in)
}

func (s *OmsServer) CompanyAddressList(ctx context.Context, in *omsclient.CompanyAddressListReq) (*omsclient.CompanyAddressListResp, error) {
	l := logic.NewCompanyAddressListLogic(ctx, s.svcCtx)
	return l.CompanyAddressList(in)
}

func (s *OmsServer) CompanyAddressUpdate(ctx context.Context, in *omsclient.CompanyAddressUpdateReq) (*omsclient.CompanyAddressUpdateResp, error) {
	l := logic.NewCompanyAddressUpdateLogic(ctx, s.svcCtx)
	return l.CompanyAddressUpdate(in)
}

func (s *OmsServer) CompanyAddressDelete(ctx context.Context, in *omsclient.CompanyAddressDeleteReq) (*omsclient.CompanyAddressDeleteResp, error) {
	l := logic.NewCompanyAddressDeleteLogic(ctx, s.svcCtx)
	return l.CompanyAddressDelete(in)
}

func (s *OmsServer) OrderItemAdd(ctx context.Context, in *omsclient.OrderItemAddReq) (*omsclient.OrderItemAddResp, error) {
	l := logic.NewOrderItemAddLogic(ctx, s.svcCtx)
	return l.OrderItemAdd(in)
}

func (s *OmsServer) OrderItemList(ctx context.Context, in *omsclient.OrderItemListReq) (*omsclient.OrderItemListResp, error) {
	l := logic.NewOrderItemListLogic(ctx, s.svcCtx)
	return l.OrderItemList(in)
}

func (s *OmsServer) OrderItemUpdate(ctx context.Context, in *omsclient.OrderItemUpdateReq) (*omsclient.OrderItemUpdateResp, error) {
	l := logic.NewOrderItemUpdateLogic(ctx, s.svcCtx)
	return l.OrderItemUpdate(in)
}

func (s *OmsServer) OrderItemDelete(ctx context.Context, in *omsclient.OrderItemDeleteReq) (*omsclient.OrderItemDeleteResp, error) {
	l := logic.NewOrderItemDeleteLogic(ctx, s.svcCtx)
	return l.OrderItemDelete(in)
}

func (s *OmsServer) OrderOperateHistoryAdd(ctx context.Context, in *omsclient.OrderOperateHistoryAddReq) (*omsclient.OrderOperateHistoryAddResp, error) {
	l := logic.NewOrderOperateHistoryAddLogic(ctx, s.svcCtx)
	return l.OrderOperateHistoryAdd(in)
}

func (s *OmsServer) OrderOperateHistoryList(ctx context.Context, in *omsclient.OrderOperateHistoryListReq) (*omsclient.OrderOperateHistoryListResp, error) {
	l := logic.NewOrderOperateHistoryListLogic(ctx, s.svcCtx)
	return l.OrderOperateHistoryList(in)
}

func (s *OmsServer) OrderOperateHistoryUpdate(ctx context.Context, in *omsclient.OrderOperateHistoryUpdateReq) (*omsclient.OrderOperateHistoryUpdateResp, error) {
	l := logic.NewOrderOperateHistoryUpdateLogic(ctx, s.svcCtx)
	return l.OrderOperateHistoryUpdate(in)
}

func (s *OmsServer) OrderOperateHistoryDelete(ctx context.Context, in *omsclient.OrderOperateHistoryDeleteReq) (*omsclient.OrderOperateHistoryDeleteResp, error) {
	l := logic.NewOrderOperateHistoryDeleteLogic(ctx, s.svcCtx)
	return l.OrderOperateHistoryDelete(in)
}

func (s *OmsServer) OrderReturnApplyAdd(ctx context.Context, in *omsclient.OrderReturnApplyAddReq) (*omsclient.OrderReturnApplyAddResp, error) {
	l := logic.NewOrderReturnApplyAddLogic(ctx, s.svcCtx)
	return l.OrderReturnApplyAdd(in)
}

func (s *OmsServer) OrderReturnApplyList(ctx context.Context, in *omsclient.OrderReturnApplyListReq) (*omsclient.OrderReturnApplyListResp, error) {
	l := logic.NewOrderReturnApplyListLogic(ctx, s.svcCtx)
	return l.OrderReturnApplyList(in)
}

func (s *OmsServer) OrderReturnApplyUpdate(ctx context.Context, in *omsclient.OrderReturnApplyUpdateReq) (*omsclient.OrderReturnApplyUpdateResp, error) {
	l := logic.NewOrderReturnApplyUpdateLogic(ctx, s.svcCtx)
	return l.OrderReturnApplyUpdate(in)
}

func (s *OmsServer) OrderReturnApplyDelete(ctx context.Context, in *omsclient.OrderReturnApplyDeleteReq) (*omsclient.OrderReturnApplyDeleteResp, error) {
	l := logic.NewOrderReturnApplyDeleteLogic(ctx, s.svcCtx)
	return l.OrderReturnApplyDelete(in)
}

func (s *OmsServer) OrderReturnReasonAdd(ctx context.Context, in *omsclient.OrderReturnReasonAddReq) (*omsclient.OrderReturnReasonAddResp, error) {
	l := logic.NewOrderReturnReasonAddLogic(ctx, s.svcCtx)
	return l.OrderReturnReasonAdd(in)
}

func (s *OmsServer) OrderReturnReasonList(ctx context.Context, in *omsclient.OrderReturnReasonListReq) (*omsclient.OrderReturnReasonListResp, error) {
	l := logic.NewOrderReturnReasonListLogic(ctx, s.svcCtx)
	return l.OrderReturnReasonList(in)
}

func (s *OmsServer) OrderReturnReasonUpdate(ctx context.Context, in *omsclient.OrderReturnReasonUpdateReq) (*omsclient.OrderReturnReasonUpdateResp, error) {
	l := logic.NewOrderReturnReasonUpdateLogic(ctx, s.svcCtx)
	return l.OrderReturnReasonUpdate(in)
}

func (s *OmsServer) OrderReturnReasonDelete(ctx context.Context, in *omsclient.OrderReturnReasonDeleteReq) (*omsclient.OrderReturnReasonDeleteResp, error) {
	l := logic.NewOrderReturnReasonDeleteLogic(ctx, s.svcCtx)
	return l.OrderReturnReasonDelete(in)
}

func (s *OmsServer) OrderSettingAdd(ctx context.Context, in *omsclient.OrderSettingAddReq) (*omsclient.OrderSettingAddResp, error) {
	l := logic.NewOrderSettingAddLogic(ctx, s.svcCtx)
	return l.OrderSettingAdd(in)
}

func (s *OmsServer) OrderSettingList(ctx context.Context, in *omsclient.OrderSettingListReq) (*omsclient.OrderSettingListResp, error) {
	l := logic.NewOrderSettingListLogic(ctx, s.svcCtx)
	return l.OrderSettingList(in)
}

func (s *OmsServer) OrderSettingUpdate(ctx context.Context, in *omsclient.OrderSettingUpdateReq) (*omsclient.OrderSettingUpdateResp, error) {
	l := logic.NewOrderSettingUpdateLogic(ctx, s.svcCtx)
	return l.OrderSettingUpdate(in)
}

func (s *OmsServer) OrderSettingDelete(ctx context.Context, in *omsclient.OrderSettingDeleteReq) (*omsclient.OrderSettingDeleteResp, error) {
	l := logic.NewOrderSettingDeleteLogic(ctx, s.svcCtx)
	return l.OrderSettingDelete(in)
}
