// Code generated by goctl. DO NOT EDIT.
// Source: cms.proto

package server

import (
	"context"

	"github.com/feihua/zero-admin/rpc/cms/cmsclient"
	"github.com/feihua/zero-admin/rpc/cms/internal/logic/topicservice"
	"github.com/feihua/zero-admin/rpc/cms/internal/svc"
)

type TopicServiceServer struct {
	svcCtx *svc.ServiceContext
	cmsclient.UnimplementedTopicServiceServer
}

func NewTopicServiceServer(svcCtx *svc.ServiceContext) *TopicServiceServer {
	return &TopicServiceServer{
		svcCtx: svcCtx,
	}
}

// 添加话题
func (s *TopicServiceServer) AddTopic(ctx context.Context, in *cmsclient.AddTopicReq) (*cmsclient.AddTopicResp, error) {
	l := topicservicelogic.NewAddTopicLogic(ctx, s.svcCtx)
	return l.AddTopic(in)
}

// 删除话题
func (s *TopicServiceServer) DeleteTopic(ctx context.Context, in *cmsclient.DeleteTopicReq) (*cmsclient.DeleteTopicResp, error) {
	l := topicservicelogic.NewDeleteTopicLogic(ctx, s.svcCtx)
	return l.DeleteTopic(in)
}

// 更新话题
func (s *TopicServiceServer) UpdateTopic(ctx context.Context, in *cmsclient.UpdateTopicReq) (*cmsclient.UpdateTopicResp, error) {
	l := topicservicelogic.NewUpdateTopicLogic(ctx, s.svcCtx)
	return l.UpdateTopic(in)
}

// 更新话题状态
func (s *TopicServiceServer) UpdateTopicStatus(ctx context.Context, in *cmsclient.UpdateTopicStatusReq) (*cmsclient.UpdateTopicStatusResp, error) {
	l := topicservicelogic.NewUpdateTopicStatusLogic(ctx, s.svcCtx)
	return l.UpdateTopicStatus(in)
}

// 查询话题详情
func (s *TopicServiceServer) QueryTopicDetail(ctx context.Context, in *cmsclient.QueryTopicDetailReq) (*cmsclient.QueryTopicDetailResp, error) {
	l := topicservicelogic.NewQueryTopicDetailLogic(ctx, s.svcCtx)
	return l.QueryTopicDetail(in)
}

// 查询话题列表
func (s *TopicServiceServer) QueryTopicList(ctx context.Context, in *cmsclient.QueryTopicListReq) (*cmsclient.QueryTopicListResp, error) {
	l := topicservicelogic.NewQueryTopicListLogic(ctx, s.svcCtx)
	return l.QueryTopicList(in)
}
