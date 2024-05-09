package userservicelogic

import (
	"context"
	"errors"
	"fmt"
	"github.com/feihua/zero-admin/rpc/sys/gen/model"
	"github.com/feihua/zero-admin/rpc/sys/gen/query"
	"github.com/feihua/zero-admin/rpc/sys/internal/svc"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"strings"
)

// UserInfoLogic
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
	//根据id查询用户信息
	q := query.SysUser
	userInfo, err := q.WithContext(l.ctx).Where(q.ID.Eq(in.UserId)).First()

	switch err {
	case nil:
	case sqlc.ErrNotFound:
		logc.Errorf(l.ctx, "用户不存在userId: %s", in.UserId)
		return nil, errors.New(fmt.Sprintf("用户不存在userId: %d", in.UserId))
	default:
		return nil, err
	}

	var list []*sysclient.MenuListTree
	var listUrls []string

	//id为1是系统预留超级管理员,它获取的是全部权限
	if in.UserId == 1 {
		menus, _ := query.SysMenu.WithContext(l.ctx).Find()
		list, listUrls = listTrees(menus, list, listUrls)
		logc.Infof(l.ctx, "超级管理员: %s登录,菜单: %+v", userInfo.Name, list)
	} else {
		var result []*model.SysMenu
		sql := "select sm.* from sys_user_role sur left join sys_role sr on sur.role_id = sr.id left join sys_role_menu srm on sr.id = srm.role_id left join sys_menu sm on srm.menu_id = sm.id where sur.user_id=? order by sm.id"
		db := l.svcCtx.DB
		_ = db.WithContext(l.ctx).Raw(sql, in.UserId).Scan(&result).Error
		list, listUrls = listTrees(result, list, listUrls)
		logc.Infof(l.ctx, "普通管理员: %s登录,菜单: %+v", userInfo.Name, list)
	}

	return &sysclient.InfoResp{
		Avatar:         "https://gw.alipayobjects.com/zos/antfincdn/XAosXuNZyF/BiazfanxmamNRoxxVxka.png",
		Name:           userInfo.Name,
		MenuListTree:   list,
		BackgroundUrls: listUrls,
	}, nil
}

func listTrees(menus []*model.SysMenu, list []*sysclient.MenuListTree, listUrls []string) ([]*sysclient.MenuListTree, []string) {
	for _, menu := range menus {
		if menu.Type == 1 || menu.Type == 0 {
			list = append(list, &sysclient.MenuListTree{
				Id:           menu.ID,
				Name:         menu.Name,
				Icon:         *menu.Icon,
				ParentId:     menu.ParentID,
				Path:         *menu.URL,
				VuePath:      *menu.VuePath,
				VueComponent: *menu.VueComponent,
				VueIcon:      *menu.VueIcon,
				VueRedirect:  *menu.VueRedirect,
			})
		}

		if len(strings.TrimSpace(menu.BackgroundURL)) != 0 {
			listUrls = append(listUrls, menu.BackgroundURL)
		}

	}
	return list, listUrls
}
