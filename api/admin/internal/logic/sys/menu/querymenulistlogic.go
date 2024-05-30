package menu

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
)

// QueryMenuListLogic 查询菜单列表
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

// QueryMenuList 查询菜单列表
func (l *QueryMenuListLogic) QueryMenuList(req *types.QueryMenuListReq) (*types.QueryMenuListResp, error) {
	result, err := l.svcCtx.MenuService.QueryMenuList(l.ctx, &sysclient.QueryMenuListReq{})

	if err != nil {
		logc.Errorf(l.ctx, "查询菜单列表,参数: %+v,异常:%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	var list []*types.QueryMenuListData

	for _, menu := range result.List {
		menuItem := &types.QueryMenuListData{
			BackgroundUrl: menu.BackgroundUrl,
			CreateBy:      menu.CreateBy,
			CreateTime:    menu.CreateTime,
			Id:            menu.Id,
			MenuIcon:      menu.MenuIcon,
			MenuName:      menu.MenuName,
			MenuPath:      menu.MenuPath,
			MenuPerms:     menu.MenuPerms,
			MenuSort:      menu.MenuSort,
			MenuStatus:    menu.MenuStatus,
			MenuType:      menu.MenuType,
			ParentId:      menu.ParentId,
			Remark:        menu.Remark,
			UpdateBy:      menu.UpdateBy,
			UpdateTime:    menu.UpdateTime,
			VueComponent:  menu.VueComponent,
			VueIcon:       menu.VueIcon,
			VuePath:       menu.VuePath,
			VueRedirect:   menu.VueRedirect,
			IsVisible:     menu.IsVisible,
		}

		list = append(list, menuItem)
	}

	return &types.QueryMenuListResp{
		Code:    "000000",
		Message: "查询菜单成功",
		Data:    list,
		Success: true,
		Total:   result.Total,
	}, nil
}
