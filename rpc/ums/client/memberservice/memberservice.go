// Code generated by goctl. DO NOT EDIT.
// goctl 1.8.1
// Source: ums.proto

package memberservice

import (
	"context"

	"github.com/feihua/zero-admin/rpc/ums/umsclient"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	AddGrowthChangeHistoryReq                  = umsclient.AddGrowthChangeHistoryReq
	AddGrowthChangeHistoryResp                 = umsclient.AddGrowthChangeHistoryResp
	AddIntegrationChangeHistoryReq             = umsclient.AddIntegrationChangeHistoryReq
	AddIntegrationChangeHistoryResp            = umsclient.AddIntegrationChangeHistoryResp
	AddIntegrationConsumeSettingReq            = umsclient.AddIntegrationConsumeSettingReq
	AddIntegrationConsumeSettingResp           = umsclient.AddIntegrationConsumeSettingResp
	AddMemberBrandAttentionReq                 = umsclient.AddMemberBrandAttentionReq
	AddMemberBrandAttentionResp                = umsclient.AddMemberBrandAttentionResp
	AddMemberLevelReq                          = umsclient.AddMemberLevelReq
	AddMemberLevelResp                         = umsclient.AddMemberLevelResp
	AddMemberMemberTagRelationReq              = umsclient.AddMemberMemberTagRelationReq
	AddMemberMemberTagRelationResp             = umsclient.AddMemberMemberTagRelationResp
	AddMemberProductCategoryRelationReq        = umsclient.AddMemberProductCategoryRelationReq
	AddMemberProductCategoryRelationResp       = umsclient.AddMemberProductCategoryRelationResp
	AddMemberProductCollectionReq              = umsclient.AddMemberProductCollectionReq
	AddMemberProductCollectionResp             = umsclient.AddMemberProductCollectionResp
	AddMemberReadHistoryReq                    = umsclient.AddMemberReadHistoryReq
	AddMemberReadHistoryResp                   = umsclient.AddMemberReadHistoryResp
	AddMemberReceiveAddressReq                 = umsclient.AddMemberReceiveAddressReq
	AddMemberReceiveAddressResp                = umsclient.AddMemberReceiveAddressResp
	AddMemberReq                               = umsclient.AddMemberReq
	AddMemberResp                              = umsclient.AddMemberResp
	AddMemberRuleSettingReq                    = umsclient.AddMemberRuleSettingReq
	AddMemberRuleSettingResp                   = umsclient.AddMemberRuleSettingResp
	AddMemberStatisticsInfoReq                 = umsclient.AddMemberStatisticsInfoReq
	AddMemberStatisticsInfoResp                = umsclient.AddMemberStatisticsInfoResp
	AddMemberTagReq                            = umsclient.AddMemberTagReq
	AddMemberTagResp                           = umsclient.AddMemberTagResp
	AddMemberTaskReq                           = umsclient.AddMemberTaskReq
	AddMemberTaskResp                          = umsclient.AddMemberTaskResp
	DeleteGrowthChangeHistoryReq               = umsclient.DeleteGrowthChangeHistoryReq
	DeleteGrowthChangeHistoryResp              = umsclient.DeleteGrowthChangeHistoryResp
	DeleteIntegrationChangeHistoryReq          = umsclient.DeleteIntegrationChangeHistoryReq
	DeleteIntegrationChangeHistoryResp         = umsclient.DeleteIntegrationChangeHistoryResp
	DeleteIntegrationConsumeSettingReq         = umsclient.DeleteIntegrationConsumeSettingReq
	DeleteIntegrationConsumeSettingResp        = umsclient.DeleteIntegrationConsumeSettingResp
	DeleteMemberBrandAttentionReq              = umsclient.DeleteMemberBrandAttentionReq
	DeleteMemberBrandAttentionResp             = umsclient.DeleteMemberBrandAttentionResp
	DeleteMemberLevelReq                       = umsclient.DeleteMemberLevelReq
	DeleteMemberLevelResp                      = umsclient.DeleteMemberLevelResp
	DeleteMemberLoginLogReq                    = umsclient.DeleteMemberLoginLogReq
	DeleteMemberLoginLogResp                   = umsclient.DeleteMemberLoginLogResp
	DeleteMemberMemberTagRelationReq           = umsclient.DeleteMemberMemberTagRelationReq
	DeleteMemberMemberTagRelationResp          = umsclient.DeleteMemberMemberTagRelationResp
	DeleteMemberProductCollectionReq           = umsclient.DeleteMemberProductCollectionReq
	DeleteMemberProductCollectionResp          = umsclient.DeleteMemberProductCollectionResp
	DeleteMemberReadHistoryReq                 = umsclient.DeleteMemberReadHistoryReq
	DeleteMemberReadHistoryResp                = umsclient.DeleteMemberReadHistoryResp
	DeleteMemberReceiveAddressReq              = umsclient.DeleteMemberReceiveAddressReq
	DeleteMemberReceiveAddressResp             = umsclient.DeleteMemberReceiveAddressResp
	DeleteMemberReq                            = umsclient.DeleteMemberReq
	DeleteMemberResp                           = umsclient.DeleteMemberResp
	DeleteMemberRuleSettingReq                 = umsclient.DeleteMemberRuleSettingReq
	DeleteMemberRuleSettingResp                = umsclient.DeleteMemberRuleSettingResp
	DeleteMemberTagReq                         = umsclient.DeleteMemberTagReq
	DeleteMemberTagResp                        = umsclient.DeleteMemberTagResp
	DeleteMemberTaskReq                        = umsclient.DeleteMemberTaskReq
	DeleteMemberTaskResp                       = umsclient.DeleteMemberTaskResp
	GrowthChangeHistoryListData                = umsclient.GrowthChangeHistoryListData
	IntegrationChangeHistoryListData           = umsclient.IntegrationChangeHistoryListData
	IntegrationConsumeSettingListData          = umsclient.IntegrationConsumeSettingListData
	MemberBrandAttentionListData               = umsclient.MemberBrandAttentionListData
	MemberLevelListData                        = umsclient.MemberLevelListData
	MemberListData                             = umsclient.MemberListData
	MemberLoginLogListData                     = umsclient.MemberLoginLogListData
	MemberLoginReq                             = umsclient.MemberLoginReq
	MemberLoginResp                            = umsclient.MemberLoginResp
	MemberMemberTagRelationListData            = umsclient.MemberMemberTagRelationListData
	MemberProductCategoryRelationListData      = umsclient.MemberProductCategoryRelationListData
	MemberProductCollectionListData            = umsclient.MemberProductCollectionListData
	MemberReadHistoryListData                  = umsclient.MemberReadHistoryListData
	MemberReceiveAddressListData               = umsclient.MemberReceiveAddressListData
	MemberRuleSettingListData                  = umsclient.MemberRuleSettingListData
	MemberTagListData                          = umsclient.MemberTagListData
	MemberTaskListData                         = umsclient.MemberTaskListData
	QueryGrowthChangeHistoryDetailReq          = umsclient.QueryGrowthChangeHistoryDetailReq
	QueryGrowthChangeHistoryDetailResp         = umsclient.QueryGrowthChangeHistoryDetailResp
	QueryGrowthChangeHistoryListReq            = umsclient.QueryGrowthChangeHistoryListReq
	QueryGrowthChangeHistoryListResp           = umsclient.QueryGrowthChangeHistoryListResp
	QueryIntegrationChangeHistoryDetailReq     = umsclient.QueryIntegrationChangeHistoryDetailReq
	QueryIntegrationChangeHistoryDetailResp    = umsclient.QueryIntegrationChangeHistoryDetailResp
	QueryIntegrationChangeHistoryListReq       = umsclient.QueryIntegrationChangeHistoryListReq
	QueryIntegrationChangeHistoryListResp      = umsclient.QueryIntegrationChangeHistoryListResp
	QueryIntegrationConsumeSettingDetailReq    = umsclient.QueryIntegrationConsumeSettingDetailReq
	QueryIntegrationConsumeSettingDetailResp   = umsclient.QueryIntegrationConsumeSettingDetailResp
	QueryIntegrationConsumeSettingListReq      = umsclient.QueryIntegrationConsumeSettingListReq
	QueryIntegrationConsumeSettingListResp     = umsclient.QueryIntegrationConsumeSettingListResp
	QueryMemberBrandAttentionDetailReq         = umsclient.QueryMemberBrandAttentionDetailReq
	QueryMemberBrandAttentionDetailResp        = umsclient.QueryMemberBrandAttentionDetailResp
	QueryMemberBrandAttentionListReq           = umsclient.QueryMemberBrandAttentionListReq
	QueryMemberBrandAttentionListResp          = umsclient.QueryMemberBrandAttentionListResp
	QueryMemberDetailReq                       = umsclient.QueryMemberDetailReq
	QueryMemberDetailResp                      = umsclient.QueryMemberDetailResp
	QueryMemberLevelDetailReq                  = umsclient.QueryMemberLevelDetailReq
	QueryMemberLevelDetailResp                 = umsclient.QueryMemberLevelDetailResp
	QueryMemberLevelListReq                    = umsclient.QueryMemberLevelListReq
	QueryMemberLevelListResp                   = umsclient.QueryMemberLevelListResp
	QueryMemberListReq                         = umsclient.QueryMemberListReq
	QueryMemberListResp                        = umsclient.QueryMemberListResp
	QueryMemberLoginLogListReq                 = umsclient.QueryMemberLoginLogListReq
	QueryMemberLoginLogListResp                = umsclient.QueryMemberLoginLogListResp
	QueryMemberMemberTagRelationListReq        = umsclient.QueryMemberMemberTagRelationListReq
	QueryMemberMemberTagRelationListResp       = umsclient.QueryMemberMemberTagRelationListResp
	QueryMemberProductCategoryRelationListReq  = umsclient.QueryMemberProductCategoryRelationListReq
	QueryMemberProductCategoryRelationListResp = umsclient.QueryMemberProductCategoryRelationListResp
	QueryMemberProductCollectionDetailReq      = umsclient.QueryMemberProductCollectionDetailReq
	QueryMemberProductCollectionDetailResp     = umsclient.QueryMemberProductCollectionDetailResp
	QueryMemberProductCollectionListReq        = umsclient.QueryMemberProductCollectionListReq
	QueryMemberProductCollectionListResp       = umsclient.QueryMemberProductCollectionListResp
	QueryMemberReadHistoryDetailReq            = umsclient.QueryMemberReadHistoryDetailReq
	QueryMemberReadHistoryDetailResp           = umsclient.QueryMemberReadHistoryDetailResp
	QueryMemberReadHistoryListReq              = umsclient.QueryMemberReadHistoryListReq
	QueryMemberReadHistoryListResp             = umsclient.QueryMemberReadHistoryListResp
	QueryMemberReceiveAddressDetailReq         = umsclient.QueryMemberReceiveAddressDetailReq
	QueryMemberReceiveAddressDetailResp        = umsclient.QueryMemberReceiveAddressDetailResp
	QueryMemberReceiveAddressListReq           = umsclient.QueryMemberReceiveAddressListReq
	QueryMemberReceiveAddressListResp          = umsclient.QueryMemberReceiveAddressListResp
	QueryMemberRuleSettingDetailReq            = umsclient.QueryMemberRuleSettingDetailReq
	QueryMemberRuleSettingDetailResp           = umsclient.QueryMemberRuleSettingDetailResp
	QueryMemberRuleSettingListReq              = umsclient.QueryMemberRuleSettingListReq
	QueryMemberRuleSettingListResp             = umsclient.QueryMemberRuleSettingListResp
	QueryMemberStatisticsInfoDetailReq         = umsclient.QueryMemberStatisticsInfoDetailReq
	QueryMemberStatisticsInfoDetailResp        = umsclient.QueryMemberStatisticsInfoDetailResp
	QueryMemberStatisticsInfoListData          = umsclient.QueryMemberStatisticsInfoListData
	QueryMemberStatisticsInfoListReq           = umsclient.QueryMemberStatisticsInfoListReq
	QueryMemberStatisticsInfoListResp          = umsclient.QueryMemberStatisticsInfoListResp
	QueryMemberTagDetailReq                    = umsclient.QueryMemberTagDetailReq
	QueryMemberTagDetailResp                   = umsclient.QueryMemberTagDetailResp
	QueryMemberTagListReq                      = umsclient.QueryMemberTagListReq
	QueryMemberTagListResp                     = umsclient.QueryMemberTagListResp
	QueryMemberTaskDetailReq                   = umsclient.QueryMemberTaskDetailReq
	QueryMemberTaskDetailResp                  = umsclient.QueryMemberTaskDetailResp
	QueryMemberTaskListReq                     = umsclient.QueryMemberTaskListReq
	QueryMemberTaskListResp                    = umsclient.QueryMemberTaskListResp
	UpdateIntegrationConsumeSettingReq         = umsclient.UpdateIntegrationConsumeSettingReq
	UpdateIntegrationConsumeSettingResp        = umsclient.UpdateIntegrationConsumeSettingResp
	UpdateIntegrationConsumeSettingStatusReq   = umsclient.UpdateIntegrationConsumeSettingStatusReq
	UpdateIntegrationConsumeSettingStatusResp  = umsclient.UpdateIntegrationConsumeSettingStatusResp
	UpdateMemberIntegrationReq                 = umsclient.UpdateMemberIntegrationReq
	UpdateMemberIntegrationResp                = umsclient.UpdateMemberIntegrationResp
	UpdateMemberLevelReq                       = umsclient.UpdateMemberLevelReq
	UpdateMemberLevelResp                      = umsclient.UpdateMemberLevelResp
	UpdateMemberLevelStatusReq                 = umsclient.UpdateMemberLevelStatusReq
	UpdateMemberLevelStatusResp                = umsclient.UpdateMemberLevelStatusResp
	UpdateMemberReceiveAddressReq              = umsclient.UpdateMemberReceiveAddressReq
	UpdateMemberReceiveAddressResp             = umsclient.UpdateMemberReceiveAddressResp
	UpdateMemberReceiveAddressStatusReq        = umsclient.UpdateMemberReceiveAddressStatusReq
	UpdateMemberReceiveAddressStatusResp       = umsclient.UpdateMemberReceiveAddressStatusResp
	UpdateMemberReq                            = umsclient.UpdateMemberReq
	UpdateMemberResp                           = umsclient.UpdateMemberResp
	UpdateMemberRuleSettingReq                 = umsclient.UpdateMemberRuleSettingReq
	UpdateMemberRuleSettingResp                = umsclient.UpdateMemberRuleSettingResp
	UpdateMemberRuleSettingStatusReq           = umsclient.UpdateMemberRuleSettingStatusReq
	UpdateMemberRuleSettingStatusResp          = umsclient.UpdateMemberRuleSettingStatusResp
	UpdateMemberStatusReq                      = umsclient.UpdateMemberStatusReq
	UpdateMemberStatusResp                     = umsclient.UpdateMemberStatusResp
	UpdateMemberTagReq                         = umsclient.UpdateMemberTagReq
	UpdateMemberTagResp                        = umsclient.UpdateMemberTagResp
	UpdateMemberTagStatusReq                   = umsclient.UpdateMemberTagStatusReq
	UpdateMemberTagStatusResp                  = umsclient.UpdateMemberTagStatusResp
	UpdateMemberTaskReq                        = umsclient.UpdateMemberTaskReq
	UpdateMemberTaskResp                       = umsclient.UpdateMemberTaskResp
	UpdateMemberTaskStatusReq                  = umsclient.UpdateMemberTaskStatusReq
	UpdateMemberTaskStatusResp                 = umsclient.UpdateMemberTaskStatusResp

	MemberService interface {
		// 添加会员表
		AddMember(ctx context.Context, in *AddMemberReq, opts ...grpc.CallOption) (*AddMemberResp, error)
		// 删除会员表
		DeleteMember(ctx context.Context, in *DeleteMemberReq, opts ...grpc.CallOption) (*DeleteMemberResp, error)
		// 更新会员表
		UpdateMember(ctx context.Context, in *UpdateMemberReq, opts ...grpc.CallOption) (*UpdateMemberResp, error)
		// 更新会员表状态
		UpdateMemberStatus(ctx context.Context, in *UpdateMemberStatusReq, opts ...grpc.CallOption) (*UpdateMemberStatusResp, error)
		// 查询会员表详情
		QueryMemberDetail(ctx context.Context, in *QueryMemberDetailReq, opts ...grpc.CallOption) (*QueryMemberDetailResp, error)
		// 查询会员表列表
		QueryMemberList(ctx context.Context, in *QueryMemberListReq, opts ...grpc.CallOption) (*QueryMemberListResp, error)
		// 会员登录
		MemberLogin(ctx context.Context, in *MemberLoginReq, opts ...grpc.CallOption) (*MemberLoginResp, error)
		// 更新会员积分
		UpdateMemberIntegration(ctx context.Context, in *UpdateMemberIntegrationReq, opts ...grpc.CallOption) (*UpdateMemberIntegrationResp, error)
	}

	defaultMemberService struct {
		cli zrpc.Client
	}
)

func NewMemberService(cli zrpc.Client) MemberService {
	return &defaultMemberService{
		cli: cli,
	}
}

// 添加会员表
func (m *defaultMemberService) AddMember(ctx context.Context, in *AddMemberReq, opts ...grpc.CallOption) (*AddMemberResp, error) {
	client := umsclient.NewMemberServiceClient(m.cli.Conn())
	return client.AddMember(ctx, in, opts...)
}

// 删除会员表
func (m *defaultMemberService) DeleteMember(ctx context.Context, in *DeleteMemberReq, opts ...grpc.CallOption) (*DeleteMemberResp, error) {
	client := umsclient.NewMemberServiceClient(m.cli.Conn())
	return client.DeleteMember(ctx, in, opts...)
}

// 更新会员表
func (m *defaultMemberService) UpdateMember(ctx context.Context, in *UpdateMemberReq, opts ...grpc.CallOption) (*UpdateMemberResp, error) {
	client := umsclient.NewMemberServiceClient(m.cli.Conn())
	return client.UpdateMember(ctx, in, opts...)
}

// 更新会员表状态
func (m *defaultMemberService) UpdateMemberStatus(ctx context.Context, in *UpdateMemberStatusReq, opts ...grpc.CallOption) (*UpdateMemberStatusResp, error) {
	client := umsclient.NewMemberServiceClient(m.cli.Conn())
	return client.UpdateMemberStatus(ctx, in, opts...)
}

// 查询会员表详情
func (m *defaultMemberService) QueryMemberDetail(ctx context.Context, in *QueryMemberDetailReq, opts ...grpc.CallOption) (*QueryMemberDetailResp, error) {
	client := umsclient.NewMemberServiceClient(m.cli.Conn())
	return client.QueryMemberDetail(ctx, in, opts...)
}

// 查询会员表列表
func (m *defaultMemberService) QueryMemberList(ctx context.Context, in *QueryMemberListReq, opts ...grpc.CallOption) (*QueryMemberListResp, error) {
	client := umsclient.NewMemberServiceClient(m.cli.Conn())
	return client.QueryMemberList(ctx, in, opts...)
}

// 会员登录
func (m *defaultMemberService) MemberLogin(ctx context.Context, in *MemberLoginReq, opts ...grpc.CallOption) (*MemberLoginResp, error) {
	client := umsclient.NewMemberServiceClient(m.cli.Conn())
	return client.MemberLogin(ctx, in, opts...)
}

// 更新会员积分
func (m *defaultMemberService) UpdateMemberIntegration(ctx context.Context, in *UpdateMemberIntegrationReq, opts ...grpc.CallOption) (*UpdateMemberIntegrationResp, error) {
	client := umsclient.NewMemberServiceClient(m.cli.Conn())
	return client.UpdateMemberIntegration(ctx, in, opts...)
}
