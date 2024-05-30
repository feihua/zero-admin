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

// AddRoleLogic 新增角色
/*
Author: LiuFeiHua
Date: 2023/12/18 15:34
*/
type AddRoleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) AddRoleLogic {
	return AddRoleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// AddRole 新增角色
func (l *AddRoleLogic) AddRole(req *types.AddRoleReq) (*types.AddRoleResp, error) {
	roleAddReq := sysclient.AddRoleReq{
		CreateBy:   l.ctx.Value("userName").(string),
		DataScope:  req.DataScope,
		IsAdmin:    req.IsAdmin,
		Remark:     req.Remark,
		RoleKey:    req.RoleKey,
		RoleName:   req.RoleName,
		RoleSort:   req.RoleSort,
		RoleStatus: req.RoleStatus,
	}

	_, err := l.svcCtx.RoleService.AddRole(l.ctx, &roleAddReq)

	if err != nil {
		logc.Errorf(l.ctx, "添加角色信息失败,参数:%+v,异常:%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	return &types.AddRoleResp{
		Code:    "000000",
		Message: "添加角色成功",
	}, nil
}
