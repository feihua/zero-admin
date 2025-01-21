package user

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
)

// QueryUserRoleListLogic 查询用户与角色的关联
/*
Author: LiuFeiHua
Date: 2024/5/23 17:33
*/
type QueryUserRoleListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryUserRoleListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryUserRoleListLogic {
	return &QueryUserRoleListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// QueryUserRoleList 查询用户与角色的关联
func (l *QueryUserRoleListLogic) QueryUserRoleList(req *types.QueryUserRoleListReq) (resp *types.QueryUserRoleListResp, err error) {
	result, err := l.svcCtx.UserService.QueryUserRoleList(l.ctx, &sysclient.QueryUserRoleListReq{
		Current:  req.Current,
		PageSize: req.PageSize,
		UserId:   req.UserId,
	})
	if err != nil {
		logc.Errorf(l.ctx, "查询角色信息失败,参数:%+v,异常:%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	var roleList []types.RoleListData

	for _, detail := range result.List {
		roleList = append(roleList, types.RoleListData{
			Id:         detail.Id,         // 编号
			RoleName:   detail.RoleName,   // 角色名称
			RoleKey:    detail.RoleKey,    // 权限字符
			RoleStatus: detail.RoleStatus, // 角色状态
			RoleSort:   detail.RoleSort,   // 角色排序
			DataScope:  detail.DataScope,  // 数据权限
			IsDeleted:  detail.IsDeleted,  // 是否删除  0：否  1：是
			IsAdmin:    detail.IsAdmin,    // 是否超级管理员:  0：否  1：是
			Remark:     detail.Remark,     // 备注
			CreateBy:   detail.CreateBy,   // 创建者
			CreateTime: detail.CreateTime, // 创建时间
			UpdateBy:   detail.UpdateBy,   // 更新者
			UpdateTime: detail.UpdateTime, // 更新时间
		})
	}

	return &types.QueryUserRoleListResp{
		Data: types.UserRoleListData{
			RoleList: roleList,
			RoleIds:  result.RoleIds,
		},
		Code:    "000000",
		Message: "查询用户角色成功",
	}, nil
}
