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
			Id:         detail.Id,         // 编号
			UserName:   detail.UserName,   // 用户名
			NickName:   detail.NickName,   // 昵称
			Avatar:     detail.Avatar,     // 头像
			Email:      detail.Email,      // 邮箱
			Mobile:     detail.Mobile,     // 手机号
			UserStatus: detail.UserStatus, // 帐号状态（0正常 1停用）
			DeptId:     detail.DeptId,     // 部门id
			Remark:     detail.Remark,     // 备注信息
			CreateBy:   detail.CreateBy,   // 创建者
			CreateTime: detail.CreateTime, // 创建时间
			UpdateBy:   detail.UpdateBy,   // 更新者
			UpdateTime: detail.UpdateTime, // 更新时间
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
