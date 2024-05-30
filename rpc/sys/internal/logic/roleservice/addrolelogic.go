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
// 1.根据角色名称查询角色是否已存在
// 2.如果角色已存在,则直接返回
// 3.角色不存在时,则直接添加角色
func (l *AddRoleLogic) AddRole(in *sysclient.AddRoleReq) (*sysclient.AddRoleResp, error) {
	q := query.SysRole.WithContext(l.ctx)
	// 1.根据角色名称查询角色是否已存在
	name := in.RoleName
	count, err := q.Where(query.SysRole.RoleName.Eq(name)).Count()

	if err != nil {
		logc.Errorf(l.ctx, "根据角色名称：%s,查询角色信息失败,异常:%s", name, err.Error())
		return nil, errors.New(fmt.Sprintf("查询角色信息失败"))
	}

	//2.如果角色已存在,则直接返回
	if count > 0 {
		logc.Errorf(l.ctx, "角色信息已存在：%+v", in)
		return nil, errors.New(fmt.Sprintf("角色：%s,已存在", name))
	}

	//3.角色不存在时,则直接添加角色
	role := &model.SysRole{
		RoleName:   in.RoleName,
		RoleKey:    in.RoleKey,
		RoleStatus: in.RoleStatus,
		RoleSort:   in.RoleSort,
		DataScope:  in.DataScope,
		IsDeleted:  0,
		IsAdmin:    in.IsAdmin,
		Remark:     in.Remark,
		CreateBy:   in.CreateBy,
	}

	err = q.Create(role)

	if err != nil {
		logc.Errorf(l.ctx, "新增角色信息失败,参数:%+v,异常:%s", role, err.Error())
		return nil, errors.New("新增角色信息失败")
	}
	return &sysclient.AddRoleResp{}, nil
}
