package role

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc/status"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryRoleUserListLogic 根据roleId查询角色的用户
/*
Author: LiuFeiHua
Date: 2024/6/02 19:25
*/
type QueryRoleUserListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryRoleUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryRoleUserListLogic {
	return &QueryRoleUserListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// QueryRoleUserList 根据roleId查询角色的用户
func (l *QueryRoleUserListLogic) QueryRoleUserList(req *types.QueryRoleUserListReq) (resp *types.QueryRoleUserListResp, err error) {
	result, err := l.svcCtx.RoleService.QueryRoleUserList(l.ctx, &sysclient.QueryRoleUserListReq{
		PageNum:  req.Current,
		PageSize: req.PageSize,
		RoleId:   req.RoleId,   // 角色id
		IsExist:  req.IsExist,  // 授权标志
		Mobile:   req.Mobile,   // 手机号
		UserName: req.UserName, // 用户名
	})

	if err != nil {
		logc.Errorf(l.ctx, "查询用户列表信息失败,参数:%+v,异常:%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	var list []*types.QueryRoleUserListData

	for _, detail := range result.List {
		list = append(list, &types.QueryRoleUserListData{
			Id:            detail.Id,            // 用户id
			Mobile:        detail.Mobile,        // 手机号码
			UserName:      detail.UserName,      // 用户账号
			NickName:      detail.NickName,      // 用户昵称
			UserType:      detail.UserType,      // 用户类型（00系统用户）
			Avatar:        detail.Avatar,        // 头像路径
			Email:         detail.Email,         // 用户邮箱
			Status:        detail.Status,        // 状态(1:正常，0:禁用)
			DeptId:        detail.DeptId,        // 部门ID
			LoginIp:       detail.LoginIp,       // 最后登录IP
			LoginDate:     detail.LoginDate,     // 最后登录时间
			LoginBrowser:  detail.LoginBrowser,  // 浏览器类型
			LoginOs:       detail.LoginOs,       // 操作系统
			PwdUpdateDate: detail.PwdUpdateDate, // 密码最后更新时间
			Remark:        detail.Remark,        // 备注
			DelFlag:       detail.DelFlag,       // 删除标志（0代表删除 1代表存在）
			CreateBy:      detail.CreateBy,      // 创建者
			CreateTime:    detail.CreateTime,    // 创建时间
			UpdateBy:      detail.UpdateBy,      // 更新者
			UpdateTime:    detail.UpdateTime,    // 更新时间
		})
	}

	return &types.QueryRoleUserListResp{
		Code:     "000000",
		Message:  "查询角色用户列表成功",
		Current:  req.Current,
		Data:     list,
		PageSize: req.PageSize,
		Success:  true,
		Total:    result.Total,
	}, nil

}
