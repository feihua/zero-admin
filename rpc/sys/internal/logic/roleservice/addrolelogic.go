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

// AddRoleLogic 新增角色
/*
Author: LiuFeiHua
Date: 2023/12/18 15:54
*/
type AddRoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddRoleLogic {
	return &AddRoleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// AddRole 新增角色
// 1.查询角色名称是否存在
// 2.查询权限字符是否存在
// 3.角色不存在时,则直接添加角色
func (l *AddRoleLogic) AddRole(in *sysclient.AddRoleReq) (*sysclient.AddRoleResp, error) {
	q := query.SysRole.WithContext(l.ctx)

	// 1.查询角色名称是否存在
	name := in.RoleName
	count, err := q.Where(query.SysRole.RoleName.Eq(name)).Count()

	if err != nil {
		logc.Errorf(l.ctx, "根据角色名称：%s,查询角色失败,异常:%s", name, err.Error())
		return nil, errors.New(fmt.Sprintf("新增角色失败"))
	}

	if count > 0 {
		return nil, errors.New(fmt.Sprintf("角色名称：%s,已存在", name))
	}

	// 2.查询权限字符是否存在
	roleKey := in.RoleKey
	count, err = q.Where(query.SysRole.RoleKey.Eq(roleKey)).Count()

	if err != nil {
		logc.Errorf(l.ctx, "根据角色权限：%s,查询角色失败,异常:%s", roleKey, err.Error())
		return nil, errors.New(fmt.Sprintf("新增角色失败"))
	}

	if count > 0 {
		return nil, errors.New(fmt.Sprintf("权限字符：%s,已存在", roleKey))
	}

	// 3.角色不存在时,则直接添加角色
	role := &model.SysRole{
		RoleName:  in.RoleName,  // 名称
		RoleKey:   in.RoleKey,   // 角色权限字符串
		DataScope: in.DataScope, // 数据范围（1：全部数据权限 2：自定数据权限 3：本部门数据权限 4：本部门及以下数据权限）
		Status:    in.Status,    // 状态(1:正常，0:禁用)
		Remark:    in.Remark,    // 备注
		DelFlag:   in.DelFlag,   // 删除标志（0代表删除 1代表存在）
		CreateBy:  in.CreateBy,  // 创建者
	}

	err = q.Create(role)

	if err != nil {
		logc.Errorf(l.ctx, "新增角色失败,参数:%+v,异常:%s", role, err.Error())
		return nil, errors.New("新增角色失败")
	}
	return &sysclient.AddRoleResp{}, nil
}
