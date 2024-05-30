package roleservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/sys/gen/query"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/rpc/sys/internal/svc"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateRoleStatusLogic 更新角色信息状态
/*
Author: LiuFeiHua
Date: 2024/5/30 14:20
*/
type UpdateRoleStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateRoleStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateRoleStatusLogic {
	return &UpdateRoleStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateRoleStatus 更新角色信息状态
// 1.排除超级管理员
// 2.更新角色信息状态
func (l *UpdateRoleStatusLogic) UpdateRoleStatus(in *sysclient.UpdateRoleStatusReq) (*sysclient.UpdateRoleStatusResp, error) {
	role := query.SysRole

	// 1.排除超级管理员
	var roleIds []int64
	for _, roleId := range in.Ids {

		count, _ := role.WithContext(l.ctx).Where(role.ID.Eq(roleId), role.IsAdmin.Eq(1)).Count()
		if count > 0 {
			continue
		}

		roleIds = append(roleIds, roleId)
	}

	if len(roleIds) == 0 {
		logc.Errorf(l.ctx, "更新角色信息状态失败,参数:%+v,异常:%s", in, "超级管理员和已使用的角色不能被删除")
		return nil, errors.New("超级管理员和已使用的角色不能被更新状态")
	}

	// 2.更新角色信息状态
	_, err := role.WithContext(l.ctx).Where(role.ID.In(roleIds...)).Update(role.RoleStatus, in.RoleStatus)

	if err != nil {
		logc.Errorf(l.ctx, "更新角色信息状态失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("更新角色信息状态失败")
	}
	return &sysclient.UpdateRoleStatusResp{}, nil
}
