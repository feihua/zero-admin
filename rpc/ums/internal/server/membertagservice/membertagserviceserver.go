// Code generated by goctl. DO NOT EDIT.
// goctl 1.8.1
// Source: ums.proto

package server

import (
	"context"

	"github.com/feihua/zero-admin/rpc/ums/internal/logic/membertagservice"
	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"
)

type MemberTagServiceServer struct {
	svcCtx *svc.ServiceContext
	umsclient.UnimplementedMemberTagServiceServer
}

func NewMemberTagServiceServer(svcCtx *svc.ServiceContext) *MemberTagServiceServer {
	return &MemberTagServiceServer{
		svcCtx: svcCtx,
	}
}

// 添加用户标签表
func (s *MemberTagServiceServer) AddMemberTag(ctx context.Context, in *umsclient.AddMemberTagReq) (*umsclient.AddMemberTagResp, error) {
	l := membertagservicelogic.NewAddMemberTagLogic(ctx, s.svcCtx)
	return l.AddMemberTag(in)
}

// 删除用户标签表
func (s *MemberTagServiceServer) DeleteMemberTag(ctx context.Context, in *umsclient.DeleteMemberTagReq) (*umsclient.DeleteMemberTagResp, error) {
	l := membertagservicelogic.NewDeleteMemberTagLogic(ctx, s.svcCtx)
	return l.DeleteMemberTag(in)
}

// 更新用户标签表
func (s *MemberTagServiceServer) UpdateMemberTag(ctx context.Context, in *umsclient.UpdateMemberTagReq) (*umsclient.UpdateMemberTagResp, error) {
	l := membertagservicelogic.NewUpdateMemberTagLogic(ctx, s.svcCtx)
	return l.UpdateMemberTag(in)
}

// 更新用户标签表状态
func (s *MemberTagServiceServer) UpdateMemberTagStatus(ctx context.Context, in *umsclient.UpdateMemberTagStatusReq) (*umsclient.UpdateMemberTagStatusResp, error) {
	l := membertagservicelogic.NewUpdateMemberTagStatusLogic(ctx, s.svcCtx)
	return l.UpdateMemberTagStatus(in)
}

// 查询用户标签表详情
func (s *MemberTagServiceServer) QueryMemberTagDetail(ctx context.Context, in *umsclient.QueryMemberTagDetailReq) (*umsclient.QueryMemberTagDetailResp, error) {
	l := membertagservicelogic.NewQueryMemberTagDetailLogic(ctx, s.svcCtx)
	return l.QueryMemberTagDetail(in)
}

// 查询用户标签表列表
func (s *MemberTagServiceServer) QueryMemberTagList(ctx context.Context, in *umsclient.QueryMemberTagListReq) (*umsclient.QueryMemberTagListResp, error) {
	l := membertagservicelogic.NewQueryMemberTagListLogic(ctx, s.svcCtx)
	return l.QueryMemberTagList(in)
}
