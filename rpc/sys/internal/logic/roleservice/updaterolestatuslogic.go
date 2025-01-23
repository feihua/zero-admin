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
func (l *UpdateRoleStatusLogic) UpdateRoleStatus(in *sysclient.UpdateRoleStatusReq) (*sysclient.UpdateRoleStatusResp, error) {
	role := query.SysRole

	for _, roleId := range in.Ids {

		if roleId == 1 {
			return nil, errors.New("删除角色失败,不允许操作超级管理员角色")
		}
		count, err := role.WithContext(l.ctx).Where(role.ID.Eq(roleId)).Count()

		if err != nil {
			return nil, errors.New("查询角色失败")
		}

		if count == 0 {
			return nil, errors.New("角色不存在")
		}
	}

	_, err := role.WithContext(l.ctx).Where(role.ID.In(in.Ids...)).Update(role.RoleStatus, in.RoleStatus)

	if err != nil {
		logc.Errorf(l.ctx, "更新角色信息状态失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("更新角色信息状态失败")
	}

	return &sysclient.UpdateRoleStatusResp{}, nil
}
