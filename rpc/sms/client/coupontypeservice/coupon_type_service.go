// Code generated by goctl. DO NOT EDIT.
// goctl 1.8.4
// Source: sms.proto

package coupontypeservice

import (
	"context"

	"github.com/feihua/zero-admin/rpc/sms/smsclient"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	AddCouponRecordReq                 = smsclient.AddCouponRecordReq
	AddCouponRecordResp                = smsclient.AddCouponRecordResp
	AddCouponReq                       = smsclient.AddCouponReq
	AddCouponResp                      = smsclient.AddCouponResp
	AddCouponScopeData                 = smsclient.AddCouponScopeData
	AddCouponScopeReq                  = smsclient.AddCouponScopeReq
	AddCouponScopeResp                 = smsclient.AddCouponScopeResp
	AddCouponTypeReq                   = smsclient.AddCouponTypeReq
	AddCouponTypeResp                  = smsclient.AddCouponTypeResp
	AddHomeAdvertiseReq                = smsclient.AddHomeAdvertiseReq
	AddHomeAdvertiseResp               = smsclient.AddHomeAdvertiseResp
	AddSeckillActivityReq              = smsclient.AddSeckillActivityReq
	AddSeckillActivityResp             = smsclient.AddSeckillActivityResp
	AddSeckillProductData              = smsclient.AddSeckillProductData
	AddSeckillProductReq               = smsclient.AddSeckillProductReq
	AddSeckillProductResp              = smsclient.AddSeckillProductResp
	AddSeckillReservationReq           = smsclient.AddSeckillReservationReq
	AddSeckillReservationResp          = smsclient.AddSeckillReservationResp
	AddSeckillSessionReq               = smsclient.AddSeckillSessionReq
	AddSeckillSessionResp              = smsclient.AddSeckillSessionResp
	CouponListData                     = smsclient.CouponListData
	CouponRecordListData               = smsclient.CouponRecordListData
	CouponScopeListData                = smsclient.CouponScopeListData
	CouponTypeListData                 = smsclient.CouponTypeListData
	DeleteCouponRecordReq              = smsclient.DeleteCouponRecordReq
	DeleteCouponRecordResp             = smsclient.DeleteCouponRecordResp
	DeleteCouponReq                    = smsclient.DeleteCouponReq
	DeleteCouponResp                   = smsclient.DeleteCouponResp
	DeleteCouponScopeReq               = smsclient.DeleteCouponScopeReq
	DeleteCouponScopeResp              = smsclient.DeleteCouponScopeResp
	DeleteCouponTypeReq                = smsclient.DeleteCouponTypeReq
	DeleteCouponTypeResp               = smsclient.DeleteCouponTypeResp
	DeleteHomeAdvertiseReq             = smsclient.DeleteHomeAdvertiseReq
	DeleteHomeAdvertiseResp            = smsclient.DeleteHomeAdvertiseResp
	DeleteSeckillActivityReq           = smsclient.DeleteSeckillActivityReq
	DeleteSeckillActivityResp          = smsclient.DeleteSeckillActivityResp
	DeleteSeckillProductReq            = smsclient.DeleteSeckillProductReq
	DeleteSeckillProductResp           = smsclient.DeleteSeckillProductResp
	DeleteSeckillReservationReq        = smsclient.DeleteSeckillReservationReq
	DeleteSeckillReservationResp       = smsclient.DeleteSeckillReservationResp
	DeleteSeckillSessionReq            = smsclient.DeleteSeckillSessionReq
	DeleteSeckillSessionResp           = smsclient.DeleteSeckillSessionResp
	HomeAdvertiseListData              = smsclient.HomeAdvertiseListData
	QueryCouponByCodeReq               = smsclient.QueryCouponByCodeReq
	QueryCouponByCodeResp              = smsclient.QueryCouponByCodeResp
	QueryCouponByScopeIdReq            = smsclient.QueryCouponByScopeIdReq
	QueryCouponByScopeIdResp           = smsclient.QueryCouponByScopeIdResp
	QueryCouponData                    = smsclient.QueryCouponData
	QueryCouponDetailReq               = smsclient.QueryCouponDetailReq
	QueryCouponDetailResp              = smsclient.QueryCouponDetailResp
	QueryCouponListReq                 = smsclient.QueryCouponListReq
	QueryCouponListResp                = smsclient.QueryCouponListResp
	QueryCouponRecordDetailReq         = smsclient.QueryCouponRecordDetailReq
	QueryCouponRecordDetailResp        = smsclient.QueryCouponRecordDetailResp
	QueryCouponRecordListReq           = smsclient.QueryCouponRecordListReq
	QueryCouponRecordListResp          = smsclient.QueryCouponRecordListResp
	QueryCouponScopeDetailReq          = smsclient.QueryCouponScopeDetailReq
	QueryCouponScopeDetailResp         = smsclient.QueryCouponScopeDetailResp
	QueryCouponScopeListReq            = smsclient.QueryCouponScopeListReq
	QueryCouponScopeListResp           = smsclient.QueryCouponScopeListResp
	QueryCouponTypeDetailReq           = smsclient.QueryCouponTypeDetailReq
	QueryCouponTypeDetailResp          = smsclient.QueryCouponTypeDetailResp
	QueryCouponTypeListReq             = smsclient.QueryCouponTypeListReq
	QueryCouponTypeListResp            = smsclient.QueryCouponTypeListResp
	QueryFlashPromotionListByDateResp  = smsclient.QueryFlashPromotionListByDateResp
	QueryHomeAdvertiseDetailReq        = smsclient.QueryHomeAdvertiseDetailReq
	QueryHomeAdvertiseDetailResp       = smsclient.QueryHomeAdvertiseDetailResp
	QueryHomeAdvertiseListReq          = smsclient.QueryHomeAdvertiseListReq
	QueryHomeAdvertiseListResp         = smsclient.QueryHomeAdvertiseListResp
	QueryMemberCouponListReq           = smsclient.QueryMemberCouponListReq
	QueryMemberCouponListResp          = smsclient.QueryMemberCouponListResp
	QuerySeckillActivityDetailReq      = smsclient.QuerySeckillActivityDetailReq
	QuerySeckillActivityDetailResp     = smsclient.QuerySeckillActivityDetailResp
	QuerySeckillActivityListByDateReq  = smsclient.QuerySeckillActivityListByDateReq
	QuerySeckillActivityListReq        = smsclient.QuerySeckillActivityListReq
	QuerySeckillActivityListResp       = smsclient.QuerySeckillActivityListResp
	QuerySeckillProductBySkuIdReq      = smsclient.QuerySeckillProductBySkuIdReq
	QuerySeckillProductDetailReq       = smsclient.QuerySeckillProductDetailReq
	QuerySeckillProductDetailResp      = smsclient.QuerySeckillProductDetailResp
	QuerySeckillProductListReq         = smsclient.QuerySeckillProductListReq
	QuerySeckillProductListResp        = smsclient.QuerySeckillProductListResp
	QuerySeckillReservationDetailReq   = smsclient.QuerySeckillReservationDetailReq
	QuerySeckillReservationDetailResp  = smsclient.QuerySeckillReservationDetailResp
	QuerySeckillReservationListReq     = smsclient.QuerySeckillReservationListReq
	QuerySeckillReservationListResp    = smsclient.QuerySeckillReservationListResp
	QuerySeckillSessionDetailReq       = smsclient.QuerySeckillSessionDetailReq
	QuerySeckillSessionDetailResp      = smsclient.QuerySeckillSessionDetailResp
	QuerySeckillSessionListByTimeReq   = smsclient.QuerySeckillSessionListByTimeReq
	QuerySeckillSessionListByTimeResp  = smsclient.QuerySeckillSessionListByTimeResp
	QuerySeckillSessionListReq         = smsclient.QuerySeckillSessionListReq
	QuerySeckillSessionListResp        = smsclient.QuerySeckillSessionListResp
	SeckillActivityListData            = smsclient.SeckillActivityListData
	SeckillProductListData             = smsclient.SeckillProductListData
	SeckillReservationListData         = smsclient.SeckillReservationListData
	SeckillSessionListData             = smsclient.SeckillSessionListData
	UpdateCouponRecordReq              = smsclient.UpdateCouponRecordReq
	UpdateCouponRecordResp             = smsclient.UpdateCouponRecordResp
	UpdateCouponReq                    = smsclient.UpdateCouponReq
	UpdateCouponResp                   = smsclient.UpdateCouponResp
	UpdateCouponScopeReq               = smsclient.UpdateCouponScopeReq
	UpdateCouponScopeResp              = smsclient.UpdateCouponScopeResp
	UpdateCouponStatusReq              = smsclient.UpdateCouponStatusReq
	UpdateCouponStatusResp             = smsclient.UpdateCouponStatusResp
	UpdateCouponTypeReq                = smsclient.UpdateCouponTypeReq
	UpdateCouponTypeResp               = smsclient.UpdateCouponTypeResp
	UpdateCouponTypeStatusReq          = smsclient.UpdateCouponTypeStatusReq
	UpdateCouponTypeStatusResp         = smsclient.UpdateCouponTypeStatusResp
	UpdateHomeAdvertiseReq             = smsclient.UpdateHomeAdvertiseReq
	UpdateHomeAdvertiseResp            = smsclient.UpdateHomeAdvertiseResp
	UpdateHomeAdvertiseStatusReq       = smsclient.UpdateHomeAdvertiseStatusReq
	UpdateHomeAdvertiseStatusResp      = smsclient.UpdateHomeAdvertiseStatusResp
	UpdateSeckillActivityReq           = smsclient.UpdateSeckillActivityReq
	UpdateSeckillActivityResp          = smsclient.UpdateSeckillActivityResp
	UpdateSeckillActivityStatusReq     = smsclient.UpdateSeckillActivityStatusReq
	UpdateSeckillActivityStatusResp    = smsclient.UpdateSeckillActivityStatusResp
	UpdateSeckillProductReq            = smsclient.UpdateSeckillProductReq
	UpdateSeckillProductResp           = smsclient.UpdateSeckillProductResp
	UpdateSeckillProductStatusReq      = smsclient.UpdateSeckillProductStatusReq
	UpdateSeckillProductStatusResp     = smsclient.UpdateSeckillProductStatusResp
	UpdateSeckillReservationReq        = smsclient.UpdateSeckillReservationReq
	UpdateSeckillReservationResp       = smsclient.UpdateSeckillReservationResp
	UpdateSeckillReservationStatusReq  = smsclient.UpdateSeckillReservationStatusReq
	UpdateSeckillReservationStatusResp = smsclient.UpdateSeckillReservationStatusResp
	UpdateSeckillSessionReq            = smsclient.UpdateSeckillSessionReq
	UpdateSeckillSessionResp           = smsclient.UpdateSeckillSessionResp
	UpdateSeckillSessionStatusReq      = smsclient.UpdateSeckillSessionStatusReq
	UpdateSeckillSessionStatusResp     = smsclient.UpdateSeckillSessionStatusResp

	CouponTypeService interface {
		// 添加优惠券类型
		AddCouponType(ctx context.Context, in *AddCouponTypeReq, opts ...grpc.CallOption) (*AddCouponTypeResp, error)
		// 删除优惠券类型
		DeleteCouponType(ctx context.Context, in *DeleteCouponTypeReq, opts ...grpc.CallOption) (*DeleteCouponTypeResp, error)
		// 更新优惠券类型
		UpdateCouponType(ctx context.Context, in *UpdateCouponTypeReq, opts ...grpc.CallOption) (*UpdateCouponTypeResp, error)
		// 更新优惠券类型状态
		UpdateCouponTypeStatus(ctx context.Context, in *UpdateCouponTypeStatusReq, opts ...grpc.CallOption) (*UpdateCouponTypeStatusResp, error)
		// 查询优惠券类型详情
		QueryCouponTypeDetail(ctx context.Context, in *QueryCouponTypeDetailReq, opts ...grpc.CallOption) (*QueryCouponTypeDetailResp, error)
		// 查询优惠券类型列表
		QueryCouponTypeList(ctx context.Context, in *QueryCouponTypeListReq, opts ...grpc.CallOption) (*QueryCouponTypeListResp, error)
	}

	defaultCouponTypeService struct {
		cli zrpc.Client
	}
)

func NewCouponTypeService(cli zrpc.Client) CouponTypeService {
	return &defaultCouponTypeService{
		cli: cli,
	}
}

// 添加优惠券类型
func (m *defaultCouponTypeService) AddCouponType(ctx context.Context, in *AddCouponTypeReq, opts ...grpc.CallOption) (*AddCouponTypeResp, error) {
	client := smsclient.NewCouponTypeServiceClient(m.cli.Conn())
	return client.AddCouponType(ctx, in, opts...)
}

// 删除优惠券类型
func (m *defaultCouponTypeService) DeleteCouponType(ctx context.Context, in *DeleteCouponTypeReq, opts ...grpc.CallOption) (*DeleteCouponTypeResp, error) {
	client := smsclient.NewCouponTypeServiceClient(m.cli.Conn())
	return client.DeleteCouponType(ctx, in, opts...)
}

// 更新优惠券类型
func (m *defaultCouponTypeService) UpdateCouponType(ctx context.Context, in *UpdateCouponTypeReq, opts ...grpc.CallOption) (*UpdateCouponTypeResp, error) {
	client := smsclient.NewCouponTypeServiceClient(m.cli.Conn())
	return client.UpdateCouponType(ctx, in, opts...)
}

// 更新优惠券类型状态
func (m *defaultCouponTypeService) UpdateCouponTypeStatus(ctx context.Context, in *UpdateCouponTypeStatusReq, opts ...grpc.CallOption) (*UpdateCouponTypeStatusResp, error) {
	client := smsclient.NewCouponTypeServiceClient(m.cli.Conn())
	return client.UpdateCouponTypeStatus(ctx, in, opts...)
}

// 查询优惠券类型详情
func (m *defaultCouponTypeService) QueryCouponTypeDetail(ctx context.Context, in *QueryCouponTypeDetailReq, opts ...grpc.CallOption) (*QueryCouponTypeDetailResp, error) {
	client := smsclient.NewCouponTypeServiceClient(m.cli.Conn())
	return client.QueryCouponTypeDetail(ctx, in, opts...)
}

// 查询优惠券类型列表
func (m *defaultCouponTypeService) QueryCouponTypeList(ctx context.Context, in *QueryCouponTypeListReq, opts ...grpc.CallOption) (*QueryCouponTypeListResp, error) {
	client := smsclient.NewCouponTypeServiceClient(m.cli.Conn())
	return client.QueryCouponTypeList(ctx, in, opts...)
}
