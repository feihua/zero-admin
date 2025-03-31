// Code generated by goctl. DO NOT EDIT.
// goctl 1.8.1
// Source: sys.proto

package loginlogservice

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

	LoginLogService interface {
		// 删除系统登录日志
		DeleteLoginLog(ctx context.Context, in *DeleteLoginLogReq, opts ...grpc.CallOption) (*DeleteLoginLogResp, error)
		// 查询系统登录日志详情
		QueryLoginLogDetail(ctx context.Context, in *QueryLoginLogDetailReq, opts ...grpc.CallOption) (*QueryLoginLogDetailResp, error)
		// 查询系统登录日志列表
		QueryLoginLogList(ctx context.Context, in *QueryLoginLogListReq, opts ...grpc.CallOption) (*QueryLoginLogListResp, error)
	}

	defaultLoginLogService struct {
		cli zrpc.Client
	}
)

func NewLoginLogService(cli zrpc.Client) LoginLogService {
	return &defaultLoginLogService{
		cli: cli,
	}
}

// 删除系统登录日志
func (m *defaultLoginLogService) DeleteLoginLog(ctx context.Context, in *DeleteLoginLogReq, opts ...grpc.CallOption) (*DeleteLoginLogResp, error) {
	client := sysclient.NewLoginLogServiceClient(m.cli.Conn())
	return client.DeleteLoginLog(ctx, in, opts...)
}

// 查询系统登录日志详情
func (m *defaultLoginLogService) QueryLoginLogDetail(ctx context.Context, in *QueryLoginLogDetailReq, opts ...grpc.CallOption) (*QueryLoginLogDetailResp, error) {
	client := sysclient.NewLoginLogServiceClient(m.cli.Conn())
	return client.QueryLoginLogDetail(ctx, in, opts...)
}

// 查询系统登录日志列表
func (m *defaultLoginLogService) QueryLoginLogList(ctx context.Context, in *QueryLoginLogListReq, opts ...grpc.CallOption) (*QueryLoginLogListResp, error) {
	client := sysclient.NewLoginLogServiceClient(m.cli.Conn())
	return client.QueryLoginLogList(ctx, in, opts...)
}
