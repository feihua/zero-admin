// Code generated by goctl. DO NOT EDIT.
// goctl 1.8.1
// Source: sys.proto

package dicttypeservice

import (
	"context"

	"github.com/feihua/zero-admin/rpc/sys/sysclient"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	AddDeptReq                = sysclient.AddDeptReq
	AddDeptResp               = sysclient.AddDeptResp
	AddDictItemReq            = sysclient.AddDictItemReq
	AddDictItemResp           = sysclient.AddDictItemResp
	AddDictTypeReq            = sysclient.AddDictTypeReq
	AddDictTypeResp           = sysclient.AddDictTypeResp
	AddMenuReq                = sysclient.AddMenuReq
	AddMenuResp               = sysclient.AddMenuResp
	AddOperateLogReq          = sysclient.AddOperateLogReq
	AddOperateLogResp         = sysclient.AddOperateLogResp
	AddPostReq                = sysclient.AddPostReq
	AddPostResp               = sysclient.AddPostResp
	AddRoleReq                = sysclient.AddRoleReq
	AddRoleResp               = sysclient.AddRoleResp
	AddUserReq                = sysclient.AddUserReq
	AddUserResp               = sysclient.AddUserResp
	CancelAuthorizationReq    = sysclient.CancelAuthorizationReq
	CancelAuthorizationResp   = sysclient.CancelAuthorizationResp
	DeleteDeptReq             = sysclient.DeleteDeptReq
	DeleteDeptResp            = sysclient.DeleteDeptResp
	DeleteDictItemReq         = sysclient.DeleteDictItemReq
	DeleteDictItemResp        = sysclient.DeleteDictItemResp
	DeleteDictTypeReq         = sysclient.DeleteDictTypeReq
	DeleteDictTypeResp        = sysclient.DeleteDictTypeResp
	DeleteLoginLogReq         = sysclient.DeleteLoginLogReq
	DeleteLoginLogResp        = sysclient.DeleteLoginLogResp
	DeleteMenuReq             = sysclient.DeleteMenuReq
	DeleteMenuResp            = sysclient.DeleteMenuResp
	DeleteOperateLogReq       = sysclient.DeleteOperateLogReq
	DeleteOperateLogResp      = sysclient.DeleteOperateLogResp
	DeletePostReq             = sysclient.DeletePostReq
	DeletePostResp            = sysclient.DeletePostResp
	DeleteRoleReq             = sysclient.DeleteRoleReq
	DeleteRoleResp            = sysclient.DeleteRoleResp
	DeleteUserReq             = sysclient.DeleteUserReq
	DeleteUserResp            = sysclient.DeleteUserResp
	DeptData                  = sysclient.DeptData
	DeptListData              = sysclient.DeptListData
	DictItemListData          = sysclient.DictItemListData
	DictTypeListData          = sysclient.DictTypeListData
	InfoReq                   = sysclient.InfoReq
	InfoResp                  = sysclient.InfoResp
	LoginLogListData          = sysclient.LoginLogListData
	LoginReq                  = sysclient.LoginReq
	LoginResp                 = sysclient.LoginResp
	MenuData                  = sysclient.MenuData
	MenuListData              = sysclient.MenuListData
	MenuListTree              = sysclient.MenuListTree
	OperateLogListData        = sysclient.OperateLogListData
	PostData                  = sysclient.PostData
	PostListData              = sysclient.PostListData
	QueryDeptAndPostListReq   = sysclient.QueryDeptAndPostListReq
	QueryDeptAndPostListResp  = sysclient.QueryDeptAndPostListResp
	QueryDeptDetailReq        = sysclient.QueryDeptDetailReq
	QueryDeptDetailResp       = sysclient.QueryDeptDetailResp
	QueryDeptListReq          = sysclient.QueryDeptListReq
	QueryDeptListResp         = sysclient.QueryDeptListResp
	QueryDictItemDetailReq    = sysclient.QueryDictItemDetailReq
	QueryDictItemDetailResp   = sysclient.QueryDictItemDetailResp
	QueryDictItemListReq      = sysclient.QueryDictItemListReq
	QueryDictItemListResp     = sysclient.QueryDictItemListResp
	QueryDictTypeDetailReq    = sysclient.QueryDictTypeDetailReq
	QueryDictTypeDetailResp   = sysclient.QueryDictTypeDetailResp
	QueryDictTypeListReq      = sysclient.QueryDictTypeListReq
	QueryDictTypeListResp     = sysclient.QueryDictTypeListResp
	QueryLoginLogDetailReq    = sysclient.QueryLoginLogDetailReq
	QueryLoginLogDetailResp   = sysclient.QueryLoginLogDetailResp
	QueryLoginLogListReq      = sysclient.QueryLoginLogListReq
	QueryLoginLogListResp     = sysclient.QueryLoginLogListResp
	QueryMenuDetailReq        = sysclient.QueryMenuDetailReq
	QueryMenuDetailResp       = sysclient.QueryMenuDetailResp
	QueryMenuListReq          = sysclient.QueryMenuListReq
	QueryMenuListResp         = sysclient.QueryMenuListResp
	QueryOperateLogDetailReq  = sysclient.QueryOperateLogDetailReq
	QueryOperateLogDetailResp = sysclient.QueryOperateLogDetailResp
	QueryOperateLogListReq    = sysclient.QueryOperateLogListReq
	QueryOperateLogListResp   = sysclient.QueryOperateLogListResp
	QueryPostDetailReq        = sysclient.QueryPostDetailReq
	QueryPostDetailResp       = sysclient.QueryPostDetailResp
	QueryPostListReq          = sysclient.QueryPostListReq
	QueryPostListResp         = sysclient.QueryPostListResp
	QueryRoleDetailReq        = sysclient.QueryRoleDetailReq
	QueryRoleDetailResp       = sysclient.QueryRoleDetailResp
	QueryRoleListReq          = sysclient.QueryRoleListReq
	QueryRoleListResp         = sysclient.QueryRoleListResp
	QueryRoleMenuListReq      = sysclient.QueryRoleMenuListReq
	QueryRoleMenuListResp     = sysclient.QueryRoleMenuListResp
	QueryRoleUserListReq      = sysclient.QueryRoleUserListReq
	QueryRoleUserListResp     = sysclient.QueryRoleUserListResp
	QueryUserDetailReq        = sysclient.QueryUserDetailReq
	QueryUserDetailResp       = sysclient.QueryUserDetailResp
	QueryUserListReq          = sysclient.QueryUserListReq
	QueryUserListResp         = sysclient.QueryUserListResp
	QueryUserRoleListReq      = sysclient.QueryUserRoleListReq
	QueryUserRoleListResp     = sysclient.QueryUserRoleListResp
	ReSetPasswordReq          = sysclient.ReSetPasswordReq
	ReSetPasswordResp         = sysclient.ReSetPasswordResp
	RoleData                  = sysclient.RoleData
	RoleListData              = sysclient.RoleListData
	UpdateDeptReq             = sysclient.UpdateDeptReq
	UpdateDeptResp            = sysclient.UpdateDeptResp
	UpdateDeptStatusReq       = sysclient.UpdateDeptStatusReq
	UpdateDeptStatusResp      = sysclient.UpdateDeptStatusResp
	UpdateDictItemReq         = sysclient.UpdateDictItemReq
	UpdateDictItemResp        = sysclient.UpdateDictItemResp
	UpdateDictItemStatusReq   = sysclient.UpdateDictItemStatusReq
	UpdateDictItemStatusResp  = sysclient.UpdateDictItemStatusResp
	UpdateDictTypeReq         = sysclient.UpdateDictTypeReq
	UpdateDictTypeResp        = sysclient.UpdateDictTypeResp
	UpdateDictTypeStatusReq   = sysclient.UpdateDictTypeStatusReq
	UpdateDictTypeStatusResp  = sysclient.UpdateDictTypeStatusResp
	UpdateMenuReq             = sysclient.UpdateMenuReq
	UpdateMenuResp            = sysclient.UpdateMenuResp
	UpdateMenuStatusReq       = sysclient.UpdateMenuStatusReq
	UpdateMenuStatusResp      = sysclient.UpdateMenuStatusResp
	UpdatePostReq             = sysclient.UpdatePostReq
	UpdatePostResp            = sysclient.UpdatePostResp
	UpdatePostStatusReq       = sysclient.UpdatePostStatusReq
	UpdatePostStatusResp      = sysclient.UpdatePostStatusResp
	UpdateRoleMenuReq         = sysclient.UpdateRoleMenuReq
	UpdateRoleMenuResp        = sysclient.UpdateRoleMenuResp
	UpdateRoleReq             = sysclient.UpdateRoleReq
	UpdateRoleResp            = sysclient.UpdateRoleResp
	UpdateRoleStatusReq       = sysclient.UpdateRoleStatusReq
	UpdateRoleStatusResp      = sysclient.UpdateRoleStatusResp
	UpdateUserReq             = sysclient.UpdateUserReq
	UpdateUserResp            = sysclient.UpdateUserResp
	UpdateUserRoleListReq     = sysclient.UpdateUserRoleListReq
	UpdateUserRoleListResp    = sysclient.UpdateUserRoleListResp
	UpdateUserStatusReq       = sysclient.UpdateUserStatusReq
	UpdateUserStatusResp      = sysclient.UpdateUserStatusResp
	UserData                  = sysclient.UserData
	UserListData              = sysclient.UserListData

	DictTypeService interface {
		// 添加字典类型
		AddDictType(ctx context.Context, in *AddDictTypeReq, opts ...grpc.CallOption) (*AddDictTypeResp, error)
		// 删除字典类型
		DeleteDictType(ctx context.Context, in *DeleteDictTypeReq, opts ...grpc.CallOption) (*DeleteDictTypeResp, error)
		// 更新字典类型
		UpdateDictType(ctx context.Context, in *UpdateDictTypeReq, opts ...grpc.CallOption) (*UpdateDictTypeResp, error)
		// 更新字典类型状态
		UpdateDictTypeStatus(ctx context.Context, in *UpdateDictTypeStatusReq, opts ...grpc.CallOption) (*UpdateDictTypeStatusResp, error)
		// 查询字典类型详情
		QueryDictTypeDetail(ctx context.Context, in *QueryDictTypeDetailReq, opts ...grpc.CallOption) (*QueryDictTypeDetailResp, error)
		// 查询字典类型列表
		QueryDictTypeList(ctx context.Context, in *QueryDictTypeListReq, opts ...grpc.CallOption) (*QueryDictTypeListResp, error)
	}

	defaultDictTypeService struct {
		cli zrpc.Client
	}
)

func NewDictTypeService(cli zrpc.Client) DictTypeService {
	return &defaultDictTypeService{
		cli: cli,
	}
}

// 添加字典类型
func (m *defaultDictTypeService) AddDictType(ctx context.Context, in *AddDictTypeReq, opts ...grpc.CallOption) (*AddDictTypeResp, error) {
	client := sysclient.NewDictTypeServiceClient(m.cli.Conn())
	return client.AddDictType(ctx, in, opts...)
}

// 删除字典类型
func (m *defaultDictTypeService) DeleteDictType(ctx context.Context, in *DeleteDictTypeReq, opts ...grpc.CallOption) (*DeleteDictTypeResp, error) {
	client := sysclient.NewDictTypeServiceClient(m.cli.Conn())
	return client.DeleteDictType(ctx, in, opts...)
}

// 更新字典类型
func (m *defaultDictTypeService) UpdateDictType(ctx context.Context, in *UpdateDictTypeReq, opts ...grpc.CallOption) (*UpdateDictTypeResp, error) {
	client := sysclient.NewDictTypeServiceClient(m.cli.Conn())
	return client.UpdateDictType(ctx, in, opts...)
}

// 更新字典类型状态
func (m *defaultDictTypeService) UpdateDictTypeStatus(ctx context.Context, in *UpdateDictTypeStatusReq, opts ...grpc.CallOption) (*UpdateDictTypeStatusResp, error) {
	client := sysclient.NewDictTypeServiceClient(m.cli.Conn())
	return client.UpdateDictTypeStatus(ctx, in, opts...)
}

// 查询字典类型详情
func (m *defaultDictTypeService) QueryDictTypeDetail(ctx context.Context, in *QueryDictTypeDetailReq, opts ...grpc.CallOption) (*QueryDictTypeDetailResp, error) {
	client := sysclient.NewDictTypeServiceClient(m.cli.Conn())
	return client.QueryDictTypeDetail(ctx, in, opts...)
}

// 查询字典类型列表
func (m *defaultDictTypeService) QueryDictTypeList(ctx context.Context, in *QueryDictTypeListReq, opts ...grpc.CallOption) (*QueryDictTypeListResp, error) {
	client := sysclient.NewDictTypeServiceClient(m.cli.Conn())
	return client.QueryDictTypeList(ctx, in, opts...)
}
