// Code generated by goctl. DO NOT EDIT.
// goctl 1.8.4
// Source: ums.proto

package server

import (
	"context"

	"github.com/feihua/zero-admin/rpc/ums/internal/logic/membertaskrelationservice"
	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"
)

type MemberTaskRelationServiceServer struct {
	svcCtx *svc.ServiceContext
	umsclient.UnimplementedMemberTaskRelationServiceServer
}

func NewMemberTaskRelationServiceServer(svcCtx *svc.ServiceContext) *MemberTaskRelationServiceServer {
	return &MemberTaskRelationServiceServer{
		svcCtx: svcCtx,
	}
}

// 添加会员任务关联
func (s *MemberTaskRelationServiceServer) AddMemberTaskRelation(ctx context.Context, in *umsclient.AddMemberTaskRelationReq) (*umsclient.AddMemberTaskRelationResp, error) {
	l := membertaskrelationservicelogic.NewAddMemberTaskRelationLogic(ctx, s.svcCtx)
	return l.AddMemberTaskRelation(in)
}

// 查询会员任务关联详情
func (s *MemberTaskRelationServiceServer) QueryMemberTaskRelationDetail(ctx context.Context, in *umsclient.QueryMemberTaskRelationDetailReq) (*umsclient.QueryMemberTaskRelationDetailResp, error) {
	l := membertaskrelationservicelogic.NewQueryMemberTaskRelationDetailLogic(ctx, s.svcCtx)
	return l.QueryMemberTaskRelationDetail(in)
}

// 查询会员任务关联列表
func (s *MemberTaskRelationServiceServer) QueryMemberTaskRelationList(ctx context.Context, in *umsclient.QueryMemberTaskRelationListReq) (*umsclient.QueryMemberTaskRelationListResp, error) {
	l := membertaskrelationservicelogic.NewQueryMemberTaskRelationListLogic(ctx, s.svcCtx)
	return l.QueryMemberTaskRelationList(in)
}
