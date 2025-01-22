package userservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/sys/gen/model"
	"github.com/feihua/zero-admin/rpc/sys/gen/query"
	"github.com/feihua/zero-admin/rpc/sys/internal/logic/common"
	"github.com/feihua/zero-admin/rpc/sys/internal/svc"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
	"strings"
)

// UserInfoLogic 获取用户信息
/*
Author: LiuFeiHua
Date: 2023/12/18 14:30
*/
type UserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UserInfo 获取用户信息
// 1.根据id查询用户信息
// 2.查询用户菜单和权限
func (l *UserInfoLogic) UserInfo(in *sysclient.InfoReq) (*sysclient.InfoResp, error) {
	// 1.根据id查询用户信息
	q := query.SysUser
	info, err := q.WithContext(l.ctx).Where(q.ID.Eq(in.UserId)).First()

	if errors.Is(err, gorm.ErrRecordNotFound) {
		logc.Errorf(l.ctx, "用户不存在,参数：%+v,异常:%s", in, err.Error())
		return nil, errors.New("用户不存在")
	}

	if err != nil {
		logc.Errorf(l.ctx, "查询用户信息,参数：%+v,异常:%s", in, err.Error())
		return nil, errors.New("查询用户信息异常")
	}

	// 2.查询用户菜单和权限
	menuList, apiUrls := l.queryApis(in.UserId)

	return &sysclient.InfoResp{
		Avatar:         info.Avatar,   // 头像
		Name:           info.UserName, // 用户名
		MenuListTree:   menuList,      // 菜单
		BackgroundUrls: apiUrls,       // 权限
	}, nil
}

// 查询用户菜单和权限
func (l *UserInfoLogic) queryApis(userId int64) ([]*sysclient.MenuListTree, []string) {
	var result []*model.SysMenu
	if common.IsAdmin(l.ctx, userId, l.svcCtx.DB) {
		result, _ = query.SysMenu.WithContext(l.ctx).Where(query.SysMenu.IsVisible.Eq(1)).Find()
	} else {
		sql := `
				select sm.*
				from sys_user_role sur
						 left join sys_role sr on sur.role_id = sr.id
						 left join sys_role_menu srm on sr.id = srm.role_id
						 left join sys_menu sm on srm.menu_id = sm.id
				where sur.user_id = ? and sm.is_visible=1
				order by sm.id
				`
		db := l.svcCtx.DB
		db.WithContext(l.ctx).Raw(sql, userId).Scan(&result)
	}
	return buildMenuTree(result)
}

// 构建返回值
func buildMenuTree(menus []*model.SysMenu) ([]*sysclient.MenuListTree, []string) {
	var menuListTrees = make([]*sysclient.MenuListTree, 0, len(menus))
	var urls = make([]string, 0, len(menus))

	for _, menu := range menus {
		if menu.MenuType == 1 || menu.MenuType == 0 {
			menuListTrees = append(menuListTrees, &sysclient.MenuListTree{
				Id:           menu.ID,           // 菜单id
				Name:         menu.MenuName,     // 菜单名称
				Icon:         menu.MenuIcon,     // 图标
				ParentId:     menu.ParentID,     // 父级id
				Path:         menu.MenuPath,     // 路径
				VuePath:      menu.VuePath,      // vue路径
				VueComponent: menu.VueComponent, // vue组件
				VueIcon:      menu.VueIcon,      // vue图标
				VueRedirect:  menu.VueRedirect,  // vue的路由重定向
			})
		}

		if len(strings.TrimSpace(menu.BackgroundURL)) != 0 {
			urls = append(urls, menu.BackgroundURL) // 接口地址
		}

	}
	return menuListTrees, urls
}
