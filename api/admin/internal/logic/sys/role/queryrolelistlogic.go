package role

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
	"strings"
)

// QueryRoleListLogic 查询角色列表
/*
Author: LiuFeiHua
Date: 2023/12/18 15:39
*/
type QueryRoleListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryRoleListLogic(ctx context.Context, svcCtx *svc.ServiceContext) QueryRoleListLogic {
	return QueryRoleListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// QueryRoleList 查询角色列表
func (l *QueryRoleListLogic) QueryRoleList(req *types.QueryRoleListReq) (*types.QueryRoleListResp, error) {
	result, err := l.svcCtx.RoleService.QueryRoleList(l.ctx, &sysclient.QueryRoleListReq{
		PageNum:    req.Current,
		PageSize:   req.PageSize,
		IsAdmin:    req.IsAdmin,
		RoleKey:    strings.TrimSpace(req.RoleKey),
		RoleName:   strings.TrimSpace(req.RoleName),
		RoleStatus: req.RoleStatus,
	})

	if err != nil {
		logc.Errorf(l.ctx, "查询角色列表,参数: %+v,异常:%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	var list []*types.QueryRoleListData

	for _, detail := range result.List {
		list = append(list, &types.QueryRoleListData{
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

	return &types.QueryRoleListResp{
		Code:     "000000",
		Message:  "查询角色成功",
		Current:  req.Current,
		Data:     list,
		PageSize: req.PageSize,
		Success:  true,
		Total:    result.Total,
	}, nil
}
