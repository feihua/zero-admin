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

// UpdateMenuRoleListLogic 更新角色与菜单的关联
/*
Author: LiuFeiHua
Date: 2024/5/24 15:31
*/
type UpdateMenuRoleListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateMenuRoleListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateMenuRoleListLogic {
	return &UpdateMenuRoleListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateMenuRoleList 更新角色与菜单的关联
// 1.不允许操作超级管理员角色
// 2.判断角色是否存在
// 3.删除角色与菜单的关联
// 4.添加角色与菜单的关联
func (l *UpdateMenuRoleListLogic) UpdateMenuRoleList(in *sysclient.UpdateMenuRoleReq) (*sysclient.UpdateMenuRoleResp, error) {
	// 1.不允许操作超级管理员角色
	if in.RoleId == 1 {
		return nil, errors.New(fmt.Sprintf("不允许操作超级管理员角色"))
	}

	// 2.判断角色是否存在
	role := query.SysRole
	count, err := role.WithContext(l.ctx).Where(role.ID.Eq(in.RoleId)).Count()

	if err != nil {
		return nil, errors.New("查询角色失败")
	}

	if count == 0 {
		return nil, errors.New("角色不存在")
	}

	err = query.Q.Transaction(func(tx *query.Query) error {

		q := tx.SysRoleMenu
		// 3.删除角色与菜单的关联
		if _, err = q.WithContext(l.ctx).Where(q.RoleID.Eq(in.RoleId)).Delete(); err != nil {
			logc.Errorf(l.ctx, "删除角色与菜单的关联失败,参数:%+v,异常:%s", in, err.Error())
			return err
		}

		var roleMenus []*model.SysRoleMenu
		for _, menuId := range in.MenuIds {
			roleMenus = append(roleMenus, &model.SysRoleMenu{
				RoleID: in.RoleId,
				MenuID: menuId,
			})
		}

		// 4.添加角色与菜单的关联
		if err = q.WithContext(l.ctx).CreateInBatches(roleMenus, len(roleMenus)); err != nil {
			logc.Errorf(l.ctx, "添加角色与菜单的关联失败,参数:%+v,异常:%s", in, err.Error())
			return err
		}

		return nil
	})

	if err != nil {
		logc.Errorf(l.ctx, "更新角色与菜单的关联失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("更新角色与菜单的关联失败")
	}

	return &sysclient.UpdateMenuRoleResp{}, nil
}
