// Code generated by goctl. DO NOT EDIT.
// goctl 1.8.1
// Source: ums.proto

package server

import (
	"context"

	"github.com/feihua/zero-admin/rpc/ums/internal/logic/memberrulesettingservice"
	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"
)

type MemberRuleSettingServiceServer struct {
	svcCtx *svc.ServiceContext
	umsclient.UnimplementedMemberRuleSettingServiceServer
}

func NewMemberRuleSettingServiceServer(svcCtx *svc.ServiceContext) *MemberRuleSettingServiceServer {
	return &MemberRuleSettingServiceServer{
		svcCtx: svcCtx,
	}
}

// 添加会员积分成长规则表
func (s *MemberRuleSettingServiceServer) AddMemberRuleSetting(ctx context.Context, in *umsclient.AddMemberRuleSettingReq) (*umsclient.AddMemberRuleSettingResp, error) {
	l := memberrulesettingservicelogic.NewAddMemberRuleSettingLogic(ctx, s.svcCtx)
	return l.AddMemberRuleSetting(in)
}

// 删除会员积分成长规则表
func (s *MemberRuleSettingServiceServer) DeleteMemberRuleSetting(ctx context.Context, in *umsclient.DeleteMemberRuleSettingReq) (*umsclient.DeleteMemberRuleSettingResp, error) {
	l := memberrulesettingservicelogic.NewDeleteMemberRuleSettingLogic(ctx, s.svcCtx)
	return l.DeleteMemberRuleSetting(in)
}

// 更新会员积分成长规则表
func (s *MemberRuleSettingServiceServer) UpdateMemberRuleSetting(ctx context.Context, in *umsclient.UpdateMemberRuleSettingReq) (*umsclient.UpdateMemberRuleSettingResp, error) {
	l := memberrulesettingservicelogic.NewUpdateMemberRuleSettingLogic(ctx, s.svcCtx)
	return l.UpdateMemberRuleSetting(in)
}

// 更新会员积分成长规则表状态
func (s *MemberRuleSettingServiceServer) UpdateMemberRuleSettingStatus(ctx context.Context, in *umsclient.UpdateMemberRuleSettingStatusReq) (*umsclient.UpdateMemberRuleSettingStatusResp, error) {
	l := memberrulesettingservicelogic.NewUpdateMemberRuleSettingStatusLogic(ctx, s.svcCtx)
	return l.UpdateMemberRuleSettingStatus(in)
}

// 查询会员积分成长规则表详情
func (s *MemberRuleSettingServiceServer) QueryMemberRuleSettingDetail(ctx context.Context, in *umsclient.QueryMemberRuleSettingDetailReq) (*umsclient.QueryMemberRuleSettingDetailResp, error) {
	l := memberrulesettingservicelogic.NewQueryMemberRuleSettingDetailLogic(ctx, s.svcCtx)
	return l.QueryMemberRuleSettingDetail(in)
}

// 查询会员积分成长规则表列表
func (s *MemberRuleSettingServiceServer) QueryMemberRuleSettingList(ctx context.Context, in *umsclient.QueryMemberRuleSettingListReq) (*umsclient.QueryMemberRuleSettingListResp, error) {
	l := memberrulesettingservicelogic.NewQueryMemberRuleSettingListLogic(ctx, s.svcCtx)
	return l.QueryMemberRuleSettingList(in)
}
