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
func (l *UserInfoLogic) UserInfo(in *sysclient.InfoReq) (*sysclient.InfoResp, error) {
	//1.根据id查询用户信息
	q := query.SysUser
	info, err := q.WithContext(l.ctx).Where(q.ID.Eq(in.UserId)).First()

	// 2.判断用户是否存在
	if errors.Is(err, gorm.ErrRecordNotFound) {
		logc.Errorf(l.ctx, "用户不存在,参数：%+v,异常:%s", in, err.Error())
		return nil, errors.New("用户不存在")
	}

	if err != nil {
		logc.Errorf(l.ctx, "查询用户信息,参数：%+v,异常:%s", in, err.Error())
		return nil, errors.New("查询用户信息异常")
	}

	//3.查询用户菜单和权限
	menuList, apiUrls := l.queryUserMenuAndApiUrls(in.UserId)

	return &sysclient.InfoResp{
		Avatar:         info.Avatar,
		Name:           info.UserName,
		MenuListTree:   menuList,
		BackgroundUrls: apiUrls,
	}, nil
}

// 查询用户菜单和权限
func (l *UserInfoLogic) queryUserMenuAndApiUrls(userId int64) ([]*sysclient.MenuListTree, []string) {
	var result []*model.SysMenu
	if common.IsAdmin(l.ctx, userId, l.svcCtx.DB) {
		result, _ = query.SysMenu.WithContext(l.ctx).Find()
	} else {
		sql := `
				select sm.*
				from sys_user_role sur
						 left join sys_role sr on sur.role_id = sr.id
						 left join sys_role_menu srm on sr.id = srm.role_id
						 left join sys_menu sm on srm.menu_id = sm.id
				where sur.user_id = ?
				order by sm.id
				`
		db := l.svcCtx.DB
		db.WithContext(l.ctx).Raw(sql, userId).Scan(&result)
	}
	return buildMenuTree(result)
}

// 构建返回值
func buildMenuTree(menus []*model.SysMenu) ([]*sysclient.MenuListTree, []string) {
	var menuListTrees []*sysclient.MenuListTree
	var urls []string
	for _, menu := range menus {
		if menu.MenuType == 1 || menu.MenuType == 0 {
			menuListTrees = append(menuListTrees, &sysclient.MenuListTree{
				Id:           menu.ID,
				Name:         menu.MenuName,
				Icon:         menu.MenuIcon,
				ParentId:     menu.ParentID,
				Path:         menu.MenuPath,
				VuePath:      menu.VuePath,
				VueComponent: menu.VueComponent,
				VueIcon:      menu.VueIcon,
				VueRedirect:  menu.VueRedirect,
			})
		}

		if len(strings.TrimSpace(menu.BackgroundURL)) != 0 {
			urls = append(urls, menu.BackgroundURL)
		}

	}
	return menuListTrees, urls
}
