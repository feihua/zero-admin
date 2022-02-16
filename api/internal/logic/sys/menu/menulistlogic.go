package logic

import (
	"context"
	"encoding/json"
	"strconv"
	"zero-admin/api/internal/common/errorx"
	"zero-admin/api/internal/svc"
	"zero-admin/api/internal/types"
	"zero-admin/rpc/sys/sysclient"

	"github.com/zeromicro/go-zero/core/logx"
)

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

func (l *MenuListLogic) MenuList(req types.ListMenuReq) (*types.ListMenuResp, error) {
	resp, err := l.svcCtx.Sys.MenuList(l.ctx, &sysclient.MenuListReq{
		Name: req.Name,
		Url:  req.Url,
	})

	if err != nil {
		data, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("参数: %s,查询菜单列表异常:%s", string(data), err.Error())
		return nil, errorx.NewDefaultError("查询菜单失败")
	}

	var list []*types.ListtMenuData

	for _, menu := range resp.List {
		list = append(list, &types.ListtMenuData{
			Id:             menu.Id,
			Key:            strconv.FormatInt(menu.Id, 10),
			Name:           menu.Name,
			Title:          menu.Name,
			ParentId:       menu.ParentId,
			Url:            menu.Url,
			Perms:          menu.Perms,
			Type:           menu.Type,
			Icon:           menu.Icon,
			OrderNum:       menu.OrderNum,
			CreateBy:       menu.CreateBy,
			CreateTime:     menu.CreateTime,
			LastUpdateBy:   menu.LastUpdateBy,
			LastUpdateTime: menu.LastUpdateTime,
			DelFlag:        menu.DelFlag,
			VuePath:        menu.VuePath,
			VueComponent:   menu.VueComponent,
			VueIcon:        menu.VueIcon,
			VueRedirect:    menu.VueRedirect,
		})
	}

	return &types.ListMenuResp{
		Code:    "000000",
		Message: "查询菜单成功",
		Data:    list,
		Success: true,
		Total:   resp.Total,
	}, nil
}
