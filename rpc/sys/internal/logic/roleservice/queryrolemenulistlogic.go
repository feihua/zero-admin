package roleservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/sys/gen/query"
	"github.com/feihua/zero-admin/rpc/sys/internal/svc"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryRoleMenuListLogic 查询角色与菜单的关联
/*
Author: LiuFeiHua
Date: 2024/5/24 14:42
*/
type QueryRoleMenuListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryRoleMenuListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryRoleMenuListLogic {
	return &QueryRoleMenuListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryRoleMenuList 查询角色与菜单的关联
// 1.查询所有菜单
// 2.如果角色不是admin则根据roleId查询菜单
func (l *QueryRoleMenuListLogic) QueryRoleMenuList(in *sysclient.QueryRoleMenuListReq) (*sysclient.QueryRoleMenuListResp, error) {
	// 1.查询所有菜单
	menus, err := query.SysMenu.WithContext(l.ctx).Find()
	if err != nil {
		logc.Errorf(l.ctx, "查询菜单列表失败,异常:%s", err.Error())
		return nil, errors.New("查询菜单列表失败")
	}

	var menuList = make([]*sysclient.MenuData, 0, len(menus))
	var menuIds = make([]int64, 0, len(menus))

	for _, menu := range menus {
		menuList = append(menuList, &sysclient.MenuData{
			Id:       menu.ID,       // 编号
			MenuName: menu.MenuName, // 菜单名称
			ParentId: menu.ParentID, // 父菜单ID，一级菜单为0
		})
		menuIds = append(menuIds, menu.ID)
	}

	// 2.如果角色不是admin则根据roleId查询菜单
	if in.RoleId != 1 {
		var ids []int64
		q := query.SysRoleMenu
		_ = q.WithContext(l.ctx).Select(q.MenuID).Where(q.RoleID.Eq(in.RoleId)).Scan(&ids)
		menuIds = ids
	}

	return &sysclient.QueryRoleMenuListResp{
		List:    menuList,
		MenuIds: menuIds,
	}, nil
}
