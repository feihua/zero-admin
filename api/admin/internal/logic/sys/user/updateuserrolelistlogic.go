package user

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateUserRoleListLogic 更新用户与角色的关联
/*
Author: LiuFeiHua
Date: 2024/5/23 17:33
*/
type UpdateUserRoleListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateUserRoleListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserRoleListLogic {
	return &UpdateUserRoleListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// UpdateUserRoleList 更新用户与角色的关联
func (l *UpdateUserRoleListLogic) UpdateUserRoleList(req *types.UpdateUserRoleReq) (resp *types.UpdateUserRoleResp, err error) {
	_, err = l.svcCtx.UserService.UpdateUserRole(l.ctx, &sysclient.UpdateUserRoleReq{
		UserId:  req.UserId,
		RoleIds: req.RoleIds,
	})

	if err != nil {
		logc.Errorf(l.ctx, "更新用户与角色的关联失败,参数:%+v,异常:%s", req, err.Error())
		return nil, errorx.NewDefaultError("更新角色菜单失败")
	}

	return &types.UpdateUserRoleResp{
		Code:    "000000",
		Message: "更新用户与角色的关联成功",
	}, nil
}
