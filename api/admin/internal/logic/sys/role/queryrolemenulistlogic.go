package role

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"
	"github.com/zeromicro/go-zero/core/logc"
	"strconv"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryRoleMenuListLogic 根据roleId查询角色的权限
/*
Author: LiuFeiHua
Date: 2024/5/23 17:25
*/
type QueryRoleMenuListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryRoleMenuListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryRoleMenuListLogic {
	return &QueryRoleMenuListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// QueryRoleMenuList 根据roleId查询角色的权限
func (l *QueryRoleMenuListLogic) QueryRoleMenuList(req *types.QueryRoleMenuListReq) (resp *types.QueryRoleMenuListResp, err error) {
	//查询所有菜单
	result, err := l.svcCtx.MenuService.MenuList(l.ctx, &sysclient.MenuListReq{})

	if err != nil {
		logc.Errorf(l.ctx, "查询菜单信息失败,参数:%+v,异常:%s", req, err.Error())
		return nil, errorx.NewDefaultError("查询菜单信息失败")
	}

	var menuList []types.MenuTreeListData
	var menuIds []int64

	for _, menu := range result.List {
		menuList = append(menuList, types.MenuTreeListData{
			Key:      strconv.FormatInt(menu.Id, 10),
			Title:    menu.Name,
			ParentId: menu.ParentId,
			Id:       menu.Id,
			Label:    menu.Name,
		})
		//admin账号全部权限
		menuIds = append(menuIds, menu.Id)
	}

	//如果角色不是admin则根据roleId查询菜单
	if req.RoleId != 1 {
		QueryMenu, err1 := l.svcCtx.RoleService.QueryMenuByRoleId(l.ctx, &sysclient.QueryMenuByRoleIdReq{
			Id: req.RoleId,
		})
		if err1 != nil {
			logc.Errorf(l.ctx, "根据roleId查询菜单失败,参数:%+v,异常:%s", req, err1.Error())
			return nil, errorx.NewDefaultError("根据roleId查询菜单失败")
		}
		menuIds = QueryMenu.Ids
	}

	return &types.QueryRoleMenuListResp{
		Data: types.RoleMenuListData{
			MenuList: menuList,
			MenuIds:  menuIds,
		},
		Code:    "000000",
		Message: "根据角色id查询菜单成功",
	}, nil
}
