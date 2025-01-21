package roleservicelogic

import (
	"context"
	"errors"
	"fmt"
	"github.com/feihua/zero-admin/rpc/sys/gen/model"
	"github.com/feihua/zero-admin/rpc/sys/gen/query"
	"github.com/feihua/zero-admin/rpc/sys/internal/svc"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateRoleLogic 更新角色
/*
Author: LiuFeiHua
Date: 2023/12/18 15:57
*/
type UpdateRoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateRoleLogic {
	return &UpdateRoleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateRole 更新角色(id为1的是系统预留超级管理员角色,不能更新)
// 1.根据角色id查询角色是否已存在
// 2.角色存在时,则直接更新角色
func (l *UpdateRoleLogic) UpdateRole(in *sysclient.UpdateRoleReq) (*sysclient.UpdateRoleResp, error) {
	role := query.SysRole
	q := role.WithContext(l.ctx)

	if in.Id == 1 {
		return nil, errors.New(fmt.Sprintf("更新角色信息失败,不允许操作超级管理员角色"))
	}

	// 1.查询角色角色是否已存在
	_, err := q.Where(role.ID.Eq(in.Id)).First()

	if err != nil {
		logc.Errorf(l.ctx, "根据角色id：%d,查询角色信息失败,异常:%s", in.Id, err.Error())
		return nil, errors.New(fmt.Sprintf("更新角色信息失败,角色信息不存在"))
	}

	// 2.查询角色名称是否存在
	name := in.RoleName
	count, err := q.Where(role.ID.Neq(in.Id), role.RoleName.Eq(name)).Count()

	if err != nil {
		logc.Errorf(l.ctx, "根据角色名称：%s,查询角色信息失败,异常:%s", name, err.Error())
		return nil, errors.New(fmt.Sprintf("更新角色失败"))
	}

	if count > 0 {
		return nil, errors.New(fmt.Sprintf("更新角色信息失败,角色名称已存在"))
	}

	// 3.查询权限字符是否存在
	roleKey := in.RoleKey
	count, err = q.Where(role.ID.Neq(in.Id), role.RoleKey.Eq(roleKey)).Count()

	if err != nil {
		logc.Errorf(l.ctx, "根据角色权限：%s,查询角色信息失败,异常:%s", roleKey, err.Error())
		return nil, errors.New(fmt.Sprintf("更新角色失败"))
	}

	if count > 0 {
		return nil, errors.New(fmt.Sprintf("更新角色信息失败,角色权限已存在"))
	}

	// 3.角色存在时,则直接更新角色
	sysRole := &model.SysRole{
		ID:         in.Id,         // 编号
		RoleName:   in.RoleName,   // 角色名称
		RoleKey:    in.RoleKey,    // 权限字符
		RoleStatus: in.RoleStatus, // 角色状态
		RoleSort:   in.RoleSort,   // 角色排序
		DataScope:  in.DataScope,  // 数据权限
		IsAdmin:    in.IsAdmin,    // 是否超级管理员:  0：否  1：是
		Remark:     in.Remark,     // 备注
		UpdateBy:   in.UpdateBy,   // 更新者
	}

	_, err = q.Updates(sysRole)

	if err != nil {
		logc.Errorf(l.ctx, "更新角色信息失败,参数:%+v,异常:%s", role, err.Error())
		return nil, errors.New("更新角色信息失败")
	}
	return &sysclient.UpdateRoleResp{}, nil
}
