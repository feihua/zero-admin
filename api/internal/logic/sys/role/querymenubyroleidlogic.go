package logic

import (
	"context"
	"strconv"
	"zero-admin/rpc/sys/sysclient"

	"zero-admin/api/internal/svc"
	"zero-admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryMenuByRoleIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryMenuByRoleIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) QueryMenuByRoleIdLogic {
	return QueryMenuByRoleIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// QueryMenuByRoleId 根据roleId查询角色的权限
func (l *QueryMenuByRoleIdLogic) QueryMenuByRoleId(req types.RoleMenuReq) (*types.RoleMenuResp, error) {

	//查询所有菜单
	resp, _ := l.svcCtx.Sys.MenuList(l.ctx, &sysclient.MenuListReq{
		Name: "",
		Url:  "",
	})

	var list []*types.ListMenuData
	var listIds []int64

	for _, menu := range resp.List {
		list = append(list, &types.ListMenuData{
			Key:      strconv.FormatInt(menu.Id, 10),
			Title:    menu.Name,
			ParentId: menu.ParentId,
			Id:       menu.Id,
			Label:    menu.Name,
		})
		//admin账号全部权限
		listIds = append(listIds, menu.Id)
	}

	//如果角色不是admin则根据roleId查询菜单
	if req.Id != 1 {
		QueryMenu, _ := l.svcCtx.Sys.QueryMenuByRoleId(l.ctx, &sysclient.QueryMenuByRoleIdReq{
			Id: req.Id,
		})
		listIds = QueryMenu.Ids
	}

	return &types.RoleMenuResp{
		AllData:  list,
		RoleData: listIds,
		Code:     "000000",
		Message:  "根据角色id查询菜单成功",
	}, nil
}
