// Code generated by goctl. DO NOT EDIT.
// goctl 1.8.4
// Source: cms.proto

package server

import (
	"context"

	"github.com/feihua/zero-admin/rpc/cms/cmsclient"
	"github.com/feihua/zero-admin/rpc/cms/internal/logic/subjectcommentservice"
	"github.com/feihua/zero-admin/rpc/cms/internal/svc"
)

type SubjectCommentServiceServer struct {
	svcCtx *svc.ServiceContext
	cmsclient.UnimplementedSubjectCommentServiceServer
}

func NewSubjectCommentServiceServer(svcCtx *svc.ServiceContext) *SubjectCommentServiceServer {
	return &SubjectCommentServiceServer{
		svcCtx: svcCtx,
	}
}

// 添加专题评论
func (s *SubjectCommentServiceServer) AddSubjectComment(ctx context.Context, in *cmsclient.AddSubjectCommentReq) (*cmsclient.AddSubjectCommentResp, error) {
	l := subjectcommentservicelogic.NewAddSubjectCommentLogic(ctx, s.svcCtx)
	return l.AddSubjectComment(in)
}

// 删除专题评论
func (s *SubjectCommentServiceServer) DeleteSubjectComment(ctx context.Context, in *cmsclient.DeleteSubjectCommentReq) (*cmsclient.DeleteSubjectCommentResp, error) {
	l := subjectcommentservicelogic.NewDeleteSubjectCommentLogic(ctx, s.svcCtx)
	return l.DeleteSubjectComment(in)
}

// 更新专题评论
func (s *SubjectCommentServiceServer) UpdateSubjectComment(ctx context.Context, in *cmsclient.UpdateSubjectCommentReq) (*cmsclient.UpdateSubjectCommentResp, error) {
	l := subjectcommentservicelogic.NewUpdateSubjectCommentLogic(ctx, s.svcCtx)
	return l.UpdateSubjectComment(in)
}

// 更新专题评论状态
func (s *SubjectCommentServiceServer) UpdateSubjectCommentStatus(ctx context.Context, in *cmsclient.UpdateSubjectCommentStatusReq) (*cmsclient.UpdateSubjectCommentStatusResp, error) {
	l := subjectcommentservicelogic.NewUpdateSubjectCommentStatusLogic(ctx, s.svcCtx)
	return l.UpdateSubjectCommentStatus(in)
}

// 查询专题评论详情
func (s *SubjectCommentServiceServer) QuerySubjectCommentDetail(ctx context.Context, in *cmsclient.QuerySubjectCommentDetailReq) (*cmsclient.QuerySubjectCommentDetailResp, error) {
	l := subjectcommentservicelogic.NewQuerySubjectCommentDetailLogic(ctx, s.svcCtx)
	return l.QuerySubjectCommentDetail(in)
}

// 查询专题评论列表
func (s *SubjectCommentServiceServer) QuerySubjectCommentList(ctx context.Context, in *cmsclient.QuerySubjectCommentListReq) (*cmsclient.QuerySubjectCommentListResp, error) {
	l := subjectcommentservicelogic.NewQuerySubjectCommentListLogic(ctx, s.svcCtx)
	return l.QuerySubjectCommentList(in)
}
