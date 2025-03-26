// Code generated by goctl. DO NOT EDIT.
// goctl 1.8.1
// Source: ums.proto

package server

import (
	"context"

	"github.com/feihua/zero-admin/rpc/ums/internal/logic/membertaskservice"
	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"
)

type MemberTaskServiceServer struct {
	svcCtx *svc.ServiceContext
	umsclient.UnimplementedMemberTaskServiceServer
}

func NewMemberTaskServiceServer(svcCtx *svc.ServiceContext) *MemberTaskServiceServer {
	return &MemberTaskServiceServer{
		svcCtx: svcCtx,
	}
}

// 添加会员任务表
func (s *MemberTaskServiceServer) AddMemberTask(ctx context.Context, in *umsclient.AddMemberTaskReq) (*umsclient.AddMemberTaskResp, error) {
	l := membertaskservicelogic.NewAddMemberTaskLogic(ctx, s.svcCtx)
	return l.AddMemberTask(in)
}

// 删除会员任务表
func (s *MemberTaskServiceServer) DeleteMemberTask(ctx context.Context, in *umsclient.DeleteMemberTaskReq) (*umsclient.DeleteMemberTaskResp, error) {
	l := membertaskservicelogic.NewDeleteMemberTaskLogic(ctx, s.svcCtx)
	return l.DeleteMemberTask(in)
}

// 更新会员任务表
func (s *MemberTaskServiceServer) UpdateMemberTask(ctx context.Context, in *umsclient.UpdateMemberTaskReq) (*umsclient.UpdateMemberTaskResp, error) {
	l := membertaskservicelogic.NewUpdateMemberTaskLogic(ctx, s.svcCtx)
	return l.UpdateMemberTask(in)
}

// 更新会员任务表状态
func (s *MemberTaskServiceServer) UpdateMemberTaskStatus(ctx context.Context, in *umsclient.UpdateMemberTaskStatusReq) (*umsclient.UpdateMemberTaskStatusResp, error) {
	l := membertaskservicelogic.NewUpdateMemberTaskStatusLogic(ctx, s.svcCtx)
	return l.UpdateMemberTaskStatus(in)
}

// 查询会员任务表详情
func (s *MemberTaskServiceServer) QueryMemberTaskDetail(ctx context.Context, in *umsclient.QueryMemberTaskDetailReq) (*umsclient.QueryMemberTaskDetailResp, error) {
	l := membertaskservicelogic.NewQueryMemberTaskDetailLogic(ctx, s.svcCtx)
	return l.QueryMemberTaskDetail(in)
}

// 查询会员任务表列表
func (s *MemberTaskServiceServer) QueryMemberTaskList(ctx context.Context, in *umsclient.QueryMemberTaskListReq) (*umsclient.QueryMemberTaskListResp, error) {
	l := membertaskservicelogic.NewQueryMemberTaskListLogic(ctx, s.svcCtx)
	return l.QueryMemberTaskList(in)
}
