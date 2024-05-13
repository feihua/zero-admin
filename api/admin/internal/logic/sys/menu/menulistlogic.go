package menu

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"
	"github.com/zeromicro/go-zero/core/logc"
	"strconv"

	"github.com/zeromicro/go-zero/core/logx"
)

// MenuListLogic
/*
Author: LiuFeiHua
Date: 2023/12/18 15:27
*/
type MenuListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMenuListLogic(ctx context.Context, svcCtx *svc.ServiceContext) MenuListLogic {
	return MenuListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// MenuList 菜单列表
func (l *MenuListLogic) MenuList(req types.ListMenuReq) (*types.ListMenuResp, error) {
	resp, err := l.svcCtx.MenuService.MenuList(l.ctx, &sysclient.MenuListReq{
		Name: req.Name,
		Url:  req.Url,
	})

	if err != nil {
		logc.Errorf(l.ctx, "参数: %+v,查询菜单列表异常:%s", req, err.Error())
		return nil, errorx.NewDefaultError("查询菜单失败")
	}

	var list []*types.ListMenuData

	for _, menu := range resp.List {
		menuItem := &types.ListMenuData{
			Id:            menu.Id,
			Key:           strconv.FormatInt(menu.Id, 10),
			Name:          menu.Name,
			Title:         menu.Name,
			ParentId:      menu.ParentId,
			Url:           menu.Url,
			Perms:         menu.Perms,
			Type:          menu.Type,
			Icon:          menu.Icon,
			OrderNum:      menu.OrderNum,
			CreateBy:      menu.CreateBy,
			CreateTime:    menu.CreateTime,
			UpdateBy:      menu.UpdateBy,
			UpdateTime:    menu.UpdateTime,
			DelFlag:       menu.DelFlag,
			VuePath:       menu.VuePath,
			VueComponent:  menu.VueComponent,
			VueIcon:       menu.VueIcon,
			VueRedirect:   menu.VueRedirect,
			BackgroundUrl: menu.BackgroundUrl,
		}

		list = append(list, menuItem)
	}

	return &types.ListMenuResp{
		Code:    "000000",
		Message: "查询菜单成功",
		Data:    list,
		Success: true,
		Total:   resp.Total,
	}, nil
}
