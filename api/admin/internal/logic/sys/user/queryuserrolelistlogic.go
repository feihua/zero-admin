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
			Id:         detail.Id,         // 角色id
			RoleName:   detail.RoleName,   // 名称
			RoleKey:    detail.RoleKey,    // 角色权限字符串
			DataScope:  detail.DataScope,  // 数据范围（1：全部数据权限 2：自定数据权限 3：本部门数据权限 4：本部门及以下数据权限）
			Status:     detail.Status,     // 状态(1:正常，0:禁用)
			Remark:     detail.Remark,     // 备注
			DelFlag:    detail.DelFlag,    // 删除标志（0代表删除 1代表存在）
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
