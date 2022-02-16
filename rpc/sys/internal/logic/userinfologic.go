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
	userInfo, err := l.svcCtx.UserModel.FindOne(in.UserId)

	switch err {
	case nil:
	case sqlc.ErrNotFound:
		logx.WithContext(l.ctx).Infof("用户不存在userId: %s", in.UserId)
		return nil, errors.New(fmt.Sprintf("用户不存在userId: %s", strconv.FormatInt(in.UserId, 10)))
	default:
		return nil, err
	}

	var list []*sys.MenuListTree
	var listUrls []string

	if in.UserId == 1 {
		menus, _ := l.svcCtx.MenuModel.FindAll(1, 1000)
		list, listUrls = listTrees(menus, list, listUrls)
		logx.WithContext(l.ctx).Infof("超级管理员: %s登录,菜单: %+v", userInfo.Name, list)
	} else {
		menus, _ := l.svcCtx.MenuModel.FindAllByUserId(in.UserId)
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
		list = append(list, &sys.MenuListTree{
			Id:           menu.Id,
			Name:         menu.Name,
			Icon:         menu.Icon,
			ParentId:     menu.ParentId,
			Path:         menu.Url,
			VuePath:      menu.VuePath,
			VueComponent: menu.VueComponent,
			VueIcon:      menu.VueIcon,
			VueRedirect:  menu.VueRedirect,
		})

		if len(strings.TrimSpace(menu.BackgroundUrl)) != 0 {
			listUrls = append(listUrls, menu.BackgroundUrl)
		}

	}
	return list, listUrls
}
