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

	for _, role := range result.List {
		list = append(list, &types.QueryRoleListData{
			CreateBy:   role.CreateBy,
			CreateTime: role.CreateTime,
			DataScope:  role.DataScope,
			Id:         role.Id,
			IsAdmin:    role.IsAdmin,
			Remark:     role.Remark,
			RoleKey:    role.RoleKey,
			RoleName:   role.RoleName,
			RoleSort:   role.RoleSort,
			RoleStatus: role.RoleStatus,
			UpdateBy:   role.UpdateBy,
			UpdateTime: role.UpdateTime,
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
