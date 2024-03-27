package role

import (
	"context"
	"github.com/feihua/zero-admin/api/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"
	"github.com/zeromicro/go-zero/core/logc"
	"strconv"

	"github.com/feihua/zero-admin/api/internal/svc"
	"github.com/feihua/zero-admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryMenuByRoleIdLogic
/*
Author: LiuFeiHua
Date: 2023/12/18 15:32
*/
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
	resp, err := l.svcCtx.MenuService.MenuList(l.ctx, &sysclient.MenuListReq{
		Name: "",
		Url:  "",
	})

	if err != nil {
		logc.Errorf(l.ctx, "查询菜单信息失败,参数:%+v,异常:%s", req, err.Error())
		return nil, errorx.NewDefaultError("查询菜单信息失败")
	}

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
		QueryMenu, err1 := l.svcCtx.RoleService.QueryMenuByRoleId(l.ctx, &sysclient.QueryMenuByRoleIdReq{
			Id: req.Id,
		})
		if err1 != nil {
			logc.Errorf(l.ctx, "根据roleId查询菜单失败,参数:%+v,异常:%s", req, err1.Error())
			return nil, errorx.NewDefaultError("根据roleId查询菜单失败")
		}
		listIds = QueryMenu.Ids
	}

	return &types.RoleMenuResp{
		AllData:  list,
		RoleData: listIds,
		Code:     "000000",
		Message:  "根据角色id查询菜单成功",
	}, nil
}
