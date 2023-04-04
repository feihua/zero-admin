// Code generated by goctl. DO NOT EDIT.
// Source: ums.proto

package ums

import (
	"context"

	"zero-admin/rpc/ums/umsclient"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	GenerateTokenReq                        = umsclient.GenerateTokenReq
	GenerateTokenResp                       = umsclient.GenerateTokenResp
	GrowthChangeHistoryAddReq               = umsclient.GrowthChangeHistoryAddReq
	GrowthChangeHistoryAddResp              = umsclient.GrowthChangeHistoryAddResp
	GrowthChangeHistoryDeleteReq            = umsclient.GrowthChangeHistoryDeleteReq
	GrowthChangeHistoryDeleteResp           = umsclient.GrowthChangeHistoryDeleteResp
	GrowthChangeHistoryListData             = umsclient.GrowthChangeHistoryListData
	GrowthChangeHistoryListReq              = umsclient.GrowthChangeHistoryListReq
	GrowthChangeHistoryListResp             = umsclient.GrowthChangeHistoryListResp
	GrowthChangeHistoryUpdateReq            = umsclient.GrowthChangeHistoryUpdateReq
	GrowthChangeHistoryUpdateResp           = umsclient.GrowthChangeHistoryUpdateResp
	IntegrationChangeHistoryAddReq          = umsclient.IntegrationChangeHistoryAddReq
	IntegrationChangeHistoryAddResp         = umsclient.IntegrationChangeHistoryAddResp
	IntegrationChangeHistoryDeleteReq       = umsclient.IntegrationChangeHistoryDeleteReq
	IntegrationChangeHistoryDeleteResp      = umsclient.IntegrationChangeHistoryDeleteResp
	IntegrationChangeHistoryListData        = umsclient.IntegrationChangeHistoryListData
	IntegrationChangeHistoryListReq         = umsclient.IntegrationChangeHistoryListReq
	IntegrationChangeHistoryListResp        = umsclient.IntegrationChangeHistoryListResp
	IntegrationChangeHistoryUpdateReq       = umsclient.IntegrationChangeHistoryUpdateReq
	IntegrationChangeHistoryUpdateResp      = umsclient.IntegrationChangeHistoryUpdateResp
	IntegrationConsumeSettingAddReq         = umsclient.IntegrationConsumeSettingAddReq
	IntegrationConsumeSettingAddResp        = umsclient.IntegrationConsumeSettingAddResp
	IntegrationConsumeSettingDeleteReq      = umsclient.IntegrationConsumeSettingDeleteReq
	IntegrationConsumeSettingDeleteResp     = umsclient.IntegrationConsumeSettingDeleteResp
	IntegrationConsumeSettingListData       = umsclient.IntegrationConsumeSettingListData
	IntegrationConsumeSettingListReq        = umsclient.IntegrationConsumeSettingListReq
	IntegrationConsumeSettingListResp       = umsclient.IntegrationConsumeSettingListResp
	IntegrationConsumeSettingUpdateReq      = umsclient.IntegrationConsumeSettingUpdateReq
	IntegrationConsumeSettingUpdateResp     = umsclient.IntegrationConsumeSettingUpdateResp
	LoginReq                                = umsclient.LoginReq
	LoginResp                               = umsclient.LoginResp
	Member                                  = umsclient.Member
	MemberAddReq                            = umsclient.MemberAddReq
	MemberAddResp                           = umsclient.MemberAddResp
	MemberAuth                              = umsclient.MemberAuth
	MemberAuthByAuthKeyReq                  = umsclient.MemberAuthByAuthKeyReq
	MemberAuthByAuthKeyResp                 = umsclient.MemberAuthByAuthKeyResp
	MemberAuthByMemberIdReq                 = umsclient.MemberAuthByMemberIdReq
	MemberAuthyMemberIdResp                 = umsclient.MemberAuthyMemberIdResp
	MemberDeleteReq                         = umsclient.MemberDeleteReq
	MemberDeleteResp                        = umsclient.MemberDeleteResp
	MemberInfoReq                           = umsclient.MemberInfoReq
	MemberInfoResp                          = umsclient.MemberInfoResp
	MemberLevelAddReq                       = umsclient.MemberLevelAddReq
	MemberLevelAddResp                      = umsclient.MemberLevelAddResp
	MemberLevelDeleteReq                    = umsclient.MemberLevelDeleteReq
	MemberLevelDeleteResp                   = umsclient.MemberLevelDeleteResp
	MemberLevelListData                     = umsclient.MemberLevelListData
	MemberLevelListReq                      = umsclient.MemberLevelListReq
	MemberLevelListResp                     = umsclient.MemberLevelListResp
	MemberLevelUpdateReq                    = umsclient.MemberLevelUpdateReq
	MemberLevelUpdateResp                   = umsclient.MemberLevelUpdateResp
	MemberListData                          = umsclient.MemberListData
	MemberListReq                           = umsclient.MemberListReq
	MemberListResp                          = umsclient.MemberListResp
	MemberLoginLogAddReq                    = umsclient.MemberLoginLogAddReq
	MemberLoginLogAddResp                   = umsclient.MemberLoginLogAddResp
	MemberLoginLogDeleteReq                 = umsclient.MemberLoginLogDeleteReq
	MemberLoginLogDeleteResp                = umsclient.MemberLoginLogDeleteResp
	MemberLoginLogListData                  = umsclient.MemberLoginLogListData
	MemberLoginLogListReq                   = umsclient.MemberLoginLogListReq
	MemberLoginLogListResp                  = umsclient.MemberLoginLogListResp
	MemberLoginLogUpdateReq                 = umsclient.MemberLoginLogUpdateReq
	MemberLoginLogUpdateResp                = umsclient.MemberLoginLogUpdateResp
	MemberLoginReq                          = umsclient.MemberLoginReq
	MemberLoginResp                         = umsclient.MemberLoginResp
	MemberMemberTagRelationAddReq           = umsclient.MemberMemberTagRelationAddReq
	MemberMemberTagRelationAddResp          = umsclient.MemberMemberTagRelationAddResp
	MemberMemberTagRelationDeleteReq        = umsclient.MemberMemberTagRelationDeleteReq
	MemberMemberTagRelationDeleteResp       = umsclient.MemberMemberTagRelationDeleteResp
	MemberMemberTagRelationListData         = umsclient.MemberMemberTagRelationListData
	MemberMemberTagRelationListReq          = umsclient.MemberMemberTagRelationListReq
	MemberMemberTagRelationListResp         = umsclient.MemberMemberTagRelationListResp
	MemberMemberTagRelationUpdateReq        = umsclient.MemberMemberTagRelationUpdateReq
	MemberMemberTagRelationUpdateResp       = umsclient.MemberMemberTagRelationUpdateResp
	MemberPrepaidCardRelationAddReq         = umsclient.MemberPrepaidCardRelationAddReq
	MemberPrepaidCardRelationAddResp        = umsclient.MemberPrepaidCardRelationAddResp
	MemberPrepaidCardRelationDeleteReq      = umsclient.MemberPrepaidCardRelationDeleteReq
	MemberPrepaidCardRelationDeleteResp     = umsclient.MemberPrepaidCardRelationDeleteResp
	MemberPrepaidCardRelationListData       = umsclient.MemberPrepaidCardRelationListData
	MemberPrepaidCardRelationListReq        = umsclient.MemberPrepaidCardRelationListReq
	MemberPrepaidCardRelationListResp       = umsclient.MemberPrepaidCardRelationListResp
	MemberPrepaidCardRelationUpdateReq      = umsclient.MemberPrepaidCardRelationUpdateReq
	MemberPrepaidCardRelationUpdateResp     = umsclient.MemberPrepaidCardRelationUpdateResp
	MemberProductCategoryRelationAddReq     = umsclient.MemberProductCategoryRelationAddReq
	MemberProductCategoryRelationAddResp    = umsclient.MemberProductCategoryRelationAddResp
	MemberProductCategoryRelationDeleteReq  = umsclient.MemberProductCategoryRelationDeleteReq
	MemberProductCategoryRelationDeleteResp = umsclient.MemberProductCategoryRelationDeleteResp
	MemberProductCategoryRelationListData   = umsclient.MemberProductCategoryRelationListData
	MemberProductCategoryRelationListReq    = umsclient.MemberProductCategoryRelationListReq
	MemberProductCategoryRelationListResp   = umsclient.MemberProductCategoryRelationListResp
	MemberProductCategoryRelationUpdateReq  = umsclient.MemberProductCategoryRelationUpdateReq
	MemberProductCategoryRelationUpdateResp = umsclient.MemberProductCategoryRelationUpdateResp
	MemberReceiveAddressAddReq              = umsclient.MemberReceiveAddressAddReq
	MemberReceiveAddressAddResp             = umsclient.MemberReceiveAddressAddResp
	MemberReceiveAddressDeleteReq           = umsclient.MemberReceiveAddressDeleteReq
	MemberReceiveAddressDeleteResp          = umsclient.MemberReceiveAddressDeleteResp
	MemberReceiveAddressListData            = umsclient.MemberReceiveAddressListData
	MemberReceiveAddressListReq             = umsclient.MemberReceiveAddressListReq
	MemberReceiveAddressListResp            = umsclient.MemberReceiveAddressListResp
	MemberReceiveAddressQueryDetailReq      = umsclient.MemberReceiveAddressQueryDetailReq
	MemberReceiveAddressQueryDetailResp     = umsclient.MemberReceiveAddressQueryDetailResp
	MemberReceiveAddressUpdateReq           = umsclient.MemberReceiveAddressUpdateReq
	MemberReceiveAddressUpdateResp          = umsclient.MemberReceiveAddressUpdateResp
	MemberRuleSettingAddReq                 = umsclient.MemberRuleSettingAddReq
	MemberRuleSettingAddResp                = umsclient.MemberRuleSettingAddResp
	MemberRuleSettingDeleteReq              = umsclient.MemberRuleSettingDeleteReq
	MemberRuleSettingDeleteResp             = umsclient.MemberRuleSettingDeleteResp
	MemberRuleSettingListData               = umsclient.MemberRuleSettingListData
	MemberRuleSettingListReq                = umsclient.MemberRuleSettingListReq
	MemberRuleSettingListResp               = umsclient.MemberRuleSettingListResp
	MemberRuleSettingUpdateReq              = umsclient.MemberRuleSettingUpdateReq
	MemberRuleSettingUpdateResp             = umsclient.MemberRuleSettingUpdateResp
	MemberStatisticsInfoAddReq              = umsclient.MemberStatisticsInfoAddReq
	MemberStatisticsInfoAddResp             = umsclient.MemberStatisticsInfoAddResp
	MemberStatisticsInfoDeleteReq           = umsclient.MemberStatisticsInfoDeleteReq
	MemberStatisticsInfoDeleteResp          = umsclient.MemberStatisticsInfoDeleteResp
	MemberStatisticsInfoListData            = umsclient.MemberStatisticsInfoListData
	MemberStatisticsInfoListReq             = umsclient.MemberStatisticsInfoListReq
	MemberStatisticsInfoListResp            = umsclient.MemberStatisticsInfoListResp
	MemberStatisticsInfoUpdateReq           = umsclient.MemberStatisticsInfoUpdateReq
	MemberStatisticsInfoUpdateResp          = umsclient.MemberStatisticsInfoUpdateResp
	MemberTagAddReq                         = umsclient.MemberTagAddReq
	MemberTagAddResp                        = umsclient.MemberTagAddResp
	MemberTagDeleteReq                      = umsclient.MemberTagDeleteReq
	MemberTagDeleteResp                     = umsclient.MemberTagDeleteResp
	MemberTagListData                       = umsclient.MemberTagListData
	MemberTagListReq                        = umsclient.MemberTagListReq
	MemberTagListResp                       = umsclient.MemberTagListResp
	MemberTagUpdateReq                      = umsclient.MemberTagUpdateReq
	MemberTagUpdateResp                     = umsclient.MemberTagUpdateResp
	MemberTaskAddReq                        = umsclient.MemberTaskAddReq
	MemberTaskAddResp                       = umsclient.MemberTaskAddResp
	MemberTaskDeleteReq                     = umsclient.MemberTaskDeleteReq
	MemberTaskDeleteResp                    = umsclient.MemberTaskDeleteResp
	MemberTaskListData                      = umsclient.MemberTaskListData
	MemberTaskListReq                       = umsclient.MemberTaskListReq
	MemberTaskListResp                      = umsclient.MemberTaskListResp
	MemberTaskUpdateReq                     = umsclient.MemberTaskUpdateReq
	MemberTaskUpdateResp                    = umsclient.MemberTaskUpdateResp
	MemberUpdateReq                         = umsclient.MemberUpdateReq
	MemberUpdateResp                        = umsclient.MemberUpdateResp
	RegisterReq                             = umsclient.RegisterReq
	RegisterResp                            = umsclient.RegisterResp

	Ums interface {
		MemberAdd(ctx context.Context, in *MemberAddReq, opts ...grpc.CallOption) (*MemberAddResp, error)
		MemberLogin(ctx context.Context, in *MemberLoginReq, opts ...grpc.CallOption) (*MemberLoginResp, error)
		MemberList(ctx context.Context, in *MemberListReq, opts ...grpc.CallOption) (*MemberListResp, error)
		MemberUpdate(ctx context.Context, in *MemberUpdateReq, opts ...grpc.CallOption) (*MemberUpdateResp, error)
		MemberDelete(ctx context.Context, in *MemberDeleteReq, opts ...grpc.CallOption) (*MemberDeleteResp, error)
		Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error)
		Register(ctx context.Context, in *RegisterReq, opts ...grpc.CallOption) (*RegisterResp, error)
		MemberInfo(ctx context.Context, in *MemberInfoReq, opts ...grpc.CallOption) (*MemberInfoResp, error)
		MemberAuthByAuthKey(ctx context.Context, in *MemberAuthByAuthKeyReq, opts ...grpc.CallOption) (*MemberAuthByAuthKeyResp, error)
		MemberAuthByMemberId(ctx context.Context, in *MemberAuthByMemberIdReq, opts ...grpc.CallOption) (*MemberAuthyMemberIdResp, error)
		GenerateToken(ctx context.Context, in *GenerateTokenReq, opts ...grpc.CallOption) (*GenerateTokenResp, error)
		GrowthChangeHistoryAdd(ctx context.Context, in *GrowthChangeHistoryAddReq, opts ...grpc.CallOption) (*GrowthChangeHistoryAddResp, error)
		GrowthChangeHistoryList(ctx context.Context, in *GrowthChangeHistoryListReq, opts ...grpc.CallOption) (*GrowthChangeHistoryListResp, error)
		GrowthChangeHistoryUpdate(ctx context.Context, in *GrowthChangeHistoryUpdateReq, opts ...grpc.CallOption) (*GrowthChangeHistoryUpdateResp, error)
		GrowthChangeHistoryDelete(ctx context.Context, in *GrowthChangeHistoryDeleteReq, opts ...grpc.CallOption) (*GrowthChangeHistoryDeleteResp, error)
		IntegrationChangeHistoryAdd(ctx context.Context, in *IntegrationChangeHistoryAddReq, opts ...grpc.CallOption) (*IntegrationChangeHistoryAddResp, error)
		IntegrationChangeHistoryList(ctx context.Context, in *IntegrationChangeHistoryListReq, opts ...grpc.CallOption) (*IntegrationChangeHistoryListResp, error)
		IntegrationChangeHistoryUpdate(ctx context.Context, in *IntegrationChangeHistoryUpdateReq, opts ...grpc.CallOption) (*IntegrationChangeHistoryUpdateResp, error)
		IntegrationChangeHistoryDelete(ctx context.Context, in *IntegrationChangeHistoryDeleteReq, opts ...grpc.CallOption) (*IntegrationChangeHistoryDeleteResp, error)
		IntegrationConsumeSettingAdd(ctx context.Context, in *IntegrationConsumeSettingAddReq, opts ...grpc.CallOption) (*IntegrationConsumeSettingAddResp, error)
		IntegrationConsumeSettingList(ctx context.Context, in *IntegrationConsumeSettingListReq, opts ...grpc.CallOption) (*IntegrationConsumeSettingListResp, error)
		IntegrationConsumeSettingUpdate(ctx context.Context, in *IntegrationConsumeSettingUpdateReq, opts ...grpc.CallOption) (*IntegrationConsumeSettingUpdateResp, error)
		IntegrationConsumeSettingDelete(ctx context.Context, in *IntegrationConsumeSettingDeleteReq, opts ...grpc.CallOption) (*IntegrationConsumeSettingDeleteResp, error)
		MemberLevelAdd(ctx context.Context, in *MemberLevelAddReq, opts ...grpc.CallOption) (*MemberLevelAddResp, error)
		MemberLevelList(ctx context.Context, in *MemberLevelListReq, opts ...grpc.CallOption) (*MemberLevelListResp, error)
		MemberLevelUpdate(ctx context.Context, in *MemberLevelUpdateReq, opts ...grpc.CallOption) (*MemberLevelUpdateResp, error)
		MemberLevelDelete(ctx context.Context, in *MemberLevelDeleteReq, opts ...grpc.CallOption) (*MemberLevelDeleteResp, error)
		MemberLoginLogAdd(ctx context.Context, in *MemberLoginLogAddReq, opts ...grpc.CallOption) (*MemberLoginLogAddResp, error)
		MemberLoginLogList(ctx context.Context, in *MemberLoginLogListReq, opts ...grpc.CallOption) (*MemberLoginLogListResp, error)
		MemberLoginLogUpdate(ctx context.Context, in *MemberLoginLogUpdateReq, opts ...grpc.CallOption) (*MemberLoginLogUpdateResp, error)
		MemberLoginLogDelete(ctx context.Context, in *MemberLoginLogDeleteReq, opts ...grpc.CallOption) (*MemberLoginLogDeleteResp, error)
		MemberMemberTagRelationAdd(ctx context.Context, in *MemberMemberTagRelationAddReq, opts ...grpc.CallOption) (*MemberMemberTagRelationAddResp, error)
		MemberMemberTagRelationList(ctx context.Context, in *MemberMemberTagRelationListReq, opts ...grpc.CallOption) (*MemberMemberTagRelationListResp, error)
		MemberMemberTagRelationUpdate(ctx context.Context, in *MemberMemberTagRelationUpdateReq, opts ...grpc.CallOption) (*MemberMemberTagRelationUpdateResp, error)
		MemberMemberTagRelationDelete(ctx context.Context, in *MemberMemberTagRelationDeleteReq, opts ...grpc.CallOption) (*MemberMemberTagRelationDeleteResp, error)
		MemberPrepaidCardRelationAdd(ctx context.Context, in *MemberPrepaidCardRelationAddReq, opts ...grpc.CallOption) (*MemberPrepaidCardRelationAddResp, error)
		MemberPrepaidCardRelationList(ctx context.Context, in *MemberPrepaidCardRelationListReq, opts ...grpc.CallOption) (*MemberPrepaidCardRelationListResp, error)
		MemberPrepaidCardRelationUpdate(ctx context.Context, in *MemberPrepaidCardRelationUpdateReq, opts ...grpc.CallOption) (*MemberPrepaidCardRelationUpdateResp, error)
		MemberPrepaidCardRelationDelete(ctx context.Context, in *MemberPrepaidCardRelationDeleteReq, opts ...grpc.CallOption) (*MemberPrepaidCardRelationDeleteResp, error)
		MemberProductCategoryRelationAdd(ctx context.Context, in *MemberProductCategoryRelationAddReq, opts ...grpc.CallOption) (*MemberProductCategoryRelationAddResp, error)
		MemberProductCategoryRelationList(ctx context.Context, in *MemberProductCategoryRelationListReq, opts ...grpc.CallOption) (*MemberProductCategoryRelationListResp, error)
		MemberProductCategoryRelationUpdate(ctx context.Context, in *MemberProductCategoryRelationUpdateReq, opts ...grpc.CallOption) (*MemberProductCategoryRelationUpdateResp, error)
		MemberProductCategoryRelationDelete(ctx context.Context, in *MemberProductCategoryRelationDeleteReq, opts ...grpc.CallOption) (*MemberProductCategoryRelationDeleteResp, error)
		MemberReceiveAddressAdd(ctx context.Context, in *MemberReceiveAddressAddReq, opts ...grpc.CallOption) (*MemberReceiveAddressAddResp, error)
		MemberReceiveAddressList(ctx context.Context, in *MemberReceiveAddressListReq, opts ...grpc.CallOption) (*MemberReceiveAddressListResp, error)
		MemberReceiveAddressUpdate(ctx context.Context, in *MemberReceiveAddressUpdateReq, opts ...grpc.CallOption) (*MemberReceiveAddressUpdateResp, error)
		MemberReceiveAddressDelete(ctx context.Context, in *MemberReceiveAddressDeleteReq, opts ...grpc.CallOption) (*MemberReceiveAddressDeleteResp, error)
		MemberReceiveAddressQueryDetail(ctx context.Context, in *MemberReceiveAddressQueryDetailReq, opts ...grpc.CallOption) (*MemberReceiveAddressQueryDetailResp, error)
		MemberRuleSettingAdd(ctx context.Context, in *MemberRuleSettingAddReq, opts ...grpc.CallOption) (*MemberRuleSettingAddResp, error)
		MemberRuleSettingList(ctx context.Context, in *MemberRuleSettingListReq, opts ...grpc.CallOption) (*MemberRuleSettingListResp, error)
		MemberRuleSettingUpdate(ctx context.Context, in *MemberRuleSettingUpdateReq, opts ...grpc.CallOption) (*MemberRuleSettingUpdateResp, error)
		MemberRuleSettingDelete(ctx context.Context, in *MemberRuleSettingDeleteReq, opts ...grpc.CallOption) (*MemberRuleSettingDeleteResp, error)
		MemberStatisticsInfoAdd(ctx context.Context, in *MemberStatisticsInfoAddReq, opts ...grpc.CallOption) (*MemberStatisticsInfoAddResp, error)
		MemberStatisticsInfoList(ctx context.Context, in *MemberStatisticsInfoListReq, opts ...grpc.CallOption) (*MemberStatisticsInfoListResp, error)
		MemberStatisticsInfoUpdate(ctx context.Context, in *MemberStatisticsInfoUpdateReq, opts ...grpc.CallOption) (*MemberStatisticsInfoUpdateResp, error)
		MemberStatisticsInfoDelete(ctx context.Context, in *MemberStatisticsInfoDeleteReq, opts ...grpc.CallOption) (*MemberStatisticsInfoDeleteResp, error)
		MemberTagAdd(ctx context.Context, in *MemberTagAddReq, opts ...grpc.CallOption) (*MemberTagAddResp, error)
		MemberTagList(ctx context.Context, in *MemberTagListReq, opts ...grpc.CallOption) (*MemberTagListResp, error)
		MemberTagUpdate(ctx context.Context, in *MemberTagUpdateReq, opts ...grpc.CallOption) (*MemberTagUpdateResp, error)
		MemberTagDelete(ctx context.Context, in *MemberTagDeleteReq, opts ...grpc.CallOption) (*MemberTagDeleteResp, error)
		MemberTaskAdd(ctx context.Context, in *MemberTaskAddReq, opts ...grpc.CallOption) (*MemberTaskAddResp, error)
		MemberTaskList(ctx context.Context, in *MemberTaskListReq, opts ...grpc.CallOption) (*MemberTaskListResp, error)
		MemberTaskUpdate(ctx context.Context, in *MemberTaskUpdateReq, opts ...grpc.CallOption) (*MemberTaskUpdateResp, error)
		MemberTaskDelete(ctx context.Context, in *MemberTaskDeleteReq, opts ...grpc.CallOption) (*MemberTaskDeleteResp, error)
	}

	defaultUms struct {
		cli zrpc.Client
	}
)

func NewUms(cli zrpc.Client) Ums {
	return &defaultUms{
		cli: cli,
	}
}

func (m *defaultUms) MemberAdd(ctx context.Context, in *MemberAddReq, opts ...grpc.CallOption) (*MemberAddResp, error) {
	client := umsclient.NewUmsClient(m.cli.Conn())
	return client.MemberAdd(ctx, in, opts...)
}

func (m *defaultUms) MemberLogin(ctx context.Context, in *MemberLoginReq, opts ...grpc.CallOption) (*MemberLoginResp, error) {
	client := umsclient.NewUmsClient(m.cli.Conn())
	return client.MemberLogin(ctx, in, opts...)
}

func (m *defaultUms) MemberList(ctx context.Context, in *MemberListReq, opts ...grpc.CallOption) (*MemberListResp, error) {
	client := umsclient.NewUmsClient(m.cli.Conn())
	return client.MemberList(ctx, in, opts...)
}

func (m *defaultUms) MemberUpdate(ctx context.Context, in *MemberUpdateReq, opts ...grpc.CallOption) (*MemberUpdateResp, error) {
	client := umsclient.NewUmsClient(m.cli.Conn())
	return client.MemberUpdate(ctx, in, opts...)
}

func (m *defaultUms) MemberDelete(ctx context.Context, in *MemberDeleteReq, opts ...grpc.CallOption) (*MemberDeleteResp, error) {
	client := umsclient.NewUmsClient(m.cli.Conn())
	return client.MemberDelete(ctx, in, opts...)
}

func (m *defaultUms) Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error) {
	client := umsclient.NewUmsClient(m.cli.Conn())
	return client.Login(ctx, in, opts...)
}

func (m *defaultUms) Register(ctx context.Context, in *RegisterReq, opts ...grpc.CallOption) (*RegisterResp, error) {
	client := umsclient.NewUmsClient(m.cli.Conn())
	return client.Register(ctx, in, opts...)
}

func (m *defaultUms) MemberInfo(ctx context.Context, in *MemberInfoReq, opts ...grpc.CallOption) (*MemberInfoResp, error) {
	client := umsclient.NewUmsClient(m.cli.Conn())
	return client.MemberInfo(ctx, in, opts...)
}

func (m *defaultUms) MemberAuthByAuthKey(ctx context.Context, in *MemberAuthByAuthKeyReq, opts ...grpc.CallOption) (*MemberAuthByAuthKeyResp, error) {
	client := umsclient.NewUmsClient(m.cli.Conn())
	return client.MemberAuthByAuthKey(ctx, in, opts...)
}

func (m *defaultUms) MemberAuthByMemberId(ctx context.Context, in *MemberAuthByMemberIdReq, opts ...grpc.CallOption) (*MemberAuthyMemberIdResp, error) {
	client := umsclient.NewUmsClient(m.cli.Conn())
	return client.MemberAuthByMemberId(ctx, in, opts...)
}

func (m *defaultUms) GenerateToken(ctx context.Context, in *GenerateTokenReq, opts ...grpc.CallOption) (*GenerateTokenResp, error) {
	client := umsclient.NewUmsClient(m.cli.Conn())
	return client.GenerateToken(ctx, in, opts...)
}

func (m *defaultUms) GrowthChangeHistoryAdd(ctx context.Context, in *GrowthChangeHistoryAddReq, opts ...grpc.CallOption) (*GrowthChangeHistoryAddResp, error) {
	client := umsclient.NewUmsClient(m.cli.Conn())
	return client.GrowthChangeHistoryAdd(ctx, in, opts...)
}

func (m *defaultUms) GrowthChangeHistoryList(ctx context.Context, in *GrowthChangeHistoryListReq, opts ...grpc.CallOption) (*GrowthChangeHistoryListResp, error) {
	client := umsclient.NewUmsClient(m.cli.Conn())
	return client.GrowthChangeHistoryList(ctx, in, opts...)
}

func (m *defaultUms) GrowthChangeHistoryUpdate(ctx context.Context, in *GrowthChangeHistoryUpdateReq, opts ...grpc.CallOption) (*GrowthChangeHistoryUpdateResp, error) {
	client := umsclient.NewUmsClient(m.cli.Conn())
	return client.GrowthChangeHistoryUpdate(ctx, in, opts...)
}

func (m *defaultUms) GrowthChangeHistoryDelete(ctx context.Context, in *GrowthChangeHistoryDeleteReq, opts ...grpc.CallOption) (*GrowthChangeHistoryDeleteResp, error) {
	client := umsclient.NewUmsClient(m.cli.Conn())
	return client.GrowthChangeHistoryDelete(ctx, in, opts...)
}

func (m *defaultUms) IntegrationChangeHistoryAdd(ctx context.Context, in *IntegrationChangeHistoryAddReq, opts ...grpc.CallOption) (*IntegrationChangeHistoryAddResp, error) {
	client := umsclient.NewUmsClient(m.cli.Conn())
	return client.IntegrationChangeHistoryAdd(ctx, in, opts...)
}

func (m *defaultUms) IntegrationChangeHistoryList(ctx context.Context, in *IntegrationChangeHistoryListReq, opts ...grpc.CallOption) (*IntegrationChangeHistoryListResp, error) {
	client := umsclient.NewUmsClient(m.cli.Conn())
	return client.IntegrationChangeHistoryList(ctx, in, opts...)
}

func (m *defaultUms) IntegrationChangeHistoryUpdate(ctx context.Context, in *IntegrationChangeHistoryUpdateReq, opts ...grpc.CallOption) (*IntegrationChangeHistoryUpdateResp, error) {
	client := umsclient.NewUmsClient(m.cli.Conn())
	return client.IntegrationChangeHistoryUpdate(ctx, in, opts...)
}

func (m *defaultUms) IntegrationChangeHistoryDelete(ctx context.Context, in *IntegrationChangeHistoryDeleteReq, opts ...grpc.CallOption) (*IntegrationChangeHistoryDeleteResp, error) {
	client := umsclient.NewUmsClient(m.cli.Conn())
	return client.IntegrationChangeHistoryDelete(ctx, in, opts...)
}

func (m *defaultUms) IntegrationConsumeSettingAdd(ctx context.Context, in *IntegrationConsumeSettingAddReq, opts ...grpc.CallOption) (*IntegrationConsumeSettingAddResp, error) {
	client := umsclient.NewUmsClient(m.cli.Conn())
	return client.IntegrationConsumeSettingAdd(ctx, in, opts...)
}

func (m *defaultUms) IntegrationConsumeSettingList(ctx context.Context, in *IntegrationConsumeSettingListReq, opts ...grpc.CallOption) (*IntegrationConsumeSettingListResp, error) {
	client := umsclient.NewUmsClient(m.cli.Conn())
	return client.IntegrationConsumeSettingList(ctx, in, opts...)
}

func (m *defaultUms) IntegrationConsumeSettingUpdate(ctx context.Context, in *IntegrationConsumeSettingUpdateReq, opts ...grpc.CallOption) (*IntegrationConsumeSettingUpdateResp, error) {
	client := umsclient.NewUmsClient(m.cli.Conn())
	return client.IntegrationConsumeSettingUpdate(ctx, in, opts...)
}

func (m *defaultUms) IntegrationConsumeSettingDelete(ctx context.Context, in *IntegrationConsumeSettingDeleteReq, opts ...grpc.CallOption) (*IntegrationConsumeSettingDeleteResp, error) {
	client := umsclient.NewUmsClient(m.cli.Conn())
	return client.IntegrationConsumeSettingDelete(ctx, in, opts...)
}

func (m *defaultUms) MemberLevelAdd(ctx context.Context, in *MemberLevelAddReq, opts ...grpc.CallOption) (*MemberLevelAddResp, error) {
	client := umsclient.NewUmsClient(m.cli.Conn())
	return client.MemberLevelAdd(ctx, in, opts...)
}

func (m *defaultUms) MemberLevelList(ctx context.Context, in *MemberLevelListReq, opts ...grpc.CallOption) (*MemberLevelListResp, error) {
	client := umsclient.NewUmsClient(m.cli.Conn())
	return client.MemberLevelList(ctx, in, opts...)
}

func (m *defaultUms) MemberLevelUpdate(ctx context.Context, in *MemberLevelUpdateReq, opts ...grpc.CallOption) (*MemberLevelUpdateResp, error) {
	client := umsclient.NewUmsClient(m.cli.Conn())
	return client.MemberLevelUpdate(ctx, in, opts...)
}

func (m *defaultUms) MemberLevelDelete(ctx context.Context, in *MemberLevelDeleteReq, opts ...grpc.CallOption) (*MemberLevelDeleteResp, error) {
	client := umsclient.NewUmsClient(m.cli.Conn())
	return client.MemberLevelDelete(ctx, in, opts...)
}

func (m *defaultUms) MemberLoginLogAdd(ctx context.Context, in *MemberLoginLogAddReq, opts ...grpc.CallOption) (*MemberLoginLogAddResp, error) {
	client := umsclient.NewUmsClient(m.cli.Conn())
	return client.MemberLoginLogAdd(ctx, in, opts...)
}

func (m *defaultUms) MemberLoginLogList(ctx context.Context, in *MemberLoginLogListReq, opts ...grpc.CallOption) (*MemberLoginLogListResp, error) {
	client := umsclient.NewUmsClient(m.cli.Conn())
	return client.MemberLoginLogList(ctx, in, opts...)
}

func (m *defaultUms) MemberLoginLogUpdate(ctx context.Context, in *MemberLoginLogUpdateReq, opts ...grpc.CallOption) (*MemberLoginLogUpdateResp, error) {
	client := umsclient.NewUmsClient(m.cli.Conn())
	return client.MemberLoginLogUpdate(ctx, in, opts...)
}

func (m *defaultUms) MemberLoginLogDelete(ctx context.Context, in *MemberLoginLogDeleteReq, opts ...grpc.CallOption) (*MemberLoginLogDeleteResp, error) {
	client := umsclient.NewUmsClient(m.cli.Conn())
	return client.MemberLoginLogDelete(ctx, in, opts...)
}

func (m *defaultUms) MemberMemberTagRelationAdd(ctx context.Context, in *MemberMemberTagRelationAddReq, opts ...grpc.CallOption) (*MemberMemberTagRelationAddResp, error) {
	client := umsclient.NewUmsClient(m.cli.Conn())
	return client.MemberMemberTagRelationAdd(ctx, in, opts...)
}

func (m *defaultUms) MemberMemberTagRelationList(ctx context.Context, in *MemberMemberTagRelationListReq, opts ...grpc.CallOption) (*MemberMemberTagRelationListResp, error) {
	client := umsclient.NewUmsClient(m.cli.Conn())
	return client.MemberMemberTagRelationList(ctx, in, opts...)
}

func (m *defaultUms) MemberMemberTagRelationUpdate(ctx context.Context, in *MemberMemberTagRelationUpdateReq, opts ...grpc.CallOption) (*MemberMemberTagRelationUpdateResp, error) {
	client := umsclient.NewUmsClient(m.cli.Conn())
	return client.MemberMemberTagRelationUpdate(ctx, in, opts...)
}

func (m *defaultUms) MemberMemberTagRelationDelete(ctx context.Context, in *MemberMemberTagRelationDeleteReq, opts ...grpc.CallOption) (*MemberMemberTagRelationDeleteResp, error) {
	client := umsclient.NewUmsClient(m.cli.Conn())
	return client.MemberMemberTagRelationDelete(ctx, in, opts...)
}

func (m *defaultUms) MemberPrepaidCardRelationAdd(ctx context.Context, in *MemberPrepaidCardRelationAddReq, opts ...grpc.CallOption) (*MemberPrepaidCardRelationAddResp, error) {
	client := umsclient.NewUmsClient(m.cli.Conn())
	return client.MemberPrepaidCardRelationAdd(ctx, in, opts...)
}

func (m *defaultUms) MemberPrepaidCardRelationList(ctx context.Context, in *MemberPrepaidCardRelationListReq, opts ...grpc.CallOption) (*MemberPrepaidCardRelationListResp, error) {
	client := umsclient.NewUmsClient(m.cli.Conn())
	return client.MemberPrepaidCardRelationList(ctx, in, opts...)
}

func (m *defaultUms) MemberPrepaidCardRelationUpdate(ctx context.Context, in *MemberPrepaidCardRelationUpdateReq, opts ...grpc.CallOption) (*MemberPrepaidCardRelationUpdateResp, error) {
	client := umsclient.NewUmsClient(m.cli.Conn())
	return client.MemberPrepaidCardRelationUpdate(ctx, in, opts...)
}

func (m *defaultUms) MemberPrepaidCardRelationDelete(ctx context.Context, in *MemberPrepaidCardRelationDeleteReq, opts ...grpc.CallOption) (*MemberPrepaidCardRelationDeleteResp, error) {
	client := umsclient.NewUmsClient(m.cli.Conn())
	return client.MemberPrepaidCardRelationDelete(ctx, in, opts...)
}

func (m *defaultUms) MemberProductCategoryRelationAdd(ctx context.Context, in *MemberProductCategoryRelationAddReq, opts ...grpc.CallOption) (*MemberProductCategoryRelationAddResp, error) {
	client := umsclient.NewUmsClient(m.cli.Conn())
	return client.MemberProductCategoryRelationAdd(ctx, in, opts...)
}

func (m *defaultUms) MemberProductCategoryRelationList(ctx context.Context, in *MemberProductCategoryRelationListReq, opts ...grpc.CallOption) (*MemberProductCategoryRelationListResp, error) {
	client := umsclient.NewUmsClient(m.cli.Conn())
	return client.MemberProductCategoryRelationList(ctx, in, opts...)
}

func (m *defaultUms) MemberProductCategoryRelationUpdate(ctx context.Context, in *MemberProductCategoryRelationUpdateReq, opts ...grpc.CallOption) (*MemberProductCategoryRelationUpdateResp, error) {
	client := umsclient.NewUmsClient(m.cli.Conn())
	return client.MemberProductCategoryRelationUpdate(ctx, in, opts...)
}

func (m *defaultUms) MemberProductCategoryRelationDelete(ctx context.Context, in *MemberProductCategoryRelationDeleteReq, opts ...grpc.CallOption) (*MemberProductCategoryRelationDeleteResp, error) {
	client := umsclient.NewUmsClient(m.cli.Conn())
	return client.MemberProductCategoryRelationDelete(ctx, in, opts...)
}

func (m *defaultUms) MemberReceiveAddressAdd(ctx context.Context, in *MemberReceiveAddressAddReq, opts ...grpc.CallOption) (*MemberReceiveAddressAddResp, error) {
	client := umsclient.NewUmsClient(m.cli.Conn())
	return client.MemberReceiveAddressAdd(ctx, in, opts...)
}

func (m *defaultUms) MemberReceiveAddressList(ctx context.Context, in *MemberReceiveAddressListReq, opts ...grpc.CallOption) (*MemberReceiveAddressListResp, error) {
	client := umsclient.NewUmsClient(m.cli.Conn())
	return client.MemberReceiveAddressList(ctx, in, opts...)
}

func (m *defaultUms) MemberReceiveAddressUpdate(ctx context.Context, in *MemberReceiveAddressUpdateReq, opts ...grpc.CallOption) (*MemberReceiveAddressUpdateResp, error) {
	client := umsclient.NewUmsClient(m.cli.Conn())
	return client.MemberReceiveAddressUpdate(ctx, in, opts...)
}

func (m *defaultUms) MemberReceiveAddressDelete(ctx context.Context, in *MemberReceiveAddressDeleteReq, opts ...grpc.CallOption) (*MemberReceiveAddressDeleteResp, error) {
	client := umsclient.NewUmsClient(m.cli.Conn())
	return client.MemberReceiveAddressDelete(ctx, in, opts...)
}

func (m *defaultUms) MemberReceiveAddressQueryDetail(ctx context.Context, in *MemberReceiveAddressQueryDetailReq, opts ...grpc.CallOption) (*MemberReceiveAddressQueryDetailResp, error) {
	client := umsclient.NewUmsClient(m.cli.Conn())
	return client.MemberReceiveAddressQueryDetail(ctx, in, opts...)
}

func (m *defaultUms) MemberRuleSettingAdd(ctx context.Context, in *MemberRuleSettingAddReq, opts ...grpc.CallOption) (*MemberRuleSettingAddResp, error) {
	client := umsclient.NewUmsClient(m.cli.Conn())
	return client.MemberRuleSettingAdd(ctx, in, opts...)
}

func (m *defaultUms) MemberRuleSettingList(ctx context.Context, in *MemberRuleSettingListReq, opts ...grpc.CallOption) (*MemberRuleSettingListResp, error) {
	client := umsclient.NewUmsClient(m.cli.Conn())
	return client.MemberRuleSettingList(ctx, in, opts...)
}

func (m *defaultUms) MemberRuleSettingUpdate(ctx context.Context, in *MemberRuleSettingUpdateReq, opts ...grpc.CallOption) (*MemberRuleSettingUpdateResp, error) {
	client := umsclient.NewUmsClient(m.cli.Conn())
	return client.MemberRuleSettingUpdate(ctx, in, opts...)
}

func (m *defaultUms) MemberRuleSettingDelete(ctx context.Context, in *MemberRuleSettingDeleteReq, opts ...grpc.CallOption) (*MemberRuleSettingDeleteResp, error) {
	client := umsclient.NewUmsClient(m.cli.Conn())
	return client.MemberRuleSettingDelete(ctx, in, opts...)
}

func (m *defaultUms) MemberStatisticsInfoAdd(ctx context.Context, in *MemberStatisticsInfoAddReq, opts ...grpc.CallOption) (*MemberStatisticsInfoAddResp, error) {
	client := umsclient.NewUmsClient(m.cli.Conn())
	return client.MemberStatisticsInfoAdd(ctx, in, opts...)
}

func (m *defaultUms) MemberStatisticsInfoList(ctx context.Context, in *MemberStatisticsInfoListReq, opts ...grpc.CallOption) (*MemberStatisticsInfoListResp, error) {
	client := umsclient.NewUmsClient(m.cli.Conn())
	return client.MemberStatisticsInfoList(ctx, in, opts...)
}

func (m *defaultUms) MemberStatisticsInfoUpdate(ctx context.Context, in *MemberStatisticsInfoUpdateReq, opts ...grpc.CallOption) (*MemberStatisticsInfoUpdateResp, error) {
	client := umsclient.NewUmsClient(m.cli.Conn())
	return client.MemberStatisticsInfoUpdate(ctx, in, opts...)
}

func (m *defaultUms) MemberStatisticsInfoDelete(ctx context.Context, in *MemberStatisticsInfoDeleteReq, opts ...grpc.CallOption) (*MemberStatisticsInfoDeleteResp, error) {
	client := umsclient.NewUmsClient(m.cli.Conn())
	return client.MemberStatisticsInfoDelete(ctx, in, opts...)
}

func (m *defaultUms) MemberTagAdd(ctx context.Context, in *MemberTagAddReq, opts ...grpc.CallOption) (*MemberTagAddResp, error) {
	client := umsclient.NewUmsClient(m.cli.Conn())
	return client.MemberTagAdd(ctx, in, opts...)
}

func (m *defaultUms) MemberTagList(ctx context.Context, in *MemberTagListReq, opts ...grpc.CallOption) (*MemberTagListResp, error) {
	client := umsclient.NewUmsClient(m.cli.Conn())
	return client.MemberTagList(ctx, in, opts...)
}

func (m *defaultUms) MemberTagUpdate(ctx context.Context, in *MemberTagUpdateReq, opts ...grpc.CallOption) (*MemberTagUpdateResp, error) {
	client := umsclient.NewUmsClient(m.cli.Conn())
	return client.MemberTagUpdate(ctx, in, opts...)
}

func (m *defaultUms) MemberTagDelete(ctx context.Context, in *MemberTagDeleteReq, opts ...grpc.CallOption) (*MemberTagDeleteResp, error) {
	client := umsclient.NewUmsClient(m.cli.Conn())
	return client.MemberTagDelete(ctx, in, opts...)
}

func (m *defaultUms) MemberTaskAdd(ctx context.Context, in *MemberTaskAddReq, opts ...grpc.CallOption) (*MemberTaskAddResp, error) {
	client := umsclient.NewUmsClient(m.cli.Conn())
	return client.MemberTaskAdd(ctx, in, opts...)
}

func (m *defaultUms) MemberTaskList(ctx context.Context, in *MemberTaskListReq, opts ...grpc.CallOption) (*MemberTaskListResp, error) {
	client := umsclient.NewUmsClient(m.cli.Conn())
	return client.MemberTaskList(ctx, in, opts...)
}

func (m *defaultUms) MemberTaskUpdate(ctx context.Context, in *MemberTaskUpdateReq, opts ...grpc.CallOption) (*MemberTaskUpdateResp, error) {
	client := umsclient.NewUmsClient(m.cli.Conn())
	return client.MemberTaskUpdate(ctx, in, opts...)
}

func (m *defaultUms) MemberTaskDelete(ctx context.Context, in *MemberTaskDeleteReq, opts ...grpc.CallOption) (*MemberTaskDeleteResp, error) {
	client := umsclient.NewUmsClient(m.cli.Conn())
	return client.MemberTaskDelete(ctx, in, opts...)
}
