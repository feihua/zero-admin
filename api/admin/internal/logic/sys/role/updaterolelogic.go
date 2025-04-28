package role

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/api/admin/internal/common/res"
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
func (l *UpdateRoleLogic) UpdateRole(req *types.UpdateRoleReq) (*types.BaseResp, error) {
	roleReq := sysclient.UpdateRoleReq{
		Id:        req.Id,        // 角色id
		RoleName:  req.RoleName,  // 名称
		RoleKey:   req.RoleKey,   // 角色权限字符串
		DataScope: req.DataScope, // 数据范围（1：全部数据权限 2：自定数据权限 3：本部门数据权限 4：本部门及以下数据权限）
		Status:    req.Status,    // 状态(1:正常，0:禁用)
		Remark:    req.Remark,    // 备注
		UpdateBy:  l.ctx.Value("userName").(string),
	}
	_, err := l.svcCtx.RoleService.UpdateRole(l.ctx, &roleReq)

	if err != nil {
		logc.Errorf(l.ctx, "更新角色信息失败,参数:%+v,异常:%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	return res.Success()
}
