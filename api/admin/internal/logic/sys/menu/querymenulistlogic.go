package menu

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc/status"
	"strconv"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryMenuListLogic
/*
Author: LiuFeiHua
Date: 2023/12/18 15:27
*/
type QueryMenuListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryMenuListLogic(ctx context.Context, svcCtx *svc.ServiceContext) QueryMenuListLogic {
	return QueryMenuListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// QueryMenuList 菜单列表
func (l *QueryMenuListLogic) QueryMenuList(req *types.ListMenuReq) (*types.ListMenuResp, error) {
	resp, err := l.svcCtx.MenuService.MenuList(l.ctx, &sysclient.MenuListReq{
		Name: req.Name,
		Url:  req.Url,
	})

	if err != nil {
		logc.Errorf(l.ctx, "参数: %+v,查询菜单列表异常:%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
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
