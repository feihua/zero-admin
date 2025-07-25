// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.6
// Source: pms.proto

package server

import (
	"context"

	"github.com/feihua/zero-admin/rpc/pms/internal/logic/productvertifyrecordservice"
	"github.com/feihua/zero-admin/rpc/pms/internal/svc"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"
)

type ProductVertifyRecordServiceServer struct {
	svcCtx *svc.ServiceContext
	pmsclient.UnimplementedProductVertifyRecordServiceServer
}

func NewProductVertifyRecordServiceServer(svcCtx *svc.ServiceContext) *ProductVertifyRecordServiceServer {
	return &ProductVertifyRecordServiceServer{
		svcCtx: svcCtx,
	}
}

// 添加商品审核记录
func (s *ProductVertifyRecordServiceServer) AddProductVertifyRecord(ctx context.Context, in *pmsclient.AddProductVertifyRecordReq) (*pmsclient.AddProductVertifyRecordResp, error) {
	l := productvertifyrecordservicelogic.NewAddProductVertifyRecordLogic(ctx, s.svcCtx)
	return l.AddProductVertifyRecord(in)
}

// 查询商品审核记录详情
func (s *ProductVertifyRecordServiceServer) QueryProductVertifyRecordDetail(ctx context.Context, in *pmsclient.QueryProductVertifyRecordDetailReq) (*pmsclient.QueryProductVertifyRecordDetailResp, error) {
	l := productvertifyrecordservicelogic.NewQueryProductVertifyRecordDetailLogic(ctx, s.svcCtx)
	return l.QueryProductVertifyRecordDetail(in)
}

// 查询商品审核记录列表
func (s *ProductVertifyRecordServiceServer) QueryProductVertifyRecordList(ctx context.Context, in *pmsclient.QueryProductVertifyRecordListReq) (*pmsclient.QueryProductVertifyRecordListResp, error) {
	l := productvertifyrecordservicelogic.NewQueryProductVertifyRecordListLogic(ctx, s.svcCtx)
	return l.QueryProductVertifyRecordList(in)
}
