package role

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc/status"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryRoleDetailLogic 查询角色详情
/*
Author: LiuFeiHua
Date: 2024/5/29 18:02
*/
type QueryRoleDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryRoleDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryRoleDetailLogic {
	return &QueryRoleDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// QueryRoleDetail 查询角色详情
func (l *QueryRoleDetailLogic) QueryRoleDetail(req *types.QueryRoleDetailReq) (resp *types.QueryRoleDetailResp, err error) {
	role, err := l.svcCtx.RoleService.QueryRoleDetail(l.ctx, &sysclient.QueryRoleDetailReq{
		Id: req.Id,
	})

	if err != nil {
		logc.Errorf(l.ctx, "查询角色详情,参数: %+v,异常:%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	data := types.QueryRoleDetailData{
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
	}

	return &types.QueryRoleDetailResp{
		Code:    "000000",
		Message: "查询角色详情成功",
		Data:    data,
	}, nil
}
