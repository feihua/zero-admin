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

// UpdateRoleMenuListLogic 更新角色与菜单的关联
/*
Author: LiuFeiHua
Date: 2024/5/23 17:30
*/
type UpdateRoleMenuListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateRoleMenuListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateRoleMenuListLogic {
	return &UpdateRoleMenuListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// UpdateRoleMenuList 更新角色与菜单的关联
func (l *UpdateRoleMenuListLogic) UpdateRoleMenuList(req *types.UpdateRoleMenuListReq) (resp *types.UpdateRoleMenuListResp, err error) {
	_, err = l.svcCtx.RoleService.UpdateMenuRoleList(l.ctx, &sysclient.UpdateRoleMenuReq{
		RoleId:  req.RoleId,  // 角色id
		MenuIds: req.MenuIds, // 菜单id
	})

	if err != nil {
		logc.Errorf(l.ctx, "更新角色菜单失败,参数:%+v,异常:%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	return &types.UpdateRoleMenuListResp{
		Code:    "000000",
		Message: "更新角色菜单成功",
	}, nil
}
