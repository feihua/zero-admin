package userservicelogic

import (
	"context"
	"errors"
	"fmt"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"strings"
	"zero-admin/rpc/model/sysmodel"
	"zero-admin/rpc/sys/internal/svc"
	"zero-admin/rpc/sys/sysclient"
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
	userInfo, err := l.svcCtx.UserModel.FindOne(l.ctx, in.UserId)

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
		menus, _ := l.svcCtx.MenuModel.FindAll(l.ctx, 1, 1000)
		list, listUrls = listTrees(menus, list, listUrls)
		logx.WithContext(l.ctx).Infof("超级管理员: %s登录,菜单: %+v", userInfo.Name, list)
	} else {
		menus, _ := l.svcCtx.MenuModel.FindAllByUserId(l.ctx, in.UserId)
		list, listUrls = listTrees(menus, list, listUrls)
		logx.WithContext(l.ctx).Infof("普通管理员: %s登录,菜单: %+v", userInfo.Name, list)
	}

	return &sysclient.InfoResp{
		Avatar:         "https://gw.alipayobjects.com/zos/antfincdn/XAosXuNZyF/BiazfanxmamNRoxxVxka.png",
		Name:           userInfo.Name,
		MenuListTree:   list,
		BackgroundUrls: listUrls,
	}, nil
}

func listTrees(menus *[]sysmodel.SysMenu, list []*sysclient.MenuListTree, listUrls []string) ([]*sysclient.MenuListTree, []string) {
	for _, menu := range *menus {
		if menu.Type == 1 || menu.Type == 0 {
			list = append(list, &sysclient.MenuListTree{
				Id:           menu.Id,
				Name:         menu.Name,
				Icon:         menu.Icon.String,
				ParentId:     menu.ParentId,
				Path:         menu.Url.String,
				VuePath:      menu.VuePath.String,
				VueComponent: menu.VueComponent.String,
				VueIcon:      menu.VueIcon.String,
				VueRedirect:  menu.VueRedirect.String,
			})
		}

		if len(strings.TrimSpace(menu.BackgroundUrl)) != 0 {
			listUrls = append(listUrls, menu.BackgroundUrl)
		}

	}
	return list, listUrls
}
