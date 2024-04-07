package role

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateRoleMenuLogic
/*
Author: LiuFeiHua
Date: 2023/12/18 15:43
*/
type UpdateRoleMenuLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateRoleMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) UpdateRoleMenuLogic {
	return UpdateRoleMenuLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// UpdateRoleMenu 更新角色与菜单的关联
func (l *UpdateRoleMenuLogic) UpdateRoleMenu(req types.UpdateRoleMenuReq) (*types.UpdateRoleMenuResp, error) {
	_, err := l.svcCtx.RoleService.UpdateMenuRole(l.ctx, &sysclient.UpdateMenuRoleReq{
		RoleId:   req.RoleId,
		MenuIds:  req.MenuIds,
		CreateBy: l.ctx.Value("userName").(string),
	})

	if err != nil {
		logc.Errorf(l.ctx, "更新角色菜单失败,参数:%+v,异常:%s", req, err.Error())
		return nil, errorx.NewDefaultError("更新角色菜单失败")
	}

	return &types.UpdateRoleMenuResp{
		Code:    "000000",
		Message: "更新角色菜单成功",
	}, nil
}
