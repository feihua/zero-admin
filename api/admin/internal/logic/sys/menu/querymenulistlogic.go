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
			Id:            menu.Id,            // 编号
			MenuName:      menu.MenuName,      // 菜单名称
			ParentId:      menu.ParentId,      // 父菜单ID，一级菜单为0
			MenuPath:      menu.MenuPath,      // 前端路由
			MenuPerms:     menu.MenuPerms,     // 权限标识
			MenuType:      menu.MenuType,      // 类型 0：目录,1：菜单,2：按钮,3：外链
			MenuIcon:      menu.MenuIcon,      // 菜单图标
			MenuSort:      menu.MenuSort,      // 菜单排序
			CreateBy:      menu.CreateBy,      // 创建者
			CreateTime:    menu.CreateTime,    // 创建时间
			UpdateBy:      menu.UpdateBy,      // 更新者
			UpdateTime:    menu.UpdateTime,    // 更新时间
			MenuStatus:    menu.MenuStatus,    // 菜单状态
			IsDeleted:     menu.IsDeleted,     // 是否删除  0：否  1：是
			IsVisible:     menu.IsVisible,     // 是否可见  0：否  1：是
			Remark:        menu.Remark,        // 备注信息
			VuePath:       menu.VuePath,       // vue系统的path
			VueComponent:  menu.VueComponent,  // vue的页面
			VueIcon:       menu.VueIcon,       // vue的图标
			VueRedirect:   menu.VueRedirect,   // vue的路由重定向
			BackgroundUrl: menu.BackgroundUrl, // 接口地址
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
