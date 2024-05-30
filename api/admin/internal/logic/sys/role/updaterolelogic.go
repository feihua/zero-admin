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

// UpdateRoleLogic 更新角色
/*
Author: LiuFeiHua
Date: 2023/12/18 15:41
*/
type UpdateRoleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) UpdateRoleLogic {
	return UpdateRoleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// UpdateRole 更新角色
func (l *UpdateRoleLogic) UpdateRole(req *types.UpdateRoleReq) (*types.UpdateRoleResp, error) {
	roleReq := sysclient.UpdateRoleReq{
		DataScope:  req.DataScope,
		Id:         req.Id,
		IsAdmin:    req.IsAdmin,
		Remark:     req.Remark,
		RoleKey:    req.RoleKey,
		RoleName:   req.RoleName,
		RoleSort:   req.RoleSort,
		RoleStatus: req.RoleStatus,
		UpdateBy:   l.ctx.Value("userName").(string),
	}
	_, err := l.svcCtx.RoleService.UpdateRole(l.ctx, &roleReq)

	if err != nil {
		logc.Errorf(l.ctx, "更新角色信息失败,参数:%+v,异常:%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	return &types.UpdateRoleResp{
		Code:    "000000",
		Message: "更新角色成功",
	}, nil
}
