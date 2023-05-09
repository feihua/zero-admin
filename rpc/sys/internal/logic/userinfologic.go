package logic

import (
	"context"
	"errors"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"strconv"
	"strings"
	"zero-admin/rpc/model/sysmodel"
	"zero-admin/rpc/sys/internal/svc"
	"zero-admin/rpc/sys/sys"
)

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

func (l *UserInfoLogic) UserInfo(in *sys.InfoReq) (*sys.InfoResp, error) {
	userInfo, err := l.svcCtx.UserModel.FindOne(l.ctx, in.UserId)

	switch err {
	case nil:
	case sqlc.ErrNotFound:
		logx.WithContext(l.ctx).Errorf("用户不存在userId: %s", in.UserId)
		return nil, errors.New(fmt.Sprintf("用户不存在userId: %s", strconv.FormatInt(in.UserId, 10)))
	default:
		return nil, err
	}

	var list []*sys.MenuListTree
	var listUrls []string

	if in.UserId == 1 {
		menus, _ := l.svcCtx.MenuModel.FindAll(l.ctx, 1, 1000)
		list, listUrls = listTrees(menus, list, listUrls)
		logx.WithContext(l.ctx).Infof("超级管理员: %s登录,菜单: %+v", userInfo.Name, list)
	} else {
		menus, _ := l.svcCtx.MenuModel.FindAllByUserId(l.ctx, in.UserId)
		list, listUrls = listTrees(menus, list, listUrls)
		logx.WithContext(l.ctx).Infof("普通管理员: %s登录,菜单: %+v", userInfo.Name, list)
	}

	return &sys.InfoResp{
		Avatar:         "https://gw.alipayobjects.com/zos/antfincdn/XAosXuNZyF/BiazfanxmamNRoxxVxka.png",
		Name:           userInfo.Name,
		MenuListTree:   list,
		BackgroundUrls: listUrls,
	}, nil
}

func listTrees(menus *[]sysmodel.SysMenu, list []*sys.MenuListTree, listUrls []string) ([]*sys.MenuListTree, []string) {
	for _, menu := range *menus {
		if menu.Type == 1 || menu.Type == 0 {
			list = append(list, &sys.MenuListTree{
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
