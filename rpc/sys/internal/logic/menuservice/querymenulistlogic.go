package menuservicelogic

import (
	"context"
	"errors"

	"github.com/feihua/zero-admin/pkg/time_util"
	"github.com/feihua/zero-admin/rpc/sys/gen/query"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/rpc/sys/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryMenuListLogic 菜单列表
/*
Author: LiuFeiHua
Date: 2023/12/18 15:45
*/
type QueryMenuListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryMenuListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryMenuListLogic {
	return &QueryMenuListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryMenuList 菜单列表
func (l *QueryMenuListLogic) QueryMenuList(in *sysclient.QueryMenuListReq) (*sysclient.QueryMenuListResp, error) {
	result, err := query.SysMenu.WithContext(l.ctx).Where(query.SysMenu.MenuType.Neq(2)).Order(query.SysMenu.MenuSort).Find()

	if err != nil {
		logc.Errorf(l.ctx, "查询菜单列表信息失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("查询菜单列表信息失败")
	}
	var list = make([]*sysclient.MenuListData, 0, len(result))

	for _, menu := range result {
		list = append(list, &sysclient.MenuListData{
			Id:            menu.ID,                                 // 编号
			MenuName:      menu.MenuName,                           // 菜单名称
			ParentId:      menu.ParentID,                           // 父菜单ID，一级菜单为0
			MenuPath:      menu.MenuPath,                           // 前端路由
			MenuPerms:     menu.MenuPerms,                          // 权限标识
			MenuType:      menu.MenuType,                           // 类型 0：目录,1：菜单,2：按钮,3：外链
			MenuIcon:      menu.MenuIcon,                           // 菜单图标
			MenuSort:      menu.MenuSort,                           // 菜单排序
			CreateBy:      menu.CreateBy,                           // 创建者
			CreateTime:    time_util.TimeToStr(menu.CreateTime),    // 创建时间
			UpdateBy:      menu.UpdateBy,                           // 更新者
			UpdateTime:    time_util.TimeToString(menu.UpdateTime), // 更新时间
			MenuStatus:    menu.MenuStatus,                         // 菜单状态
			IsDeleted:     menu.IsDeleted,                          // 是否删除  0：否  1：是
			IsVisible:     menu.IsVisible,                          // 是否可见  0：否  1：是
			Remark:        menu.Remark,                             // 备注信息
			VuePath:       menu.VuePath,                            // vue系统的path
			VueComponent:  menu.VueComponent,                       // vue的页面
			VueIcon:       menu.VueIcon,                            // vue的图标
			VueRedirect:   menu.VueRedirect,                        // vue的路由重定向
			BackgroundUrl: menu.BackgroundURL,                      // 接口地址
		})
	}

	return &sysclient.QueryMenuListResp{
		Total: 0,
		List:  list,
	}, nil

}
