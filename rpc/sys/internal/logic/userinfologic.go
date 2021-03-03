package logic

import (
	"context"
	"errors"
	"github.com/tal-tech/go-zero/core/stores/sqlc"

	"go-zero-admin/rpc/sys/internal/svc"
	"go-zero-admin/rpc/sys/sys"

	"github.com/tal-tech/go-zero/core/logx"
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
		return nil, errors.New("用户名不存在")
	default:
		return nil, err
	}

	menus, _ := l.svcCtx.MenuModel.FindAll(1, 100)

	var list []*sys.MenuListTree
	for _, menu := range *menus {
		list = append(list, &sys.MenuListTree{
			Id:           menu.Id,
			Name:         menu.Name,
			Icon:         menu.Icon,
			ParentId:     menu.ParentId,
			Path:         menu.Url,
			VuePath:      menu.VuePath.String,
			VueComponent: menu.VueComponent.String,
			VueIcon:      menu.VueIcon.String,
			VueRedirect:  menu.VueRedirect.String,
		})
	}

	return &sys.InfoResp{
		Avatar:       "https://gw.alipayobjects.com/zos/antfincdn/XAosXuNZyF/BiazfanxmamNRoxxVxka.png",
		Name:         userInfo.Name,
		MenuListTree: list,
	}, nil
}
